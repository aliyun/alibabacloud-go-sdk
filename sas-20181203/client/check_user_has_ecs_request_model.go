// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserHasEcsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CheckUserHasEcsRequest
	GetLang() *string
}

type CheckUserHasEcsRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s CheckUserHasEcsRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckUserHasEcsRequest) GoString() string {
	return s.String()
}

func (s *CheckUserHasEcsRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckUserHasEcsRequest) SetLang(v string) *CheckUserHasEcsRequest {
	s.Lang = &v
	return s
}

func (s *CheckUserHasEcsRequest) Validate() error {
	return dara.Validate(s)
}
