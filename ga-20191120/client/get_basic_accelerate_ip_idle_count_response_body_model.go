// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicAccelerateIpIdleCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetBasicAccelerateIpIdleCountResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetBasicAccelerateIpIdleCountResponseBody
	GetTotalCount() *int64
}

type GetBasicAccelerateIpIdleCountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned, which indicates the number of idle accelerated IP addresses.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetBasicAccelerateIpIdleCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAccelerateIpIdleCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpIdleCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBasicAccelerateIpIdleCountResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetBasicAccelerateIpIdleCountResponseBody) SetRequestId(v string) *GetBasicAccelerateIpIdleCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicAccelerateIpIdleCountResponseBody) SetTotalCount(v int64) *GetBasicAccelerateIpIdleCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetBasicAccelerateIpIdleCountResponseBody) Validate() error {
	return dara.Validate(s)
}
