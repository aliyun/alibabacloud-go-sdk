// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupAliDingDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGroupAliDingDocResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateGroupAliDingDocResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreateGroupAliDingDocResponseBody
	GetGmtCreate() *string
	SetGroupId(v string) *CreateGroupAliDingDocResponseBody
	GetGroupId() *string
	SetMessage(v string) *CreateGroupAliDingDocResponseBody
	GetMessage() *string
	SetName(v string) *CreateGroupAliDingDocResponseBody
	GetName() *string
	SetRequestId(v string) *CreateGroupAliDingDocResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateGroupAliDingDocResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateGroupAliDingDocResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateGroupAliDingDocResponseBody
	GetStatus() *string
}

type CreateGroupAliDingDocResponseBody struct {
	// The response code.
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
	// 2025-11-14T02:18:27Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The project group ID.
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The operation message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The image name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	// TENANT
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The original project ID.
	//
	// example:
	//
	// src_feishu_doc_1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The task running status.
	//
	// example:
	//
	// {\\"observedGeneration\\": 7, \\"servicesInstances\\": {}, \\"observedTime\\": \\"2025-12-17T11:57:07Z\\", \\"servicesWithPendingChanges\\": [], \\"latestEnvironmentDeploymentName\\": \\"manual-1765972627273-k7GZvr\\"}
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateGroupAliDingDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupAliDingDocResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupAliDingDocResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGroupAliDingDocResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupAliDingDocResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateGroupAliDingDocResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupAliDingDocResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGroupAliDingDocResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateGroupAliDingDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupAliDingDocResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateGroupAliDingDocResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateGroupAliDingDocResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateGroupAliDingDocResponseBody) SetCode(v string) *CreateGroupAliDingDocResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupAliDingDocResponseBody) SetDirectoryId(v string) *CreateGroupAliDingDocResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupAliDingDocResponseBody) SetGmtCreate(v string) *CreateGroupAliDingDocResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateGroupAliDingDocResponseBody) SetGroupId(v string) *CreateGroupAliDingDocResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateGroupAliDingDocResponseBody) SetMessage(v string) *CreateGroupAliDingDocResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupAliDingDocResponseBody) SetName(v string) *CreateGroupAliDingDocResponseBody {
	s.Name = &v
	return s
}

func (s *CreateGroupAliDingDocResponseBody) SetRequestId(v string) *CreateGroupAliDingDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupAliDingDocResponseBody) SetScope(v string) *CreateGroupAliDingDocResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateGroupAliDingDocResponseBody) SetSourceId(v string) *CreateGroupAliDingDocResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateGroupAliDingDocResponseBody) SetStatus(v string) *CreateGroupAliDingDocResponseBody {
	s.Status = &v
	return s
}

func (s *CreateGroupAliDingDocResponseBody) Validate() error {
	return dara.Validate(s)
}
