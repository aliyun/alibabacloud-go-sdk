// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDatasetImportDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *OpenDatasetImportDataRequest
	GetDatasetId() *int64
	SetRecords(v []map[string]interface{}) *OpenDatasetImportDataRequest
	GetRecords() []map[string]interface{}
}

type OpenDatasetImportDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 730
	DatasetId *int64 `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	// This parameter is required.
	Records []map[string]interface{} `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s OpenDatasetImportDataRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenDatasetImportDataRequest) GoString() string {
	return s.String()
}

func (s *OpenDatasetImportDataRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *OpenDatasetImportDataRequest) GetRecords() []map[string]interface{} {
	return s.Records
}

func (s *OpenDatasetImportDataRequest) SetDatasetId(v int64) *OpenDatasetImportDataRequest {
	s.DatasetId = &v
	return s
}

func (s *OpenDatasetImportDataRequest) SetRecords(v []map[string]interface{}) *OpenDatasetImportDataRequest {
	s.Records = v
	return s
}

func (s *OpenDatasetImportDataRequest) Validate() error {
	return dara.Validate(s)
}
