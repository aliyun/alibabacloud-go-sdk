// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteMetaCollectionRequest
	GetId() *string
}

type DeleteMetaCollectionRequest struct {
	// The collection ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// category.123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteMetaCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaCollectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetaCollectionRequest) GetId() *string {
	return s.Id
}

func (s *DeleteMetaCollectionRequest) SetId(v string) *DeleteMetaCollectionRequest {
	s.Id = &v
	return s
}

func (s *DeleteMetaCollectionRequest) Validate() error {
	return dara.Validate(s)
}
