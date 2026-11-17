// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateShareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateShareResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateShareResponseBody
	GetMessage() *string
	SetOutputItemId(v string) *UpdateShareResponseBody
	GetOutputItemId() *string
	SetRequestId(v string) *UpdateShareResponseBody
	GetRequestId() *string
	SetSessionId(v string) *UpdateShareResponseBody
	GetSessionId() *string
	SetShareAccessType(v string) *UpdateShareResponseBody
	GetShareAccessType() *string
	SetShareEnabled(v bool) *UpdateShareResponseBody
	GetShareEnabled() *bool
	SetShareEnabledAt(v string) *UpdateShareResponseBody
	GetShareEnabledAt() *string
	SetSharePath(v string) *UpdateShareResponseBody
	GetSharePath() *string
	SetShareToken(v string) *UpdateShareResponseBody
	GetShareToken() *string
	SetShareUrl(v string) *UpdateShareResponseBody
	GetShareUrl() *string
	SetShareUserGroupIds(v []*string) *UpdateShareResponseBody
	GetShareUserGroupIds() []*string
	SetShareUserIds(v []*int64) *UpdateShareResponseBody
	GetShareUserIds() []*int64
	SetTargetType(v string) *UpdateShareResponseBody
	GetTargetType() *string
}

type UpdateShareResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The message.
	//
	// example:
	//
	// Invalid request parameters
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The output item ID.
	//
	// example:
	//
	// exampleOutputItemId
	OutputItemId *string `json:"outputItemId,omitempty" xml:"outputItemId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The unique identifier of the function session.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The current sharing scope. Valid values: PUBLIC and SPECIFIED.
	//
	// example:
	//
	// SPECIFIED
	ShareAccessType *string `json:"shareAccessType,omitempty" xml:"shareAccessType,omitempty"`
	// Specifies whether sharing is enabled.
	//
	// example:
	//
	// true
	ShareEnabled *bool `json:"shareEnabled,omitempty" xml:"shareEnabled,omitempty"`
	// The time when sharing was most recently enabled, in ISO 8601 format.
	//
	// example:
	//
	// 2026-11-16T13:00:00+00:00
	ShareEnabledAt *string `json:"shareEnabledAt,omitempty" xml:"shareEnabledAt,omitempty"`
	// The relative path for in-site sharing. This parameter is empty when sharing is disabled.
	//
	// example:
	//
	// /10001/share/output/token
	SharePath *string `json:"sharePath,omitempty" xml:"sharePath,omitempty"`
	// The sharing token, which is returned when sharing is enabled. You can use this token to access the public sharing preview API.
	//
	// example:
	//
	// exampleShareToken
	ShareToken *string `json:"shareToken,omitempty" xml:"shareToken,omitempty"`
	// The absolute URL of the current sharing link. This parameter is empty when sharing is disabled.
	//
	// example:
	//
	// https://example.com/10001/share/output/token
	ShareUrl *string `json:"shareUrl,omitempty" xml:"shareUrl,omitempty"`
	// The list of currently specified user group IDs.
	//
	// example:
	//
	// []
	ShareUserGroupIds []*string `json:"shareUserGroupIds,omitempty" xml:"shareUserGroupIds,omitempty" type:"Repeated"`
	// The list of currently specified user IDs.
	//
	// example:
	//
	// [10001]
	ShareUserIds []*int64 `json:"shareUserIds,omitempty" xml:"shareUserIds,omitempty" type:"Repeated"`
	// The type of the current permission. Valid values:
	//
	// - period: Period type.
	//
	// - objective: Objective type.
	//
	// - kr: Key result type.
	//
	// example:
	//
	// outputItem
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s UpdateShareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateShareResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateShareResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateShareResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateShareResponseBody) GetOutputItemId() *string {
	return s.OutputItemId
}

func (s *UpdateShareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateShareResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *UpdateShareResponseBody) GetShareAccessType() *string {
	return s.ShareAccessType
}

func (s *UpdateShareResponseBody) GetShareEnabled() *bool {
	return s.ShareEnabled
}

func (s *UpdateShareResponseBody) GetShareEnabledAt() *string {
	return s.ShareEnabledAt
}

func (s *UpdateShareResponseBody) GetSharePath() *string {
	return s.SharePath
}

func (s *UpdateShareResponseBody) GetShareToken() *string {
	return s.ShareToken
}

func (s *UpdateShareResponseBody) GetShareUrl() *string {
	return s.ShareUrl
}

func (s *UpdateShareResponseBody) GetShareUserGroupIds() []*string {
	return s.ShareUserGroupIds
}

func (s *UpdateShareResponseBody) GetShareUserIds() []*int64 {
	return s.ShareUserIds
}

func (s *UpdateShareResponseBody) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateShareResponseBody) SetCode(v string) *UpdateShareResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateShareResponseBody) SetMessage(v string) *UpdateShareResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateShareResponseBody) SetOutputItemId(v string) *UpdateShareResponseBody {
	s.OutputItemId = &v
	return s
}

func (s *UpdateShareResponseBody) SetRequestId(v string) *UpdateShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateShareResponseBody) SetSessionId(v string) *UpdateShareResponseBody {
	s.SessionId = &v
	return s
}

func (s *UpdateShareResponseBody) SetShareAccessType(v string) *UpdateShareResponseBody {
	s.ShareAccessType = &v
	return s
}

func (s *UpdateShareResponseBody) SetShareEnabled(v bool) *UpdateShareResponseBody {
	s.ShareEnabled = &v
	return s
}

func (s *UpdateShareResponseBody) SetShareEnabledAt(v string) *UpdateShareResponseBody {
	s.ShareEnabledAt = &v
	return s
}

func (s *UpdateShareResponseBody) SetSharePath(v string) *UpdateShareResponseBody {
	s.SharePath = &v
	return s
}

func (s *UpdateShareResponseBody) SetShareToken(v string) *UpdateShareResponseBody {
	s.ShareToken = &v
	return s
}

func (s *UpdateShareResponseBody) SetShareUrl(v string) *UpdateShareResponseBody {
	s.ShareUrl = &v
	return s
}

func (s *UpdateShareResponseBody) SetShareUserGroupIds(v []*string) *UpdateShareResponseBody {
	s.ShareUserGroupIds = v
	return s
}

func (s *UpdateShareResponseBody) SetShareUserIds(v []*int64) *UpdateShareResponseBody {
	s.ShareUserIds = v
	return s
}

func (s *UpdateShareResponseBody) SetTargetType(v string) *UpdateShareResponseBody {
	s.TargetType = &v
	return s
}

func (s *UpdateShareResponseBody) Validate() error {
	return dara.Validate(s)
}
