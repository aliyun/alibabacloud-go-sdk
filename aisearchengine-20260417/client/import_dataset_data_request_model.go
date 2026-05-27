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
	// example:
	//
	// 730
	DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
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
