// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v int64) *ListDeploymentJobRequest
	GetCertId() *int64
	SetCurrentPage(v int32) *ListDeploymentJobRequest
	GetCurrentPage() *int32
	SetInstanceId(v string) *ListDeploymentJobRequest
	GetInstanceId() *string
	SetJobType(v string) *ListDeploymentJobRequest
	GetJobType() *string
	SetShowSize(v int32) *ListDeploymentJobRequest
	GetShowSize() *int32
	SetStatus(v string) *ListDeploymentJobRequest
	GetStatus() *string
}

type ListDeploymentJobRequest struct {
	// The certificate ID.
	//
	// > You can obtain this ID by calling the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation.
	//
	// example:
	//
	// 23982365
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The current page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cas-ivauto-hqito6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The task type.
	//
	// example:
	//
	// user
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The paging Settings for the number of certificates to return on each page during a paged query. Default value: **50**.
	//
	// example:
	//
	// 50
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The deployment task status.
	//
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeploymentJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentJobRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobRequest) GetCertId() *int64 {
	return s.CertId
}

func (s *ListDeploymentJobRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDeploymentJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDeploymentJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *ListDeploymentJobRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListDeploymentJobRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDeploymentJobRequest) SetCertId(v int64) *ListDeploymentJobRequest {
	s.CertId = &v
	return s
}

func (s *ListDeploymentJobRequest) SetCurrentPage(v int32) *ListDeploymentJobRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDeploymentJobRequest) SetInstanceId(v string) *ListDeploymentJobRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDeploymentJobRequest) SetJobType(v string) *ListDeploymentJobRequest {
	s.JobType = &v
	return s
}

func (s *ListDeploymentJobRequest) SetShowSize(v int32) *ListDeploymentJobRequest {
	s.ShowSize = &v
	return s
}

func (s *ListDeploymentJobRequest) SetStatus(v string) *ListDeploymentJobRequest {
	s.Status = &v
	return s
}

func (s *ListDeploymentJobRequest) Validate() error {
	return dara.Validate(s)
}
