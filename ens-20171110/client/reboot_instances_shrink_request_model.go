// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *RebootInstancesShrinkRequest
	GetInstanceIdsShrink() *string
}

type RebootInstancesShrinkRequest struct {
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s RebootInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RebootInstancesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *RebootInstancesShrinkRequest) SetInstanceIdsShrink(v string) *RebootInstancesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *RebootInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
