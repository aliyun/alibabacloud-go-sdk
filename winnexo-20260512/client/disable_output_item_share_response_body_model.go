// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableOutputItemShareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableOutputItemShareResponseBody
	GetCode() *string
	SetMessage(v string) *DisableOutputItemShareResponseBody
	GetMessage() *string
	SetOutputItemId(v string) *DisableOutputItemShareResponseBody
	GetOutputItemId() *string
	SetRequestId(v string) *DisableOutputItemShareResponseBody
	GetRequestId() *string
	SetShareAccessType(v string) *DisableOutputItemShareResponseBody
	GetShareAccessType() *string
	SetShareEnabled(v bool) *DisableOutputItemShareResponseBody
	GetShareEnabled() *bool
	SetShareEnabledAt(v string) *DisableOutputItemShareResponseBody
	GetShareEnabledAt() *string
	SetSharePath(v string) *DisableOutputItemShareResponseBody
	GetSharePath() *string
	SetShareToken(v string) *DisableOutputItemShareResponseBody
	GetShareToken() *string
	SetShareUrl(v string) *DisableOutputItemShareResponseBody
	GetShareUrl() *string
	SetShareUserGroupIds(v []*string) *DisableOutputItemShareResponseBody
	GetShareUserGroupIds() []*string
	SetShareUserIds(v []*int64) *DisableOutputItemShareResponseBody
	GetShareUserIds() []*int64
}

type DisableOutputItemShareResponseBody struct {
	// The status code.
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
	// The most recently configured sharing scope.
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
	// The time when sharing was most recently enabled, in ISO 8601 format.
	//
	// example:
	//
	// 2026-11-16T13:00:00+00:00
	ShareEnabledAt *string `json:"shareEnabledAt,omitempty" xml:"shareEnabledAt,omitempty"`
	// The relative path for internal sharing. This value is empty when sharing is disabled.
	//
	// example:
	//
	// null
	SharePath *string `json:"sharePath,omitempty" xml:"sharePath,omitempty"`
	// The sharing token. This token exists when sharing is enabled and can be used to access the public sharing preview API operation.
	//
	// example:
	//
	// null
	ShareToken *string `json:"shareToken,omitempty" xml:"shareToken,omitempty"`
	// The current absolute sharing URL. This value is empty when sharing is disabled.
	//
	// example:
	//
	// null
	ShareUrl *string `json:"shareUrl,omitempty" xml:"shareUrl,omitempty"`
	// The list of specified user group IDs from the most recent configuration.
	//
	// example:
	//
	// []
	ShareUserGroupIds []*string `json:"shareUserGroupIds,omitempty" xml:"shareUserGroupIds,omitempty" type:"Repeated"`
	// The list of specified user IDs from the most recent configuration.
	//
	// example:
	//
	// []
	ShareUserIds []*int64 `json:"shareUserIds,omitempty" xml:"shareUserIds,omitempty" type:"Repeated"`
}

func (s DisableOutputItemShareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableOutputItemShareResponseBody) GoString() string {
	return s.String()
}

func (s *DisableOutputItemShareResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableOutputItemShareResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableOutputItemShareResponseBody) GetOutputItemId() *string {
	return s.OutputItemId
}

func (s *DisableOutputItemShareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableOutputItemShareResponseBody) GetShareAccessType() *string {
	return s.ShareAccessType
}

func (s *DisableOutputItemShareResponseBody) GetShareEnabled() *bool {
	return s.ShareEnabled
}

func (s *DisableOutputItemShareResponseBody) GetShareEnabledAt() *string {
	return s.ShareEnabledAt
}

func (s *DisableOutputItemShareResponseBody) GetSharePath() *string {
	return s.SharePath
}

func (s *DisableOutputItemShareResponseBody) GetShareToken() *string {
	return s.ShareToken
}

func (s *DisableOutputItemShareResponseBody) GetShareUrl() *string {
	return s.ShareUrl
}

func (s *DisableOutputItemShareResponseBody) GetShareUserGroupIds() []*string {
	return s.ShareUserGroupIds
}

func (s *DisableOutputItemShareResponseBody) GetShareUserIds() []*int64 {
	return s.ShareUserIds
}

func (s *DisableOutputItemShareResponseBody) SetCode(v string) *DisableOutputItemShareResponseBody {
	s.Code = &v
	return s
}

func (s *DisableOutputItemShareResponseBody) SetMessage(v string) *DisableOutputItemShareResponseBody {
	s.Message = &v
	return s
}

func (s *DisableOutputItemShareResponseBody) SetOutputItemId(v string) *DisableOutputItemShareResponseBody {
	s.OutputItemId = &v
	return s
}

func (s *DisableOutputItemShareResponseBody) SetRequestId(v string) *DisableOutputItemShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableOutputItemShareResponseBody) SetShareAccessType(v string) *DisableOutputItemShareResponseBody {
	s.ShareAccessType = &v
	return s
}

func (s *DisableOutputItemShareResponseBody) SetShareEnabled(v bool) *DisableOutputItemShareResponseBody {
	s.ShareEnabled = &v
	return s
}

func (s *DisableOutputItemShareResponseBody) SetShareEnabledAt(v string) *DisableOutputItemShareResponseBody {
	s.ShareEnabledAt = &v
	return s
}

func (s *DisableOutputItemShareResponseBody) SetSharePath(v string) *DisableOutputItemShareResponseBody {
	s.SharePath = &v
	return s
}

func (s *DisableOutputItemShareResponseBody) SetShareToken(v string) *DisableOutputItemShareResponseBody {
	s.ShareToken = &v
	return s
}

func (s *DisableOutputItemShareResponseBody) SetShareUrl(v string) *DisableOutputItemShareResponseBody {
	s.ShareUrl = &v
	return s
}

func (s *DisableOutputItemShareResponseBody) SetShareUserGroupIds(v []*string) *DisableOutputItemShareResponseBody {
	s.ShareUserGroupIds = v
	return s
}

func (s *DisableOutputItemShareResponseBody) SetShareUserIds(v []*int64) *DisableOutputItemShareResponseBody {
	s.ShareUserIds = v
	return s
}

func (s *DisableOutputItemShareResponseBody) Validate() error {
	return dara.Validate(s)
}
