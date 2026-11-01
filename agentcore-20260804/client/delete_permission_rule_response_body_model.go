// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePermissionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePermissionRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeletePermissionRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeletePermissionRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePermissionRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePermissionRuleResponseBody
	GetSuccess() *bool
}

type DeletePermissionRuleResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The message about the request processing result.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeletePermissionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePermissionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePermissionRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePermissionRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeletePermissionRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePermissionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePermissionRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePermissionRuleResponseBody) SetCode(v string) *DeletePermissionRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePermissionRuleResponseBody) SetHttpStatusCode(v int32) *DeletePermissionRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeletePermissionRuleResponseBody) SetMessage(v string) *DeletePermissionRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePermissionRuleResponseBody) SetRequestId(v string) *DeletePermissionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePermissionRuleResponseBody) SetSuccess(v bool) *DeletePermissionRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePermissionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
