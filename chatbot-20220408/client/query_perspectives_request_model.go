// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPerspectivesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *QueryPerspectivesRequest
	GetAgentKey() *string
}

type QueryPerspectivesRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s QueryPerspectivesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPerspectivesRequest) GoString() string {
	return s.String()
}

func (s *QueryPerspectivesRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *QueryPerspectivesRequest) SetAgentKey(v string) *QueryPerspectivesRequest {
	s.AgentKey = &v
	return s
}

func (s *QueryPerspectivesRequest) Validate() error {
	return dara.Validate(s)
}
