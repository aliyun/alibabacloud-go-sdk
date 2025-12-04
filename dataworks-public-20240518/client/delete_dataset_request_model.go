// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteDatasetRequest
	GetId() *string
}

type DeleteDatasetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dataworks-dataset:3pXXXb8o0ngr07njhps1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetRequest) GetId() *string {
	return s.Id
}

func (s *DeleteDatasetRequest) SetId(v string) *DeleteDatasetRequest {
	s.Id = &v
	return s
}

func (s *DeleteDatasetRequest) Validate() error {
	return dara.Validate(s)
}
