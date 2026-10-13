// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMqConsumerTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetMqConsumerTagResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetMqConsumerTagResponseBody
	GetCode() *int32
	SetData(v string) *GetMqConsumerTagResponseBody
	GetData() *string
	SetMessage(v string) *GetMqConsumerTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMqConsumerTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMqConsumerTagResponseBody
	GetSuccess() *bool
}

type GetMqConsumerTagResponseBody struct {
	// The details of the permission error. An empty string is returned if you have the required permissions.
	//
	// example:
	//
	// {\\"AuthAction\\":\\"ram:GetMqConsumerTag\\",\\"AuthPrincipalDisplayName\\":\\"2146216584788xxxxx\\",\\"AuthPrincipalOwnerId\\":\\"1135850448xxxxx\\",\\"AuthPrincipalType\\":\\"SubUser\\",\\"NoPermissionType\\":\\"ImplicitDeny\\",\\"PolicyType\\":\\"AccountLevelIdentityBasedPolicy\\"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The return code of the API request.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return value. This is the unique tag for your Message Queue (MQ) consumer.
	//
	// example:
	//
	// 46434030
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message. This parameter is returned when \\`success\\` is \\`false\\`.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request.
	//
	// example:
	//
	// CA6204AC-6AA9-4CFA-9310-7DFD20C19EBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of \\`true\\` indicates success. A value of \\`false\\` indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMqConsumerTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMqConsumerTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetMqConsumerTagResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetMqConsumerTagResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetMqConsumerTagResponseBody) GetData() *string {
	return s.Data
}

func (s *GetMqConsumerTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMqConsumerTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMqConsumerTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMqConsumerTagResponseBody) SetAccessDeniedDetail(v string) *GetMqConsumerTagResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetMqConsumerTagResponseBody) SetCode(v int32) *GetMqConsumerTagResponseBody {
	s.Code = &v
	return s
}

func (s *GetMqConsumerTagResponseBody) SetData(v string) *GetMqConsumerTagResponseBody {
	s.Data = &v
	return s
}

func (s *GetMqConsumerTagResponseBody) SetMessage(v string) *GetMqConsumerTagResponseBody {
	s.Message = &v
	return s
}

func (s *GetMqConsumerTagResponseBody) SetRequestId(v string) *GetMqConsumerTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMqConsumerTagResponseBody) SetSuccess(v bool) *GetMqConsumerTagResponseBody {
	s.Success = &v
	return s
}

func (s *GetMqConsumerTagResponseBody) Validate() error {
	return dara.Validate(s)
}
