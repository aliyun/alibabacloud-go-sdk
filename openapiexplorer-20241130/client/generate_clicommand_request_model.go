// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCLICommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApi(v string) *GenerateCLICommandRequest
	GetApi() *string
	SetApiParams(v map[string]interface{}) *GenerateCLICommandRequest
	GetApiParams() map[string]interface{}
	SetApiVersion(v string) *GenerateCLICommandRequest
	GetApiVersion() *string
	SetProduct(v string) *GenerateCLICommandRequest
	GetProduct() *string
	SetRegionId(v string) *GenerateCLICommandRequest
	GetRegionId() *string
}

type GenerateCLICommandRequest struct {
	// This parameter is required.
	Api       *string                `json:"api,omitempty" xml:"api,omitempty"`
	ApiParams map[string]interface{} `json:"apiParams,omitempty" xml:"apiParams,omitempty"`
	// This parameter is required.
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// This parameter is required.
	Product  *string `json:"product,omitempty" xml:"product,omitempty"`
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GenerateCLICommandRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCLICommandRequest) GoString() string {
	return s.String()
}

func (s *GenerateCLICommandRequest) GetApi() *string {
	return s.Api
}

func (s *GenerateCLICommandRequest) GetApiParams() map[string]interface{} {
	return s.ApiParams
}

func (s *GenerateCLICommandRequest) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GenerateCLICommandRequest) GetProduct() *string {
	return s.Product
}

func (s *GenerateCLICommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateCLICommandRequest) SetApi(v string) *GenerateCLICommandRequest {
	s.Api = &v
	return s
}

func (s *GenerateCLICommandRequest) SetApiParams(v map[string]interface{}) *GenerateCLICommandRequest {
	s.ApiParams = v
	return s
}

func (s *GenerateCLICommandRequest) SetApiVersion(v string) *GenerateCLICommandRequest {
	s.ApiVersion = &v
	return s
}

func (s *GenerateCLICommandRequest) SetProduct(v string) *GenerateCLICommandRequest {
	s.Product = &v
	return s
}

func (s *GenerateCLICommandRequest) SetRegionId(v string) *GenerateCLICommandRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateCLICommandRequest) Validate() error {
	return dara.Validate(s)
}
