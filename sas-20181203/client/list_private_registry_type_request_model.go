// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateRegistryTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListPrivateRegistryTypeRequest
	GetLang() *string
}

type ListPrivateRegistryTypeRequest struct {
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

func (s ListPrivateRegistryTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateRegistryTypeRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateRegistryTypeRequest) GetLang() *string {
	return s.Lang
}

func (s *ListPrivateRegistryTypeRequest) SetLang(v string) *ListPrivateRegistryTypeRequest {
	s.Lang = &v
	return s
}

func (s *ListPrivateRegistryTypeRequest) Validate() error {
	return dara.Validate(s)
}
