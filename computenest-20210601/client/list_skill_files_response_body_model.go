// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSkillFilesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSkillFilesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSkillFilesResponseBody
	GetRequestId() *string
	SetSkillFiles(v []*ListSkillFilesResponseBodySkillFiles) *ListSkillFilesResponseBody
	GetSkillFiles() []*ListSkillFilesResponseBodySkillFiles
	SetTotalCount(v int32) *ListSkillFilesResponseBody
	GetTotalCount() *int32
}

type ListSkillFilesResponseBody struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2849EE73-AFFA-5AFD-9575-12FA886451DA
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SkillFiles []*ListSkillFilesResponseBodySkillFiles `json:"SkillFiles,omitempty" xml:"SkillFiles,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSkillFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillFilesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSkillFilesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSkillFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillFilesResponseBody) GetSkillFiles() []*ListSkillFilesResponseBodySkillFiles {
	return s.SkillFiles
}

func (s *ListSkillFilesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSkillFilesResponseBody) SetMaxResults(v int32) *ListSkillFilesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSkillFilesResponseBody) SetNextToken(v string) *ListSkillFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSkillFilesResponseBody) SetRequestId(v string) *ListSkillFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillFilesResponseBody) SetSkillFiles(v []*ListSkillFilesResponseBodySkillFiles) *ListSkillFilesResponseBody {
	s.SkillFiles = v
	return s
}

func (s *ListSkillFilesResponseBody) SetTotalCount(v int32) *ListSkillFilesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSkillFilesResponseBody) Validate() error {
	if s.SkillFiles != nil {
		for _, item := range s.SkillFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSkillFilesResponseBodySkillFiles struct {
	// example:
	//
	// SKILL.md
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// https://embedding-pic.oss-cn-beijing-internal.aliyuncs.com/SKILL.md
	SignedUrl *string `json:"SignedUrl,omitempty" xml:"SignedUrl,omitempty"`
}

func (s ListSkillFilesResponseBodySkillFiles) String() string {
	return dara.Prettify(s)
}

func (s ListSkillFilesResponseBodySkillFiles) GoString() string {
	return s.String()
}

func (s *ListSkillFilesResponseBodySkillFiles) GetFilePath() *string {
	return s.FilePath
}

func (s *ListSkillFilesResponseBodySkillFiles) GetSignedUrl() *string {
	return s.SignedUrl
}

func (s *ListSkillFilesResponseBodySkillFiles) SetFilePath(v string) *ListSkillFilesResponseBodySkillFiles {
	s.FilePath = &v
	return s
}

func (s *ListSkillFilesResponseBodySkillFiles) SetSignedUrl(v string) *ListSkillFilesResponseBodySkillFiles {
	s.SignedUrl = &v
	return s
}

func (s *ListSkillFilesResponseBodySkillFiles) Validate() error {
	return dara.Validate(s)
}
