// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *QueryDatasetInfoRequest
	GetDatasetId() *string
}

type QueryDatasetInfoRequest struct {
	// Queries information about a specified dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// a201c85c-******
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
}

func (s QueryDatasetInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *QueryDatasetInfoRequest) SetDatasetId(v string) *QueryDatasetInfoRequest {
	s.DatasetId = &v
	return s
}

func (s *QueryDatasetInfoRequest) Validate() error {
	return dara.Validate(s)
}
