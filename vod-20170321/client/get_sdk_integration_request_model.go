// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSdkIntegrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *GetSdkIntegrationRequest
	GetBusinessType() *string
	SetIntegrationType(v string) *GetSdkIntegrationRequest
	GetIntegrationType() *string
	SetPlatform(v string) *GetSdkIntegrationRequest
	GetPlatform() *string
	SetProduct(v string) *GetSdkIntegrationRequest
	GetProduct() *string
	SetSdkCodeId(v int32) *GetSdkIntegrationRequest
	GetSdkCodeId() *int32
}

type GetSdkIntegrationRequest struct {
	BusinessType    *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	IntegrationType *string `json:"IntegrationType,omitempty" xml:"IntegrationType,omitempty"`
	Platform        *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Product         *string `json:"Product,omitempty" xml:"Product,omitempty"`
	SdkCodeId       *int32  `json:"SdkCodeId,omitempty" xml:"SdkCodeId,omitempty"`
}

func (s GetSdkIntegrationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSdkIntegrationRequest) GoString() string {
	return s.String()
}

func (s *GetSdkIntegrationRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetSdkIntegrationRequest) GetIntegrationType() *string {
	return s.IntegrationType
}

func (s *GetSdkIntegrationRequest) GetPlatform() *string {
	return s.Platform
}

func (s *GetSdkIntegrationRequest) GetProduct() *string {
	return s.Product
}

func (s *GetSdkIntegrationRequest) GetSdkCodeId() *int32 {
	return s.SdkCodeId
}

func (s *GetSdkIntegrationRequest) SetBusinessType(v string) *GetSdkIntegrationRequest {
	s.BusinessType = &v
	return s
}

func (s *GetSdkIntegrationRequest) SetIntegrationType(v string) *GetSdkIntegrationRequest {
	s.IntegrationType = &v
	return s
}

func (s *GetSdkIntegrationRequest) SetPlatform(v string) *GetSdkIntegrationRequest {
	s.Platform = &v
	return s
}

func (s *GetSdkIntegrationRequest) SetProduct(v string) *GetSdkIntegrationRequest {
	s.Product = &v
	return s
}

func (s *GetSdkIntegrationRequest) SetSdkCodeId(v int32) *GetSdkIntegrationRequest {
	s.SdkCodeId = &v
	return s
}

func (s *GetSdkIntegrationRequest) Validate() error {
	return dara.Validate(s)
}
