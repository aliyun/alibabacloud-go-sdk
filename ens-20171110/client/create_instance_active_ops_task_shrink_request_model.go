// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceActiveOpsTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *CreateInstanceActiveOpsTaskShrinkRequest
	GetInstanceIdsShrink() *string
}

type CreateInstanceActiveOpsTaskShrinkRequest struct {
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s CreateInstanceActiveOpsTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceActiveOpsTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceActiveOpsTaskShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *CreateInstanceActiveOpsTaskShrinkRequest) SetInstanceIdsShrink(v string) *CreateInstanceActiveOpsTaskShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *CreateInstanceActiveOpsTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
