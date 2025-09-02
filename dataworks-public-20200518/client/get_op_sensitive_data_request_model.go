// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpSensitiveDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *GetOpSensitiveDataRequest
	GetDate() *string
	SetName(v string) *GetOpSensitiveDataRequest
	GetName() *string
	SetOpType(v string) *GetOpSensitiveDataRequest
	GetOpType() *string
	SetPageNo(v int32) *GetOpSensitiveDataRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetOpSensitiveDataRequest
	GetPageSize() *int32
}

type GetOpSensitiveDataRequest struct {
	// The date on which access records were generated. Specify the value in the yyyyMMdd format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20210116
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The parameters that you can configure to query the access records. Valid values:
	//
	// 	- dbType
	//
	// 	- instanceName
	//
	// 	- databaseName
	//
	// 	- projectName
	//
	// 	- clusterName
	//
	// The following example shows the parameters configured to query the access records of the sensitive data in the abc database of the Hologres instance ABC: [ {"dbType":"hologres","instanceName":"ABC","databaseName":"abc"} ]
	//
	// You must configure the parameters based on the compute engine that you use in your business.
	//
	// This parameter is required.
	//
	// example:
	//
	// [  {"dbType":"hologres","instanceName":"ABC","databaseName":"abc"},  {"dbType":"ODPS.ODPS","projectName":"adbc"}  ]
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation that is performed on the data. Valid values:
	//
	// 	- SQL_SELECT: specifies the data access operation. For example, execute a SELECT statement to query data.
	//
	// 	- TUNNEL_DOWNLOAD: specifies the data download operation. For example, run a Tunnel command to download data.
	//
	// example:
	//
	// SQL_SELECT
	OpType *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	// The page number. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Minimum value: 1. Maximum value: 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetOpSensitiveDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOpSensitiveDataRequest) GoString() string {
	return s.String()
}

func (s *GetOpSensitiveDataRequest) GetDate() *string {
	return s.Date
}

func (s *GetOpSensitiveDataRequest) GetName() *string {
	return s.Name
}

func (s *GetOpSensitiveDataRequest) GetOpType() *string {
	return s.OpType
}

func (s *GetOpSensitiveDataRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetOpSensitiveDataRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOpSensitiveDataRequest) SetDate(v string) *GetOpSensitiveDataRequest {
	s.Date = &v
	return s
}

func (s *GetOpSensitiveDataRequest) SetName(v string) *GetOpSensitiveDataRequest {
	s.Name = &v
	return s
}

func (s *GetOpSensitiveDataRequest) SetOpType(v string) *GetOpSensitiveDataRequest {
	s.OpType = &v
	return s
}

func (s *GetOpSensitiveDataRequest) SetPageNo(v int32) *GetOpSensitiveDataRequest {
	s.PageNo = &v
	return s
}

func (s *GetOpSensitiveDataRequest) SetPageSize(v int32) *GetOpSensitiveDataRequest {
	s.PageSize = &v
	return s
}

func (s *GetOpSensitiveDataRequest) Validate() error {
	return dara.Validate(s)
}
