// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactBuildRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *GetArtifactBuildRuleResponseBody
	GetArtifactType() *string
	SetBuildRuleId(v string) *GetArtifactBuildRuleResponseBody
	GetBuildRuleId() *string
	SetCode(v string) *GetArtifactBuildRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *GetArtifactBuildRuleResponseBody
	GetIsSuccess() *bool
	SetParameters(v *GetArtifactBuildRuleResponseBodyParameters) *GetArtifactBuildRuleResponseBody
	GetParameters() *GetArtifactBuildRuleResponseBodyParameters
	SetRequestId(v string) *GetArtifactBuildRuleResponseBody
	GetRequestId() *string
	SetScopeId(v string) *GetArtifactBuildRuleResponseBody
	GetScopeId() *string
	SetScopeType(v string) *GetArtifactBuildRuleResponseBody
	GetScopeType() *string
}

type GetArtifactBuildRuleResponseBody struct {
	// example:
	//
	// ACCELERATED_IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// example:
	//
	// crabr-o2670wqz2n70****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	IsSuccess  *bool                                       `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	Parameters *GetArtifactBuildRuleResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// example:
	//
	// 7A3E98F6-296C-54AC-A612-B75E7777D4C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// crr-8dz3aedjqlmk****
	ScopeId *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	// example:
	//
	// REPOSITORY
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s GetArtifactBuildRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactBuildRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildRuleResponseBody) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *GetArtifactBuildRuleResponseBody) GetBuildRuleId() *string {
	return s.BuildRuleId
}

func (s *GetArtifactBuildRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetArtifactBuildRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetArtifactBuildRuleResponseBody) GetParameters() *GetArtifactBuildRuleResponseBodyParameters {
	return s.Parameters
}

func (s *GetArtifactBuildRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetArtifactBuildRuleResponseBody) GetScopeId() *string {
	return s.ScopeId
}

func (s *GetArtifactBuildRuleResponseBody) GetScopeType() *string {
	return s.ScopeType
}

func (s *GetArtifactBuildRuleResponseBody) SetArtifactType(v string) *GetArtifactBuildRuleResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) SetBuildRuleId(v string) *GetArtifactBuildRuleResponseBody {
	s.BuildRuleId = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) SetCode(v string) *GetArtifactBuildRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) SetIsSuccess(v bool) *GetArtifactBuildRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) SetParameters(v *GetArtifactBuildRuleResponseBodyParameters) *GetArtifactBuildRuleResponseBody {
	s.Parameters = v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) SetRequestId(v string) *GetArtifactBuildRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) SetScopeId(v string) *GetArtifactBuildRuleResponseBody {
	s.ScopeId = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) SetScopeType(v string) *GetArtifactBuildRuleResponseBody {
	s.ScopeType = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetArtifactBuildRuleResponseBodyParameters struct {
	ImageIndexOnly *bool   `json:"ImageIndexOnly,omitempty" xml:"ImageIndexOnly,omitempty"`
	PriorityFile   *string `json:"PriorityFile,omitempty" xml:"PriorityFile,omitempty"`
}

func (s GetArtifactBuildRuleResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactBuildRuleResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildRuleResponseBodyParameters) GetImageIndexOnly() *bool {
	return s.ImageIndexOnly
}

func (s *GetArtifactBuildRuleResponseBodyParameters) GetPriorityFile() *string {
	return s.PriorityFile
}

func (s *GetArtifactBuildRuleResponseBodyParameters) SetImageIndexOnly(v bool) *GetArtifactBuildRuleResponseBodyParameters {
	s.ImageIndexOnly = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBodyParameters) SetPriorityFile(v string) *GetArtifactBuildRuleResponseBodyParameters {
	s.PriorityFile = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBodyParameters) Validate() error {
	return dara.Validate(s)
}
