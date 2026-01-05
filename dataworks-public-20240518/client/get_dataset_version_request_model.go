// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetDatasetVersionRequest
	GetId() *string
}

type GetDatasetVersionRequest struct {
	// The dataset version ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dataworks-datasetVersion:3pXXXb8o0ngr07njhps1
	//
	// :2
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDatasetVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetVersionRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetVersionRequest) GetId() *string {
	return s.Id
}

func (s *GetDatasetVersionRequest) SetId(v string) *GetDatasetVersionRequest {
	s.Id = &v
	return s
}

func (s *GetDatasetVersionRequest) Validate() error {
	return dara.Validate(s)
}
