// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeApplicationMonitorRequest
	GetClientToken() *string
	SetRegionId(v string) *DescribeApplicationMonitorRequest
	GetRegionId() *string
	SetTaskId(v string) *DescribeApplicationMonitorRequest
	GetTaskId() *string
}

type DescribeApplicationMonitorRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the origin probing task.
	//
	// This parameter is required.
	//
	// example:
	//
	// sm-bp1fpdjfju9k8yr1y****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeApplicationMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationMonitorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeApplicationMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationMonitorRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeApplicationMonitorRequest) SetClientToken(v string) *DescribeApplicationMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeApplicationMonitorRequest) SetRegionId(v string) *DescribeApplicationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationMonitorRequest) SetTaskId(v string) *DescribeApplicationMonitorRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeApplicationMonitorRequest) Validate() error {
	return dara.Validate(s)
}
