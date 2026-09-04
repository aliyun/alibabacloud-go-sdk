// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckHealthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthSource(v string) *CheckHealthResponseBody
	GetAuthSource() *string
	SetCallerType(v string) *CheckHealthResponseBody
	GetCallerType() *string
	SetCode(v string) *CheckHealthResponseBody
	GetCode() *string
	SetDigitalEmployeeName(v string) *CheckHealthResponseBody
	GetDigitalEmployeeName() *string
	SetMessage(v string) *CheckHealthResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckHealthResponseBody
	GetRequestId() *string
	SetTenantId(v int64) *CheckHealthResponseBody
	GetTenantId() *int64
	SetUserId(v int64) *CheckHealthResponseBody
	GetUserId() *int64
}

type CheckHealthResponseBody struct {
	// The authentication source. Valid values: bearer and aliyun_gateway.
	//
	// example:
	//
	// aliyun_gateway
	AuthSource *string `json:"authSource,omitempty" xml:"authSource,omitempty"`
	// The caller type. Valid values: user, aliyun_main, aliyun_ram, and service.
	//
	// example:
	//
	// aliyun_main
	CallerType *string `json:"callerType,omitempty" xml:"callerType,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The name of the currently effective digital employee. This value is empty if not configured.
	//
	// example:
	//
	// pcitc-magent
	DigitalEmployeeName *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F4A9EB1C-6952-5CCC-B1DC-355576FC82A7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The effective tenant ID.
	//
	// example:
	//
	// 21577
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The platform user ID.
	//
	// example:
	//
	// 10001
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CheckHealthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckHealthResponseBody) GoString() string {
	return s.String()
}

func (s *CheckHealthResponseBody) GetAuthSource() *string {
	return s.AuthSource
}

func (s *CheckHealthResponseBody) GetCallerType() *string {
	return s.CallerType
}

func (s *CheckHealthResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckHealthResponseBody) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *CheckHealthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckHealthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckHealthResponseBody) GetTenantId() *int64 {
	return s.TenantId
}

func (s *CheckHealthResponseBody) GetUserId() *int64 {
	return s.UserId
}

func (s *CheckHealthResponseBody) SetAuthSource(v string) *CheckHealthResponseBody {
	s.AuthSource = &v
	return s
}

func (s *CheckHealthResponseBody) SetCallerType(v string) *CheckHealthResponseBody {
	s.CallerType = &v
	return s
}

func (s *CheckHealthResponseBody) SetCode(v string) *CheckHealthResponseBody {
	s.Code = &v
	return s
}

func (s *CheckHealthResponseBody) SetDigitalEmployeeName(v string) *CheckHealthResponseBody {
	s.DigitalEmployeeName = &v
	return s
}

func (s *CheckHealthResponseBody) SetMessage(v string) *CheckHealthResponseBody {
	s.Message = &v
	return s
}

func (s *CheckHealthResponseBody) SetRequestId(v string) *CheckHealthResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckHealthResponseBody) SetTenantId(v int64) *CheckHealthResponseBody {
	s.TenantId = &v
	return s
}

func (s *CheckHealthResponseBody) SetUserId(v int64) *CheckHealthResponseBody {
	s.UserId = &v
	return s
}

func (s *CheckHealthResponseBody) Validate() error {
	return dara.Validate(s)
}
