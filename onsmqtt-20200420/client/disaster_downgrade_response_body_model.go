// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisasterDowngradeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DisasterDowngradeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *DisasterDowngradeResponseBody
	GetCode() *int32
	SetMessage(v string) *DisasterDowngradeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisasterDowngradeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisasterDowngradeResponseBody
	GetSuccess() *bool
}

type DisasterDowngradeResponseBody struct {
	// The detailed information about the access denial. This field is returned only when RAM authentication fails.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The return code of the operation. A value of 200 indicates success. Other values indicate error codes. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result message of the call.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82B9E503-F4A1-4F30-976F-C6999FF9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisasterDowngradeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisasterDowngradeResponseBody) GoString() string {
	return s.String()
}

func (s *DisasterDowngradeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DisasterDowngradeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DisasterDowngradeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisasterDowngradeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisasterDowngradeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisasterDowngradeResponseBody) SetAccessDeniedDetail(v string) *DisasterDowngradeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DisasterDowngradeResponseBody) SetCode(v int32) *DisasterDowngradeResponseBody {
	s.Code = &v
	return s
}

func (s *DisasterDowngradeResponseBody) SetMessage(v string) *DisasterDowngradeResponseBody {
	s.Message = &v
	return s
}

func (s *DisasterDowngradeResponseBody) SetRequestId(v string) *DisasterDowngradeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisasterDowngradeResponseBody) SetSuccess(v bool) *DisasterDowngradeResponseBody {
	s.Success = &v
	return s
}

func (s *DisasterDowngradeResponseBody) Validate() error {
	return dara.Validate(s)
}
