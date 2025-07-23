// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIds(v string) *UpdateDeploymentJobRequest
	GetCertIds() *string
	SetContactIds(v string) *UpdateDeploymentJobRequest
	GetContactIds() *string
	SetJobId(v int64) *UpdateDeploymentJobRequest
	GetJobId() *int64
	SetName(v string) *UpdateDeploymentJobRequest
	GetName() *string
	SetResourceIds(v string) *UpdateDeploymentJobRequest
	GetResourceIds() *string
	SetScheduleTime(v int64) *UpdateDeploymentJobRequest
	GetScheduleTime() *int64
}

type UpdateDeploymentJobRequest struct {
	// The ID of the certificate. Separate multiple certificate IDs with commas (,). You can call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to obtain the ID of a certificate from the **CertificateId*	- parameter.
	//
	// example:
	//
	// 9957285,12067351,12304338,12342649
	CertIds *string `json:"CertIds,omitempty" xml:"CertIds,omitempty"`
	// The ID of the contact. Separate multiple contact IDs with commas (,). You can call the [ListContact](https://help.aliyun.com/document_detail/2712221.html) operation to obtain the ID of a contact from the **ContactId*	- parameter.
	//
	// example:
	//
	// 9957285,12067351,12304338,12342649
	ContactIds *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	// The ID of the deployment task. You can call the [CreateDeploymentJob](https://help.aliyun.com/document_detail/2712234.html) operation to obtain the ID of a deployment task from the JobId parameter. You can also call the [ListDeploymentJob](https://help.aliyun.com/document_detail/2712223.html) operation to obtain the ID of a deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the deployment task.
	//
	// example:
	//
	// cert-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the cloud resource. Separate multiple resource IDs with commas (,). You can call the [ListCloudResources](https://help.aliyun.com/document_detail/2712230.html) operation to obtain the ID of a cloud resource from the **Id*	- parameter.
	//
	// example:
	//
	// 9957285,12067351,12304338,12342649
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The time when the task starts. The value is a UNIX timestamp. If you do not specify this parameter, the system immediately starts the task after it is created.
	//
	// example:
	//
	// 1606482979000
	ScheduleTime *int64 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
}

func (s UpdateDeploymentJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentJobRequest) GetCertIds() *string {
	return s.CertIds
}

func (s *UpdateDeploymentJobRequest) GetContactIds() *string {
	return s.ContactIds
}

func (s *UpdateDeploymentJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *UpdateDeploymentJobRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDeploymentJobRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *UpdateDeploymentJobRequest) GetScheduleTime() *int64 {
	return s.ScheduleTime
}

func (s *UpdateDeploymentJobRequest) SetCertIds(v string) *UpdateDeploymentJobRequest {
	s.CertIds = &v
	return s
}

func (s *UpdateDeploymentJobRequest) SetContactIds(v string) *UpdateDeploymentJobRequest {
	s.ContactIds = &v
	return s
}

func (s *UpdateDeploymentJobRequest) SetJobId(v int64) *UpdateDeploymentJobRequest {
	s.JobId = &v
	return s
}

func (s *UpdateDeploymentJobRequest) SetName(v string) *UpdateDeploymentJobRequest {
	s.Name = &v
	return s
}

func (s *UpdateDeploymentJobRequest) SetResourceIds(v string) *UpdateDeploymentJobRequest {
	s.ResourceIds = &v
	return s
}

func (s *UpdateDeploymentJobRequest) SetScheduleTime(v int64) *UpdateDeploymentJobRequest {
	s.ScheduleTime = &v
	return s
}

func (s *UpdateDeploymentJobRequest) Validate() error {
	return dara.Validate(s)
}
