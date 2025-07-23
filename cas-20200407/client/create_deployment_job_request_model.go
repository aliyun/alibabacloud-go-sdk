// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIds(v string) *CreateDeploymentJobRequest
	GetCertIds() *string
	SetContactIds(v string) *CreateDeploymentJobRequest
	GetContactIds() *string
	SetJobType(v string) *CreateDeploymentJobRequest
	GetJobType() *string
	SetName(v string) *CreateDeploymentJobRequest
	GetName() *string
	SetResourceIds(v string) *CreateDeploymentJobRequest
	GetResourceIds() *string
	SetScheduleTime(v int64) *CreateDeploymentJobRequest
	GetScheduleTime() *int64
}

type CreateDeploymentJobRequest struct {
	// The ID of the certificate. Separate multiple certificate IDs with commas (,). You can call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to obtain the IDs of certificates from the **CertificateId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12342649,12304338,12067351,9957285
	CertIds *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	// The ID of the contact. Separate multiple contact IDs with commas (,). You can call the [ListContact](https://help.aliyun.com/document_detail/2712221.html) operation to obtain the IDs of contacts from the **ContactId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2
	ContactIds *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	// The type of the deployment task.
	//
	// Valid values:
	//
	// 	- cloud: multi-cloud deployment task.
	//
	// 	- user: cloud service deployment task. This type of task does not support cloud servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The name of the deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// jobName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the cloud resource. Separate multiple resource IDs with commas (,). You can call the [ListCloudResources](https://help.aliyun.com/document_detail/2712230.html) operation to obtain the IDs of cloud resources from the **Id*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6591316,6585549,6190248,5811894,5775176,5772504
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The time when the task starts. The value is a UNIX timestamp. If you do not specify this parameter, the system immediately starts the task after the task is in the pending state.
	//
	// example:
	//
	// 1706613560008
	ScheduleTime *int64 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
}

func (s CreateDeploymentJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentJobRequest) GetCertIds() *string {
	return s.CertIds
}

func (s *CreateDeploymentJobRequest) GetContactIds() *string {
	return s.ContactIds
}

func (s *CreateDeploymentJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *CreateDeploymentJobRequest) GetName() *string {
	return s.Name
}

func (s *CreateDeploymentJobRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *CreateDeploymentJobRequest) GetScheduleTime() *int64 {
	return s.ScheduleTime
}

func (s *CreateDeploymentJobRequest) SetCertIds(v string) *CreateDeploymentJobRequest {
	s.CertIds = &v
	return s
}

func (s *CreateDeploymentJobRequest) SetContactIds(v string) *CreateDeploymentJobRequest {
	s.ContactIds = &v
	return s
}

func (s *CreateDeploymentJobRequest) SetJobType(v string) *CreateDeploymentJobRequest {
	s.JobType = &v
	return s
}

func (s *CreateDeploymentJobRequest) SetName(v string) *CreateDeploymentJobRequest {
	s.Name = &v
	return s
}

func (s *CreateDeploymentJobRequest) SetResourceIds(v string) *CreateDeploymentJobRequest {
	s.ResourceIds = &v
	return s
}

func (s *CreateDeploymentJobRequest) SetScheduleTime(v int64) *CreateDeploymentJobRequest {
	s.ScheduleTime = &v
	return s
}

func (s *CreateDeploymentJobRequest) Validate() error {
	return dara.Validate(s)
}
