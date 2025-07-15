// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEnterpriseVocAnalysisTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetApiKey() *string
	SetContentTagsShrink(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetContentTagsShrink() *string
	SetContentsShrink(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetContentsShrink() *string
	SetFileKey(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetFileKey() *string
	SetFilterTagsShrink(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetFilterTagsShrink() *string
	SetMaterialType(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetMaterialType() *string
	SetModelId(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetModelId() *string
	SetPositiveSample(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetPositiveSample() *string
	SetPositiveSampleFileKey(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetPositiveSampleFileKey() *string
	SetTaskType(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetTaskType() *string
	SetWorkspaceId(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest
	GetWorkspaceId() *string
}

type SubmitEnterpriseVocAnalysisTaskShrinkRequest struct {
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// This parameter is required.
	ContentTagsShrink *string `json:"ContentTags,omitempty" xml:"ContentTags,omitempty"`
	ContentsShrink    *string `json:"Contents,omitempty" xml:"Contents,omitempty"`
	// example:
	//
	// oss://default/bucket-name/materialDocument/tenant_agent/fileName
	FileKey          *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	FilterTagsShrink *string `json:"FilterTags,omitempty" xml:"FilterTags,omitempty"`
	// example:
	//
	// shortContent
	MaterialType *string `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	ModelId        *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	PositiveSample *string `json:"PositiveSample,omitempty" xml:"PositiveSample,omitempty"`
	// example:
	//
	// oss://default/bucket-name/path/xxx.xlsx
	PositiveSampleFileKey *string `json:"PositiveSampleFileKey,omitempty" xml:"PositiveSampleFileKey,omitempty"`
	// example:
	//
	// lightAppSass
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitEnterpriseVocAnalysisTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitEnterpriseVocAnalysisTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetContentTagsShrink() *string {
	return s.ContentTagsShrink
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetContentsShrink() *string {
	return s.ContentsShrink
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetFilterTagsShrink() *string {
	return s.FilterTagsShrink
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetMaterialType() *string {
	return s.MaterialType
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetPositiveSample() *string {
	return s.PositiveSample
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetPositiveSampleFileKey() *string {
	return s.PositiveSampleFileKey
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetApiKey(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.ApiKey = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetContentTagsShrink(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.ContentTagsShrink = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetContentsShrink(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.ContentsShrink = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetFileKey(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.FileKey = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetFilterTagsShrink(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.FilterTagsShrink = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetMaterialType(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.MaterialType = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetModelId(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetPositiveSample(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.PositiveSample = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetPositiveSampleFileKey(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.PositiveSampleFileKey = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetTaskType(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) SetWorkspaceId(v string) *SubmitEnterpriseVocAnalysisTaskShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitEnterpriseVocAnalysisTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
