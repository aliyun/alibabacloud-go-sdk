// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisasterRecoveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DisasterRecoveryResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *DisasterRecoveryResponseBody
	GetCode() *int32
	SetMessage(v string) *DisasterRecoveryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisasterRecoveryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisasterRecoveryResponseBody
	GetSuccess() *bool
}

type DisasterRecoveryResponseBody struct {
	// The details that are returned when the access is denied. This field is returned only when the RAM authentication fails.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The return code. A value of 200 indicates that the call is successful. Other values indicate error codes. For more information about error codes, see Error codes.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 135F8639-F262-4417-98D1-4DE4595C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates that the call was successful. A value of false indicates that the call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisasterRecoveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisasterRecoveryResponseBody) GoString() string {
	return s.String()
}

func (s *DisasterRecoveryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DisasterRecoveryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DisasterRecoveryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisasterRecoveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisasterRecoveryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisasterRecoveryResponseBody) SetAccessDeniedDetail(v string) *DisasterRecoveryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DisasterRecoveryResponseBody) SetCode(v int32) *DisasterRecoveryResponseBody {
	s.Code = &v
	return s
}

func (s *DisasterRecoveryResponseBody) SetMessage(v string) *DisasterRecoveryResponseBody {
	s.Message = &v
	return s
}

func (s *DisasterRecoveryResponseBody) SetRequestId(v string) *DisasterRecoveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisasterRecoveryResponseBody) SetSuccess(v bool) *DisasterRecoveryResponseBody {
	s.Success = &v
	return s
}

func (s *DisasterRecoveryResponseBody) Validate() error {
	return dara.Validate(s)
}
