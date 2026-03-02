// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppToken interface {
	dara.Model
	String() string
	GoString() string
	SetToken(v string) *NeuronAppToken
	GetToken() *string
}

type NeuronAppToken struct {
	// example:
	//
	// 1343424sdadsa
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s NeuronAppToken) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppToken) GoString() string {
	return s.String()
}

func (s *NeuronAppToken) GetToken() *string {
	return s.Token
}

func (s *NeuronAppToken) SetToken(v string) *NeuronAppToken {
	s.Token = &v
	return s
}

func (s *NeuronAppToken) Validate() error {
	return dara.Validate(s)
}
