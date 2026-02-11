// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegrationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductType(v string) *GetIntegrationTokenRequest
	GetProductType() *string
	SetRegionId(v string) *GetIntegrationTokenRequest
	GetRegionId() *string
}

type GetIntegrationTokenRequest struct {
	// This parameter is required.
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetIntegrationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationTokenRequest) GoString() string {
	return s.String()
}

func (s *GetIntegrationTokenRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetIntegrationTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetIntegrationTokenRequest) SetProductType(v string) *GetIntegrationTokenRequest {
	s.ProductType = &v
	return s
}

func (s *GetIntegrationTokenRequest) SetRegionId(v string) *GetIntegrationTokenRequest {
	s.RegionId = &v
	return s
}

func (s *GetIntegrationTokenRequest) Validate() error {
	return dara.Validate(s)
}
