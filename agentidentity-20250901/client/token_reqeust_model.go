// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTokenReqeust interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoint(v string) *TokenReqeust
	GetEndpoint() *string
}

type TokenReqeust struct {
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
}

func (s TokenReqeust) String() string {
	return dara.Prettify(s)
}

func (s TokenReqeust) GoString() string {
	return s.String()
}

func (s *TokenReqeust) GetEndpoint() *string {
	return s.Endpoint
}

func (s *TokenReqeust) SetEndpoint(v string) *TokenReqeust {
	s.Endpoint = &v
	return s
}

func (s *TokenReqeust) Validate() error {
	return dara.Validate(s)
}
