// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceMyProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetDataServiceMyProjectsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetDataServiceMyProjectsRequest
	GetOpUserId() *string
}

type GetDataServiceMyProjectsRequest struct {
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

func (s GetDataServiceMyProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceMyProjectsRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceMyProjectsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataServiceMyProjectsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetDataServiceMyProjectsRequest) SetOpTenantId(v int64) *GetDataServiceMyProjectsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceMyProjectsRequest) SetOpUserId(v string) *GetDataServiceMyProjectsRequest {
	s.OpUserId = &v
	return s
}

func (s *GetDataServiceMyProjectsRequest) Validate() error {
	return dara.Validate(s)
}
