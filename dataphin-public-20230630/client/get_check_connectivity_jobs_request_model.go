// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckConnectivityJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v int64) *GetCheckConnectivityJobsRequest
	GetDataSourceId() *int64
	SetOpTenantId(v int64) *GetCheckConnectivityJobsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetCheckConnectivityJobsRequest
	GetOpUserId() *string
}

type GetCheckConnectivityJobsRequest struct {
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 462935472785
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetCheckConnectivityJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckConnectivityJobsRequest) GoString() string {
	return s.String()
}

func (s *GetCheckConnectivityJobsRequest) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *GetCheckConnectivityJobsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetCheckConnectivityJobsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetCheckConnectivityJobsRequest) SetDataSourceId(v int64) *GetCheckConnectivityJobsRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetCheckConnectivityJobsRequest) SetOpTenantId(v int64) *GetCheckConnectivityJobsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetCheckConnectivityJobsRequest) SetOpUserId(v string) *GetCheckConnectivityJobsRequest {
	s.OpUserId = &v
	return s
}

func (s *GetCheckConnectivityJobsRequest) Validate() error {
	return dara.Validate(s)
}
