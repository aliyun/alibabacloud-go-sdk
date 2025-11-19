// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTemplateResult
	GetCode() *string
	SetData(v *Template) *DeleteTemplateResult
	GetData() *Template
	SetRequestId(v string) *DeleteTemplateResult
	GetRequestId() *string
}

type DeleteTemplateResult struct {
	// SUCCESS 为成功
	Code *string   `json:"code,omitempty" xml:"code,omitempty"`
	Data *Template `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteTemplateResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateResult) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResult) GetCode() *string {
	return s.Code
}

func (s *DeleteTemplateResult) GetData() *Template {
	return s.Data
}

func (s *DeleteTemplateResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTemplateResult) SetCode(v string) *DeleteTemplateResult {
	s.Code = &v
	return s
}

func (s *DeleteTemplateResult) SetData(v *Template) *DeleteTemplateResult {
	s.Data = v
	return s
}

func (s *DeleteTemplateResult) SetRequestId(v string) *DeleteTemplateResult {
	s.RequestId = &v
	return s
}

func (s *DeleteTemplateResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
