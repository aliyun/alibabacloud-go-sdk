// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppTokenQuery interface {
	dara.Model
	String() string
	GoString() string
	SetToken(v string) *NeuronAppTokenQuery
	GetToken() *string
}

type NeuronAppTokenQuery struct {
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s NeuronAppTokenQuery) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppTokenQuery) GoString() string {
	return s.String()
}

func (s *NeuronAppTokenQuery) GetToken() *string {
	return s.Token
}

func (s *NeuronAppTokenQuery) SetToken(v string) *NeuronAppTokenQuery {
	s.Token = &v
	return s
}

func (s *NeuronAppTokenQuery) Validate() error {
	return dara.Validate(s)
}
