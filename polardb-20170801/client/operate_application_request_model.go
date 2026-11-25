// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *OperateApplicationRequest
	GetApplicationId() *string
	SetComponentTypes(v string) *OperateApplicationRequest
	GetComponentTypes() *string
	SetGroupNames(v string) *OperateApplicationRequest
	GetGroupNames() *string
	SetOperation(v string) *OperateApplicationRequest
	GetOperation() *string
}

type OperateApplicationRequest struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The Ray component type. Supported only for the restart operation of Ray applications. The value is a comma-separated list of components. Valid components are head, jupyterlab, and grafana, which can be combined. worker or gpuworker is not accepted. To select workers, you must specify GroupNames. If only head is specified, other components are not actively restarted, but connections and jobs may be affected. To restore dependencies, explicitly add jupyterlab or select the entire cluster. When provided together with GroupNames, the union of both is used and deduplicated. If neither selection parameter is provided, all resident components are restarted, but scaled-out workers are not restarted. Explicit empty strings, pure whitespace, comma-only strings, unknown types, or non-Ray applications are rejected and do not fall back to a full-cluster operation. DisableHeadScheduling and EnableHeadScheduling always apply to the head of the entire application. Specifying this parameter for these operations is rejected.
	//
	// example:
	//
	// head,jupyterlab,grafana
	ComponentTypes *string `json:"ComponentTypes,omitempty" xml:"ComponentTypes,omitempty"`
	// The Ray resident worker group name. Supported only for the restart operation of Ray applications. The value is a comma-separated list of exact workerGroup names, which are trimmed, deduplicated, and case-sensitive. Only resident CPU or GPU workers in the specified groups are restarted, excluding scaled-out workers. If only this parameter is specified, the head, other groups, or auxiliary components are not restarted. Unknown group names, explicit empty strings, pure whitespace, comma-only strings, or non-Ray applications are rejected. If any explicitly selected group exists but has no resident targets, the entire request is rejected and does not expand to scaled-out workers or a full-cluster operation. Full-cluster requests without selection parameters skip groups with no resident targets. DisableHeadScheduling and EnableHeadScheduling do not restart any workers. Specifying this parameter for these operations is rejected.
	//
	// example:
	//
	// cpuGroup,gpuGroup
	GroupNames *string `json:"GroupNames,omitempty" xml:"GroupNames,omitempty"`
	// The operation type.
	//
	// 	- **restart**: Restart.
	//
	// 	- **stop**: Stop.
	//
	// 	- **start**: Start.
	//
	// example:
	//
	// restart
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s OperateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateApplicationRequest) GoString() string {
	return s.String()
}

func (s *OperateApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *OperateApplicationRequest) GetComponentTypes() *string {
	return s.ComponentTypes
}

func (s *OperateApplicationRequest) GetGroupNames() *string {
	return s.GroupNames
}

func (s *OperateApplicationRequest) GetOperation() *string {
	return s.Operation
}

func (s *OperateApplicationRequest) SetApplicationId(v string) *OperateApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *OperateApplicationRequest) SetComponentTypes(v string) *OperateApplicationRequest {
	s.ComponentTypes = &v
	return s
}

func (s *OperateApplicationRequest) SetGroupNames(v string) *OperateApplicationRequest {
	s.GroupNames = &v
	return s
}

func (s *OperateApplicationRequest) SetOperation(v string) *OperateApplicationRequest {
	s.Operation = &v
	return s
}

func (s *OperateApplicationRequest) Validate() error {
	return dara.Validate(s)
}
