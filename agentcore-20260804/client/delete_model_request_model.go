// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteModelRequest
	GetClientToken() *string
}

type DeleteModelRequest struct {
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s DeleteModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteModelRequest) SetClientToken(v string) *DeleteModelRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteModelRequest) Validate() error {
	return dara.Validate(s)
}
