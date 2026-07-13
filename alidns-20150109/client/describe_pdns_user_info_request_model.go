// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePdnsUserInfoRequest
	GetLang() *string
}

type DescribePdnsUserInfoRequest struct {
	// The language of the request and response. The default value is **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribePdnsUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsUserInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribePdnsUserInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePdnsUserInfoRequest) SetLang(v string) *DescribePdnsUserInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribePdnsUserInfoRequest) Validate() error {
	return dara.Validate(s)
}
