// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCLICommandShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApi(v string) *GenerateCLICommandShrinkRequest
	GetApi() *string
	SetApiParamsShrink(v string) *GenerateCLICommandShrinkRequest
	GetApiParamsShrink() *string
	SetApiVersion(v string) *GenerateCLICommandShrinkRequest
	GetApiVersion() *string
	SetProduct(v string) *GenerateCLICommandShrinkRequest
	GetProduct() *string
	SetRegionId(v string) *GenerateCLICommandShrinkRequest
	GetRegionId() *string
}

type GenerateCLICommandShrinkRequest struct {
	// This parameter is required.
	Api             *string `json:"api,omitempty" xml:"api,omitempty"`
	ApiParamsShrink *string `json:"apiParams,omitempty" xml:"apiParams,omitempty"`
	// This parameter is required.
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// This parameter is required.
	Product  *string `json:"product,omitempty" xml:"product,omitempty"`
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GenerateCLICommandShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCLICommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateCLICommandShrinkRequest) GetApi() *string {
	return s.Api
}

func (s *GenerateCLICommandShrinkRequest) GetApiParamsShrink() *string {
	return s.ApiParamsShrink
}

func (s *GenerateCLICommandShrinkRequest) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GenerateCLICommandShrinkRequest) GetProduct() *string {
	return s.Product
}

func (s *GenerateCLICommandShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateCLICommandShrinkRequest) SetApi(v string) *GenerateCLICommandShrinkRequest {
	s.Api = &v
	return s
}

func (s *GenerateCLICommandShrinkRequest) SetApiParamsShrink(v string) *GenerateCLICommandShrinkRequest {
	s.ApiParamsShrink = &v
	return s
}

func (s *GenerateCLICommandShrinkRequest) SetApiVersion(v string) *GenerateCLICommandShrinkRequest {
	s.ApiVersion = &v
	return s
}

func (s *GenerateCLICommandShrinkRequest) SetProduct(v string) *GenerateCLICommandShrinkRequest {
	s.Product = &v
	return s
}

func (s *GenerateCLICommandShrinkRequest) SetRegionId(v string) *GenerateCLICommandShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateCLICommandShrinkRequest) Validate() error {
	return dara.Validate(s)
}
