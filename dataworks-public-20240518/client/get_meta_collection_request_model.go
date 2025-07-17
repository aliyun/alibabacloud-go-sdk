// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetMetaCollectionRequest
	GetId() *string
}

type GetMetaCollectionRequest struct {
	// The collection ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// category.123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetMetaCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCollectionRequest) GoString() string {
	return s.String()
}

func (s *GetMetaCollectionRequest) GetId() *string {
	return s.Id
}

func (s *GetMetaCollectionRequest) SetId(v string) *GetMetaCollectionRequest {
	s.Id = &v
	return s
}

func (s *GetMetaCollectionRequest) Validate() error {
	return dara.Validate(s)
}
