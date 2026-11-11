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
	// The configuration version, such as v1. The initial value is v1. The version number increments each time the stack is updated or refreshed and the configuration changes.
	//
	// example:
	//
	// v1
	ConfigVersion *string `json:"configVersion,omitempty" xml:"configVersion,omitempty"`
	// The name of the deployment.
	//
	// example:
	//
	// production
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
	// The deployment number. The deployment number of each stack starts from 1 and increments each time a deployment is successfully triggered.
	//
	// example:
	//
	// 1
	DeploymentNo *string `json:"deploymentNo,omitempty" xml:"deploymentNo,omitempty"`
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20. Minimum value: 1. Maximum value: 200.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The deployment status. Valid values:
	//
	// | Name | Description |
	//
	// |------|------|
	//
	// | Pending | The initial status after the deployment is created. |
	//
	// | PriorityQueued | The deployment is queued with priority. |
	//
	// | PlanQueued | The deployment is queued and waiting for an available workflow after it is created. |
	//
	// | ApplyQueued | The deployment is queued and waiting for an available workflow during execution. |
	//
	// | Planning | The deployment is in the plan phase. |
	//
	// | Planned | The deployment has completed the plan phase. |
	//
	// | ConfigProactiveInProgress | The compliance precheck is in progress. |
	//
	// | ConfigProactiveSuccess | The compliance precheck is successful. |
	//
	// | DetectInProgress | Drift detection is in progress. |
	//
	// | ImportQueued | The deployment is queued and waiting for an available workflow during the import phase. |
	//
	// | Importing | The deployment is in the import phase. |
	//
	// | Imported | The deployment has completed the import phase. |
	//
	// | StateQueued | The deployment is queued and waiting for an available workflow during the state command execution. |
	//
	// | Stating | The deployment is executing the state command. |
	//
	// | Stated | The deployment has completed the execution of the state command. |
	//
	// | Confirmed | The deployment is confirmed after the plan phase. |
	//
	// | PlannedAndFinished | The deployment is in the final state because no differences are found after the plan phase. |
	//
	// | Applying | The deployment is in the apply phase. |
	//
	// | Applied | The deployment has completed the apply phase. |
	//
	// | Discarded | The deployment is discarded and enters the final state. |
	//
	// | Errored | An error occurred during the deployment execution, and the deployment enters the final state. |
	//
	// | ConfigProactiveFailure | The compliance precheck failed. |
	//
	// | Canceled | The deployment execution is canceled, and the deployment enters the final state. |
	//
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
