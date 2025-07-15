// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListVersionsRequest
	GetAgentKey() *string
}

type ListVersionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 33a2658aaabf4c24b45d50e575125311_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s ListVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListVersionsRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListVersionsRequest) SetAgentKey(v string) *ListVersionsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListVersionsRequest) Validate() error {
	return dara.Validate(s)
}
