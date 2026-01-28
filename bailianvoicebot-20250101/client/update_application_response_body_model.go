// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateApplicationResponseBody
	GetCode() *string
	SetData(v string) *UpdateApplicationResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateApplicationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApplicationResponseBody
	GetRequestId() *string
}

type UpdateApplicationResponseBody struct {
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

func (s UpdateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateApplicationResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateApplicationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationResponseBody) SetCode(v string) *UpdateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApplicationResponseBody) SetData(v string) *UpdateApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateApplicationResponseBody) SetHttpStatusCode(v int32) *UpdateApplicationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateApplicationResponseBody) SetMessage(v string) *UpdateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApplicationResponseBody) SetRequestId(v string) *UpdateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
