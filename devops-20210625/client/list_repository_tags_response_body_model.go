// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListRepositoryTagsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListRepositoryTagsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListRepositoryTagsResponseBody
	GetRequestId() *string
	SetResult(v []*ListRepositoryTagsResponseBodyResult) *ListRepositoryTagsResponseBody
	GetResult() []*ListRepositoryTagsResponseBodyResult
	SetSuccess(v bool) *ListRepositoryTagsResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListRepositoryTagsResponseBody
	GetTotal() *int64
}

type ListRepositoryTagsResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListRepositoryTagsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 30
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListRepositoryTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListRepositoryTagsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListRepositoryTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepositoryTagsResponseBody) GetResult() []*ListRepositoryTagsResponseBodyResult {
	return s.Result
}

func (s *ListRepositoryTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRepositoryTagsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListRepositoryTagsResponseBody) SetErrorCode(v string) *ListRepositoryTagsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetErrorMessage(v string) *ListRepositoryTagsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetRequestId(v string) *ListRepositoryTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetResult(v []*ListRepositoryTagsResponseBodyResult) *ListRepositoryTagsResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetSuccess(v bool) *ListRepositoryTagsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetTotal(v int64) *ListRepositoryTagsResponseBody {
	s.Total = &v
	return s
}

func (s *ListRepositoryTagsResponseBody) Validate() error {
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

type ListRepositoryTagsResponseBodyResult struct {
	Commit *ListRepositoryTagsResponseBodyResultCommit `json:"commit,omitempty" xml:"commit,omitempty" type:"Struct"`
	// example:
	//
	// 9a494e7b88ca35cde00579af2df4ab46136c022e
	Id      *string `json:"id,omitempty" xml:"id,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// tag v1.0
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRepositoryTagsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryTagsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsResponseBodyResult) GetCommit() *ListRepositoryTagsResponseBodyResultCommit {
	return s.Commit
}

func (s *ListRepositoryTagsResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListRepositoryTagsResponseBodyResult) GetMessage() *string {
	return s.Message
}

func (s *ListRepositoryTagsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListRepositoryTagsResponseBodyResult) SetCommit(v *ListRepositoryTagsResponseBodyResultCommit) *ListRepositoryTagsResponseBodyResult {
	s.Commit = v
	return s
}

func (s *ListRepositoryTagsResponseBodyResult) SetId(v string) *ListRepositoryTagsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResult) SetMessage(v string) *ListRepositoryTagsResponseBodyResult {
	s.Message = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResult) SetName(v string) *ListRepositoryTagsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResult) Validate() error {
	if s.Commit != nil {
		if err := s.Commit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRepositoryTagsResponseBodyResultCommit struct {
	// example:
	//
	// username@example.com
	AuthorEmail *string `json:"authorEmail,omitempty" xml:"authorEmail,omitempty"`
	AuthorName  *string `json:"authorName,omitempty" xml:"authorName,omitempty"`
	// example:
	//
	// 2022-03-18 10:00:00
	AuthoredDate *string `json:"authoredDate,omitempty" xml:"authoredDate,omitempty"`
	// example:
	//
	// 2022-03-18 09:00:00
	CommittedDate *string `json:"committedDate,omitempty" xml:"committedDate,omitempty"`
	// example:
	//
	// username@example.com
	CommitterEmail *string `json:"committerEmail,omitempty" xml:"committerEmail,omitempty"`
	CommitterName  *string `json:"committerName,omitempty" xml:"committerName,omitempty"`
	// example:
	//
	// 2022-03-18 10:00:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// de02b625ba8488f92eb204bcb3773a40c1b4ddac
	Id        *string   `json:"id,omitempty" xml:"id,omitempty"`
	Message   *string   `json:"message,omitempty" xml:"message,omitempty"`
	ParentIds []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
	// example:
	//
	// de02b625
	ShortId   *string                                              `json:"shortId,omitempty" xml:"shortId,omitempty"`
	Signature *ListRepositoryTagsResponseBodyResultCommitSignature `json:"signature,omitempty" xml:"signature,omitempty" type:"Struct"`
	Title     *string                                              `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListRepositoryTagsResponseBodyResultCommit) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryTagsResponseBodyResultCommit) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsResponseBodyResultCommit) GetAuthorEmail() *string {
	return s.AuthorEmail
}

func (s *ListRepositoryTagsResponseBodyResultCommit) GetAuthorName() *string {
	return s.AuthorName
}

func (s *ListRepositoryTagsResponseBodyResultCommit) GetAuthoredDate() *string {
	return s.AuthoredDate
}

func (s *ListRepositoryTagsResponseBodyResultCommit) GetCommittedDate() *string {
	return s.CommittedDate
}

func (s *ListRepositoryTagsResponseBodyResultCommit) GetCommitterEmail() *string {
	return s.CommitterEmail
}

func (s *ListRepositoryTagsResponseBodyResultCommit) GetCommitterName() *string {
	return s.CommitterName
}

func (s *ListRepositoryTagsResponseBodyResultCommit) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListRepositoryTagsResponseBodyResultCommit) GetId() *string {
	return s.Id
}

func (s *ListRepositoryTagsResponseBodyResultCommit) GetMessage() *string {
	return s.Message
}

func (s *ListRepositoryTagsResponseBodyResultCommit) GetParentIds() []*string {
	return s.ParentIds
}

func (s *ListRepositoryTagsResponseBodyResultCommit) GetShortId() *string {
	return s.ShortId
}

func (s *ListRepositoryTagsResponseBodyResultCommit) GetSignature() *ListRepositoryTagsResponseBodyResultCommitSignature {
	return s.Signature
}

func (s *ListRepositoryTagsResponseBodyResultCommit) GetTitle() *string {
	return s.Title
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetAuthorEmail(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.AuthorEmail = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetAuthorName(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.AuthorName = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetAuthoredDate(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.AuthoredDate = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetCommittedDate(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.CommittedDate = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetCommitterEmail(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.CommitterEmail = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetCommitterName(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.CommitterName = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetCreatedAt(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetId(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.Id = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetMessage(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.Message = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetParentIds(v []*string) *ListRepositoryTagsResponseBodyResultCommit {
	s.ParentIds = v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetShortId(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.ShortId = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetSignature(v *ListRepositoryTagsResponseBodyResultCommitSignature) *ListRepositoryTagsResponseBodyResultCommit {
	s.Signature = v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetTitle(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.Title = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) Validate() error {
	if s.Signature != nil {
		if err := s.Signature.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRepositoryTagsResponseBodyResultCommitSignature struct {
	// example:
	//
	// ""
	GpgKeyId *string `json:"gpgKeyId,omitempty" xml:"gpgKeyId,omitempty"`
	// example:
	//
	// verified
	VerificationStatus *string `json:"verificationStatus,omitempty" xml:"verificationStatus,omitempty"`
}

func (s ListRepositoryTagsResponseBodyResultCommitSignature) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryTagsResponseBodyResultCommitSignature) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsResponseBodyResultCommitSignature) GetGpgKeyId() *string {
	return s.GpgKeyId
}

func (s *ListRepositoryTagsResponseBodyResultCommitSignature) GetVerificationStatus() *string {
	return s.VerificationStatus
}

func (s *ListRepositoryTagsResponseBodyResultCommitSignature) SetGpgKeyId(v string) *ListRepositoryTagsResponseBodyResultCommitSignature {
	s.GpgKeyId = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommitSignature) SetVerificationStatus(v string) *ListRepositoryTagsResponseBodyResultCommitSignature {
	s.VerificationStatus = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommitSignature) Validate() error {
	return dara.Validate(s)
}
