// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveEnvShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvShrink(v string) *SaveEnvShrinkRequest
	GetEnvShrink() *string
}

type SaveEnvShrinkRequest struct {
	// The JMeter environment.
	//
	// This parameter is required.
	EnvShrink *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s SaveEnvShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveEnvShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveEnvShrinkRequest) GetEnvShrink() *string {
	return s.EnvShrink
}

func (s *SaveEnvShrinkRequest) SetEnvShrink(v string) *SaveEnvShrinkRequest {
	s.EnvShrink = &v
	return s
}

func (s *SaveEnvShrinkRequest) Validate() error {
	return dara.Validate(s)
}
