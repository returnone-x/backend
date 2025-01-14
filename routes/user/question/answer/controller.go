package questionAnswer

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app fiber.Router) {

	app.Post("/new/:question_id", NewAnswer)
	app.Delete("/delete/:answer_id", DeleteAnswer)

	app.Post("/upvote/:answer_id", UpVote)
	app.Post("/downvote/:answer_id", DownVote)
	app.Delete("/deletevote/:answer_id", DeleteVote)

	app.Get("/", GetQuestion)
}
