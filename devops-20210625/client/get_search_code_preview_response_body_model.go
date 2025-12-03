// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSearchCodePreviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetSearchCodePreviewResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetSearchCodePreviewResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetSearchCodePreviewResponseBody
	GetRequestId() *string
	SetResult(v *GetSearchCodePreviewResponseBodyResult) *GetSearchCodePreviewResponseBody
	GetResult() *GetSearchCodePreviewResponseBodyResult
	SetSuccess(v bool) *GetSearchCodePreviewResponseBody
	GetSuccess() *bool
}

type GetSearchCodePreviewResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// A7586FEB-E48D-5579-983F-74981FBFF627
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetSearchCodePreviewResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSearchCodePreviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSearchCodePreviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetSearchCodePreviewResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSearchCodePreviewResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSearchCodePreviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSearchCodePreviewResponseBody) GetResult() *GetSearchCodePreviewResponseBodyResult {
	return s.Result
}

func (s *GetSearchCodePreviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSearchCodePreviewResponseBody) SetErrorCode(v string) *GetSearchCodePreviewResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSearchCodePreviewResponseBody) SetErrorMessage(v string) *GetSearchCodePreviewResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSearchCodePreviewResponseBody) SetRequestId(v string) *GetSearchCodePreviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSearchCodePreviewResponseBody) SetResult(v *GetSearchCodePreviewResponseBodyResult) *GetSearchCodePreviewResponseBody {
	s.Result = v
	return s
}

func (s *GetSearchCodePreviewResponseBody) SetSuccess(v bool) *GetSearchCodePreviewResponseBody {
	s.Success = &v
	return s
}

