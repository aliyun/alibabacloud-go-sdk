// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanTaskStatusesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIds(v []*string) *ListVirusScanTaskStatusesRequest
	GetTaskIds() []*string
}

type ListVirusScanTaskStatusesRequest struct {
	// This parameter is required.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s ListVirusScanTaskStatusesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTaskStatusesRequest) GoString() string {
	return s.String()
}

func (s *ListVirusScanTaskStatusesRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *ListVirusScanTaskStatusesRequest) SetTaskIds(v []*string) *ListVirusScanTaskStatusesRequest {
	s.TaskIds = v
	return s
}

func (s *ListVirusScanTaskStatusesRequest) Validate() error {
	return dara.Validate(s)
}
