// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceExpireTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceExpireTimeResponseBody
	GetCode() *string
	SetFound(v bool) *GetInstanceExpireTimeResponseBody
	GetFound() *bool
	SetInstanceExpireTime(v string) *GetInstanceExpireTimeResponseBody
	GetInstanceExpireTime() *string
	SetInstanceId(v string) *GetInstanceExpireTimeResponseBody
	GetInstanceId() *string
	SetInstanceStatus(v string) *GetInstanceExpireTimeResponseBody
	GetInstanceStatus() *string
	SetMessage(v string) *GetInstanceExpireTimeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceExpireTimeResponseBody
	GetRequestId() *string
	SetTenantId(v int64) *GetInstanceExpireTimeResponseBody
	GetTenantId() *int64
}

type GetInstanceExpireTimeResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether a standard package instance is found.
	//
	// example:
	//
	// true
	Found *bool `json:"found,omitempty" xml:"found,omitempty"`
	// The expiration time of the instance in ISO format.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	InstanceExpireTime *string `json:"instanceExpireTime,omitempty" xml:"instanceExpireTime,omitempty"`
	// The instance ID. This parameter is required.
	//
	// example:
	//
	// exampleInstanceId
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance status. Valid values:
	//
	// - RUNNING: Running.
	//
	// - TERMINATED: Terminated.
	//
	// - COMPLETED: Completed.
	//
	// - ERROR: Error.
	//
	// example:
	//
	// string_value
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetInstanceExpireTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceExpireTimeResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceExpireTimeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceExpireTimeResponseBody) GetFound() *bool {
	return s.Found
}

func (s *GetInstanceExpireTimeResponseBody) GetInstanceExpireTime() *string {
	return s.InstanceExpireTime
}

func (s *GetInstanceExpireTimeResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceExpireTimeResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetInstanceExpireTimeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceExpireTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceExpireTimeResponseBody) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetInstanceExpireTimeResponseBody) SetCode(v string) *GetInstanceExpireTimeResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceExpireTimeResponseBody) SetFound(v bool) *GetInstanceExpireTimeResponseBody {
	s.Found = &v
	return s
}

func (s *GetInstanceExpireTimeResponseBody) SetInstanceExpireTime(v string) *GetInstanceExpireTimeResponseBody {
	s.InstanceExpireTime = &v
	return s
}

func (s *GetInstanceExpireTimeResponseBody) SetInstanceId(v string) *GetInstanceExpireTimeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceExpireTimeResponseBody) SetInstanceStatus(v string) *GetInstanceExpireTimeResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceExpireTimeResponseBody) SetMessage(v string) *GetInstanceExpireTimeResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceExpireTimeResponseBody) SetRequestId(v string) *GetInstanceExpireTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceExpireTimeResponseBody) SetTenantId(v int64) *GetInstanceExpireTimeResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetInstanceExpireTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
