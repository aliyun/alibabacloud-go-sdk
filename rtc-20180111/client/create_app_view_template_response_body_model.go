// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppViewTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAppViewTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *CreateAppViewTemplateResponseBody
	GetTemplateId() *string
}

type CreateAppViewTemplateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 30D41049-D02D-1C21-86AE-B3E5FD805C27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// bc5v****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateAppViewTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppViewTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppViewTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppViewTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateAppViewTemplateResponseBody) SetRequestId(v string) *CreateAppViewTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppViewTemplateResponseBody) SetTemplateId(v string) *CreateAppViewTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateAppViewTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
