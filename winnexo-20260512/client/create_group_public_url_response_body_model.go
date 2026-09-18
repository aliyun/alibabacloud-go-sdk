// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupPublicUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGroupPublicUrlResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateGroupPublicUrlResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreateGroupPublicUrlResponseBody
	GetGmtCreate() *string
	SetGroupId(v string) *CreateGroupPublicUrlResponseBody
	GetGroupId() *string
	SetMessage(v string) *CreateGroupPublicUrlResponseBody
	GetMessage() *string
	SetName(v string) *CreateGroupPublicUrlResponseBody
	GetName() *string
	SetOriginalUrl(v string) *CreateGroupPublicUrlResponseBody
	GetOriginalUrl() *string
	SetRequestId(v string) *CreateGroupPublicUrlResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateGroupPublicUrlResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateGroupPublicUrlResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateGroupPublicUrlResponseBody
	GetStatus() *string
}

type CreateGroupPublicUrlResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-11-12T03:08:56Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The project group ID.
	//
	// example:
	//
	// group_delivery
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The knowledge base name.
	//
	// example:
	//
	// p-toolset-80a4520e-b35c-4e8b-acf7-3a01c7307522
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The URL of the web page.
	//
	// example:
	//
	// https://mp.weixin.qq.com/s/iHqLKhkJcOyHNCOGejO32A
	OriginalUrl *string `json:"originalUrl,omitempty" xml:"originalUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The permission scope.
	//
	// example:
	//
	// PERSONAL
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The source ID.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The refund status. If a refund is in progress, query to confirm the refund status. Valid values:
	//
	// - SUCCESS: All succeeded.
	//
	// - FAIL: Failed.
	//
	// - WAIT_PAY: Waiting for refund.
	//
	// - EXPIRE: Expired.
	//
	// - PAYING: Refund in progress.
	//
	// - TERMINATE: Refund terminated.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateGroupPublicUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupPublicUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupPublicUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGroupPublicUrlResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupPublicUrlResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateGroupPublicUrlResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupPublicUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGroupPublicUrlResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateGroupPublicUrlResponseBody) GetOriginalUrl() *string {
	return s.OriginalUrl
}

func (s *CreateGroupPublicUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupPublicUrlResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateGroupPublicUrlResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateGroupPublicUrlResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateGroupPublicUrlResponseBody) SetCode(v string) *CreateGroupPublicUrlResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupPublicUrlResponseBody) SetDirectoryId(v string) *CreateGroupPublicUrlResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupPublicUrlResponseBody) SetGmtCreate(v string) *CreateGroupPublicUrlResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateGroupPublicUrlResponseBody) SetGroupId(v string) *CreateGroupPublicUrlResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateGroupPublicUrlResponseBody) SetMessage(v string) *CreateGroupPublicUrlResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupPublicUrlResponseBody) SetName(v string) *CreateGroupPublicUrlResponseBody {
	s.Name = &v
	return s
}

func (s *CreateGroupPublicUrlResponseBody) SetOriginalUrl(v string) *CreateGroupPublicUrlResponseBody {
	s.OriginalUrl = &v
	return s
}

func (s *CreateGroupPublicUrlResponseBody) SetRequestId(v string) *CreateGroupPublicUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupPublicUrlResponseBody) SetScope(v string) *CreateGroupPublicUrlResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateGroupPublicUrlResponseBody) SetSourceId(v string) *CreateGroupPublicUrlResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateGroupPublicUrlResponseBody) SetStatus(v string) *CreateGroupPublicUrlResponseBody {
	s.Status = &v
	return s
}

func (s *CreateGroupPublicUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
