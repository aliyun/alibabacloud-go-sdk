// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDatasetDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *ImportDatasetDataRequest
	GetDatasetId() *string
	SetRecords(v []map[string]interface{}) *ImportDatasetDataRequest
	GetRecords() []map[string]interface{}
}

type ImportDatasetDataRequest struct {
	// The dataset ID. You can view this ID in the dataset list in the console.
	//
	// example:
	//
	// 730
	DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	// The list of data records to add or update. A maximum of 100 records can be included in a single batch operation.
	//
	// Note: The records must strictly follow the schema configured for the target dataset in the console. The add or update logic depends on the primary key type of the target dataset. For detailed example requests, see the Request Description section below.
	//
	// This parameter is required.
	Records []map[string]interface{} `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s ImportDatasetDataRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportDatasetDataRequest) GoString() string {
	return s.String()
}

func (s *ImportDatasetDataRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *ImportDatasetDataRequest) GetRecords() []map[string]interface{} {
	return s.Records
}

func (s *ImportDatasetDataRequest) SetDatasetId(v string) *ImportDatasetDataRequest {
	s.DatasetId = &v
	return s
}

func (s *ImportDatasetDataRequest) SetRecords(v []map[string]interface{}) *ImportDatasetDataRequest {
	s.Records = v
	return s
}

func (s *ImportDatasetDataRequest) Validate() error {
	return dara.Validate(s)
}
