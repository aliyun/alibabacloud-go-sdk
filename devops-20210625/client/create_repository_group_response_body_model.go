// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepositoryGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateRepositoryGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateRepositoryGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateRepositoryGroupResponseBody
	GetRequestId() *string
	SetResult(v *CreateRepositoryGroupResponseBodyResult) *CreateRepositoryGroupResponseBody
	GetResult() *CreateRepositoryGroupResponseBodyResult
	SetSuccess(v bool) *CreateRepositoryGroupResponseBody
	GetSuccess() *bool
}

type CreateRepositoryGroupResponseBody struct {
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
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateRepositoryGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateRepositoryGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRepositoryGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepositoryGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateRepositoryGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateRepositoryGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRepositoryGroupResponseBody) GetResult() *CreateRepositoryGroupResponseBodyResult {
	return s.Result
}

func (s *CreateRepositoryGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRepositoryGroupResponseBody) SetErrorCode(v string) *CreateRepositoryGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateRepositoryGroupResponseBody) SetErrorMessage(v string) *CreateRepositoryGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRepositoryGroupResponseBody) SetRequestId(v string) *CreateRepositoryGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepositoryGroupResponseBody) SetResult(v *CreateRepositoryGroupResponseBodyResult) *CreateRepositoryGroupResponseBody {
	s.Result = v
	return s
}

func (s *CreateRepositoryGroupResponseBody) SetSuccess(v bool) *CreateRepositoryGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRepositoryGroupResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRepositoryGroupResponseBodyResult struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl   *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 18685
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test-create-group
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
	// test-create-group
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// codeup-test-org/test-create-group
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

func (s CreateRepositoryGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateRepositoryGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRepositoryGroupResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *CreateRepositoryGroupResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *CreateRepositoryGroupResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *CreateRepositoryGroupResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateRepositoryGroupResponseBodyResult) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *CreateRepositoryGroupResponseBodyResult) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRepositoryGroupResponseBodyResult) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateRepositoryGroupResponseBodyResult) GetPath() *string {
	return s.Path
}

func (s *CreateRepositoryGroupResponseBodyResult) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *CreateRepositoryGroupResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *CreateRepositoryGroupResponseBodyResult) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *CreateRepositoryGroupResponseBodyResult) GetWebUrl() *string {
	return s.WebUrl
}

func (s *CreateRepositoryGroupResponseBodyResult) SetAvatarUrl(v string) *CreateRepositoryGroupResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetDescription(v string) *CreateRepositoryGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetId(v int64) *CreateRepositoryGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetName(v string) *CreateRepositoryGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetNameWithNamespace(v string) *CreateRepositoryGroupResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetOwnerId(v int64) *CreateRepositoryGroupResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetParentId(v int64) *CreateRepositoryGroupResponseBodyResult {
	s.ParentId = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetPath(v string) *CreateRepositoryGroupResponseBodyResult {
	s.Path = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetPathWithNamespace(v string) *CreateRepositoryGroupResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetType(v string) *CreateRepositoryGroupResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetVisibilityLevel(v int32) *CreateRepositoryGroupResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetWebUrl(v string) *CreateRepositoryGroupResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
