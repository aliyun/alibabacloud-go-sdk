// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateApplicationVersionResponseBody
	GetCode() *string
	SetData(v string) *CreateApplicationVersionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateApplicationVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateApplicationVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApplicationVersionResponseBody
	GetRequestId() *string
}

type CreateApplicationVersionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 20904943-f711-494f-9f1f-e7f340f37707
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

func (s CreateApplicationVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateApplicationVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateApplicationVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateApplicationVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApplicationVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationVersionResponseBody) SetCode(v string) *CreateApplicationVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApplicationVersionResponseBody) SetData(v string) *CreateApplicationVersionResponseBody {
	s.Data = &v
	return s
}

func (s *CreateApplicationVersionResponseBody) SetHttpStatusCode(v int32) *CreateApplicationVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateApplicationVersionResponseBody) SetMessage(v string) *CreateApplicationVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApplicationVersionResponseBody) SetRequestId(v string) *CreateApplicationVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
