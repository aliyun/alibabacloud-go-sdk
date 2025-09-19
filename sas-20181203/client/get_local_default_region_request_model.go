// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLocalDefaultRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVendor(v string) *GetLocalDefaultRegionRequest
	GetVendor() *string
}

type GetLocalDefaultRegionRequest struct {
	// The cloud service provider. Valid values:
	//
	// 	- **Tencent**: Tencent Cloud.
	//
	// 	- **HUAWEICLOUD**: Huawei Cloud.
	//
	// 	- **Azure**: Microsoft Azure.
	//
	// 	- **AWS**: Amazon Web Services (AWS).
	//
	// example:
	//
	// Tencent
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetLocalDefaultRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLocalDefaultRegionRequest) GoString() string {
	return s.String()
}

func (s *GetLocalDefaultRegionRequest) GetVendor() *string {
	return s.Vendor
}

func (s *GetLocalDefaultRegionRequest) SetVendor(v string) *GetLocalDefaultRegionRequest {
	s.Vendor = &v
	return s
}

func (s *GetLocalDefaultRegionRequest) Validate() error {
	return dara.Validate(s)
}
