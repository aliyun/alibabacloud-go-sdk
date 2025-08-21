// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSDGShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSDGIdShrink(v string) *DeleteSDGShrinkRequest
	GetSDGIdShrink() *string
}

type DeleteSDGShrinkRequest struct {
	// The IDs of the SDGs that you want to delete.
	//
	// This parameter is required.
	SDGIdShrink *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s DeleteSDGShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSDGShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteSDGShrinkRequest) GetSDGIdShrink() *string {
	return s.SDGIdShrink
}

func (s *DeleteSDGShrinkRequest) SetSDGIdShrink(v string) *DeleteSDGShrinkRequest {
	s.SDGIdShrink = &v
	return s
}

func (s *DeleteSDGShrinkRequest) Validate() error {
	return dara.Validate(s)
}
