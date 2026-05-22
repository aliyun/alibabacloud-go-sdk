// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushTimesUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *PushTimesUsageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *PushTimesUsageResponseBody
	GetCode() *string
	SetForceFatal(v bool) *PushTimesUsageResponseBody
	GetForceFatal() *bool
	SetHttpStateCode(v string) *PushTimesUsageResponseBody
	GetHttpStateCode() *string
	SetMessage(v string) *PushTimesUsageResponseBody
	GetMessage() *string
	SetRequestId(v string) *PushTimesUsageResponseBody
	GetRequestId() *string
	SetResult(v bool) *PushTimesUsageResponseBody
	GetResult() *bool
	SetSuccess(v bool) *PushTimesUsageResponseBody
	GetSuccess() *bool
}

type PushTimesUsageResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ForceFatal *bool   `json:"ForceFatal,omitempty" xml:"ForceFatal,omitempty"`
	// example:
	//
	// 200
	HttpStateCode *string `json:"HttpStateCode,omitempty" xml:"HttpStateCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushTimesUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushTimesUsageResponseBody) GoString() string {
	return s.String()
}

func (s *PushTimesUsageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *PushTimesUsageResponseBody) GetCode() *string {
	return s.Code
}

func (s *PushTimesUsageResponseBody) GetForceFatal() *bool {
	return s.ForceFatal
}

func (s *PushTimesUsageResponseBody) GetHttpStateCode() *string {
	return s.HttpStateCode
}

func (s *PushTimesUsageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PushTimesUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushTimesUsageResponseBody) GetResult() *bool {
	return s.Result
}

func (s *PushTimesUsageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PushTimesUsageResponseBody) SetAccessDeniedDetail(v string) *PushTimesUsageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *PushTimesUsageResponseBody) SetCode(v string) *PushTimesUsageResponseBody {
	s.Code = &v
	return s
}

func (s *PushTimesUsageResponseBody) SetForceFatal(v bool) *PushTimesUsageResponseBody {
	s.ForceFatal = &v
	return s
}

func (s *PushTimesUsageResponseBody) SetHttpStateCode(v string) *PushTimesUsageResponseBody {
	s.HttpStateCode = &v
	return s
}

func (s *PushTimesUsageResponseBody) SetMessage(v string) *PushTimesUsageResponseBody {
	s.Message = &v
	return s
}

func (s *PushTimesUsageResponseBody) SetRequestId(v string) *PushTimesUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushTimesUsageResponseBody) SetResult(v bool) *PushTimesUsageResponseBody {
	s.Result = &v
	return s
}

func (s *PushTimesUsageResponseBody) SetSuccess(v bool) *PushTimesUsageResponseBody {
	s.Success = &v
	return s
}

func (s *PushTimesUsageResponseBody) Validate() error {
	return dara.Validate(s)
}
