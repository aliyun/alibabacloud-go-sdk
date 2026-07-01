// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRCSTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateCode(v string) *QueryRCSTemplateRequest
	GetTemplateCode() *string
}

type QueryRCSTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s QueryRCSTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRCSTemplateRequest) GoString() string {
	return s.String()
}

func (s *QueryRCSTemplateRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *QueryRCSTemplateRequest) SetTemplateCode(v string) *QueryRCSTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *QueryRCSTemplateRequest) Validate() error {
	return dara.Validate(s)
}
