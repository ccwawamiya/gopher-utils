// Copyright © 2016年 kuaifazs.com. All rights reserved.
// 
// @Author: wuchengshuang@kuaifzs.com
// @Date: 2017/2/20
// @Version: -
// @Desc: -

package gopher_utils

import (
    "regexp"
    "errors"
)

type Account struct {
    Column string
    Value  string
}

func CheckAccount(accountname string) (*Account, error) {
    account := Account{Value:accountname}
    if mobile(accountname) {
        account.Column = "mobile"
    }else if email(accountname) {
        account.Column = "mail"
    }else if username(accountname) {
        account.Column = "user_name"
    }else {
        return nil,errors.New("error account")
    }
    return &account,nil
}

func username(username string) bool {
    reg := regexp.MustCompile(`^[0-9a-zA-Z]{4,20}$`)
    return reg.MatchString(username)
}

func mobile(mobile string) bool {
    reg := regexp.MustCompile(`^(13[0-9]|15[0-9]|18[0-9]|14[0-9]|17[0-9])\d{8}$`)
    return reg.MatchString(mobile)
}

func email(email string) bool {
    reg := regexp.MustCompile(`^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$`)
    return reg.MatchString(email)
}