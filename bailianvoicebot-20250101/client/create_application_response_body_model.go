// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateApplicationResponseBody
	GetCode() *string
	SetData(v string) *CreateApplicationResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateApplicationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApplicationResponseBody
	GetRequestId() *string
}

type CreateApplicationResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// a395011f-a247-400f-bc69-28796749fd52
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CF6D3484-19A1-5C77-863B-AC8B5754D37C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateApplicationResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateApplicationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationResponseBody) SetCode(v string) *CreateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApplicationResponseBody) SetData(v string) *CreateApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *CreateApplicationResponseBody) SetHttpStatusCode(v int32) *CreateApplicationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateApplicationResponseBody) SetMessage(v string) *CreateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
