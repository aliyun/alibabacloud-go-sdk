// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventOperationStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceOwnerId(v int64) *DescribeSecurityEventOperationStatusRequest
	GetResourceOwnerId() *int64
	SetSecurityEventIds(v []*string) *DescribeSecurityEventOperationStatusRequest
	GetSecurityEventIds() []*string
	SetSourceIp(v string) *DescribeSecurityEventOperationStatusRequest
	GetSourceIp() *string
	SetTaskId(v int64) *DescribeSecurityEventOperationStatusRequest
	GetTaskId() *int64
}

type DescribeSecurityEventOperationStatusRequest struct {
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of the alert events.
	//
	// >  You must specify at least one of the TaskId and SecurityEventIds parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["909361"]
	SecurityEventIds []*string `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty" type:"Repeated"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the task that handles the alert events.
	//
	// >  You must specify at least one of the TaskId and SecurityEventIds parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12121
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeSecurityEventOperationStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSecurityEventOperationStatusRequest) GetSecurityEventIds() []*string {
	return s.SecurityEventIds
}

func (s *DescribeSecurityEventOperationStatusRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSecurityEventOperationStatusRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeSecurityEventOperationStatusRequest) SetResourceOwnerId(v int64) *DescribeSecurityEventOperationStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) SetSecurityEventIds(v []*string) *DescribeSecurityEventOperationStatusRequest {
	s.SecurityEventIds = v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) SetSourceIp(v string) *DescribeSecurityEventOperationStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) SetTaskId(v int64) *DescribeSecurityEventOperationStatusRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) Validate() error {
	return dara.Validate(s)
}
