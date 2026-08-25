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
	// The deployment name.
	//
	// example:
	//
	// production
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
	// The deployment number. The deployment number for each stack starts from 1 and increments each time a deployment is successfully triggered.
	//
	// example:
	//
	// 1
	DeploymentNo *string `json:"deploymentNo,omitempty" xml:"deploymentNo,omitempty"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of results returned per page. Default value: 20. Minimum value: 1. Maximum value: 200.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The deployment status.
	//
	// | Name | Description |
	//
	// |------|------|
	//
	// | Pending | The initial status after the deployment is created. |
	//
	// | PriorityQueued | Priority queuing in progress. |
	//
	// | PlanQueued | The deployment is queuing because no workflow is available after creation. |
	//
	// | ApplyQueued | The deployment is queuing because no workflow is available during execution. |
	//
	// | Planning | The resource deployment is in the Plan phase. |
	//
	// | Planned | The resource deployment has completed the Plan phase. |
	//
	// | ConfigProactiveInProgress | Compliance pre-check in progress. |
	//
	// | ConfigProactiveSuccess | Compliance pre-check succeeded. |
	//
	// | DetectInProgress | Drift detection in progress. |
	//
	// | ImportQueued | The deployment is queuing because no workflow is available during Import execution. |
	//
	// | Importing | The resource deployment is in the Import phase. |
	//
	// | Imported | The resource deployment has completed the Import phase. |
	//
	// | StateQueued | The deployment is queuing because no workflow is available during state command execution. |
	//
	// | Stating | The resource deployment is executing the state command. |
	//
	// | Stated | The resource deployment has completed the state command execution. |
	//
	// | Confirmed | The resource deployment has been confirmed after the Plan phase. |
	//
	// | PlannedAndFinished | No diff was found after the Plan phase. The deployment is in a final status. |
	//
	// | Applying | The resource deployment is in the Apply phase. |
	//
	// | Applied | The resource deployment has completed the Apply phase. |
	//
	// | Discarded | The resource deployment has been discarded and is in a final status. |
	//
	// | Errored | The deployment execution encountered an error and is in a final status. |
	//
	// | ConfigProactiveFailure | Compliance pre-check failed. |
	//
	// | Canceled | The deployment execution has been canceled and is in a final status. |
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
