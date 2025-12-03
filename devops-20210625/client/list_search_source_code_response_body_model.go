// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchSourceCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListSearchSourceCodeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListSearchSourceCodeResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListSearchSourceCodeResponseBody
	GetRequestId() *string
	SetResult(v []*ListSearchSourceCodeResponseBodyResult) *ListSearchSourceCodeResponseBody
	GetResult() []*ListSearchSourceCodeResponseBodyResult
	SetSuccess(v bool) *ListSearchSourceCodeResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListSearchSourceCodeResponseBody
	GetTotal() *int64
}

type ListSearchSourceCodeResponseBody struct {
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
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListSearchSourceCodeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListSearchSourceCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSearchSourceCodeResponseBody) GoString() string {
	return s.String()
}

func (s *ListSearchSourceCodeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSearchSourceCodeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSearchSourceCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSearchSourceCodeResponseBody) GetResult() []*ListSearchSourceCodeResponseBodyResult {
	return s.Result
}

func (s *ListSearchSourceCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSearchSourceCodeResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListSearchSourceCodeResponseBody) SetErrorCode(v string) *ListSearchSourceCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSearchSourceCodeResponseBody) SetErrorMessage(v string) *ListSearchSourceCodeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSearchSourceCodeResponseBody) SetRequestId(v string) *ListSearchSourceCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSearchSourceCodeResponseBody) SetResult(v []*ListSearchSourceCodeResponseBodyResult) *ListSearchSourceCodeResponseBody {
	s.Result = v
	return s
}

func (s *ListSearchSourceCodeResponseBody) SetSuccess(v bool) *ListSearchSourceCodeResponseBody {
	s.Success = &v
	return s
}

func (s *ListSearchSourceCodeResponseBody) SetTotal(v int64) *ListSearchSourceCodeResponseBody {
	s.Total = &v
	return s
}

