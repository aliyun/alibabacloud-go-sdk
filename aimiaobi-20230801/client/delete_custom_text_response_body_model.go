// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCustomTextResponseBody
	GetCode() *string
	SetData(v bool) *DeleteCustomTextResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *DeleteCustomTextResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCustomTextResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomTextResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomTextResponseBody
	GetSuccess() *bool
}

type DeleteCustomTextResponseBody struct {
	// The status code.
	//
	// example:
	//
	// successful
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the text was deleted successfully.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error description.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. A value of true indicates success. A value of false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTextResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomTextResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCustomTextResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteCustomTextResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCustomTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomTextResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomTextResponseBody) SetCode(v string) *DeleteCustomTextResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomTextResponseBody) SetData(v bool) *DeleteCustomTextResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCustomTextResponseBody) SetHttpStatusCode(v int32) *DeleteCustomTextResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCustomTextResponseBody) SetMessage(v string) *DeleteCustomTextResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomTextResponseBody) SetRequestId(v string) *DeleteCustomTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomTextResponseBody) SetSuccess(v bool) *DeleteCustomTextResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomTextResponseBody) Validate() error {
	return dara.Validate(s)
}
