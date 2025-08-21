// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCollectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartCollectorRequest
	GetClientToken() *string
}

type StartCollectorRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s StartCollectorRequest) String() string {
	return dara.Prettify(s)
}

func (s StartCollectorRequest) GoString() string {
	return s.String()
}

func (s *StartCollectorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartCollectorRequest) SetClientToken(v string) *StartCollectorRequest {
	s.ClientToken = &v
	return s
}

func (s *StartCollectorRequest) Validate() error {
	return dara.Validate(s)
}
