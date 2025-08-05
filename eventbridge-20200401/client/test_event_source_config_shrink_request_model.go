// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestEventSourceConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceMySQLParametersShrink(v string) *TestEventSourceConfigShrinkRequest
	GetSourceMySQLParametersShrink() *string
}

type TestEventSourceConfigShrinkRequest struct {
	// The parameters that are configured if you specify MySQL as the event source.
	SourceMySQLParametersShrink *string `json:"SourceMySQLParameters,omitempty" xml:"SourceMySQLParameters,omitempty"`
}

func (s TestEventSourceConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TestEventSourceConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *TestEventSourceConfigShrinkRequest) GetSourceMySQLParametersShrink() *string {
	return s.SourceMySQLParametersShrink
}

func (s *TestEventSourceConfigShrinkRequest) SetSourceMySQLParametersShrink(v string) *TestEventSourceConfigShrinkRequest {
	s.SourceMySQLParametersShrink = &v
	return s
}

func (s *TestEventSourceConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
