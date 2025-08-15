// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChangeResourceGroupResponseBody
	GetCode() *string
	SetData(v bool) *ChangeResourceGroupResponseBody
	GetData() *bool
	SetDynamicCode(v string) *ChangeResourceGroupResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ChangeResourceGroupResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ChangeResourceGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ChangeResourceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeResourceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeResourceGroupResponseBody
	GetSuccess() *bool
}

type ChangeResourceGroupResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// Instance.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeResourceGroupResponseBody) GetData() *bool {
	return s.Data
}

func (s *ChangeResourceGroupResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ChangeResourceGroupResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ChangeResourceGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ChangeResourceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeResourceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeResourceGroupResponseBody) SetCode(v string) *ChangeResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetData(v bool) *ChangeResourceGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetDynamicCode(v string) *ChangeResourceGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetDynamicMessage(v string) *ChangeResourceGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetHttpStatusCode(v int32) *ChangeResourceGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetMessage(v string) *ChangeResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetSuccess(v bool) *ChangeResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
