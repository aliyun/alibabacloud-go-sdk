// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSaasPermissionGroupInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListSaasPermissionGroupInfosRequest
	GetAgentKey() *string
}

type ListSaasPermissionGroupInfosRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s ListSaasPermissionGroupInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSaasPermissionGroupInfosRequest) GoString() string {
	return s.String()
}

func (s *ListSaasPermissionGroupInfosRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListSaasPermissionGroupInfosRequest) SetAgentKey(v string) *ListSaasPermissionGroupInfosRequest {
	s.AgentKey = &v
	return s
}

func (s *ListSaasPermissionGroupInfosRequest) Validate() error {
	return dara.Validate(s)
}
