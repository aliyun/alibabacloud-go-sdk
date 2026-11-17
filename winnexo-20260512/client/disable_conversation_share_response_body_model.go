// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableConversationShareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableConversationShareResponseBody
	GetCode() *string
	SetMessage(v string) *DisableConversationShareResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableConversationShareResponseBody
	GetRequestId() *string
	SetSessionId(v string) *DisableConversationShareResponseBody
	GetSessionId() *string
	SetShareAccessType(v string) *DisableConversationShareResponseBody
	GetShareAccessType() *string
	SetShareEnabled(v bool) *DisableConversationShareResponseBody
	GetShareEnabled() *bool
	SetShareEnabledAt(v string) *DisableConversationShareResponseBody
	GetShareEnabledAt() *string
	SetSharePath(v string) *DisableConversationShareResponseBody
	GetSharePath() *string
	SetShareToken(v string) *DisableConversationShareResponseBody
	GetShareToken() *string
	SetShareUrl(v string) *DisableConversationShareResponseBody
	GetShareUrl() *string
	SetShareUserGroupIds(v []*string) *DisableConversationShareResponseBody
	GetShareUserGroupIds() []*string
	SetShareUserIds(v []*int64) *DisableConversationShareResponseBody
	GetShareUserIds() []*int64
}

type DisableConversationShareResponseBody struct {
	// The result code. A value of SUCCESS indicates success. For failures, the corresponding error type is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// Invalid request parameters
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the source one-on-one session.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The share access type.
	//
	// example:
	//
	// PUBLIC
	ShareAccessType *string `json:"shareAccessType,omitempty" xml:"shareAccessType,omitempty"`
	// Indicates whether sharing is enabled.
	//
	// example:
	//
	// false
	ShareEnabled *bool `json:"shareEnabled,omitempty" xml:"shareEnabled,omitempty"`
	// The time when sharing was most recently enabled.
	//
	// example:
	//
	// 2026-11-16T13:00:00+00:00
	ShareEnabledAt *string `json:"shareEnabledAt,omitempty" xml:"shareEnabledAt,omitempty"`
	// The share path.
	//
	// example:
	//
	// null
	SharePath *string `json:"sharePath,omitempty" xml:"sharePath,omitempty"`
	// The share token. This token is present when sharing is enabled and can be used to access the public share preview API operation.
	//
	// example:
	//
	// null
	ShareToken *string `json:"shareToken,omitempty" xml:"shareToken,omitempty"`
	// The share URL.
	//
	// example:
	//
	// null
	ShareUrl *string `json:"shareUrl,omitempty" xml:"shareUrl,omitempty"`
	// The currently specified user group IDs.
	//
	// example:
	//
	// []
	ShareUserGroupIds []*string `json:"shareUserGroupIds,omitempty" xml:"shareUserGroupIds,omitempty" type:"Repeated"`
	// The list of most recently specified user IDs.
	//
	// example:
	//
	// []
	ShareUserIds []*int64 `json:"shareUserIds,omitempty" xml:"shareUserIds,omitempty" type:"Repeated"`
}

func (s DisableConversationShareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableConversationShareResponseBody) GoString() string {
	return s.String()
}

func (s *DisableConversationShareResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableConversationShareResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableConversationShareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableConversationShareResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *DisableConversationShareResponseBody) GetShareAccessType() *string {
	return s.ShareAccessType
}

func (s *DisableConversationShareResponseBody) GetShareEnabled() *bool {
	return s.ShareEnabled
}

func (s *DisableConversationShareResponseBody) GetShareEnabledAt() *string {
	return s.ShareEnabledAt
}

func (s *DisableConversationShareResponseBody) GetSharePath() *string {
	return s.SharePath
}

func (s *DisableConversationShareResponseBody) GetShareToken() *string {
	return s.ShareToken
}

func (s *DisableConversationShareResponseBody) GetShareUrl() *string {
	return s.ShareUrl
}

func (s *DisableConversationShareResponseBody) GetShareUserGroupIds() []*string {
	return s.ShareUserGroupIds
}

func (s *DisableConversationShareResponseBody) GetShareUserIds() []*int64 {
	return s.ShareUserIds
}

func (s *DisableConversationShareResponseBody) SetCode(v string) *DisableConversationShareResponseBody {
	s.Code = &v
	return s
}

func (s *DisableConversationShareResponseBody) SetMessage(v string) *DisableConversationShareResponseBody {
	s.Message = &v
	return s
}

func (s *DisableConversationShareResponseBody) SetRequestId(v string) *DisableConversationShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableConversationShareResponseBody) SetSessionId(v string) *DisableConversationShareResponseBody {
	s.SessionId = &v
	return s
}

func (s *DisableConversationShareResponseBody) SetShareAccessType(v string) *DisableConversationShareResponseBody {
	s.ShareAccessType = &v
	return s
}

func (s *DisableConversationShareResponseBody) SetShareEnabled(v bool) *DisableConversationShareResponseBody {
	s.ShareEnabled = &v
	return s
}

func (s *DisableConversationShareResponseBody) SetShareEnabledAt(v string) *DisableConversationShareResponseBody {
	s.ShareEnabledAt = &v
	return s
}

func (s *DisableConversationShareResponseBody) SetSharePath(v string) *DisableConversationShareResponseBody {
	s.SharePath = &v
	return s
}

func (s *DisableConversationShareResponseBody) SetShareToken(v string) *DisableConversationShareResponseBody {
	s.ShareToken = &v
	return s
}

func (s *DisableConversationShareResponseBody) SetShareUrl(v string) *DisableConversationShareResponseBody {
	s.ShareUrl = &v
	return s
}

func (s *DisableConversationShareResponseBody) SetShareUserGroupIds(v []*string) *DisableConversationShareResponseBody {
	s.ShareUserGroupIds = v
	return s
}

func (s *DisableConversationShareResponseBody) SetShareUserIds(v []*int64) *DisableConversationShareResponseBody {
	s.ShareUserIds = v
	return s
}

func (s *DisableConversationShareResponseBody) Validate() error {
	return dara.Validate(s)
}
