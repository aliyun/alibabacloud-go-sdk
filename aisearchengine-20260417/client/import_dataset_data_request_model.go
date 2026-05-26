// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDatasetDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *ImportDatasetDataRequest
	GetDatasetId() *int64
	SetRecords(v []map[string]interface{}) *ImportDatasetDataRequest
	GetRecords() []map[string]interface{}
}

type ImportDatasetDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 730
	DatasetId *int64 `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	// This parameter is required.
	Records []map[string]interface{} `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s ImportDatasetDataRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportDatasetDataRequest) GoString() string {
	return s.String()
}

func (s *ImportDatasetDataRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *ImportDatasetDataRequest) GetRecords() []map[string]interface{} {
	return s.Records
}

func (s *ImportDatasetDataRequest) SetDatasetId(v int64) *ImportDatasetDataRequest {
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
