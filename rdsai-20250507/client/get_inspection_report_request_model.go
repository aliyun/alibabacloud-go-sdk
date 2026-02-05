// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInspectionReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInspectionReportRequest
	GetInstanceId() *string
	SetTaskId(v string) *GetInspectionReportRequest
	GetTaskId() *string
}

type GetInspectionReportRequest struct {
	// example:
	//
	// rm-2zep6e5u6l2yu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9d246af2-a0cd-4f69-857d-3785048f****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetInspectionReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInspectionReportRequest) GoString() string {
	return s.String()
}

func (s *GetInspectionReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInspectionReportRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetInspectionReportRequest) SetInstanceId(v string) *GetInspectionReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInspectionReportRequest) SetTaskId(v string) *GetInspectionReportRequest {
	s.TaskId = &v
	return s
}

func (s *GetInspectionReportRequest) Validate() error {
	return dara.Validate(s)
}
