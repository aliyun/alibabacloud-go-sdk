// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataServiceId(v string) *DeleteAIDBClusterDatasetResponseBody
	GetDataServiceId() *string
	SetDatasetId(v string) *DeleteAIDBClusterDatasetResponseBody
	GetDatasetId() *string
	SetRequestId(v string) *DeleteAIDBClusterDatasetResponseBody
	GetRequestId() *string
}

type DeleteAIDBClusterDatasetResponseBody struct {
	// example:
	//
	// pcs-2zeei***
	DataServiceId *string `json:"DataServiceId,omitempty" xml:"DataServiceId,omitempty"`
	// example:
	//
	// pds-xxxxxxxxxxxxxxxx
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CD35F3-F3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAIDBClusterDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterDatasetResponseBody) GetDataServiceId() *string {
	return s.DataServiceId
}

func (s *DeleteAIDBClusterDatasetResponseBody) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DeleteAIDBClusterDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAIDBClusterDatasetResponseBody) SetDataServiceId(v string) *DeleteAIDBClusterDatasetResponseBody {
	s.DataServiceId = &v
	return s
}

func (s *DeleteAIDBClusterDatasetResponseBody) SetDatasetId(v string) *DeleteAIDBClusterDatasetResponseBody {
	s.DatasetId = &v
	return s
}

func (s *DeleteAIDBClusterDatasetResponseBody) SetRequestId(v string) *DeleteAIDBClusterDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAIDBClusterDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
