// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataSourceConnectivityByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CheckDataSourceConnectivityByIdRequest
	GetId() *int64
	SetOpTenantId(v int64) *CheckDataSourceConnectivityByIdRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CheckDataSourceConnectivityByIdRequest
	GetOpUserId() *string
}

type CheckDataSourceConnectivityByIdRequest struct {
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s CheckDataSourceConnectivityByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityByIdRequest) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityByIdRequest) GetId() *int64 {
	return s.Id
}

func (s *CheckDataSourceConnectivityByIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CheckDataSourceConnectivityByIdRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CheckDataSourceConnectivityByIdRequest) SetId(v int64) *CheckDataSourceConnectivityByIdRequest {
	s.Id = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdRequest) SetOpTenantId(v int64) *CheckDataSourceConnectivityByIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdRequest) SetOpUserId(v string) *CheckDataSourceConnectivityByIdRequest {
	s.OpUserId = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdRequest) Validate() error {
	return dara.Validate(s)
}
