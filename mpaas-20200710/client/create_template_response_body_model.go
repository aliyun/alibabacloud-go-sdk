// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTemplateResponseBody
	GetCode() *string
	SetData(v string) *CreateTemplateResponseBody
	GetData() *string
	SetMsg(v string) *CreateTemplateResponseBody
	GetMsg() *string
	SetRequestId(v string) *CreateTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTemplateResponseBody
	GetSuccess() *bool
}

type CreateTemplateResponseBody struct {
	// example:
	//
	// 100
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// SUCCESS
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0CC8A9CB-9BA3-13FD-A404-6E2E7461881A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTemplateResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateTemplateResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CreateTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTemplateResponseBody) SetCode(v string) *CreateTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTemplateResponseBody) SetData(v string) *CreateTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTemplateResponseBody) SetMsg(v string) *CreateTemplateResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetSuccess(v bool) *CreateTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
