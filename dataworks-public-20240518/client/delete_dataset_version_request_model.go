// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteDatasetVersionRequest
	GetId() *string
}

type DeleteDatasetVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dataworks-datasetVersion:3pXXXb8o0ngr07njhps1
	//
	// :2
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDatasetVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetVersionRequest) GetId() *string {
	return s.Id
}

func (s *DeleteDatasetVersionRequest) SetId(v string) *DeleteDatasetVersionRequest {
	s.Id = &v
	return s
}

func (s *DeleteDatasetVersionRequest) Validate() error {
	return dara.Validate(s)
}
