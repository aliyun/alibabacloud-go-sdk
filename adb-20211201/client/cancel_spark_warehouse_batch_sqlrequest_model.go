// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSparkWarehouseBatchSQLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgency(v string) *CancelSparkWarehouseBatchSQLRequest
	GetAgency() *string
	SetDBClusterId(v string) *CancelSparkWarehouseBatchSQLRequest
	GetDBClusterId() *string
	SetQueryId(v string) *CancelSparkWarehouseBatchSQLRequest
	GetQueryId() *string
}

type CancelSparkWarehouseBatchSQLRequest struct {
	// The name of the client, which can be up to 16 characters in length. Specify a descriptive name that makes it easy to identify.
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
	// The query ID of the Spark SQL statement.
	//
	// This parameter is required.
	//
	// example:
	//
	// sq202501011001s****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s CancelSparkWarehouseBatchSQLRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelSparkWarehouseBatchSQLRequest) GoString() string {
	return s.String()
}

func (s *CancelSparkWarehouseBatchSQLRequest) GetAgency() *string {
	return s.Agency
}

func (s *CancelSparkWarehouseBatchSQLRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CancelSparkWarehouseBatchSQLRequest) GetQueryId() *string {
	return s.QueryId
}

func (s *CancelSparkWarehouseBatchSQLRequest) SetAgency(v string) *CancelSparkWarehouseBatchSQLRequest {
	s.Agency = &v
	return s
}

func (s *CancelSparkWarehouseBatchSQLRequest) SetDBClusterId(v string) *CancelSparkWarehouseBatchSQLRequest {
	s.DBClusterId = &v
	return s
}

func (s *CancelSparkWarehouseBatchSQLRequest) SetQueryId(v string) *CancelSparkWarehouseBatchSQLRequest {
	s.QueryId = &v
	return s
}

func (s *CancelSparkWarehouseBatchSQLRequest) Validate() error {
	return dara.Validate(s)
}
