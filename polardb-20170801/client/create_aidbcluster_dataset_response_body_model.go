// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIDBClusterDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateAIDBClusterDatasetResponseBody
	GetDBClusterId() *string
	SetDataServiceId(v string) *CreateAIDBClusterDatasetResponseBody
	GetDataServiceId() *string
	SetDatasetId(v string) *CreateAIDBClusterDatasetResponseBody
	GetDatasetId() *string
	SetDatasetName(v string) *CreateAIDBClusterDatasetResponseBody
	GetDatasetName() *string
	SetPath(v string) *CreateAIDBClusterDatasetResponseBody
	GetPath() *string
	SetRequestId(v string) *CreateAIDBClusterDatasetResponseBody
	GetRequestId() *string
}

type CreateAIDBClusterDatasetResponseBody struct {
	// The ID of the PolarDB database cluster.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The dataset management service ID.
	//
	// example:
	//
	// pcs-2zeei***
	DataServiceId *string `json:"DataServiceId,omitempty" xml:"DataServiceId,omitempty"`
	// The dataset ID.
	//
	// example:
	//
	// pds-2ze88***
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The dataset name.
	//
	// example:
	//
	// dataset01
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The path to the dataset.
	//
	// example:
	//
	// polardb_ai/datasets/train/sft/dataset01
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CD35F3-F3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAIDBClusterDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAIDBClusterDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAIDBClusterDatasetResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAIDBClusterDatasetResponseBody) GetDataServiceId() *string {
	return s.DataServiceId
}

func (s *CreateAIDBClusterDatasetResponseBody) GetDatasetId() *string {
	return s.DatasetId
}

func (s *CreateAIDBClusterDatasetResponseBody) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateAIDBClusterDatasetResponseBody) GetPath() *string {
	return s.Path
}

func (s *CreateAIDBClusterDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAIDBClusterDatasetResponseBody) SetDBClusterId(v string) *CreateAIDBClusterDatasetResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateAIDBClusterDatasetResponseBody) SetDataServiceId(v string) *CreateAIDBClusterDatasetResponseBody {
	s.DataServiceId = &v
	return s
}

func (s *CreateAIDBClusterDatasetResponseBody) SetDatasetId(v string) *CreateAIDBClusterDatasetResponseBody {
	s.DatasetId = &v
	return s
}

func (s *CreateAIDBClusterDatasetResponseBody) SetDatasetName(v string) *CreateAIDBClusterDatasetResponseBody {
	s.DatasetName = &v
	return s
}

func (s *CreateAIDBClusterDatasetResponseBody) SetPath(v string) *CreateAIDBClusterDatasetResponseBody {
	s.Path = &v
	return s
}

func (s *CreateAIDBClusterDatasetResponseBody) SetRequestId(v string) *CreateAIDBClusterDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAIDBClusterDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