func (s *GetSearchCodePreviewResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSearchCodePreviewResponseBodyResult struct {
	// example:
	//
	// xxx
	DocId            *string                                                 `json:"docId,omitempty" xml:"docId,omitempty"`
	HighlightTextMap *GetSearchCodePreviewResponseBodyResultHighlightTextMap `json:"highlightTextMap,omitempty" xml:"highlightTextMap,omitempty" type:"Struct"`
	Source           *GetSearchCodePreviewResponseBodyResultSource           `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s GetSearchCodePreviewResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetSearchCodePreviewResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSearchCodePreviewResponseBodyResult) GetDocId() *string {
	return s.DocId
}

func (s *GetSearchCodePreviewResponseBodyResult) GetHighlightTextMap() *GetSearchCodePreviewResponseBodyResultHighlightTextMap {
	return s.HighlightTextMap
}

func (s *GetSearchCodePreviewResponseBodyResult) GetSource() *GetSearchCodePreviewResponseBodyResultSource {
	return s.Source
}

func (s *GetSearchCodePreviewResponseBodyResult) SetDocId(v string) *GetSearchCodePreviewResponseBodyResult {
	s.DocId = &v
	return s
}

func (s *GetSearchCodePreviewResponseBodyResult) SetHighlightTextMap(v *GetSearchCodePreviewResponseBodyResultHighlightTextMap) *GetSearchCodePreviewResponseBodyResult {
	s.HighlightTextMap = v
	return s
}

func (s *GetSearchCodePreviewResponseBodyResult) SetSource(v *GetSearchCodePreviewResponseBodyResultSource) *GetSearchCodePreviewResponseBodyResult {
	s.Source = v
	return s
}

func (s *GetSearchCodePreviewResponseBodyResult) Validate() error {
	if s.HighlightTextMap != nil {
		if err := s.HighlightTextMap.Validate(); err != nil {
			return err
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSearchCodePreviewResponseBodyResultHighlightTextMap struct {
	// example:
	//
	// xxx
	Clob *string `json:"clob,omitempty" xml:"clob,omitempty"`
	// example:
	//
	// test.rb
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 5ffd468b1e45db3c1cc26ad6
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s GetSearchCodePreviewResponseBodyResultHighlightTextMap) String() string {
	return dara.Prettify(s)
}

func (s GetSearchCodePreviewResponseBodyResultHighlightTextMap) GoString() string {
	return s.String()
}

func (s *GetSearchCodePreviewResponseBodyResultHighlightTextMap) GetClob() *string {
	return s.Clob
}

func (s *GetSearchCodePreviewResponseBodyResultHighlightTextMap) GetFileName() *string {
	return s.FileName
}

func (s *GetSearchCodePreviewResponseBodyResultHighlightTextMap) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetSearchCodePreviewResponseBodyResultHighlightTextMap) SetClob(v string) *GetSearchCodePreviewResponseBodyResultHighlightTextMap {
	s.Clob = &v
	return s
}

func (s *GetSearchCodePreviewResponseBodyResultHighlightTextMap) SetFileName(v string) *GetSearchCodePreviewResponseBodyResultHighlightTextMap {
	s.FileName = &v
	return s
}

func (s *GetSearchCodePreviewResponseBodyResultHighlightTextMap) SetOrganizationId(v string) *GetSearchCodePreviewResponseBodyResultHighlightTextMap {
	s.OrganizationId = &v
	return s
}

func (s *GetSearchCodePreviewResponseBodyResultHighlightTextMap) Validate() error {
	return dara.Validate(s)
}

type GetSearchCodePreviewResponseBodyResultSource struct {
	// example:
	//
	// master
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// 2022-12-12 12:12:12
	CheckinDate *string `json:"checkinDate,omitempty" xml:"checkinDate,omitempty"`
	// example:
	//
	// test.rb
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// config/environments/test.rb
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// example:
	//
	// Ruby
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 5f9f9f6122a8c7ff3934f99a
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// codeup/test-repo
	RepoPath *string `json:"repoPath,omitempty" xml:"repoPath,omitempty"`
}

func (s GetSearchCodePreviewResponseBodyResultSource) String() string {
	return dara.Prettify(s)
}

func (s GetSearchCodePreviewResponseBodyResultSource) GoString() string {
	return s.String()
}

func (s *GetSearchCodePreviewResponseBodyResultSource) GetBranch() *string {
	return s.Branch
}

func (s *GetSearchCodePreviewResponseBodyResultSource) GetCheckinDate() *string {
	return s.CheckinDate
}

func (s *GetSearchCodePreviewResponseBodyResultSource) GetFileName() *string {
	return s.FileName
}

func (s *GetSearchCodePreviewResponseBodyResultSource) GetFilePath() *string {
	return s.FilePath
}

func (s *GetSearchCodePreviewResponseBodyResultSource) GetLanguage() *string {
	return s.Language
}

func (s *GetSearchCodePreviewResponseBodyResultSource) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetSearchCodePreviewResponseBodyResultSource) GetRepoPath() *string {
	return s.RepoPath
}

func (s *GetSearchCodePreviewResponseBodyResultSource) SetBranch(v string) *GetSearchCodePreviewResponseBodyResultSource {
	s.Branch = &v
	return s
}

func (s *GetSearchCodePreviewResponseBodyResultSource) SetCheckinDate(v string) *GetSearchCodePreviewResponseBodyResultSource {
	s.CheckinDate = &v
	return s
}

func (s *GetSearchCodePreviewResponseBodyResultSource) SetFileName(v string) *GetSearchCodePreviewResponseBodyResultSource {
	s.FileName = &v
	return s
}

func (s *GetSearchCodePreviewResponseBodyResultSource) SetFilePath(v string) *GetSearchCodePreviewResponseBodyResultSource {
	s.FilePath = &v
	return s
}

func (s *GetSearchCodePreviewResponseBodyResultSource) SetLanguage(v string) *GetSearchCodePreviewResponseBodyResultSource {
	s.Language = &v
	return s
}

func (s *GetSearchCodePreviewResponseBodyResultSource) SetOrganizationId(v string) *GetSearchCodePreviewResponseBodyResultSource {
	s.OrganizationId = &v
	return s
}

func (s *GetSearchCodePreviewResponseBodyResultSource) SetRepoPath(v string) *GetSearchCodePreviewResponseBodyResultSource {
	s.RepoPath = &v
	return s
}

func (s *GetSearchCodePreviewResponseBodyResultSource) Validate() error {
	return dara.Validate(s)
}
