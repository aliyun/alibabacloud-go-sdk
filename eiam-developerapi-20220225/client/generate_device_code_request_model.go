// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDeviceCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScope(v string) *GenerateDeviceCodeRequest
	GetScope() *string
}

type GenerateDeviceCodeRequest struct {
	// The authorization scope.
	//
	// example:
	//
	// xxx
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s GenerateDeviceCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateDeviceCodeRequest) GoString() string {
	return s.String()
}

func (s *GenerateDeviceCodeRequest) GetScope() *string {
	return s.Scope
}

func (s *GenerateDeviceCodeRequest) SetScope(v string) *GenerateDeviceCodeRequest {
	s.Scope = &v
	return s
}

func (s *GenerateDeviceCodeRequest) Validate() error {
	return dara.Validate(s)
}
