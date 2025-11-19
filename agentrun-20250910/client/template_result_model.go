// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TemplateResult
	GetCode() *string
	SetData(v *Template) *TemplateResult
	GetData() *Template
	SetRequestId(v string) *TemplateResult
	GetRequestId() *string
}

type TemplateResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 模板的详细信息
	//
	// This parameter is required.
	Data *Template `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s TemplateResult) String() string {
	return dara.Prettify(s)
}

func (s TemplateResult) GoString() string {
	return s.String()
}

func (s *TemplateResult) GetCode() *string {
	return s.Code
}

func (s *TemplateResult) GetData() *Template {
	return s.Data
}

func (s *TemplateResult) GetRequestId() *string {
	return s.RequestId
}

func (s *TemplateResult) SetCode(v string) *TemplateResult {
	s.Code = &v
	return s
}

func (s *TemplateResult) SetData(v *Template) *TemplateResult {
	s.Data = v
	return s
}

func (s *TemplateResult) SetRequestId(v string) *TemplateResult {
	s.RequestId = &v
	return s
}

func (s *TemplateResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
