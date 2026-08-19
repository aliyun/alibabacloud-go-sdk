// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteMetaEntitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*string) *BatchDeleteMetaEntitiesRequest
	GetIds() []*string
}

type BatchDeleteMetaEntitiesRequest struct {
	// The list of entity IDs to delete. A maximum of 10 IDs are supported. All entities in the same batch must be of the same entityType.
	//
	// This parameter is required.
	Ids []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s BatchDeleteMetaEntitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteMetaEntitiesRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteMetaEntitiesRequest) GetIds() []*string {
	return s.Ids
}

func (s *BatchDeleteMetaEntitiesRequest) SetIds(v []*string) *BatchDeleteMetaEntitiesRequest {
	s.Ids = v
	return s
}

func (s *BatchDeleteMetaEntitiesRequest) Validate() error {
	return dara.Validate(s)
}
