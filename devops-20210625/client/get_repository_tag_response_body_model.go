// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepositoryTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetRepositoryTagResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetRepositoryTagResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetRepositoryTagResponseBody
	GetRequestId() *string
	SetResult(v *GetRepositoryTagResponseBodyResult) *GetRepositoryTagResponseBody
	GetResult() *GetRepositoryTagResponseBodyResult
	SetSuccess(v bool) *GetRepositoryTagResponseBody
	GetSuccess() *bool
}

type GetRepositoryTagResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// CE7353E3-F989-56A9-B97C-897ABBDB9A01
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetRepositoryTagResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetRepositoryTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetRepositoryTagResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetRepositoryTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRepositoryTagResponseBody) GetResult() *GetRepositoryTagResponseBodyResult {
	return s.Result
}

func (s *GetRepositoryTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRepositoryTagResponseBody) SetErrorCode(v string) *GetRepositoryTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRepositoryTagResponseBody) SetErrorMessage(v string) *GetRepositoryTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRepositoryTagResponseBody) SetRequestId(v string) *GetRepositoryTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepositoryTagResponseBody) SetResult(v *GetRepositoryTagResponseBodyResult) *GetRepositoryTagResponseBody {
	s.Result = v
	return s
}

func (s *GetRepositoryTagResponseBody) SetSuccess(v bool) *GetRepositoryTagResponseBody {
	s.Success = &v
	return s
}

func (s *GetRepositoryTagResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRepositoryTagResponseBodyResult struct {
	Commit *GetRepositoryTagResponseBodyResultCommit `json:"commit,omitempty" xml:"commit,omitempty" type:"Struct"`
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

func (s GetRepositoryTagResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagResponseBodyResult) GetCommit() *GetRepositoryTagResponseBodyResultCommit {
	return s.Commit
}

func (s *GetRepositoryTagResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *GetRepositoryTagResponseBodyResult) GetMessage() *string {
	return s.Message
}

func (s *GetRepositoryTagResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetRepositoryTagResponseBodyResult) SetCommit(v *GetRepositoryTagResponseBodyResultCommit) *GetRepositoryTagResponseBodyResult {
	s.Commit = v
	return s
}

func (s *GetRepositoryTagResponseBodyResult) SetId(v string) *GetRepositoryTagResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResult) SetMessage(v string) *GetRepositoryTagResponseBodyResult {
	s.Message = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResult) SetName(v string) *GetRepositoryTagResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResult) Validate() error {
	if s.Commit != nil {
		if err := s.Commit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRepositoryTagResponseBodyResultCommit struct {
	// example:
	//
	// username@example.com
	AuthorEmail *string `json:"authorEmail,omitempty" xml:"authorEmail,omitempty"`
	AuthorName  *string `json:"authorName,omitempty" xml:"authorName,omitempty"`
	// example:
	//
	// 2022-03-18 08:00:00
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
	// e0297d8fb0393c833a8531e7cc8832739e3cba6d
	Id        *string   `json:"id,omitempty" xml:"id,omitempty"`
	Message   *string   `json:"message,omitempty" xml:"message,omitempty"`
	ParentIds []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
	// example:
	//
	// e0297d8f
	ShortId   *string                                            `json:"shortId,omitempty" xml:"shortId,omitempty"`
	Signature *GetRepositoryTagResponseBodyResultCommitSignature `json:"signature,omitempty" xml:"signature,omitempty" type:"Struct"`
	Title     *string                                            `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetRepositoryTagResponseBodyResultCommit) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryTagResponseBodyResultCommit) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagResponseBodyResultCommit) GetAuthorEmail() *string {
	return s.AuthorEmail
}

func (s *GetRepositoryTagResponseBodyResultCommit) GetAuthorName() *string {
	return s.AuthorName
}

func (s *GetRepositoryTagResponseBodyResultCommit) GetAuthoredDate() *string {
	return s.AuthoredDate
}

func (s *GetRepositoryTagResponseBodyResultCommit) GetCommittedDate() *string {
	return s.CommittedDate
}

func (s *GetRepositoryTagResponseBodyResultCommit) GetCommitterEmail() *string {
	return s.CommitterEmail
}

func (s *GetRepositoryTagResponseBodyResultCommit) GetCommitterName() *string {
	return s.CommitterName
}

func (s *GetRepositoryTagResponseBodyResultCommit) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetRepositoryTagResponseBodyResultCommit) GetId() *string {
	return s.Id
}

func (s *GetRepositoryTagResponseBodyResultCommit) GetMessage() *string {
	return s.Message
}

func (s *GetRepositoryTagResponseBodyResultCommit) GetParentIds() []*string {
	return s.ParentIds
}

func (s *GetRepositoryTagResponseBodyResultCommit) GetShortId() *string {
	return s.ShortId
}

func (s *GetRepositoryTagResponseBodyResultCommit) GetSignature() *GetRepositoryTagResponseBodyResultCommitSignature {
	return s.Signature
}

func (s *GetRepositoryTagResponseBodyResultCommit) GetTitle() *string {
	return s.Title
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetAuthorEmail(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.AuthorEmail = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetAuthorName(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.AuthorName = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetAuthoredDate(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.AuthoredDate = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetCommittedDate(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.CommittedDate = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetCommitterEmail(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.CommitterEmail = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetCommitterName(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.CommitterName = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetCreatedAt(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.CreatedAt = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetId(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.Id = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetMessage(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.Message = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetParentIds(v []*string) *GetRepositoryTagResponseBodyResultCommit {
	s.ParentIds = v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetShortId(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.ShortId = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetSignature(v *GetRepositoryTagResponseBodyResultCommitSignature) *GetRepositoryTagResponseBodyResultCommit {
	s.Signature = v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetTitle(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.Title = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) Validate() error {
	if s.Signature != nil {
		if err := s.Signature.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRepositoryTagResponseBodyResultCommitSignature struct {
	// example:
	//
	// ""
	GpgKeyId *string `json:"gpgKeyId,omitempty" xml:"gpgKeyId,omitempty"`
	// example:
	//
	// verified
	VerificationStatus *string `json:"verificationStatus,omitempty" xml:"verificationStatus,omitempty"`
}

func (s GetRepositoryTagResponseBodyResultCommitSignature) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryTagResponseBodyResultCommitSignature) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagResponseBodyResultCommitSignature) GetGpgKeyId() *string {
	return s.GpgKeyId
}

func (s *GetRepositoryTagResponseBodyResultCommitSignature) GetVerificationStatus() *string {
	return s.VerificationStatus
}

func (s *GetRepositoryTagResponseBodyResultCommitSignature) SetGpgKeyId(v string) *GetRepositoryTagResponseBodyResultCommitSignature {
	s.GpgKeyId = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommitSignature) SetVerificationStatus(v string) *GetRepositoryTagResponseBodyResultCommitSignature {
	s.VerificationStatus = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommitSignature) Validate() error {
	return dara.Validate(s)
}
