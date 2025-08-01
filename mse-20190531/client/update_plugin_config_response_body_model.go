// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePluginConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdatePluginConfigResponseBody
	GetCode() *int32
	SetData(v int64) *UpdatePluginConfigResponseBody
	GetData() *int64
	SetDynamicMessage(v string) *UpdatePluginConfigResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *UpdatePluginConfigResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *UpdatePluginConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdatePluginConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePluginConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePluginConfigResponseBody
	GetSuccess() *bool
}

type UpdatePluginConfigResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the plug-in configuration.
	//
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// code
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code that is returned.
	//
	// example:
	//
	// 500
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 03A3E2F4-6804-5663-9D5D-2EC47A1*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePluginConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePluginConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePluginConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdatePluginConfigResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdatePluginConfigResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdatePluginConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdatePluginConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdatePluginConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePluginConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePluginConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePluginConfigResponseBody) SetCode(v int32) *UpdatePluginConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePluginConfigResponseBody) SetData(v int64) *UpdatePluginConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePluginConfigResponseBody) SetDynamicMessage(v string) *UpdatePluginConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdatePluginConfigResponseBody) SetErrorCode(v string) *UpdatePluginConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePluginConfigResponseBody) SetHttpStatusCode(v int32) *UpdatePluginConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdatePluginConfigResponseBody) SetMessage(v string) *UpdatePluginConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePluginConfigResponseBody) SetRequestId(v string) *UpdatePluginConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePluginConfigResponseBody) SetSuccess(v bool) *UpdatePluginConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePluginConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
