// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentIMChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteAgentIMChannelRequest
	GetClientToken() *string
}

type DeleteAgentIMChannelRequest struct {
	// A reserved idempotence token. The backend does not provide persistent idempotence guarantee in the current version.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s DeleteAgentIMChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentIMChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentIMChannelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAgentIMChannelRequest) SetClientToken(v string) *DeleteAgentIMChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAgentIMChannelRequest) Validate() error {
	return dara.Validate(s)
}
