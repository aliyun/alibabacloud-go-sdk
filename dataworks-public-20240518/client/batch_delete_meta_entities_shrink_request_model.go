// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteMetaEntitiesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdsShrink(v string) *BatchDeleteMetaEntitiesShrinkRequest
	GetIdsShrink() *string
}

type BatchDeleteMetaEntitiesShrinkRequest struct {
	// An array of IDs for the meta entities to delete. You can specify up to 10 IDs in a single request. All entities in the batch must have the same EntityType.
	//
	// This parameter is required.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s BatchDeleteMetaEntitiesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteMetaEntitiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteMetaEntitiesShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *BatchDeleteMetaEntitiesShrinkRequest) SetIdsShrink(v string) *BatchDeleteMetaEntitiesShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *BatchDeleteMetaEntitiesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
