// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteModelConnectionRequest
	GetClientToken() *string
}

type DeleteModelConnectionRequest struct {
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s DeleteModelConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteModelConnectionRequest) SetClientToken(v string) *DeleteModelConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteModelConnectionRequest) Validate() error {
	return dara.Validate(s)
}
