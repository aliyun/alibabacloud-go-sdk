// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkWarehouseBatchSQLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgency(v string) *GetSparkWarehouseBatchSQLRequest
	GetAgency() *string
	SetDBClusterId(v string) *GetSparkWarehouseBatchSQLRequest
	GetDBClusterId() *string
	SetQueryId(v string) *GetSparkWarehouseBatchSQLRequest
	GetQueryId() *string
}

type GetSparkWarehouseBatchSQLRequest struct {
	// The client name for auxiliary logging. The value is a meaningful string of up to 16 characters.
	//
	// example:
	//
	// DMS
	Agency *string `json:"Agency,omitempty" xml:"Agency,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-uf6485635f***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The execution ID of the Spark SQL statement.
	//
	// This parameter is required.
	//
	// example:
	//
	// sq2024123*****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s GetSparkWarehouseBatchSQLRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkWarehouseBatchSQLRequest) GoString() string {
	return s.String()
}

func (s *GetSparkWarehouseBatchSQLRequest) GetAgency() *string {
	return s.Agency
}

func (s *GetSparkWarehouseBatchSQLRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkWarehouseBatchSQLRequest) GetQueryId() *string {
	return s.QueryId
}

func (s *GetSparkWarehouseBatchSQLRequest) SetAgency(v string) *GetSparkWarehouseBatchSQLRequest {
	s.Agency = &v
	return s
}

func (s *GetSparkWarehouseBatchSQLRequest) SetDBClusterId(v string) *GetSparkWarehouseBatchSQLRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkWarehouseBatchSQLRequest) SetQueryId(v string) *GetSparkWarehouseBatchSQLRequest {
	s.QueryId = &v
	return s
}

func (s *GetSparkWarehouseBatchSQLRequest) Validate() error {
	return dara.Validate(s)
}
