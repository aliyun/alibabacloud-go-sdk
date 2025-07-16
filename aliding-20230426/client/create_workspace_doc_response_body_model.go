// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *CreateWorkspaceDocResponseBody
	GetDentryUuid() *string
	SetDocKey(v string) *CreateWorkspaceDocResponseBody
	GetDocKey() *string
	SetNodeId(v string) *CreateWorkspaceDocResponseBody
	GetNodeId() *string
	SetRequestId(v string) *CreateWorkspaceDocResponseBody
	GetRequestId() *string
	SetUrl(v string) *CreateWorkspaceDocResponseBody
	GetUrl() *string
	SetVendorRequestId(v string) *CreateWorkspaceDocResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CreateWorkspaceDocResponseBody
	GetVendorType() *string
	SetWorkspaceId(v string) *CreateWorkspaceDocResponseBody
	GetWorkspaceId() *string
}

type CreateWorkspaceDocResponseBody struct {
	// example:
	//
	// YRBcvy
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	// example:
	//
	// QoJGq7xxx
	DocKey *string `json:"docKey,omitempty" xml:"docKey,omitempty"`
	// example:
	//
	// YRBGv0Ye
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// https://xxx/workspaceId/docs/nodeId
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
	// example:
	//
	// YRBGvy
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateWorkspaceDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceDocResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocResponseBody) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *CreateWorkspaceDocResponseBody) GetDocKey() *string {
	return s.DocKey
}

func (s *CreateWorkspaceDocResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateWorkspaceDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkspaceDocResponseBody) GetUrl() *string {
	return s.Url
}

func (s *CreateWorkspaceDocResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CreateWorkspaceDocResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CreateWorkspaceDocResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateWorkspaceDocResponseBody) SetDentryUuid(v string) *CreateWorkspaceDocResponseBody {
	s.DentryUuid = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetDocKey(v string) *CreateWorkspaceDocResponseBody {
	s.DocKey = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetNodeId(v string) *CreateWorkspaceDocResponseBody {
	s.NodeId = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetRequestId(v string) *CreateWorkspaceDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetUrl(v string) *CreateWorkspaceDocResponseBody {
	s.Url = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetVendorRequestId(v string) *CreateWorkspaceDocResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetVendorType(v string) *CreateWorkspaceDocResponseBody {
	s.VendorType = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetWorkspaceId(v string) *CreateWorkspaceDocResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) Validate() error {
	return dara.Validate(s)
}
