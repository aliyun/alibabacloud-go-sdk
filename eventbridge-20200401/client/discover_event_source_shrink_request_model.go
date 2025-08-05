// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiscoverEventSourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceMySQLParametersShrink(v string) *DiscoverEventSourceShrinkRequest
	GetSourceMySQLParametersShrink() *string
}

type DiscoverEventSourceShrinkRequest struct {
	SourceMySQLParametersShrink *string `json:"SourceMySQLParameters,omitempty" xml:"SourceMySQLParameters,omitempty"`
}

func (s DiscoverEventSourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DiscoverEventSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DiscoverEventSourceShrinkRequest) GetSourceMySQLParametersShrink() *string {
	return s.SourceMySQLParametersShrink
}

func (s *DiscoverEventSourceShrinkRequest) SetSourceMySQLParametersShrink(v string) *DiscoverEventSourceShrinkRequest {
	s.SourceMySQLParametersShrink = &v
	return s
}

func (s *DiscoverEventSourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
