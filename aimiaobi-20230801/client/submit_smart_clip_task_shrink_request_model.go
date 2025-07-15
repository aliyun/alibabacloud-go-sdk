// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmartClipTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEditingConfigShrink(v string) *SubmitSmartClipTaskShrinkRequest
	GetEditingConfigShrink() *string
	SetExtendParam(v string) *SubmitSmartClipTaskShrinkRequest
	GetExtendParam() *string
	SetInputConfigShrink(v string) *SubmitSmartClipTaskShrinkRequest
	GetInputConfigShrink() *string
	SetOutputConfigShrink(v string) *SubmitSmartClipTaskShrinkRequest
	GetOutputConfigShrink() *string
	SetWorkspaceId(v string) *SubmitSmartClipTaskShrinkRequest
	GetWorkspaceId() *string
}

type SubmitSmartClipTaskShrinkRequest struct {
	EditingConfigShrink *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	ExtendParam         *string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// This parameter is required.
	InputConfigShrink  *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	OutputConfigShrink *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitSmartClipTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskShrinkRequest) GetEditingConfigShrink() *string {
	return s.EditingConfigShrink
}

func (s *SubmitSmartClipTaskShrinkRequest) GetExtendParam() *string {
	return s.ExtendParam
}

func (s *SubmitSmartClipTaskShrinkRequest) GetInputConfigShrink() *string {
	return s.InputConfigShrink
}

func (s *SubmitSmartClipTaskShrinkRequest) GetOutputConfigShrink() *string {
	return s.OutputConfigShrink
}

func (s *SubmitSmartClipTaskShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitSmartClipTaskShrinkRequest) SetEditingConfigShrink(v string) *SubmitSmartClipTaskShrinkRequest {
	s.EditingConfigShrink = &v
	return s
}

func (s *SubmitSmartClipTaskShrinkRequest) SetExtendParam(v string) *SubmitSmartClipTaskShrinkRequest {
	s.ExtendParam = &v
	return s
}

func (s *SubmitSmartClipTaskShrinkRequest) SetInputConfigShrink(v string) *SubmitSmartClipTaskShrinkRequest {
	s.InputConfigShrink = &v
	return s
}

func (s *SubmitSmartClipTaskShrinkRequest) SetOutputConfigShrink(v string) *SubmitSmartClipTaskShrinkRequest {
	s.OutputConfigShrink = &v
	return s
}

func (s *SubmitSmartClipTaskShrinkRequest) SetWorkspaceId(v string) *SubmitSmartClipTaskShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitSmartClipTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
