// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCollectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteCollectorRequest
	GetClientToken() *string
}

type DeleteCollectorRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteCollectorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCollectorRequest) GoString() string {
	return s.String()
}

func (s *DeleteCollectorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCollectorRequest) SetClientToken(v string) *DeleteCollectorRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCollectorRequest) Validate() error {
	return dara.Validate(s)
}
