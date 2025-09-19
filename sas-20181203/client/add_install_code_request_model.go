// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstallCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpiredDate(v int64) *AddInstallCodeRequest
	GetExpiredDate() *int64
	SetGroupId(v int64) *AddInstallCodeRequest
	GetGroupId() *int64
	SetOnlyImage(v bool) *AddInstallCodeRequest
	GetOnlyImage() *bool
	SetOs(v string) *AddInstallCodeRequest
	GetOs() *string
	SetPrivateLinkId(v int64) *AddInstallCodeRequest
	GetPrivateLinkId() *int64
	SetProxyCluster(v string) *AddInstallCodeRequest
	GetProxyCluster() *string
	SetVendorName(v string) *AddInstallCodeRequest
	GetVendorName() *string
}

type AddInstallCodeRequest struct {
	// The validity period of the installation command. The value is a 13-digit timestamp.
	//
	// >  The installation command is valid only within the validity period. An expired installation command cannot be used to install the Security Center agent.
	//
	// example:
	//
	// 1680257463853
	ExpiredDate *int64 `json:"ExpiredDate,omitempty" xml:"ExpiredDate,omitempty"`
	// The ID of the asset group to which you want to add the asset.
	//
	// > You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to query the IDs of asset groups.
	//
	// example:
	//
	// 8076980
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Specifies whether to create an image. Default value: **false**. Valid values:
	//
	// 	- **false**: does not create an image.
	//
	// 	- **true**: creates an image.
	//
	// example:
	//
	// false
	OnlyImage *bool `json:"OnlyImage,omitempty" xml:"OnlyImage,omitempty"`
	// The operating system of the asset. Default value: **linux**. Valid values:
	//
	// 	- **linux**
	//
	// 	- **windows**
	//
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The ID of the PrivateLink endpoint.
	//
	// example:
	//
	// 72845
	PrivateLinkId *int64 `json:"PrivateLinkId,omitempty" xml:"PrivateLinkId,omitempty"`
	// The name of the proxy cluster.
	//
	// example:
	//
	// proxy_test
	ProxyCluster *string `json:"ProxyCluster,omitempty" xml:"ProxyCluster,omitempty"`
	// The name of the service provider for the asset. Default value: **ALIYUN**.
	//
	// >  You can call the [DescribeVendorList](~~DescribeVendorList~~) operation to query the names of service providers.
	//
	// example:
	//
	// ALIYUN
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
}

func (s AddInstallCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s AddInstallCodeRequest) GoString() string {
	return s.String()
}

func (s *AddInstallCodeRequest) GetExpiredDate() *int64 {
	return s.ExpiredDate
}

func (s *AddInstallCodeRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *AddInstallCodeRequest) GetOnlyImage() *bool {
	return s.OnlyImage
}

func (s *AddInstallCodeRequest) GetOs() *string {
	return s.Os
}

func (s *AddInstallCodeRequest) GetPrivateLinkId() *int64 {
	return s.PrivateLinkId
}

func (s *AddInstallCodeRequest) GetProxyCluster() *string {
	return s.ProxyCluster
}

func (s *AddInstallCodeRequest) GetVendorName() *string {
	return s.VendorName
}

func (s *AddInstallCodeRequest) SetExpiredDate(v int64) *AddInstallCodeRequest {
	s.ExpiredDate = &v
	return s
}

func (s *AddInstallCodeRequest) SetGroupId(v int64) *AddInstallCodeRequest {
	s.GroupId = &v
	return s
}

func (s *AddInstallCodeRequest) SetOnlyImage(v bool) *AddInstallCodeRequest {
	s.OnlyImage = &v
	return s
}

func (s *AddInstallCodeRequest) SetOs(v string) *AddInstallCodeRequest {
	s.Os = &v
	return s
}

func (s *AddInstallCodeRequest) SetPrivateLinkId(v int64) *AddInstallCodeRequest {
	s.PrivateLinkId = &v
	return s
}

func (s *AddInstallCodeRequest) SetProxyCluster(v string) *AddInstallCodeRequest {
	s.ProxyCluster = &v
	return s
}

func (s *AddInstallCodeRequest) SetVendorName(v string) *AddInstallCodeRequest {
	s.VendorName = &v
	return s
}

func (s *AddInstallCodeRequest) Validate() error {
	return dara.Validate(s)
}
