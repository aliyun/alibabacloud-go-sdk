// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateTagResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateTagResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateTagResponseBody
	GetRequestId() *string
	SetResult(v *CreateTagResponseBodyResult) *CreateTagResponseBody
	GetResult() *CreateTagResponseBodyResult
	SetSuccess(v bool) *CreateTagResponseBody
	GetSuccess() *bool
}

type CreateTagResponseBody struct {
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
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateTagResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateTagResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTagResponseBody) GetResult() *CreateTagResponseBodyResult {
	return s.Result
}

func (s *CreateTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTagResponseBody) SetErrorCode(v string) *CreateTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTagResponseBody) SetErrorMessage(v string) *CreateTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTagResponseBody) SetRequestId(v string) *CreateTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTagResponseBody) SetResult(v *CreateTagResponseBodyResult) *CreateTagResponseBody {
	s.Result = v
	return s
}

func (s *CreateTagResponseBody) SetSuccess(v bool) *CreateTagResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTagResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTagResponseBodyResult struct {
	Commit *CreateTagResponseBodyResultCommit `json:"commit,omitempty" xml:"commit,omitempty" type:"Struct"`
	// example:
	//
	// 0e3b6aa5eab2b086b59fde74766b28d4e5faab0d
	Id      *string `json:"id,omitempty" xml:"id,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// v1.0
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateTagResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBodyResult) GetCommit() *CreateTagResponseBodyResultCommit {
	return s.Commit
}

func (s *CreateTagResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *CreateTagResponseBodyResult) GetMessage() *string {
	return s.Message
}

func (s *CreateTagResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateTagResponseBodyResult) SetCommit(v *CreateTagResponseBodyResultCommit) *CreateTagResponseBodyResult {
	s.Commit = v
	return s
}

func (s *CreateTagResponseBodyResult) SetId(v string) *CreateTagResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateTagResponseBodyResult) SetMessage(v string) *CreateTagResponseBodyResult {
	s.Message = &v
	return s
}

func (s *CreateTagResponseBodyResult) SetName(v string) *CreateTagResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateTagResponseBodyResult) Validate() error {
	if s.Commit != nil {
		if err := s.Commit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTagResponseBodyResultCommit struct {
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
	// e0297d8fb0393c833a8531e7cc8832739e3cba6d
	Id        *string   `json:"id,omitempty" xml:"id,omitempty"`
	Message   *string   `json:"message,omitempty" xml:"message,omitempty"`
	ParentIds []*string `json:"parentIds,omitempty" xml:"parentIds,omitempty" type:"Repeated"`
	// example:
	//
	// e0297d8f
	ShortId *string `json:"shortId,omitempty" xml:"shortId,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateTagResponseBodyResultCommit) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResponseBodyResultCommit) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBodyResultCommit) GetAuthorEmail() *string {
	return s.AuthorEmail
}

func (s *CreateTagResponseBodyResultCommit) GetAuthorName() *string {
	return s.AuthorName
}

func (s *CreateTagResponseBodyResultCommit) GetAuthoredDate() *string {
	return s.AuthoredDate
}

func (s *CreateTagResponseBodyResultCommit) GetCommittedDate() *string {
	return s.CommittedDate
}

func (s *CreateTagResponseBodyResultCommit) GetCommitterEmail() *string {
	return s.CommitterEmail
}

func (s *CreateTagResponseBodyResultCommit) GetCommitterName() *string {
	return s.CommitterName
}

func (s *CreateTagResponseBodyResultCommit) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateTagResponseBodyResultCommit) GetId() *string {
	return s.Id
}

func (s *CreateTagResponseBodyResultCommit) GetMessage() *string {
	return s.Message
}

func (s *CreateTagResponseBodyResultCommit) GetParentIds() []*string {
	return s.ParentIds
}

func (s *CreateTagResponseBodyResultCommit) GetShortId() *string {
	return s.ShortId
}

func (s *CreateTagResponseBodyResultCommit) GetTitle() *string {
	return s.Title
}

func (s *CreateTagResponseBodyResultCommit) SetAuthorEmail(v string) *CreateTagResponseBodyResultCommit {
	s.AuthorEmail = &v
	return s
}

func (s *CreateTagResponseBodyResultCommit) SetAuthorName(v string) *CreateTagResponseBodyResultCommit {
	s.AuthorName = &v
	return s
}

func (s *CreateTagResponseBodyResultCommit) SetAuthoredDate(v string) *CreateTagResponseBodyResultCommit {
	s.AuthoredDate = &v
	return s
}

func (s *CreateTagResponseBodyResultCommit) SetCommittedDate(v string) *CreateTagResponseBodyResultCommit {
	s.CommittedDate = &v
	return s
}

func (s *CreateTagResponseBodyResultCommit) SetCommitterEmail(v string) *CreateTagResponseBodyResultCommit {
	s.CommitterEmail = &v
	return s
}

func (s *CreateTagResponseBodyResultCommit) SetCommitterName(v string) *CreateTagResponseBodyResultCommit {
	s.CommitterName = &v
	return s
}

func (s *CreateTagResponseBodyResultCommit) SetCreatedAt(v string) *CreateTagResponseBodyResultCommit {
	s.CreatedAt = &v
	return s
}

func (s *CreateTagResponseBodyResultCommit) SetId(v string) *CreateTagResponseBodyResultCommit {
	s.Id = &v
	return s
}

func (s *CreateTagResponseBodyResultCommit) SetMessage(v string) *CreateTagResponseBodyResultCommit {
	s.Message = &v
	return s
}

func (s *CreateTagResponseBodyResultCommit) SetParentIds(v []*string) *CreateTagResponseBodyResultCommit {
	s.ParentIds = v
	return s
}

func (s *CreateTagResponseBodyResultCommit) SetShortId(v string) *CreateTagResponseBodyResultCommit {
	s.ShortId = &v
	return s
}

func (s *CreateTagResponseBodyResultCommit) SetTitle(v string) *CreateTagResponseBodyResultCommit {
	s.Title = &v
	return s
}

func (s *CreateTagResponseBodyResultCommit) Validate() error {
	return dara.Validate(s)
}
