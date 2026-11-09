// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteAgentVersionRequest
	GetClientToken() *string
}

type DeleteAgentVersionRequest struct {
	// The client-generated idempotency token. This is a reserved parameter, and the backend does not currently provide persistent idempotency guarantees.
	//
	// example:
	//
	// c2d8a6f0-3b1e-4a7d-9f25-6e4b0c1a8d37
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s DeleteAgentVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentVersionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAgentVersionRequest) SetClientToken(v string) *DeleteAgentVersionRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAgentVersionRequest) Validate() error {
	return dara.Validate(s)
}
