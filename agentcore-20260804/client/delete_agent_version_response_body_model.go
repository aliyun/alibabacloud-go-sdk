// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAgentVersionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteAgentVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteAgentVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAgentVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAgentVersionResponseBody
	GetSuccess() *bool
}

type DeleteAgentVersionResponseBody struct {
	// The business status code. The value is SUCCESS if the request is successful, or an error code if the request fails.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message. The value is success if the request is successful, or an error message if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is used to locate and troubleshoot the request.
	//
	// example:
	//
	// 5E8F3C2A-1B4D-4F6E-9A7C-0D2B8E6F4A1C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteAgentVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAgentVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteAgentVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAgentVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAgentVersionResponseBody) SetCode(v string) *DeleteAgentVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAgentVersionResponseBody) SetHttpStatusCode(v int32) *DeleteAgentVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAgentVersionResponseBody) SetMessage(v string) *DeleteAgentVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAgentVersionResponseBody) SetRequestId(v string) *DeleteAgentVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentVersionResponseBody) SetSuccess(v bool) *DeleteAgentVersionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAgentVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
