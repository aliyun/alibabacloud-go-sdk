// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceStatusResponseBody
	GetInstanceId() *string
	SetInstanceStatus(v int32) *DescribeInstanceStatusResponseBody
	GetInstanceStatus() *int32
	SetRequestId(v string) *DescribeInstanceStatusResponseBody
	GetRequestId() *string
}

type DescribeInstanceStatusResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// ddoscoo-cn-6ja1y6p5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- **1**: normal
	//
	// 	- **2**: expired
	//
	// 	- **3**: overdue
	//
	// 	- **4**: released
	//
	// example:
	//
	// 1
	InstanceStatus *int32 `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 112777CC-2AD6-46FC-A263-00B931406FCD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatusResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceStatusResponseBody) GetInstanceStatus() *int32 {
	return s.InstanceStatus
}

func (s *DescribeInstanceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceStatusResponseBody) SetInstanceId(v string) *DescribeInstanceStatusResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceStatusResponseBody) SetInstanceStatus(v int32) *DescribeInstanceStatusResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceStatusResponseBody) SetRequestId(v string) *DescribeInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
