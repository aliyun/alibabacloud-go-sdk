// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetRoutineResponseBody
	GetCreateTime() *string
	SetDefaultRelatedRecord(v string) *GetRoutineResponseBody
	GetDefaultRelatedRecord() *string
	SetDescription(v string) *GetRoutineResponseBody
	GetDescription() *string
	SetEnvs(v []*GetRoutineResponseBodyEnvs) *GetRoutineResponseBody
	GetEnvs() []*GetRoutineResponseBodyEnvs
	SetHasAssets(v bool) *GetRoutineResponseBody
	GetHasAssets() *bool
	SetRequestId(v string) *GetRoutineResponseBody
	GetRequestId() *string
}

type GetRoutineResponseBody struct {
	// The time when the routine was created.
	//
	// example:
	//
	// 2024-03-11T01:23:21Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The default record name to access.
	//
	// example:
	//
	// routine1.example.com
	DefaultRelatedRecord *string `json:"DefaultRelatedRecord,omitempty" xml:"DefaultRelatedRecord,omitempty"`
	// The description of the routine.
	//
	// example:
	//
	// ZWRpdCByb3V0aW5lIGNvbmZpZyBkZXNjcmlwdGlvbg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the environments.
	Envs      []*GetRoutineResponseBodyEnvs `json:"Envs,omitempty" xml:"Envs,omitempty" type:"Repeated"`
	HasAssets *bool                         `json:"HasAssets,omitempty" xml:"HasAssets,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EDBD3EB3-97DA-5465-AEF5-8DCA5DC5E395
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRoutineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoutineResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRoutineResponseBody) GetDefaultRelatedRecord() *string {
	return s.DefaultRelatedRecord
}

func (s *GetRoutineResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetRoutineResponseBody) GetEnvs() []*GetRoutineResponseBodyEnvs {
	return s.Envs
}

func (s *GetRoutineResponseBody) GetHasAssets() *bool {
	return s.HasAssets
}

func (s *GetRoutineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRoutineResponseBody) SetCreateTime(v string) *GetRoutineResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetRoutineResponseBody) SetDefaultRelatedRecord(v string) *GetRoutineResponseBody {
	s.DefaultRelatedRecord = &v
	return s
}

func (s *GetRoutineResponseBody) SetDescription(v string) *GetRoutineResponseBody {
	s.Description = &v
	return s
}

func (s *GetRoutineResponseBody) SetEnvs(v []*GetRoutineResponseBodyEnvs) *GetRoutineResponseBody {
	s.Envs = v
	return s
}

func (s *GetRoutineResponseBody) SetHasAssets(v bool) *GetRoutineResponseBody {
	s.HasAssets = &v
	return s
}

func (s *GetRoutineResponseBody) SetRequestId(v string) *GetRoutineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoutineResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRoutineResponseBodyEnvs struct {
	// The regions for canary release.
	CanaryAreaList []*string `json:"CanaryAreaList,omitempty" xml:"CanaryAreaList,omitempty" type:"Repeated"`
	// The version number for canary release.
	//
	// example:
	//
	// 1710120201067577628
	CanaryCodeVersion *string                               `json:"CanaryCodeVersion,omitempty" xml:"CanaryCodeVersion,omitempty"`
	CodeDeploy        *GetRoutineResponseBodyEnvsCodeDeploy `json:"CodeDeploy,omitempty" xml:"CodeDeploy,omitempty" type:"Struct"`
	// The version number of the code in the environment.
	//
	// example:
	//
	// 1710120201067577628
	CodeVersion *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	// The environment type.
	//
	// example:
	//
	// production
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s GetRoutineResponseBodyEnvs) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineResponseBodyEnvs) GoString() string {
	return s.String()
}

func (s *GetRoutineResponseBodyEnvs) GetCanaryAreaList() []*string {
	return s.CanaryAreaList
}

func (s *GetRoutineResponseBodyEnvs) GetCanaryCodeVersion() *string {
	return s.CanaryCodeVersion
}

func (s *GetRoutineResponseBodyEnvs) GetCodeDeploy() *GetRoutineResponseBodyEnvsCodeDeploy {
	return s.CodeDeploy
}

func (s *GetRoutineResponseBodyEnvs) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *GetRoutineResponseBodyEnvs) GetEnv() *string {
	return s.Env
}

func (s *GetRoutineResponseBodyEnvs) SetCanaryAreaList(v []*string) *GetRoutineResponseBodyEnvs {
	s.CanaryAreaList = v
	return s
}

func (s *GetRoutineResponseBodyEnvs) SetCanaryCodeVersion(v string) *GetRoutineResponseBodyEnvs {
	s.CanaryCodeVersion = &v
	return s
}

func (s *GetRoutineResponseBodyEnvs) SetCodeDeploy(v *GetRoutineResponseBodyEnvsCodeDeploy) *GetRoutineResponseBodyEnvs {
	s.CodeDeploy = v
	return s
}

func (s *GetRoutineResponseBodyEnvs) SetCodeVersion(v string) *GetRoutineResponseBodyEnvs {
	s.CodeVersion = &v
	return s
}

func (s *GetRoutineResponseBodyEnvs) SetEnv(v string) *GetRoutineResponseBodyEnvs {
	s.Env = &v
	return s
}

func (s *GetRoutineResponseBodyEnvs) Validate() error {
	return dara.Validate(s)
}

type GetRoutineResponseBodyEnvsCodeDeploy struct {
	CodeVersions []*GetRoutineResponseBodyEnvsCodeDeployCodeVersions `json:"CodeVersions,omitempty" xml:"CodeVersions,omitempty" type:"Repeated"`
	CreationTime *string                                             `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DeployId     *string                                             `json:"DeployId,omitempty" xml:"DeployId,omitempty"`
	Strategy     *string                                             `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s GetRoutineResponseBodyEnvsCodeDeploy) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineResponseBodyEnvsCodeDeploy) GoString() string {
	return s.String()
}

func (s *GetRoutineResponseBodyEnvsCodeDeploy) GetCodeVersions() []*GetRoutineResponseBodyEnvsCodeDeployCodeVersions {
	return s.CodeVersions
}

func (s *GetRoutineResponseBodyEnvsCodeDeploy) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetRoutineResponseBodyEnvsCodeDeploy) GetDeployId() *string {
	return s.DeployId
}

func (s *GetRoutineResponseBodyEnvsCodeDeploy) GetStrategy() *string {
	return s.Strategy
}

func (s *GetRoutineResponseBodyEnvsCodeDeploy) SetCodeVersions(v []*GetRoutineResponseBodyEnvsCodeDeployCodeVersions) *GetRoutineResponseBodyEnvsCodeDeploy {
	s.CodeVersions = v
	return s
}

func (s *GetRoutineResponseBodyEnvsCodeDeploy) SetCreationTime(v string) *GetRoutineResponseBodyEnvsCodeDeploy {
	s.CreationTime = &v
	return s
}

func (s *GetRoutineResponseBodyEnvsCodeDeploy) SetDeployId(v string) *GetRoutineResponseBodyEnvsCodeDeploy {
	s.DeployId = &v
	return s
}

func (s *GetRoutineResponseBodyEnvsCodeDeploy) SetStrategy(v string) *GetRoutineResponseBodyEnvsCodeDeploy {
	s.Strategy = &v
	return s
}

func (s *GetRoutineResponseBodyEnvsCodeDeploy) Validate() error {
	return dara.Validate(s)
}

type GetRoutineResponseBodyEnvsCodeDeployCodeVersions struct {
	CodeVersion *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Percentage  *int64  `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
}

