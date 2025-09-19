// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupportedModulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetSupportedModulesRequest
	GetLang() *string
}

type GetSupportedModulesRequest struct {
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

func (s GetSupportedModulesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSupportedModulesRequest) GoString() string {
	return s.String()
}

func (s *GetSupportedModulesRequest) GetLang() *string {
	return s.Lang
}

func (s *GetSupportedModulesRequest) SetLang(v string) *GetSupportedModulesRequest {
	s.Lang = &v
	return s
}

func (s *GetSupportedModulesRequest) Validate() error {
	return dara.Validate(s)
}
