// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepositoryMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteRepositoryMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteRepositoryMemberResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteRepositoryMemberResponseBody
	GetRequestId() *string
	SetResult(v *DeleteRepositoryMemberResponseBodyResult) *DeleteRepositoryMemberResponseBody
	GetResult() *DeleteRepositoryMemberResponseBodyResult
	SetSuccess(v bool) *DeleteRepositoryMemberResponseBody
	GetSuccess() *bool
}

type DeleteRepositoryMemberResponseBody struct {
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
	// 0E1BCF81-51E6-59D4-8D55-FF945111127A
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DeleteRepositoryMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteRepositoryMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteRepositoryMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteRepositoryMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRepositoryMemberResponseBody) GetResult() *DeleteRepositoryMemberResponseBodyResult {
	return s.Result
}

func (s *DeleteRepositoryMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRepositoryMemberResponseBody) SetErrorCode(v string) *DeleteRepositoryMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBody) SetErrorMessage(v string) *DeleteRepositoryMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBody) SetRequestId(v string) *DeleteRepositoryMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBody) SetResult(v *DeleteRepositoryMemberResponseBodyResult) *DeleteRepositoryMemberResponseBody {
	s.Result = v
	return s
}

func (s *DeleteRepositoryMemberResponseBody) SetSuccess(v bool) *DeleteRepositoryMemberResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteRepositoryMemberResponseBodyResult struct {
	// example:
	//
	// 30
	AccessLevel *int32 `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	// example:
	//
	// 2022-03-12 12:00:00
	CreateAt *string `json:"createAt,omitempty" xml:"createAt,omitempty"`
	// example:
	//
	// 30815
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2080398
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// Project
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// example:
	//
	// 2022-03-12 12:00:00
	UpdateAt *string `json:"updateAt,omitempty" xml:"updateAt,omitempty"`
	// example:
	//
	// 19280
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteRepositoryMemberResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryMemberResponseBodyResult) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *DeleteRepositoryMemberResponseBodyResult) GetCreateAt() *string {
	return s.CreateAt
}

func (s *DeleteRepositoryMemberResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *DeleteRepositoryMemberResponseBodyResult) GetSourceId() *int64 {
	return s.SourceId
}

func (s *DeleteRepositoryMemberResponseBodyResult) GetSourceType() *string {
	return s.SourceType
}

func (s *DeleteRepositoryMemberResponseBodyResult) GetUpdateAt() *string {
	return s.UpdateAt
}

func (s *DeleteRepositoryMemberResponseBodyResult) GetUserId() *int64 {
	return s.UserId
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetAccessLevel(v int32) *DeleteRepositoryMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetCreateAt(v string) *DeleteRepositoryMemberResponseBodyResult {
	s.CreateAt = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetId(v int64) *DeleteRepositoryMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetSourceId(v int64) *DeleteRepositoryMemberResponseBodyResult {
	s.SourceId = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetSourceType(v string) *DeleteRepositoryMemberResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetUpdateAt(v string) *DeleteRepositoryMemberResponseBodyResult {
	s.UpdateAt = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetUserId(v int64) *DeleteRepositoryMemberResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
