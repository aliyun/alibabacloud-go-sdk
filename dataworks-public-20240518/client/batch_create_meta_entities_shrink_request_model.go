// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateMetaEntitiesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntitiesShrink(v string) *BatchCreateMetaEntitiesShrinkRequest
	GetEntitiesShrink() *string
}

type BatchCreateMetaEntitiesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// []
	EntitiesShrink *string `json:"Entities,omitempty" xml:"Entities,omitempty"`
}

func (s BatchCreateMetaEntitiesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateMetaEntitiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateMetaEntitiesShrinkRequest) GetEntitiesShrink() *string {
	return s.EntitiesShrink
}

func (s *BatchCreateMetaEntitiesShrinkRequest) SetEntitiesShrink(v string) *BatchCreateMetaEntitiesShrinkRequest {
	s.EntitiesShrink = &v
	return s
}

func (s *BatchCreateMetaEntitiesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
