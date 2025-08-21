// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePerspectiveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DescribePerspectiveRequest
	GetAgentKey() *string
	SetPerspectiveId(v string) *DescribePerspectiveRequest
	GetPerspectiveId() *string
}

type DescribePerspectiveRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 3001
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
}

func (s DescribePerspectiveRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePerspectiveRequest) GoString() string {
	return s.String()
}

func (s *DescribePerspectiveRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DescribePerspectiveRequest) GetPerspectiveId() *string {
	return s.PerspectiveId
}

func (s *DescribePerspectiveRequest) SetAgentKey(v string) *DescribePerspectiveRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribePerspectiveRequest) SetPerspectiveId(v string) *DescribePerspectiveRequest {
	s.PerspectiveId = &v
	return s
}

func (s *DescribePerspectiveRequest) Validate() error {
	return dara.Validate(s)
}
