// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateCustomAgentShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateCustomAgentShrinkRequest
	GetClientToken() *string
}

type UpdateCustomAgentShrinkRequest struct {
	// The update configuration of the custom code or image agent. You must provide at least one top-level field. Omitted top-level fields retain their original values. When submitting `artifact`, specify `port` and the required fields for the corresponding artifact source. When submitting `runtime`, specify `compute` and `sessionPolicy`, and include the complete `hpa` that needs to be retained. `environment.variables` and other arrays are replaced as complete lists.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The idempotency token passed through query parameters. Required when explicitly submitting customImage or container.type=custom. Optional for ACR image and code package submissions. The length is 1 to 128 characters. Letters, digits, periods, underscores, colons, and hyphens are allowed. Reuse the same token for retries of the same submission. Using the same token for different parameters is rejected.
	//
	// example:
	//
	// custom-agent-request-20261010-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateCustomAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateCustomAgentShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCustomAgentShrinkRequest) SetBodyShrink(v string) *UpdateCustomAgentShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateCustomAgentShrinkRequest) SetClientToken(v string) *UpdateCustomAgentShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCustomAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
