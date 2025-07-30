// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTDEStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeInstanceTDEStatusResponseBody
	GetRequestId() *string
	SetTDEStatus(v string) *DescribeInstanceTDEStatusResponseBody
	GetTDEStatus() *string
}

type DescribeInstanceTDEStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether TDE is enabled. Valid values:
	//
	// 	- **Enabled**: TDE is enabled.
	//
	// 	- **Disable**: TDE is disabled.
	//
	// example:
	//
	// Enabled
	TDEStatus *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
}

func (s DescribeInstanceTDEStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTDEStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTDEStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceTDEStatusResponseBody) GetTDEStatus() *string {
	return s.TDEStatus
}

func (s *DescribeInstanceTDEStatusResponseBody) SetRequestId(v string) *DescribeInstanceTDEStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceTDEStatusResponseBody) SetTDEStatus(v string) *DescribeInstanceTDEStatusResponseBody {
	s.TDEStatus = &v
	return s
}

func (s *DescribeInstanceTDEStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
