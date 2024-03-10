package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/kouhei-github/auto-generate-golang/repository"
)

type registerBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SaveUserHandler(c *fiber.Ctx) error {
	requestBody := c.Body()
	var body registerBody
	if err := json.Unmarshal(requestBody, &body); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("errorだよ")
	}
	user := repository.User{Email: body.Email, Password: body.Password}
	// メールアドレスが存在するか確認
	users, err := user.FindByEmail()
	if len(users) > 0 {
		return c.Status(fiber.StatusBadRequest).SendString("メールアドレスは既に存在します")
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("データベースエラー")
	}
	// パスワードのハッシュ化

	// ユーザーを保存する処理
	if err := user.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("保存に失敗しました")
	}
	return c.Status(fiber.StatusCreated).JSON(body)
}
