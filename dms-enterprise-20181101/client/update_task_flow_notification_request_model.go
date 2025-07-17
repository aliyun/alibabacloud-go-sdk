// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowNotificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *UpdateTaskFlowNotificationRequest
	GetDagId() *int64
	SetDagNotificationFail(v bool) *UpdateTaskFlowNotificationRequest
	GetDagNotificationFail() *bool
	SetDagNotificationSla(v bool) *UpdateTaskFlowNotificationRequest
	GetDagNotificationSla() *bool
	SetDagNotificationSuccess(v bool) *UpdateTaskFlowNotificationRequest
	GetDagNotificationSuccess() *bool
	SetTid(v int64) *UpdateTaskFlowNotificationRequest
	GetTid() *int64
}

type UpdateTaskFlowNotificationRequest struct {
	// The unique ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// Specifies whether to enable notifications for failed task flows. Notifications are disabled by default. You can enable notifications based on your business requirements.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DagNotificationFail *bool `json:"DagNotificationFail,omitempty" xml:"DagNotificationFail,omitempty"`
	// Specifies whether to enable SLA global notifications for task flows. Notifications are disabled by default. You can enable notifications based on your business requirements.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DagNotificationSla *bool `json:"DagNotificationSla,omitempty" xml:"DagNotificationSla,omitempty"`
	// Specifies whether to enable notifications for successful task flows. Notifications are disabled by default. You can enable notifications based on your business requirements.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DagNotificationSuccess *bool `json:"DagNotificationSuccess,omitempty" xml:"DagNotificationSuccess,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskFlowNotificationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowNotificationRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowNotificationRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateTaskFlowNotificationRequest) GetDagNotificationFail() *bool {
	return s.DagNotificationFail
}

func (s *UpdateTaskFlowNotificationRequest) GetDagNotificationSla() *bool {
	return s.DagNotificationSla
}

func (s *UpdateTaskFlowNotificationRequest) GetDagNotificationSuccess() *bool {
	return s.DagNotificationSuccess
}

func (s *UpdateTaskFlowNotificationRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskFlowNotificationRequest) SetDagId(v int64) *UpdateTaskFlowNotificationRequest {
	s.DagId = &v
	return s
}

func (s *UpdateTaskFlowNotificationRequest) SetDagNotificationFail(v bool) *UpdateTaskFlowNotificationRequest {
	s.DagNotificationFail = &v
	return s
}

func (s *UpdateTaskFlowNotificationRequest) SetDagNotificationSla(v bool) *UpdateTaskFlowNotificationRequest {
	s.DagNotificationSla = &v
	return s
}

func (s *UpdateTaskFlowNotificationRequest) SetDagNotificationSuccess(v bool) *UpdateTaskFlowNotificationRequest {
	s.DagNotificationSuccess = &v
	return s
}

func (s *UpdateTaskFlowNotificationRequest) SetTid(v int64) *UpdateTaskFlowNotificationRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskFlowNotificationRequest) Validate() error {
	return dara.Validate(s)
}
