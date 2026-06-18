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
	// The creation time of the edge function Routine.
	//
	// example:
	//
	// 2024-03-11T01:23:21Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The default domain name for accessing the Routine.
	//
	// example:
	//
	// routine1.example.com
	DefaultRelatedRecord *string `json:"DefaultRelatedRecord,omitempty" xml:"DefaultRelatedRecord,omitempty"`
	// The description of the edge function Routine.
	//
	// example:
	//
	// ZWRpdCByb3V0aW5lIGNvbmZpZyBkZXNjcmlwdGlvbg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// A list of environments.
	Envs []*GetRoutineResponseBodyEnvs `json:"Envs,omitempty" xml:"Envs,omitempty" type:"Repeated"`
	// Indicates whether the Routine includes Assets.
	//
	// example:
	//
	// false
	HasAssets *bool `json:"HasAssets,omitempty" xml:"HasAssets,omitempty"`
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
	if s.Envs != nil {
		for _, item := range s.Envs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRoutineResponseBodyEnvs struct {
	// Details of the canary release for a code version.
	CodeDeploy *GetRoutineResponseBodyEnvsCodeDeploy `json:"CodeDeploy,omitempty" xml:"CodeDeploy,omitempty" type:"Struct"`
	// The environment name.
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

func (s *GetRoutineResponseBodyEnvs) GetCodeDeploy() *GetRoutineResponseBodyEnvsCodeDeploy {
	return s.CodeDeploy
}

func (s *GetRoutineResponseBodyEnvs) GetEnv() *string {
	return s.Env
}

func (s *GetRoutineResponseBodyEnvs) SetCodeDeploy(v *GetRoutineResponseBodyEnvsCodeDeploy) *GetRoutineResponseBodyEnvs {
	s.CodeDeploy = v
	return s
}

func (s *GetRoutineResponseBodyEnvs) SetEnv(v string) *GetRoutineResponseBodyEnvs {
	s.Env = &v
	return s
}

func (s *GetRoutineResponseBodyEnvs) Validate() error {
	if s.CodeDeploy != nil {
		if err := s.CodeDeploy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRoutineResponseBodyEnvsCodeDeploy struct {
	// A list of deployed code versions.
	CodeVersions []*GetRoutineResponseBodyEnvsCodeDeployCodeVersions `json:"CodeVersions,omitempty" xml:"CodeVersions,omitempty" type:"Repeated"`
	// The time the deployment was created.
	//
	// example:
	//
	// 2023-05-11T09:21:36Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The deployment ID.
	//
	// example:
	//
	// 589267
	DeployId *string `json:"DeployId,omitempty" xml:"DeployId,omitempty"`
	// The deployment strategy. The default value is `percentage`.
	//
	// example:
	//
	// percentage
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
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
	if s.CodeVersions != nil {
		for _, item := range s.CodeVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRoutineResponseBodyEnvsCodeDeployCodeVersions struct {
	// The code version ID.
	//
	// example:
	//
	// 1746583193971399525
	CodeVersion *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	// The creation time of the code version.
	//
	// example:
	//
	// 2025-07-23T09:01:40Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the code version.
	//
	// example:
	//
	// code version 1.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The percentage of traffic routed to this code version.
	//
	// example:
	//
	// 100
	Percentage *int64 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
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
