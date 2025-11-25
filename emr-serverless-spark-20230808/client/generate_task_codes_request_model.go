// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTaskCodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGenNum(v int32) *GenerateTaskCodesRequest
	GetGenNum() *int32
	SetProductNamespace(v string) *GenerateTaskCodesRequest
	GetProductNamespace() *string
	SetRegionId(v string) *GenerateTaskCodesRequest
	GetRegionId() *string
}

type GenerateTaskCodesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	GenNum *int32 `json:"genNum,omitempty" xml:"genNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SS
	ProductNamespace *string `json:"productNamespace,omitempty" xml:"productNamespace,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GenerateTaskCodesRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateTaskCodesRequest) GoString() string {
	return s.String()
}

func (s *GenerateTaskCodesRequest) GetGenNum() *int32 {
	return s.GenNum
}

func (s *GenerateTaskCodesRequest) GetProductNamespace() *string {
	return s.ProductNamespace
}

func (s *GenerateTaskCodesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateTaskCodesRequest) SetGenNum(v int32) *GenerateTaskCodesRequest {
	s.GenNum = &v
	return s
}

func (s *GenerateTaskCodesRequest) SetProductNamespace(v string) *GenerateTaskCodesRequest {
	s.ProductNamespace = &v
	return s
}

func (s *GenerateTaskCodesRequest) SetRegionId(v string) *GenerateTaskCodesRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateTaskCodesRequest) Validate() error {
	return dara.Validate(s)
}
