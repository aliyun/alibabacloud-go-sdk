// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportAttackPathAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeType(v string) *ListSupportAttackPathAssetRequest
	GetNodeType() *string
	SetPathName(v string) *ListSupportAttackPathAssetRequest
	GetPathName() *string
	SetPathType(v string) *ListSupportAttackPathAssetRequest
	GetPathType() *string
	SetSupportType(v string) *ListSupportAttackPathAssetRequest
	GetSupportType() *string
}

type ListSupportAttackPathAssetRequest struct {
	// The node type. Valid values:
	//
	// - **start**: start node.
	//
	// - **end**: end node.
	//
	// example:
	//
	// end
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The path name.
	//
	// > Call [ListAvailableAttackPath](~~ListAvailableAttackPath~~) to query path names.
	//
	// example:
	//
	// ecs_get_credential_by_create_login_profile
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
	// The path type.
	//
	// > Call [ListAvailableAttackPath](~~ListAvailableAttackPath~~) to query path types.
	//
	// example:
	//
	// role_escalation
	PathType *string `json:"PathType,omitempty" xml:"PathType,omitempty"`
	// The support type. Valid values:
	//
	// - **event**: attack path alert event.
	//
	// - **whitelist**: attack path whitelist.
	//
	// - **sensitive**: attack path sensitive asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// event
	SupportType *string `json:"SupportType,omitempty" xml:"SupportType,omitempty"`
}

func (s ListSupportAttackPathAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSupportAttackPathAssetRequest) GoString() string {
	return s.String()
}

func (s *ListSupportAttackPathAssetRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *ListSupportAttackPathAssetRequest) GetPathName() *string {
	return s.PathName
}

func (s *ListSupportAttackPathAssetRequest) GetPathType() *string {
	return s.PathType
}

func (s *ListSupportAttackPathAssetRequest) GetSupportType() *string {
	return s.SupportType
}

func (s *ListSupportAttackPathAssetRequest) SetNodeType(v string) *ListSupportAttackPathAssetRequest {
	s.NodeType = &v
	return s
}

func (s *ListSupportAttackPathAssetRequest) SetPathName(v string) *ListSupportAttackPathAssetRequest {
	s.PathName = &v
	return s
}

func (s *ListSupportAttackPathAssetRequest) SetPathType(v string) *ListSupportAttackPathAssetRequest {
	s.PathType = &v
	return s
}

func (s *ListSupportAttackPathAssetRequest) SetSupportType(v string) *ListSupportAttackPathAssetRequest {
	s.SupportType = &v
	return s
}

func (s *ListSupportAttackPathAssetRequest) Validate() error {
	return dara.Validate(s)
}
