// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFeishuDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGroupFeishuDocResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateGroupFeishuDocResponseBody
	GetDirectoryId() *string
	SetDocUrl(v string) *CreateGroupFeishuDocResponseBody
	GetDocUrl() *string
	SetGmtCreate(v string) *CreateGroupFeishuDocResponseBody
	GetGmtCreate() *string
	SetGroupId(v string) *CreateGroupFeishuDocResponseBody
	GetGroupId() *string
	SetMessage(v string) *CreateGroupFeishuDocResponseBody
	GetMessage() *string
	SetName(v string) *CreateGroupFeishuDocResponseBody
	GetName() *string
	SetRequestId(v string) *CreateGroupFeishuDocResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateGroupFeishuDocResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateGroupFeishuDocResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateGroupFeishuDocResponseBody
	GetStatus() *string
}

type CreateGroupFeishuDocResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// dir_group_child
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The document URL.
	//
	// example:
	//
	// https://example.feishu.cn/docx/doxcnExample
	DocUrl *string `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-08-26T10:00:00+08:00
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
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The image name.
	//
	// example:
	//
	// Project Plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The permission scope.
	//
	// example:
	//
	// GROUP
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The original project ID.
	//
	// example:
	//
	// src_feishu_doc_1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The signing status. Valid values:
	//
	// - CREATED: Created but not signed.
	//
	// - SUCCESS: Signed successfully.
	//
	// - STOP: Terminated.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateGroupFeishuDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuDocResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuDocResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGroupFeishuDocResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupFeishuDocResponseBody) GetDocUrl() *string {
	return s.DocUrl
}

func (s *CreateGroupFeishuDocResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateGroupFeishuDocResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupFeishuDocResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGroupFeishuDocResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateGroupFeishuDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupFeishuDocResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateGroupFeishuDocResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateGroupFeishuDocResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateGroupFeishuDocResponseBody) SetCode(v string) *CreateGroupFeishuDocResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupFeishuDocResponseBody) SetDirectoryId(v string) *CreateGroupFeishuDocResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupFeishuDocResponseBody) SetDocUrl(v string) *CreateGroupFeishuDocResponseBody {
	s.DocUrl = &v
	return s
}

func (s *CreateGroupFeishuDocResponseBody) SetGmtCreate(v string) *CreateGroupFeishuDocResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateGroupFeishuDocResponseBody) SetGroupId(v string) *CreateGroupFeishuDocResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateGroupFeishuDocResponseBody) SetMessage(v string) *CreateGroupFeishuDocResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupFeishuDocResponseBody) SetName(v string) *CreateGroupFeishuDocResponseBody {
	s.Name = &v
	return s
}

func (s *CreateGroupFeishuDocResponseBody) SetRequestId(v string) *CreateGroupFeishuDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupFeishuDocResponseBody) SetScope(v string) *CreateGroupFeishuDocResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateGroupFeishuDocResponseBody) SetSourceId(v string) *CreateGroupFeishuDocResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateGroupFeishuDocResponseBody) SetStatus(v string) *CreateGroupFeishuDocResponseBody {
	s.Status = &v
	return s
}

func (s *CreateGroupFeishuDocResponseBody) Validate() error {
	return dara.Validate(s)
}
