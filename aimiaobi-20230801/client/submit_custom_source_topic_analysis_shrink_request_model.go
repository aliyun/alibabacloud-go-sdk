// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomSourceTopicAnalysisShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisTypesShrink(v string) *SubmitCustomSourceTopicAnalysisShrinkRequest
	GetAnalysisTypesShrink() *string
	SetFileType(v string) *SubmitCustomSourceTopicAnalysisShrinkRequest
	GetFileType() *string
	SetFileUrl(v string) *SubmitCustomSourceTopicAnalysisShrinkRequest
	GetFileUrl() *string
	SetMaxTopicSize(v int32) *SubmitCustomSourceTopicAnalysisShrinkRequest
	GetMaxTopicSize() *int32
	SetNewsShrink(v string) *SubmitCustomSourceTopicAnalysisShrinkRequest
	GetNewsShrink() *string
	SetTopicsShrink(v string) *SubmitCustomSourceTopicAnalysisShrinkRequest
	GetTopicsShrink() *string
	SetTopicsFileUrl(v string) *SubmitCustomSourceTopicAnalysisShrinkRequest
	GetTopicsFileUrl() *string
	SetWorkspaceId(v string) *SubmitCustomSourceTopicAnalysisShrinkRequest
	GetWorkspaceId() *string
}

type SubmitCustomSourceTopicAnalysisShrinkRequest struct {
	AnalysisTypesShrink *string `json:"AnalysisTypes,omitempty" xml:"AnalysisTypes,omitempty"`
	// example:
	//
	// json
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// http://www.example.com/xxx.json
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// 50
	MaxTopicSize *int32  `json:"MaxTopicSize,omitempty" xml:"MaxTopicSize,omitempty"`
	NewsShrink   *string `json:"News,omitempty" xml:"News,omitempty"`
	TopicsShrink *string `json:"Topics,omitempty" xml:"Topics,omitempty"`
	// example:
	//
	// http://www.example.com/xxx.jsonline
	TopicsFileUrl *string `json:"TopicsFileUrl,omitempty" xml:"TopicsFileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitCustomSourceTopicAnalysisShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomSourceTopicAnalysisShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) GetAnalysisTypesShrink() *string {
	return s.AnalysisTypesShrink
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) GetFileType() *string {
	return s.FileType
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) GetMaxTopicSize() *int32 {
	return s.MaxTopicSize
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) GetNewsShrink() *string {
	return s.NewsShrink
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) GetTopicsShrink() *string {
	return s.TopicsShrink
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) GetTopicsFileUrl() *string {
	return s.TopicsFileUrl
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) SetAnalysisTypesShrink(v string) *SubmitCustomSourceTopicAnalysisShrinkRequest {
	s.AnalysisTypesShrink = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) SetFileType(v string) *SubmitCustomSourceTopicAnalysisShrinkRequest {
	s.FileType = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) SetFileUrl(v string) *SubmitCustomSourceTopicAnalysisShrinkRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) SetMaxTopicSize(v int32) *SubmitCustomSourceTopicAnalysisShrinkRequest {
	s.MaxTopicSize = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) SetNewsShrink(v string) *SubmitCustomSourceTopicAnalysisShrinkRequest {
	s.NewsShrink = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) SetTopicsShrink(v string) *SubmitCustomSourceTopicAnalysisShrinkRequest {
	s.TopicsShrink = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) SetTopicsFileUrl(v string) *SubmitCustomSourceTopicAnalysisShrinkRequest {
	s.TopicsFileUrl = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) SetWorkspaceId(v string) *SubmitCustomSourceTopicAnalysisShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitCustomSourceTopicAnalysisShrinkRequest) Validate() error {
	return dara.Validate(s)
}
