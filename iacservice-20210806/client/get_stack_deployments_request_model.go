// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackDeploymentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigVersion(v string) *GetStackDeploymentsRequest
	GetConfigVersion() *string
	SetDeploymentName(v string) *GetStackDeploymentsRequest
	GetDeploymentName() *string
	SetDeploymentNo(v string) *GetStackDeploymentsRequest
	GetDeploymentNo() *string
	SetPageNumber(v int32) *GetStackDeploymentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetStackDeploymentsRequest
	GetPageSize() *int32
	SetStatus(v string) *GetStackDeploymentsRequest
	GetStatus() *string
}

type GetStackDeploymentsRequest struct {
	// example:
	//
	// v1
	ConfigVersion *string `json:"configVersion,omitempty" xml:"configVersion,omitempty"`
	// example:
	//
	// production
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
	// example:
	//
	// 1
	DeploymentNo *string `json:"deploymentNo,omitempty" xml:"deploymentNo,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// Applied
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetStackDeploymentsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStackDeploymentsRequest) GoString() string {
	return s.String()
}

func (s *GetStackDeploymentsRequest) GetConfigVersion() *string {
	return s.ConfigVersion
}

func (s *GetStackDeploymentsRequest) GetDeploymentName() *string {
	return s.DeploymentName
}

func (s *GetStackDeploymentsRequest) GetDeploymentNo() *string {
	return s.DeploymentNo
}

func (s *GetStackDeploymentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetStackDeploymentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetStackDeploymentsRequest) GetStatus() *string {
	return s.Status
}

func (s *GetStackDeploymentsRequest) SetConfigVersion(v string) *GetStackDeploymentsRequest {
	s.ConfigVersion = &v
	return s
}

func (s *GetStackDeploymentsRequest) SetDeploymentName(v string) *GetStackDeploymentsRequest {
	s.DeploymentName = &v
	return s
}

func (s *GetStackDeploymentsRequest) SetDeploymentNo(v string) *GetStackDeploymentsRequest {
	s.DeploymentNo = &v
	return s
}

func (s *GetStackDeploymentsRequest) SetPageNumber(v int32) *GetStackDeploymentsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetStackDeploymentsRequest) SetPageSize(v int32) *GetStackDeploymentsRequest {
	s.PageSize = &v
	return s
}

func (s *GetStackDeploymentsRequest) SetStatus(v string) *GetStackDeploymentsRequest {
	s.Status = &v
	return s
}

func (s *GetStackDeploymentsRequest) Validate() error {
	return dara.Validate(s)
}
