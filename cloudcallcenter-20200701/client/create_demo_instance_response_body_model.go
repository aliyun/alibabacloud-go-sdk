// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDemoInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDemoInstanceResponseBody
	GetCode() *string
	SetData(v string) *CreateDemoInstanceResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateDemoInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDemoInstanceResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateDemoInstanceResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateDemoInstanceResponseBody
	GetRequestId() *string
}

type CreateDemoInstanceResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string   `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDemoInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDemoInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDemoInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDemoInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateDemoInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDemoInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDemoInstanceResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateDemoInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDemoInstanceResponseBody) SetCode(v string) *CreateDemoInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDemoInstanceResponseBody) SetData(v string) *CreateDemoInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDemoInstanceResponseBody) SetHttpStatusCode(v int32) *CreateDemoInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDemoInstanceResponseBody) SetMessage(v string) *CreateDemoInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDemoInstanceResponseBody) SetParams(v []*string) *CreateDemoInstanceResponseBody {
	s.Params = v
	return s
}

func (s *CreateDemoInstanceResponseBody) SetRequestId(v string) *CreateDemoInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDemoInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
