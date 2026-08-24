// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelVirusScanTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIds(v []*string) *CancelVirusScanTasksRequest
	GetTaskIds() []*string
}

type CancelVirusScanTasksRequest struct {
	// The collection of virus scan task IDs to cancel. The collection must contain at least one ID, and duplicate IDs are not allowed.
	//
	// This parameter is required.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s CancelVirusScanTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelVirusScanTasksRequest) GoString() string {
	return s.String()
}

func (s *CancelVirusScanTasksRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *CancelVirusScanTasksRequest) SetTaskIds(v []*string) *CancelVirusScanTasksRequest {
	s.TaskIds = v
	return s
}

func (s *CancelVirusScanTasksRequest) Validate() error {
	return dara.Validate(s)
}