func (s *ListSearchSourceCodeResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSearchSourceCodeResponseBodyResult struct {
	// example:
	//
	// 60d54f3daccf2bbd6659f3ad/gitlabhq/master/spec/frontend/snippets/test_utils.js
	DocId            *string                                                 `json:"docId,omitempty" xml:"docId,omitempty"`
	HighlightTextMap *ListSearchSourceCodeResponseBodyResultHighlightTextMap `json:"highlightTextMap,omitempty" xml:"highlightTextMap,omitempty" type:"Struct"`
	Source           *ListSearchSourceCodeResponseBodyResultSource           `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s ListSearchSourceCodeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSearchSourceCodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSearchSourceCodeResponseBodyResult) GetDocId() *string {
	return s.DocId
}

func (s *ListSearchSourceCodeResponseBodyResult) GetHighlightTextMap() *ListSearchSourceCodeResponseBodyResultHighlightTextMap {
	return s.HighlightTextMap
}

func (s *ListSearchSourceCodeResponseBodyResult) GetSource() *ListSearchSourceCodeResponseBodyResultSource {
	return s.Source
}

func (s *ListSearchSourceCodeResponseBodyResult) SetDocId(v string) *ListSearchSourceCodeResponseBodyResult {
	s.DocId = &v
	return s
}

func (s *ListSearchSourceCodeResponseBodyResult) SetHighlightTextMap(v *ListSearchSourceCodeResponseBodyResultHighlightTextMap) *ListSearchSourceCodeResponseBodyResult {
	s.HighlightTextMap = v
	return s
}

func (s *ListSearchSourceCodeResponseBodyResult) SetSource(v *ListSearchSourceCodeResponseBodyResultSource) *ListSearchSourceCodeResponseBodyResult {
	s.Source = v
	return s
}

func (s *ListSearchSourceCodeResponseBodyResult) Validate() error {
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

type ListSearchSourceCodeResponseBodyResultHighlightTextMap struct {
	// example:
	//
	// xxx
	Clob *string `json:"clob,omitempty" xml:"clob,omitempty"`
	// example:
	//
	// test.java
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// java
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ListSearchSourceCodeResponseBodyResultHighlightTextMap) String() string {
	return dara.Prettify(s)
}

func (s ListSearchSourceCodeResponseBodyResultHighlightTextMap) GoString() string {
	return s.String()
}

func (s *ListSearchSourceCodeResponseBodyResultHighlightTextMap) GetClob() *string {
	return s.Clob
}

func (s *ListSearchSourceCodeResponseBodyResultHighlightTextMap) GetFileName() *string {
	return s.FileName
}

func (s *ListSearchSourceCodeResponseBodyResultHighlightTextMap) GetLanguage() *string {
	return s.Language
}

func (s *ListSearchSourceCodeResponseBodyResultHighlightTextMap) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListSearchSourceCodeResponseBodyResultHighlightTextMap) SetClob(v string) *ListSearchSourceCodeResponseBodyResultHighlightTextMap {
	s.Clob = &v
	return s
}

func (s *ListSearchSourceCodeResponseBodyResultHighlightTextMap) SetFileName(v string) *ListSearchSourceCodeResponseBodyResultHighlightTextMap {
	s.FileName = &v
	return s
}

func (s *ListSearchSourceCodeResponseBodyResultHighlightTextMap) SetLanguage(v string) *ListSearchSourceCodeResponseBodyResultHighlightTextMap {
	s.Language = &v
	return s
}

func (s *ListSearchSourceCodeResponseBodyResultHighlightTextMap) SetOrganizationId(v string) *ListSearchSourceCodeResponseBodyResultHighlightTextMap {
	s.OrganizationId = &v
	return s
}

func (s *ListSearchSourceCodeResponseBodyResultHighlightTextMap) Validate() error {
	return dara.Validate(s)
}

type ListSearchSourceCodeResponseBodyResultSource struct {
	// example:
	//
	// master
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// 2022-10-10 10:00:00
	CheckinDate *string `json:"checkinDate,omitempty" xml:"checkinDate,omitempty"`
	// example:
	//
	// test_utils.js
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// spec/frontend/snippets/test_utils.js
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// example:
	//
	// JavaScript
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// codeup/test-repo
	RepoPath *string `json:"repoPath,omitempty" xml:"repoPath,omitempty"`
}

func (s ListSearchSourceCodeResponseBodyResultSource) String() string {
	return dara.Prettify(s)
}

func (s ListSearchSourceCodeResponseBodyResultSource) GoString() string {
	return s.String()
}

func (s *ListSearchSourceCodeResponseBodyResultSource) GetBranch() *string {
	return s.Branch
}

func (s *ListSearchSourceCodeResponseBodyResultSource) GetCheckinDate() *string {
	return s.CheckinDate
}

func (s *ListSearchSourceCodeResponseBodyResultSource) GetFileName() *string {
	return s.FileName
}

func (s *ListSearchSourceCodeResponseBodyResultSource) GetFilePath() *string {
	return s.FilePath
}

func (s *ListSearchSourceCodeResponseBodyResultSource) GetLanguage() *string {
	return s.Language
}

func (s *ListSearchSourceCodeResponseBodyResultSource) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListSearchSourceCodeResponseBodyResultSource) GetRepoPath() *string {
	return s.RepoPath
}

func (s *ListSearchSourceCodeResponseBodyResultSource) SetBranch(v string) *ListSearchSourceCodeResponseBodyResultSource {
	s.Branch = &v
	return s
}

func (s *ListSearchSourceCodeResponseBodyResultSource) SetCheckinDate(v string) *ListSearchSourceCodeResponseBodyResultSource {
	s.CheckinDate = &v
	return s
}

func (s *ListSearchSourceCodeResponseBodyResultSource) SetFileName(v string) *ListSearchSourceCodeResponseBodyResultSource {
	s.FileName = &v
	return s
}

func (s *ListSearchSourceCodeResponseBodyResultSource) SetFilePath(v string) *ListSearchSourceCodeResponseBodyResultSource {
	s.FilePath = &v
	return s
}

func (s *ListSearchSourceCodeResponseBodyResultSource) SetLanguage(v string) *ListSearchSourceCodeResponseBodyResultSource {
	s.Language = &v
	return s
}

func (s *ListSearchSourceCodeResponseBodyResultSource) SetOrganizationId(v string) *ListSearchSourceCodeResponseBodyResultSource {
	s.OrganizationId = &v
	return s
}

func (s *ListSearchSourceCodeResponseBodyResultSource) SetRepoPath(v string) *ListSearchSourceCodeResponseBodyResultSource {
	s.RepoPath = &v
	return s
}

func (s *ListSearchSourceCodeResponseBodyResultSource) Validate() error {
	return dara.Validate(s)
}
