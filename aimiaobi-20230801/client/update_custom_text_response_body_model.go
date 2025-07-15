// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCustomTextResponseBody
	GetCode() *string
	SetData(v int64) *UpdateCustomTextResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateCustomTextResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateCustomTextResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCustomTextResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCustomTextResponseBody
	GetSuccess() *bool
}

type UpdateCustomTextResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 48
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCustomTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomTextResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomTextResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCustomTextResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateCustomTextResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateCustomTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCustomTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomTextResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCustomTextResponseBody) SetCode(v string) *UpdateCustomTextResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCustomTextResponseBody) SetData(v int64) *UpdateCustomTextResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateCustomTextResponseBody) SetHttpStatusCode(v int32) *UpdateCustomTextResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCustomTextResponseBody) SetMessage(v string) *UpdateCustomTextResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCustomTextResponseBody) SetRequestId(v string) *UpdateCustomTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomTextResponseBody) SetSuccess(v bool) *UpdateCustomTextResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCustomTextResponseBody) Validate() error {
	return dara.Validate(s)
}
