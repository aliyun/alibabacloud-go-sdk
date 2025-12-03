// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupByPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetGroupByPathResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetGroupByPathResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetGroupByPathResponseBody
	GetRequestId() *string
	SetResult(v *GetGroupByPathResponseBodyResult) *GetGroupByPathResponseBody
	GetResult() *GetGroupByPathResponseBodyResult
	SetSuccess(v bool) *GetGroupByPathResponseBody
	GetSuccess() *bool
}

type GetGroupByPathResponseBody struct {
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
	// 6177543A-8D54-5736-A93B-E0195A1512CB
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetGroupByPathResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetGroupByPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGroupByPathResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupByPathResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetGroupByPathResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetGroupByPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGroupByPathResponseBody) GetResult() *GetGroupByPathResponseBodyResult {
	return s.Result
}

func (s *GetGroupByPathResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGroupByPathResponseBody) SetErrorCode(v string) *GetGroupByPathResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetGroupByPathResponseBody) SetErrorMessage(v string) *GetGroupByPathResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetGroupByPathResponseBody) SetRequestId(v string) *GetGroupByPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupByPathResponseBody) SetResult(v *GetGroupByPathResponseBodyResult) *GetGroupByPathResponseBody {
	s.Result = v
	return s
}

func (s *GetGroupByPathResponseBody) SetSuccess(v bool) *GetGroupByPathResponseBody {
	s.Success = &v
	return s
}

func (s *GetGroupByPathResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGroupByPathResponseBodyResult struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl   *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 30815
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test-group
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace *string `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	// example:
	//
	// 19230
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// example:
	//
	// 26842
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// test-group
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// codeup-test-org/test-group
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	// example:
	//
	// xxx
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s GetGroupByPathResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetGroupByPathResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetGroupByPathResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetGroupByPathResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *GetGroupByPathResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *GetGroupByPathResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetGroupByPathResponseBodyResult) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *GetGroupByPathResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetGroupByPathResponseBodyResult) GetParentId() *string {
	return s.ParentId
}

func (s *GetGroupByPathResponseBodyResult) GetPath() *string {
	return s.Path
}

func (s *GetGroupByPathResponseBodyResult) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *GetGroupByPathResponseBodyResult) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *GetGroupByPathResponseBodyResult) GetWebUrl() *string {
	return s.WebUrl
}

func (s *GetGroupByPathResponseBodyResult) SetAvatarUrl(v string) *GetGroupByPathResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *GetGroupByPathResponseBodyResult) SetDescription(v string) *GetGroupByPathResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetGroupByPathResponseBodyResult) SetId(v int64) *GetGroupByPathResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetGroupByPathResponseBodyResult) SetName(v string) *GetGroupByPathResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetGroupByPathResponseBodyResult) SetNameWithNamespace(v string) *GetGroupByPathResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *GetGroupByPathResponseBodyResult) SetOwnerId(v string) *GetGroupByPathResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *GetGroupByPathResponseBodyResult) SetParentId(v string) *GetGroupByPathResponseBodyResult {
	s.ParentId = &v
	return s
}

func (s *GetGroupByPathResponseBodyResult) SetPath(v string) *GetGroupByPathResponseBodyResult {
	s.Path = &v
	return s
}

func (s *GetGroupByPathResponseBodyResult) SetPathWithNamespace(v string) *GetGroupByPathResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *GetGroupByPathResponseBodyResult) SetVisibilityLevel(v int32) *GetGroupByPathResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *GetGroupByPathResponseBodyResult) SetWebUrl(v string) *GetGroupByPathResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *GetGroupByPathResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
