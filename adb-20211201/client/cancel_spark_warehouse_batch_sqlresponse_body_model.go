// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSparkWarehouseBatchSQLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SparkBatchSQL) *CancelSparkWarehouseBatchSQLResponseBody
	GetData() *SparkBatchSQL
	SetRequestId(v string) *CancelSparkWarehouseBatchSQLResponseBody
	GetRequestId() *string
}

type CancelSparkWarehouseBatchSQLResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// 123
	Data *SparkBatchSQL `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 771C5FAA-530F-52F7-B84D-EBAD45***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelSparkWarehouseBatchSQLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelSparkWarehouseBatchSQLResponseBody) GoString() string {
	return s.String()
}

func (s *CancelSparkWarehouseBatchSQLResponseBody) GetData() *SparkBatchSQL {
	return s.Data
}

func (s *CancelSparkWarehouseBatchSQLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelSparkWarehouseBatchSQLResponseBody) SetData(v *SparkBatchSQL) *CancelSparkWarehouseBatchSQLResponseBody {
	s.Data = v
	return s
}

func (s *CancelSparkWarehouseBatchSQLResponseBody) SetRequestId(v string) *CancelSparkWarehouseBatchSQLResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelSparkWarehouseBatchSQLResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
