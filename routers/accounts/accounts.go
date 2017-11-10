// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package accounts

import (
	"github.com/zhuharev/qiwi-admin/models"
	"github.com/zhuharev/qiwi-admin/pkg/context"
)

// Create an user
func Create(ctx *context.Context, af models.AuthForm) {
	if ctx.Req.Method == "POST" {
		_, err := models.CreateUser(af)
		if ctx.HasError(err) {
			return
		}
		ctx.Redirect("/dashboard")
	}
	ctx.Redirect("/dashboard")
}

// Setting shows user setting
func Setting(ctx *context.Context) {
	ctx.HTML(200, "account/setting")
}
