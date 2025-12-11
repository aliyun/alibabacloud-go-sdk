// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitIMConnectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *InitIMConnectRequest
	GetAgentKey() *string
	SetFrom(v string) *InitIMConnectRequest
	GetFrom() *string
	SetUserAccessToken(v string) *InitIMConnectRequest
	GetUserAccessToken() *string
}

type InitIMConnectRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	From            *string `json:"From,omitempty" xml:"From,omitempty"`
	UserAccessToken *string `json:"UserAccessToken,omitempty" xml:"UserAccessToken,omitempty"`
}

func (s InitIMConnectRequest) String() string {
	return dara.Prettify(s)
}

func (s InitIMConnectRequest) GoString() string {
	return s.String()
}

func (s *InitIMConnectRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *InitIMConnectRequest) GetFrom() *string {
	return s.From
}

func (s *InitIMConnectRequest) GetUserAccessToken() *string {
	return s.UserAccessToken
}

func (s *InitIMConnectRequest) SetAgentKey(v string) *InitIMConnectRequest {
	s.AgentKey = &v
	return s
}

func (s *InitIMConnectRequest) SetFrom(v string) *InitIMConnectRequest {
	s.From = &v
	return s
}

func (s *InitIMConnectRequest) SetUserAccessToken(v string) *InitIMConnectRequest {
	s.UserAccessToken = &v
	return s
}

func (s *InitIMConnectRequest) Validate() error {
	return dara.Validate(s)
}
