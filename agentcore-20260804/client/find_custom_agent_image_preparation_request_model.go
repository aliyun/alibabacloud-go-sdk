// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindCustomAgentImagePreparationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *FindCustomAgentImagePreparationRequest
	GetClientToken() *string
}

type FindCustomAgentImagePreparationRequest struct {
	// The idempotency token passed as a query parameter. This parameter is required when you explicitly submit customImage or container.type=custom, and optional for ACR image and code package submissions. The value must be 1 to 128 characters in length and can contain letters, digits, periods, underscores, colons, and hyphens. Reuse the same token for retries of the same submission. If the same token is used for different parameters, the request will be rejected.
	//
	// This parameter is required.
	//
	// example:
	//
	// image-update-20261022-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s FindCustomAgentImagePreparationRequest) String() string {
	return dara.Prettify(s)
}

func (s FindCustomAgentImagePreparationRequest) GoString() string {
	return s.String()
}

func (s *FindCustomAgentImagePreparationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *FindCustomAgentImagePreparationRequest) SetClientToken(v string) *FindCustomAgentImagePreparationRequest {
	s.ClientToken = &v
	return s
}

func (s *FindCustomAgentImagePreparationRequest) Validate() error {
	return dara.Validate(s)
}
