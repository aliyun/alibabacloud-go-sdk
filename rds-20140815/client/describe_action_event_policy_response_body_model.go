// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActionEventPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnableEventLog(v string) *DescribeActionEventPolicyResponseBody
	GetEnableEventLog() *string
	SetRegionId(v string) *DescribeActionEventPolicyResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeActionEventPolicyResponseBody
	GetRequestId() *string
}

type DescribeActionEventPolicyResponseBody struct {
	// Indicates whether the event history feature is enabled.
	//
	// example:
	//
	// True
	EnableEventLog *string `json:"EnableEventLog,omitempty" xml:"EnableEventLog,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CCECD3CD-AB2D-4F6D-BEDE-47BC90A398D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeActionEventPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeActionEventPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActionEventPolicyResponseBody) GetEnableEventLog() *string {
	return s.EnableEventLog
}

func (s *DescribeActionEventPolicyResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeActionEventPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActionEventPolicyResponseBody) SetEnableEventLog(v string) *DescribeActionEventPolicyResponseBody {
	s.EnableEventLog = &v
	return s
}

func (s *DescribeActionEventPolicyResponseBody) SetRegionId(v string) *DescribeActionEventPolicyResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeActionEventPolicyResponseBody) SetRequestId(v string) *DescribeActionEventPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActionEventPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
