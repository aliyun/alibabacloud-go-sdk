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
}

type GetDataServiceMyProjectsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *GetDataServiceMyProjectsRequest) SetOpTenantId(v int64) *GetDataServiceMyProjectsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceMyProjectsRequest) Validate() error {
	return dara.Validate(s)
}
