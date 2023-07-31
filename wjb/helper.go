/*
 * @Author: jarvenwang wang_jin_bao@163.com
 * @Date: 2023-07-31 16:36:12
 * @LastEditors: jarvenwang wang_jin_bao@163.com
 * @LastEditTime: 2023-07-31 16:39:10
 * @FilePath: /helper/helper.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */

package helper

import "fmt"

const version = "v1.0.0"

func Version() string {
	return version
}

func Test() bool {
	str := Version()
	fmt.Println(str)
	return true
}
