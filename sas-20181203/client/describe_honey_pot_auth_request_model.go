// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHoneyPotAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeHoneyPotAuthRequest
	GetLang() *string
}

type DescribeHoneyPotAuthRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeHoneyPotAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHoneyPotAuthRequest) GoString() string {
	return s.String()
}

func (s *DescribeHoneyPotAuthRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeHoneyPotAuthRequest) SetLang(v string) *DescribeHoneyPotAuthRequest {
	s.Lang = &v
	return s
}

func (s *DescribeHoneyPotAuthRequest) Validate() error {
	return dara.Validate(s)
}
