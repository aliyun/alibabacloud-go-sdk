// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetDetailInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *QueryDatasetDetailInfoRequest
	GetDatasetId() *string
}

type QueryDatasetDetailInfoRequest struct {
	// The ID of the training dataset that you want to remove from the specified custom linguistic model.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5820f58c-c734-4d8a-baf1-7979af4f****
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
}

func (s QueryDatasetDetailInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetDetailInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDatasetDetailInfoRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *QueryDatasetDetailInfoRequest) SetDatasetId(v string) *QueryDatasetDetailInfoRequest {
	s.DatasetId = &v
	return s
}

func (s *QueryDatasetDetailInfoRequest) Validate() error {
	return dara.Validate(s)
}
