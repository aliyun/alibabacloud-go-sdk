// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUpgradeStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *GetUpgradeStatusResponseBody
	GetErrorMessage() *string
	SetPrecheckReportId(v string) *GetUpgradeStatusResponseBody
	GetPrecheckReportId() *string
	SetStatus(v string) *GetUpgradeStatusResponseBody
	GetStatus() *string
	SetUpgradeStep(v string) *GetUpgradeStatusResponseBody
	GetUpgradeStep() *string
	SetUpgradeTask(v *GetUpgradeStatusResponseBodyUpgradeTask) *GetUpgradeStatusResponseBody
	GetUpgradeTask() *GetUpgradeStatusResponseBodyUpgradeTask
}

type GetUpgradeStatusResponseBody struct {
	// The error message returned during the update.
	//
	// example:
	//
	// subject to actual return
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// The ID of the precheck report.
	//
	// example:
	//
	// be4c8eb72de94d459ea7ace7c811d119
	PrecheckReportId *string `json:"precheck_report_id,omitempty" xml:"precheck_report_id,omitempty"`
	// The status of the update. Valid values:
	//
	// 	- `success`: The update is successful.
	//
	// 	- `fail`: The update failed.
	//
	// 	- `pause`: The update is paused.
	//
	// 	- `running`: The update is in progress.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The current phase of the update. Valid values:
	//
	// 	- `not_start`: The update is not started.
	//
	// 	- `prechecking`: The precheck is in progress.
	//
	// 	- `upgrading`: The cluster is being updated.
	//
	// 	- `pause`: The update is paused.
	//
	// 	- `success`: The update is successful.
	//
	// example:
	//
	// prechecking
	UpgradeStep *string `json:"upgrade_step,omitempty" xml:"upgrade_step,omitempty"`
	// The details of the update task.
	UpgradeTask *GetUpgradeStatusResponseBodyUpgradeTask `json:"upgrade_task,omitempty" xml:"upgrade_task,omitempty" type:"Struct"`
}

func (s GetUpgradeStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUpgradeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetUpgradeStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetUpgradeStatusResponseBody) GetPrecheckReportId() *string {
	return s.PrecheckReportId
}

func (s *GetUpgradeStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetUpgradeStatusResponseBody) GetUpgradeStep() *string {
	return s.UpgradeStep
}

func (s *GetUpgradeStatusResponseBody) GetUpgradeTask() *GetUpgradeStatusResponseBodyUpgradeTask {
	return s.UpgradeTask
}

func (s *GetUpgradeStatusResponseBody) SetErrorMessage(v string) *GetUpgradeStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetPrecheckReportId(v string) *GetUpgradeStatusResponseBody {
	s.PrecheckReportId = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetStatus(v string) *GetUpgradeStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetUpgradeStep(v string) *GetUpgradeStatusResponseBody {
	s.UpgradeStep = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetUpgradeTask(v *GetUpgradeStatusResponseBodyUpgradeTask) *GetUpgradeStatusResponseBody {
	s.UpgradeTask = v
	return s
}

func (s *GetUpgradeStatusResponseBody) Validate() error {
	if s.UpgradeTask != nil {
		if err := s.UpgradeTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUpgradeStatusResponseBodyUpgradeTask struct {
	// The description of the update task.
	//
	// example:
	//
	// subject to actual return
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The status of the update task. Valid values:
	//
	// 	- `running`: The update task is being executed.
	//
	// 	- `Success`: The update task is successfully executed.
	//
	// 	- `Failed`: The update task failed.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetUpgradeStatusResponseBodyUpgradeTask) String() string {
	return dara.Prettify(s)
}

func (s GetUpgradeStatusResponseBodyUpgradeTask) GoString() string {
	return s.String()
}

func (s *GetUpgradeStatusResponseBodyUpgradeTask) GetMessage() *string {
	return s.Message
}

func (s *GetUpgradeStatusResponseBodyUpgradeTask) GetStatus() *string {
	return s.Status
}

func (s *GetUpgradeStatusResponseBodyUpgradeTask) SetMessage(v string) *GetUpgradeStatusResponseBodyUpgradeTask {
	s.Message = &v
	return s
}

func (s *GetUpgradeStatusResponseBodyUpgradeTask) SetStatus(v string) *GetUpgradeStatusResponseBodyUpgradeTask {
	s.Status = &v
	return s
}

func (s *GetUpgradeStatusResponseBodyUpgradeTask) Validate() error {
	return dara.Validate(s)
}
