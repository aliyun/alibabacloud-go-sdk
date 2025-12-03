// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileLastCommitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetFileLastCommitResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetFileLastCommitResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetFileLastCommitResponseBody
	GetRequestId() *string
	SetResult(v *GetFileLastCommitResponseBodyResult) *GetFileLastCommitResponseBody
	GetResult() *GetFileLastCommitResponseBodyResult
	SetSuccess(v bool) *GetFileLastCommitResponseBody
	GetSuccess() *bool
}

type GetFileLastCommitResponseBody struct {
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
	// F7B85D1B-D1C2-140F-A039-341859F130B9
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetFileLastCommitResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetFileLastCommitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileLastCommitResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetFileLastCommitResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetFileLastCommitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileLastCommitResponseBody) GetResult() *GetFileLastCommitResponseBodyResult {
	return s.Result
}

func (s *GetFileLastCommitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFileLastCommitResponseBody) SetErrorCode(v string) *GetFileLastCommitResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFileLastCommitResponseBody) SetErrorMessage(v string) *GetFileLastCommitResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFileLastCommitResponseBody) SetRequestId(v string) *GetFileLastCommitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileLastCommitResponseBody) SetResult(v *GetFileLastCommitResponseBodyResult) *GetFileLastCommitResponseBody {
	s.Result = v
	return s
}

func (s *GetFileLastCommitResponseBody) SetSuccess(v bool) *GetFileLastCommitResponseBody {
	s.Success = &v
	return s
}

func (s *GetFileLastCommitResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileLastCommitResponseBodyResult struct {
	// example:
	//
	// 2022-08-08 18:09:09
	AuthorDate *string `json:"authorDate,omitempty" xml:"authorDate,omitempty"`
	// example:
	//
	// username@example.com
	AuthorEmail *string `json:"authorEmail,omitempty" xml:"authorEmail,omitempty"`
	// example:
	//
	// test-codeup
	AuthorName *string `json:"authorName,omitempty" xml:"authorName,omitempty"`
	// example:
	//
	// 2022-03-18 15:00:02
	CommittedDate *string `json:"committedDate,omitempty" xml:"committedDate,omitempty"`
	// example:
	//
	// username@example.com
	CommitterEmail *string `json:"committerEmail,omitempty" xml:"committerEmail,omitempty"`
	// example:
	//
	// committer-codeup
	CommitterName *string `json:"committerName,omitempty" xml:"committerName,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// ff4fb5ac6d1f44f452654336d2dba468ae6c8d04
	Id        *string   `json:"id,omitempty" xml:"id,omitempty"`
	Message   *string   `json:"message,omitempty" xml:"message,omitempty"`
	ParentIds []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
	// example:
	//
	// ff4fb5ac
	ShortId   *string                                       `json:"shortId,omitempty" xml:"shortId,omitempty"`
	Signature *GetFileLastCommitResponseBodyResultSignature `json:"signature,omitempty" xml:"signature,omitempty" type:"Struct"`
	Title     *string                                       `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetFileLastCommitResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetFileLastCommitResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitResponseBodyResult) GetAuthorDate() *string {
	return s.AuthorDate
}

func (s *GetFileLastCommitResponseBodyResult) GetAuthorEmail() *string {
	return s.AuthorEmail
}

func (s *GetFileLastCommitResponseBodyResult) GetAuthorName() *string {
	return s.AuthorName
}

func (s *GetFileLastCommitResponseBodyResult) GetCommittedDate() *string {
	return s.CommittedDate
}

func (s *GetFileLastCommitResponseBodyResult) GetCommitterEmail() *string {
	return s.CommitterEmail
}

func (s *GetFileLastCommitResponseBodyResult) GetCommitterName() *string {
	return s.CommitterName
}

func (s *GetFileLastCommitResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetFileLastCommitResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *GetFileLastCommitResponseBodyResult) GetMessage() *string {
	return s.Message
}

func (s *GetFileLastCommitResponseBodyResult) GetParentIds() []*string {
	return s.ParentIds
}

func (s *GetFileLastCommitResponseBodyResult) GetShortId() *string {
	return s.ShortId
}

func (s *GetFileLastCommitResponseBodyResult) GetSignature() *GetFileLastCommitResponseBodyResultSignature {
	return s.Signature
}

func (s *GetFileLastCommitResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *GetFileLastCommitResponseBodyResult) SetAuthorDate(v string) *GetFileLastCommitResponseBodyResult {
	s.AuthorDate = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetAuthorEmail(v string) *GetFileLastCommitResponseBodyResult {
	s.AuthorEmail = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetAuthorName(v string) *GetFileLastCommitResponseBodyResult {
	s.AuthorName = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetCommittedDate(v string) *GetFileLastCommitResponseBodyResult {
	s.CommittedDate = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetCommitterEmail(v string) *GetFileLastCommitResponseBodyResult {
	s.CommitterEmail = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetCommitterName(v string) *GetFileLastCommitResponseBodyResult {
	s.CommitterName = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetCreatedAt(v string) *GetFileLastCommitResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetId(v string) *GetFileLastCommitResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetMessage(v string) *GetFileLastCommitResponseBodyResult {
	s.Message = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetParentIds(v []*string) *GetFileLastCommitResponseBodyResult {
	s.ParentIds = v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetShortId(v string) *GetFileLastCommitResponseBodyResult {
	s.ShortId = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetSignature(v *GetFileLastCommitResponseBodyResultSignature) *GetFileLastCommitResponseBodyResult {
	s.Signature = v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetTitle(v string) *GetFileLastCommitResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) Validate() error {
	if s.Signature != nil {
		if err := s.Signature.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileLastCommitResponseBodyResultSignature struct {
	// GPG密钥ID
	//
	// example:
	//
	// 34d2c47c7ce46a5c4639c5ffe208
	GpgKeyId *string `json:"gpgKeyId,omitempty" xml:"gpgKeyId,omitempty"`
	// example:
	//
	// verified
	VerificationStatus *string `json:"verificationStatus,omitempty" xml:"verificationStatus,omitempty"`
}

func (s GetFileLastCommitResponseBodyResultSignature) String() string {
	return dara.Prettify(s)
}

func (s GetFileLastCommitResponseBodyResultSignature) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitResponseBodyResultSignature) GetGpgKeyId() *string {
	return s.GpgKeyId
}

func (s *GetFileLastCommitResponseBodyResultSignature) GetVerificationStatus() *string {
	return s.VerificationStatus
}

func (s *GetFileLastCommitResponseBodyResultSignature) SetGpgKeyId(v string) *GetFileLastCommitResponseBodyResultSignature {
	s.GpgKeyId = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResultSignature) SetVerificationStatus(v string) *GetFileLastCommitResponseBodyResultSignature {
	s.VerificationStatus = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResultSignature) Validate() error {
	return dara.Validate(s)
}
