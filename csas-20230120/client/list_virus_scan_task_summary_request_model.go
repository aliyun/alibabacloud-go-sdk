// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanTaskSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIds(v []*string) *ListVirusScanTaskSummaryRequest
	GetTaskIds() []*string
}

type ListVirusScanTaskSummaryRequest struct {
	// This parameter is required.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s ListVirusScanTaskSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTaskSummaryRequest) GoString() string {
	return s.String()
}

func (s *ListVirusScanTaskSummaryRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *ListVirusScanTaskSummaryRequest) SetTaskIds(v []*string) *ListVirusScanTaskSummaryRequest {
	s.TaskIds = v
	return s
}

func (s *ListVirusScanTaskSummaryRequest) Validate() error {
	return dara.Validate(s)
}
