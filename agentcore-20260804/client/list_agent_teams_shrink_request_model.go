// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentTeamsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *ListAgentTeamsShrinkRequest
	GetBodyShrink() *string
}

type ListAgentTeamsShrinkRequest struct {
	// The request parameters for querying the agent team list.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentTeamsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentTeamsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAgentTeamsShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *ListAgentTeamsShrinkRequest) SetBodyShrink(v string) *ListAgentTeamsShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *ListAgentTeamsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
