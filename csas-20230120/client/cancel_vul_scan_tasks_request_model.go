// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelVulScanTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIds(v []*string) *CancelVulScanTasksRequest
	GetTaskIds() []*string
}

type CancelVulScanTasksRequest struct {
	// The IDs of the vulnerability scanning tasks to cancel. The collection must contain at least 1 and at most 100 IDs. Duplicate IDs are not allowed.
	//
	// This parameter is required.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s CancelVulScanTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelVulScanTasksRequest) GoString() string {
	return s.String()
}

func (s *CancelVulScanTasksRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *CancelVulScanTasksRequest) SetTaskIds(v []*string) *CancelVulScanTasksRequest {
	s.TaskIds = v
	return s
}

func (s *CancelVulScanTasksRequest) Validate() error {
	return dara.Validate(s)
}
