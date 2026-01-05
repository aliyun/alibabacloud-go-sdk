// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetDatasetRequest
	GetId() *string
}

type GetDatasetRequest struct {
	// The dataset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dataworks-dataset:3pXXXb8o0ngr07njhps1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetRequest) GetId() *string {
	return s.Id
}

func (s *GetDatasetRequest) SetId(v string) *GetDatasetRequest {
	s.Id = &v
	return s
}

func (s *GetDatasetRequest) Validate() error {
	return dara.Validate(s)
}
