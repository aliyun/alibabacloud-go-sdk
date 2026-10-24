// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventOperationStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityEventIds(v []*int64) *DescribeSecurityEventOperationStatusRequest
	GetSecurityEventIds() []*int64
	SetTaskId(v int64) *DescribeSecurityEventOperationStatusRequest
	GetTaskId() *int64
}

type DescribeSecurityEventOperationStatusRequest struct {
	// The IDs of the security alert events.
	//
	// > You must specify either TaskId or SecurityEventIds. The API call succeeds only if one of these parameters is specified.
	//
	// This parameter is required.
	SecurityEventIds []*int64 `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty" type:"Repeated"`
	// The ID of the task that processes the security alert event.
	//
	// > You must specify either TaskId or SecurityEventIds. The API call succeeds only if one of these parameters is specified.
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

func (s *DescribeSecurityEventOperationStatusRequest) GetSecurityEventIds() []*int64 {
	return s.SecurityEventIds
}

func (s *DescribeSecurityEventOperationStatusRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeSecurityEventOperationStatusRequest) SetSecurityEventIds(v []*int64) *DescribeSecurityEventOperationStatusRequest {
	s.SecurityEventIds = v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) SetTaskId(v int64) *DescribeSecurityEventOperationStatusRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) Validate() error {
	return dara.Validate(s)
}
