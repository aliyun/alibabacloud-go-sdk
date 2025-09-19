// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageBuildRiskItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListImageBuildRiskItemRequest
	GetLang() *string
}

type ListImageBuildRiskItemRequest struct {
	// The language of the content within the request and response. Default value: zh. Valid values:
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

func (s ListImageBuildRiskItemRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageBuildRiskItemRequest) GoString() string {
	return s.String()
}

func (s *ListImageBuildRiskItemRequest) GetLang() *string {
	return s.Lang
}

func (s *ListImageBuildRiskItemRequest) SetLang(v string) *ListImageBuildRiskItemRequest {
	s.Lang = &v
	return s
}

func (s *ListImageBuildRiskItemRequest) Validate() error {
	return dara.Validate(s)
}
