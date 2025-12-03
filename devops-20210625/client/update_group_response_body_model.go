// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateGroupResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateGroupResponseBody
	GetRequestId() *string
	SetResult(v *UpdateGroupResponseBodyResult) *UpdateGroupResponseBody
	GetResult() *UpdateGroupResponseBodyResult
	SetSuccess(v bool) *UpdateGroupResponseBody
	GetSuccess() *bool
}

type UpdateGroupResponseBody struct {
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// InvalidParam.NoPermission
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGroupResponseBody) GetResult() *UpdateGroupResponseBodyResult {
	return s.Result
}

func (s *UpdateGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGroupResponseBody) SetErrorCode(v string) *UpdateGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateGroupResponseBody) SetErrorMessage(v string) *UpdateGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateGroupResponseBody) SetRequestId(v string) *UpdateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupResponseBody) SetResult(v *UpdateGroupResponseBodyResult) *UpdateGroupResponseBody {
	s.Result = v
	return s
}

func (s *UpdateGroupResponseBody) SetSuccess(v bool) *UpdateGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGroupResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGroupResponseBodyResult struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c4ef67f1bea827c4/w/100/h/100
	AvatarUrl   *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 30815
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// codeup
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
	// test-codeup
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// codeup-test-org/test-codeup
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

func (s UpdateGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *UpdateGroupResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *UpdateGroupResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *UpdateGroupResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateGroupResponseBodyResult) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *UpdateGroupResponseBodyResult) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateGroupResponseBodyResult) GetParentId() *int64 {
	return s.ParentId
}

func (s *UpdateGroupResponseBodyResult) GetPath() *string {
	return s.Path
}

func (s *UpdateGroupResponseBodyResult) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *UpdateGroupResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *UpdateGroupResponseBodyResult) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *UpdateGroupResponseBodyResult) GetWebUrl() *string {
	return s.WebUrl
}

func (s *UpdateGroupResponseBodyResult) SetAvatarUrl(v string) *UpdateGroupResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateGroupResponseBodyResult) SetDescription(v string) *UpdateGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *UpdateGroupResponseBodyResult) SetId(v int64) *UpdateGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateGroupResponseBodyResult) SetName(v string) *UpdateGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateGroupResponseBodyResult) SetNameWithNamespace(v string) *UpdateGroupResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *UpdateGroupResponseBodyResult) SetOwnerId(v int64) *UpdateGroupResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *UpdateGroupResponseBodyResult) SetParentId(v int64) *UpdateGroupResponseBodyResult {
	s.ParentId = &v
	return s
}

func (s *UpdateGroupResponseBodyResult) SetPath(v string) *UpdateGroupResponseBodyResult {
	s.Path = &v
	return s
}

func (s *UpdateGroupResponseBodyResult) SetPathWithNamespace(v string) *UpdateGroupResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *UpdateGroupResponseBodyResult) SetType(v string) *UpdateGroupResponseBodyResult {
	s.Type = &v
	return s
}

func (s *UpdateGroupResponseBodyResult) SetVisibilityLevel(v int32) *UpdateGroupResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *UpdateGroupResponseBodyResult) SetWebUrl(v string) *UpdateGroupResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *UpdateGroupResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
