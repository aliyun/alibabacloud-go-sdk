// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *StartInstancesShrinkRequest
	GetInstanceIdsShrink() *string
}

type StartInstancesShrinkRequest struct {
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s StartInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartInstancesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *StartInstancesShrinkRequest) SetInstanceIdsShrink(v string) *StartInstancesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *StartInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
