// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkWarehouseBatchSQLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SparkBatchSQL) *GetSparkWarehouseBatchSQLResponseBody
	GetData() *SparkBatchSQL
	SetRequestId(v string) *GetSparkWarehouseBatchSQLResponseBody
	GetRequestId() *string
}

type GetSparkWarehouseBatchSQLResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// {\\"DBClusterId\\": \\"amv-uf6485635fz852mn\\", \\"Statements\\": [{\\"Code\\": \\"show databases\\", \\"State\\": \\"FINISHED\\", \\"StatementId\\": \\"sq202501141001sh8e12e080006846-0000\\", \\"Result\\": \\"+--------------------+n
	Data *SparkBatchSQL `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3D44EDA0-A9DC-580D-9B5A-019370C075F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkWarehouseBatchSQLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkWarehouseBatchSQLResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkWarehouseBatchSQLResponseBody) GetData() *SparkBatchSQL {
	return s.Data
}

func (s *GetSparkWarehouseBatchSQLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkWarehouseBatchSQLResponseBody) SetData(v *SparkBatchSQL) *GetSparkWarehouseBatchSQLResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkWarehouseBatchSQLResponseBody) SetRequestId(v string) *GetSparkWarehouseBatchSQLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkWarehouseBatchSQLResponseBody) Validate() error {
	return dara.Validate(s)
}
