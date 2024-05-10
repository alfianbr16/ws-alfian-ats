package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	cek "github.com/alfianbr16/package_ats/module"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetAllmakanan_hewan(c *fiber.Ctx) error {
	ps := cek.GetAllMakanan_hewan()
	return c.JSON(ps)
	}