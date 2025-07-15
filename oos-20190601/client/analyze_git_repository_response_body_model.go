// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeGitRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisResults(v []*AnalyzeGitRepositoryResponseBodyAnalysisResults) *AnalyzeGitRepositoryResponseBody
	GetAnalysisResults() []*AnalyzeGitRepositoryResponseBodyAnalysisResults
	SetCount(v int32) *AnalyzeGitRepositoryResponseBody
	GetCount() *int32
	SetRequestId(v string) *AnalyzeGitRepositoryResponseBody
	GetRequestId() *string
}

type AnalyzeGitRepositoryResponseBody struct {
	AnalysisResults []*AnalyzeGitRepositoryResponseBodyAnalysisResults `json:"AnalysisResults,omitempty" xml:"AnalysisResults,omitempty" type:"Repeated"`
	Count           *int32                                             `json:"Count,omitempty" xml:"Count,omitempty"`
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AnalyzeGitRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeGitRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *AnalyzeGitRepositoryResponseBody) GetAnalysisResults() []*AnalyzeGitRepositoryResponseBodyAnalysisResults {
	return s.AnalysisResults
}

func (s *AnalyzeGitRepositoryResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *AnalyzeGitRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AnalyzeGitRepositoryResponseBody) SetAnalysisResults(v []*AnalyzeGitRepositoryResponseBodyAnalysisResults) *AnalyzeGitRepositoryResponseBody {
	s.AnalysisResults = v
	return s
}

func (s *AnalyzeGitRepositoryResponseBody) SetCount(v int32) *AnalyzeGitRepositoryResponseBody {
	s.Count = &v
	return s
}

func (s *AnalyzeGitRepositoryResponseBody) SetRequestId(v string) *AnalyzeGitRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnalyzeGitRepositoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type AnalyzeGitRepositoryResponseBodyAnalysisResults struct {
	BuildFiles  []*AnalyzeGitRepositoryResponseBodyAnalysisResultsBuildFiles `json:"BuildFiles,omitempty" xml:"BuildFiles,omitempty" type:"Repeated"`
	BuildType   *string                                                      `json:"BuildType,omitempty" xml:"BuildType,omitempty"`
	RuntimeType *string                                                      `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
}

func (s AnalyzeGitRepositoryResponseBodyAnalysisResults) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeGitRepositoryResponseBodyAnalysisResults) GoString() string {
	return s.String()
}

func (s *AnalyzeGitRepositoryResponseBodyAnalysisResults) GetBuildFiles() []*AnalyzeGitRepositoryResponseBodyAnalysisResultsBuildFiles {
	return s.BuildFiles
}

func (s *AnalyzeGitRepositoryResponseBodyAnalysisResults) GetBuildType() *string {
	return s.BuildType
}

func (s *AnalyzeGitRepositoryResponseBodyAnalysisResults) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *AnalyzeGitRepositoryResponseBodyAnalysisResults) SetBuildFiles(v []*AnalyzeGitRepositoryResponseBodyAnalysisResultsBuildFiles) *AnalyzeGitRepositoryResponseBodyAnalysisResults {
	s.BuildFiles = v
	return s
}

func (s *AnalyzeGitRepositoryResponseBodyAnalysisResults) SetBuildType(v string) *AnalyzeGitRepositoryResponseBodyAnalysisResults {
	s.BuildType = &v
	return s
}

func (s *AnalyzeGitRepositoryResponseBodyAnalysisResults) SetRuntimeType(v string) *AnalyzeGitRepositoryResponseBodyAnalysisResults {
	s.RuntimeType = &v
	return s
}

func (s *AnalyzeGitRepositoryResponseBodyAnalysisResults) Validate() error {
	return dara.Validate(s)
}

type AnalyzeGitRepositoryResponseBodyAnalysisResultsBuildFiles struct {
	FileType *string   `json:"FileType,omitempty" xml:"FileType,omitempty"`
	Paths    []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
}

func (s AnalyzeGitRepositoryResponseBodyAnalysisResultsBuildFiles) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeGitRepositoryResponseBodyAnalysisResultsBuildFiles) GoString() string {
	return s.String()
}

func (s *AnalyzeGitRepositoryResponseBodyAnalysisResultsBuildFiles) GetFileType() *string {
	return s.FileType
}

func (s *AnalyzeGitRepositoryResponseBodyAnalysisResultsBuildFiles) GetPaths() []*string {
	return s.Paths
}

func (s *AnalyzeGitRepositoryResponseBodyAnalysisResultsBuildFiles) SetFileType(v string) *AnalyzeGitRepositoryResponseBodyAnalysisResultsBuildFiles {
	s.FileType = &v
	return s
}

func (s *AnalyzeGitRepositoryResponseBodyAnalysisResultsBuildFiles) SetPaths(v []*string) *AnalyzeGitRepositoryResponseBodyAnalysisResultsBuildFiles {
	s.Paths = v
	return s
}

func (s *AnalyzeGitRepositoryResponseBodyAnalysisResultsBuildFiles) Validate() error {
	return dara.Validate(s)
}
