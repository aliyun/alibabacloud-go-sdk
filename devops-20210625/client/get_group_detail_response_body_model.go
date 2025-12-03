// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetGroupDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetGroupDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetGroupDetailResponseBody
	GetRequestId() *string
	SetResult(v *GetGroupDetailResponseBodyResult) *GetGroupDetailResponseBody
	GetResult() *GetGroupDetailResponseBodyResult
	SetSuccess(v bool) *GetGroupDetailResponseBody
	GetSuccess() *bool
}

type GetGroupDetailResponseBody struct {
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
	// 4CA06B0B-5867-5DE6-A0FA-9F39C97B524C
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetGroupDetailResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetGroupDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetGroupDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetGroupDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGroupDetailResponseBody) GetResult() *GetGroupDetailResponseBodyResult {
	return s.Result
}

func (s *GetGroupDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGroupDetailResponseBody) SetErrorCode(v string) *GetGroupDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetGroupDetailResponseBody) SetErrorMessage(v string) *GetGroupDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetGroupDetailResponseBody) SetRequestId(v string) *GetGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupDetailResponseBody) SetResult(v *GetGroupDetailResponseBodyResult) *GetGroupDetailResponseBody {
	s.Result = v
	return s
}

func (s *GetGroupDetailResponseBody) SetSuccess(v bool) *GetGroupDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetGroupDetailResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGroupDetailResponseBodyResult struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl   *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 36612
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test-group-detail
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace *string `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	// example:
	//
	// 19230
	OwnerId *int64 `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// example:
	//
	// 26842
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// test-group-detail
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// codeup-test-org/test-group-detail
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	// example:
	//
	// Group
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 10
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	// example:
	//
	// ""
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s GetGroupDetailResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetGroupDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetGroupDetailResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetGroupDetailResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *GetGroupDetailResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *GetGroupDetailResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetGroupDetailResponseBodyResult) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *GetGroupDetailResponseBodyResult) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetGroupDetailResponseBodyResult) GetParentId() *int64 {
	return s.ParentId
}

func (s *GetGroupDetailResponseBodyResult) GetPath() *string {
	return s.Path
}

func (s *GetGroupDetailResponseBodyResult) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *GetGroupDetailResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetGroupDetailResponseBodyResult) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *GetGroupDetailResponseBodyResult) GetWebUrl() *string {
	return s.WebUrl
}

func (s *GetGroupDetailResponseBodyResult) SetAvatarUrl(v string) *GetGroupDetailResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetDescription(v string) *GetGroupDetailResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetId(v int64) *GetGroupDetailResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetName(v string) *GetGroupDetailResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetNameWithNamespace(v string) *GetGroupDetailResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetOwnerId(v int64) *GetGroupDetailResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetParentId(v int64) *GetGroupDetailResponseBodyResult {
	s.ParentId = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetPath(v string) *GetGroupDetailResponseBodyResult {
	s.Path = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetPathWithNamespace(v string) *GetGroupDetailResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetType(v string) *GetGroupDetailResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetVisibilityLevel(v int32) *GetGroupDetailResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetWebUrl(v string) *GetGroupDetailResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
