// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteOutputDataset(v bool) *DeleteContextStoreRequest
	GetDeleteOutputDataset() *bool
}

type DeleteContextStoreRequest struct {
	// Specifies whether to simultaneously delete the memory output dataset (memory type). Default value: false.
	//
	// example:
	//
	// false
	DeleteOutputDataset *bool `json:"deleteOutputDataset,omitempty" xml:"deleteOutputDataset,omitempty"`
}

func (s DeleteContextStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextStoreRequest) GoString() string {
	return s.String()
}

func (s *DeleteContextStoreRequest) GetDeleteOutputDataset() *bool {
	return s.DeleteOutputDataset
}

func (s *DeleteContextStoreRequest) SetDeleteOutputDataset(v bool) *DeleteContextStoreRequest {
	s.DeleteOutputDataset = &v
	return s
}

func (s *DeleteContextStoreRequest) Validate() error {
	return dara.Validate(s)
}
