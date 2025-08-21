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
	//
	// example:
	//
	// 8XNBzDucJv
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// QUM4SndaY3VPMjhkQldDZUNOR0ZaTmZ5R3NBY0FKWHJ4OGc4dERZbEJzcjNIKzFiS1RyTjhXRUpBYmVpQlpsakprNDRFVkdxcy9HWVk2RXZvalU3bHhxRkJlc1NBUXZwdHFKOTE2UTNwamQ4b1U4N3dEbmhyRjc4R2hOQStvMnMrYkV2dlVpSHNvWC96SEVNZWRqMjBuMXdjNklpamJzaDNWYllnUldDZGhJPQ==
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
