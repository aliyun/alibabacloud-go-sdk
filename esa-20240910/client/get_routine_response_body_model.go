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
	Envs []*GetRoutineResponseBodyEnvs `json:"Envs,omitempty" xml:"Envs,omitempty" type:"Repeated"`
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
	CanaryCodeVersion *string `json:"CanaryCodeVersion,omitempty" xml:"CanaryCodeVersion,omitempty"`
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
