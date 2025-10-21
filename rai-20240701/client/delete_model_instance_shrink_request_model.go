// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelInstanceIdListShrink(v string) *DeleteModelInstanceShrinkRequest
	GetModelInstanceIdListShrink() *string
}

type DeleteModelInstanceShrinkRequest struct {
	ModelInstanceIdListShrink *string `json:"ModelInstanceIdList,omitempty" xml:"ModelInstanceIdList,omitempty"`
}

func (s DeleteModelInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelInstanceShrinkRequest) GetModelInstanceIdListShrink() *string {
	return s.ModelInstanceIdListShrink
}

func (s *DeleteModelInstanceShrinkRequest) SetModelInstanceIdListShrink(v string) *DeleteModelInstanceShrinkRequest {
	s.ModelInstanceIdListShrink = &v
	return s
}

func (s *DeleteModelInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
