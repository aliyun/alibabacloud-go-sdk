// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreserveHeaderFormatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PreserveHeaderFormatResponseBody
	GetCode() *int32
	SetData(v bool) *PreserveHeaderFormatResponseBody
	GetData() *bool
	SetDynamicCode(v string) *PreserveHeaderFormatResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *PreserveHeaderFormatResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *PreserveHeaderFormatResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *PreserveHeaderFormatResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PreserveHeaderFormatResponseBody
	GetMessage() *string
	SetRequestId(v string) *PreserveHeaderFormatResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PreserveHeaderFormatResponseBody
	GetSuccess() *bool
}

type PreserveHeaderFormatResponseBody struct {
	// The status code. A value of 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// >  The request parameter **DtsJobId*	- is invalid if **The Value of Input Parameter %s is not valid*	- is returned for **ErrMessage*	- and **DtsJobId*	- is returned for **DynamicMessage**.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The status code.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 69AD2AA7-DB47-449B-941B-B14409DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PreserveHeaderFormatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreserveHeaderFormatResponseBody) GoString() string {
	return s.String()
}

func (s *PreserveHeaderFormatResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PreserveHeaderFormatResponseBody) GetData() *bool {
	return s.Data
}

func (s *PreserveHeaderFormatResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *PreserveHeaderFormatResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *PreserveHeaderFormatResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PreserveHeaderFormatResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PreserveHeaderFormatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PreserveHeaderFormatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreserveHeaderFormatResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PreserveHeaderFormatResponseBody) SetCode(v int32) *PreserveHeaderFormatResponseBody {
	s.Code = &v
	return s
}

func (s *PreserveHeaderFormatResponseBody) SetData(v bool) *PreserveHeaderFormatResponseBody {
	s.Data = &v
	return s
}

func (s *PreserveHeaderFormatResponseBody) SetDynamicCode(v string) *PreserveHeaderFormatResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PreserveHeaderFormatResponseBody) SetDynamicMessage(v string) *PreserveHeaderFormatResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PreserveHeaderFormatResponseBody) SetErrorCode(v string) *PreserveHeaderFormatResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PreserveHeaderFormatResponseBody) SetHttpStatusCode(v int32) *PreserveHeaderFormatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PreserveHeaderFormatResponseBody) SetMessage(v string) *PreserveHeaderFormatResponseBody {
	s.Message = &v
	return s
}

func (s *PreserveHeaderFormatResponseBody) SetRequestId(v string) *PreserveHeaderFormatResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreserveHeaderFormatResponseBody) SetSuccess(v bool) *PreserveHeaderFormatResponseBody {
	s.Success = &v
	return s
}

func (s *PreserveHeaderFormatResponseBody) Validate() error {
	return dara.Validate(s)
}
