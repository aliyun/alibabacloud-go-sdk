// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSaasInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListSaasInfoRequest
	GetAgentKey() *string
	SetSaasGroupCodes(v string) *ListSaasInfoRequest
	GetSaasGroupCodes() *string
	SetSaasName(v string) *ListSaasInfoRequest
	GetSaasName() *string
}

type ListSaasInfoRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// DS,FAQ
	SaasGroupCodes *string `json:"SaasGroupCodes,omitempty" xml:"SaasGroupCodes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// userTest
	SaasName *string `json:"SaasName,omitempty" xml:"SaasName,omitempty"`
}

func (s ListSaasInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSaasInfoRequest) GoString() string {
	return s.String()
}

func (s *ListSaasInfoRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListSaasInfoRequest) GetSaasGroupCodes() *string {
	return s.SaasGroupCodes
}

func (s *ListSaasInfoRequest) GetSaasName() *string {
	return s.SaasName
}

func (s *ListSaasInfoRequest) SetAgentKey(v string) *ListSaasInfoRequest {
	s.AgentKey = &v
	return s
}

func (s *ListSaasInfoRequest) SetSaasGroupCodes(v string) *ListSaasInfoRequest {
	s.SaasGroupCodes = &v
	return s
}

func (s *ListSaasInfoRequest) SetSaasName(v string) *ListSaasInfoRequest {
	s.SaasName = &v
	return s
}

func (s *ListSaasInfoRequest) Validate() error {
	return dara.Validate(s)
}
