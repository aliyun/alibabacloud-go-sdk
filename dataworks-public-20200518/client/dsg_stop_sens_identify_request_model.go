// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgStopSensIdentifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *DsgStopSensIdentifyRequest
	GetJobId() *int64
	SetTenantId(v string) *DsgStopSensIdentifyRequest
	GetTenantId() *string
}

type DsgStopSensIdentifyRequest struct {
	// The ID of the sensitive data identification task. You can call the [DsgRunSensIdentify](https://help.aliyun.com/document_detail/2744039.html) operation to obtain the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000001
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10241024
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DsgStopSensIdentifyRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgStopSensIdentifyRequest) GoString() string {
	return s.String()
}

func (s *DsgStopSensIdentifyRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *DsgStopSensIdentifyRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DsgStopSensIdentifyRequest) SetJobId(v int64) *DsgStopSensIdentifyRequest {
	s.JobId = &v
	return s
}

func (s *DsgStopSensIdentifyRequest) SetTenantId(v string) *DsgStopSensIdentifyRequest {
	s.TenantId = &v
	return s
}

func (s *DsgStopSensIdentifyRequest) Validate() error {
	return dara.Validate(s)
}
