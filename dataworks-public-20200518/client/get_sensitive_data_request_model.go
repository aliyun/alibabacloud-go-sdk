// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSensitiveDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetSensitiveDataRequest
	GetName() *string
	SetPageNo(v int32) *GetSensitiveDataRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetSensitiveDataRequest
	GetPageSize() *int32
}

type GetSensitiveDataRequest struct {
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
	// The sample value shows the parameters configured to query the access records of the sensitive data in the abc database of the Hologres instance ABC. You must configure the parameters based on the compute engine that you use in your business.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ {"dbType":"hologres","instanceName":"ABC","databaseName":"abc"}, {"dbType":"ODPS.ODPS","projectName":"adbc"} ]
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetSensitiveDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveDataRequest) GoString() string {
	return s.String()
}

func (s *GetSensitiveDataRequest) GetName() *string {
	return s.Name
}

func (s *GetSensitiveDataRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetSensitiveDataRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSensitiveDataRequest) SetName(v string) *GetSensitiveDataRequest {
	s.Name = &v
	return s
}

func (s *GetSensitiveDataRequest) SetPageNo(v int32) *GetSensitiveDataRequest {
	s.PageNo = &v
	return s
}

func (s *GetSensitiveDataRequest) SetPageSize(v int32) *GetSensitiveDataRequest {
	s.PageSize = &v
	return s
}

func (s *GetSensitiveDataRequest) Validate() error {
	return dara.Validate(s)
}
