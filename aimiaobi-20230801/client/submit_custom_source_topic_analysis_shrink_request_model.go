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
	// The types of analysis for hot topic selection. Multiple values are supported. If you omit this parameter, the service analyzes all types by default. If you pass an empty array, the service performs only clustering and skips the analysis of hot topics for selection.
	//
	// `HotViewPoints`: Analyzes perspectives on hot topics.
	//
	// `WebReviewPoints`: Analyzes user viewpoints. This requires comments.
	//
	// `TimedViewPoints`: Analyzes perspectives on timeliness.
	//
	// `FreshViewPoints`: Analyzes novel perspectives.
	//
	// `TopicSummary`: Summarizes news content.
	AnalysisTypesShrink *string `json:"AnalysisTypes,omitempty" xml:"AnalysisTypes,omitempty"`
	// The file type. Valid values: `json` (JSON array) and `jsonLine` (JSON Lines).
	//
	// example:
	//
	// json
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The file URL. You must specify either `FileUrl` or `News`. For details on the file structure, see the description of the `News` parameter.
	//
	// example:
	//
	// http://www.example.com/xxx.json
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The maximum number of topics to analyze. By default, the service sorts clustered news by count in descending order and analyzes the top 50 topics. The maximum value is 200.
	//
	// example:
	//
	// 50
	MaxTopicSize *int32 `json:"MaxTopicSize,omitempty" xml:"MaxTopicSize,omitempty"`
	// A list of news articles. You must specify either `News` or `FileUrl`.
	NewsShrink *string `json:"News,omitempty" xml:"News,omitempty"`
	// A list of topics.
	TopicsShrink *string `json:"Topics,omitempty" xml:"Topics,omitempty"`
	// The URL of the file that contains the topic list. The file must be in JSON Lines format, with each line representing a single JSON object.
	//
	// example:
	//
	// http://www.example.com/xxx.jsonline
	TopicsFileUrl *string `json:"TopicsFileUrl,omitempty" xml:"TopicsFileUrl,omitempty"`
	// [The Model Studio workspace ID.](https://help.aliyun.com/document_detail/2782167.html)
	//
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
