// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListInstanceRegionRequest
	GetLang() *string
}

type ListInstanceRegionRequest struct {
	// The language used for response parameters. Set this parameter to `zh-CN`.
	//
	// example:
	//
	// zh-CN
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListInstanceRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRegionRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRegionRequest) GetLang() *string {
	return s.Lang
}

func (s *ListInstanceRegionRequest) SetLang(v string) *ListInstanceRegionRequest {
	s.Lang = &v
	return s
}

func (s *ListInstanceRegionRequest) Validate() error {
	return dara.Validate(s)
}