func (s GetRoutineResponseBodyEnvsCodeDeployCodeVersions) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineResponseBodyEnvsCodeDeployCodeVersions) GoString() string {
	return s.String()
}

func (s *GetRoutineResponseBodyEnvsCodeDeployCodeVersions) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *GetRoutineResponseBodyEnvsCodeDeployCodeVersions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRoutineResponseBodyEnvsCodeDeployCodeVersions) GetDescription() *string {
	return s.Description
}

func (s *GetRoutineResponseBodyEnvsCodeDeployCodeVersions) GetPercentage() *int64 {
	return s.Percentage
}

func (s *GetRoutineResponseBodyEnvsCodeDeployCodeVersions) SetCodeVersion(v string) *GetRoutineResponseBodyEnvsCodeDeployCodeVersions {
	s.CodeVersion = &v
	return s
}

func (s *GetRoutineResponseBodyEnvsCodeDeployCodeVersions) SetCreateTime(v string) *GetRoutineResponseBodyEnvsCodeDeployCodeVersions {
	s.CreateTime = &v
	return s
}

func (s *GetRoutineResponseBodyEnvsCodeDeployCodeVersions) SetDescription(v string) *GetRoutineResponseBodyEnvsCodeDeployCodeVersions {
	s.Description = &v
	return s
}

func (s *GetRoutineResponseBodyEnvsCodeDeployCodeVersions) SetPercentage(v int64) *GetRoutineResponseBodyEnvsCodeDeployCodeVersions {
	s.Percentage = &v
	return s
}

func (s *GetRoutineResponseBodyEnvsCodeDeployCodeVersions) Validate() error {
	return dara.Validate(s)
}
