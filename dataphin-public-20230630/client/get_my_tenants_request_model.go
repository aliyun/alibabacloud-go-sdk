// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMyTenantsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureCodeList(v []*string) *GetMyTenantsRequest
	GetFeatureCodeList() []*string
	SetOpTenantId(v int64) *GetMyTenantsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetMyTenantsRequest
	GetOpUserId() *string
}

type GetMyTenantsRequest struct {
	// The feature codes owned by the tenant.
	FeatureCodeList []*string `json:"FeatureCodeList,omitempty" xml:"FeatureCodeList,omitempty" type:"Repeated"`
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

func (s GetMyTenantsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMyTenantsRequest) GoString() string {
	return s.String()
}

func (s *GetMyTenantsRequest) GetFeatureCodeList() []*string {
	return s.FeatureCodeList
}

func (s *GetMyTenantsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetMyTenantsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetMyTenantsRequest) SetFeatureCodeList(v []*string) *GetMyTenantsRequest {
	s.FeatureCodeList = v
	return s
}

func (s *GetMyTenantsRequest) SetOpTenantId(v int64) *GetMyTenantsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetMyTenantsRequest) SetOpUserId(v string) *GetMyTenantsRequest {
	s.OpUserId = &v
	return s
}

func (s *GetMyTenantsRequest) Validate() error {
	return dara.Validate(s)
}
