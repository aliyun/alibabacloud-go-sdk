// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryMemberWithInheritedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListRepositoryMemberWithInheritedResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListRepositoryMemberWithInheritedResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListRepositoryMemberWithInheritedResponseBody
	GetRequestId() *string
	SetResult(v []*ListRepositoryMemberWithInheritedResponseBodyResult) *ListRepositoryMemberWithInheritedResponseBody
	GetResult() []*ListRepositoryMemberWithInheritedResponseBodyResult
	SetSuccess(v bool) *ListRepositoryMemberWithInheritedResponseBody
	GetSuccess() *bool
}

type ListRepositoryMemberWithInheritedResponseBody struct {
	ErrorCode    *string                                                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                                `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       []*ListRepositoryMemberWithInheritedResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success      *bool                                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListRepositoryMemberWithInheritedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListRepositoryMemberWithInheritedResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListRepositoryMemberWithInheritedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepositoryMemberWithInheritedResponseBody) GetResult() []*ListRepositoryMemberWithInheritedResponseBodyResult {
	return s.Result
}

func (s *ListRepositoryMemberWithInheritedResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetErrorCode(v string) *ListRepositoryMemberWithInheritedResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetErrorMessage(v string) *ListRepositoryMemberWithInheritedResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetRequestId(v string) *ListRepositoryMemberWithInheritedResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetResult(v []*ListRepositoryMemberWithInheritedResponseBodyResult) *ListRepositoryMemberWithInheritedResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetSuccess(v bool) *ListRepositoryMemberWithInheritedResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBody) Validate() error {
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

type ListRepositoryMemberWithInheritedResponseBodyResult struct {
	AccessLevel *int32 `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string                                                       `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Email     *string                                                       `json:"email,omitempty" xml:"email,omitempty"`
	Id        *int64                                                        `json:"id,omitempty" xml:"id,omitempty"`
	Inherited *ListRepositoryMemberWithInheritedResponseBodyResultInherited `json:"inherited,omitempty" xml:"inherited,omitempty" type:"Struct"`
	// example:
	//
	// codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// yunxiao
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListRepositoryMemberWithInheritedResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) GetEmail() *string {
	return s.Email
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) GetInherited() *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	return s.Inherited
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) GetState() *string {
	return s.State
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) GetUsername() *string {
	return s.Username
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetAccessLevel(v int32) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetAvatarUrl(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetEmail(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Email = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetId(v int64) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetInherited(v *ListRepositoryMemberWithInheritedResponseBodyResultInherited) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Inherited = v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetName(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetState(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetUsername(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Username = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) Validate() error {
	if s.Inherited != nil {
		if err := s.Inherited.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRepositoryMemberWithInheritedResponseBodyResultInherited struct {
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// codeup
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace *string `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	// example:
	//
	// codeup
	Path              *string `json:"path,omitempty" xml:"path,omitempty"`
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	Type              *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *string `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
}

func (s ListRepositoryMemberWithInheritedResponseBodyResultInherited) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedResponseBodyResultInherited) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) GetId() *int64 {
	return s.Id
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) GetName() *string {
	return s.Name
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) GetPath() *string {
	return s.Path
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) GetType() *string {
	return s.Type
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) GetVisibilityLevel() *string {
	return s.VisibilityLevel
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetId(v int64) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.Id = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetName(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.Name = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetNameWithNamespace(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.NameWithNamespace = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetPath(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.Path = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetPathWithNamespace(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.PathWithNamespace = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetType(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.Type = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetVisibilityLevel(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.VisibilityLevel = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) Validate() error {
	return dara.Validate(s)
}
