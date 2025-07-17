// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallPromClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *UninstallPromClusterRequest
	GetAliyunLang() *string
	SetClusterId(v string) *UninstallPromClusterRequest
	GetClusterId() *string
	SetRegionId(v string) *UninstallPromClusterRequest
	GetRegionId() *string
}

type UninstallPromClusterRequest struct {
	// Language environment(If left blank, defaults to zh):
	//
	// - zh
	//
	// - en
	//
	// example:
	//
	// en
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c0bad479465464e1d8c1e641b0afb****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID. Default value: cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UninstallPromClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallPromClusterRequest) GoString() string {
	return s.String()
}

func (s *UninstallPromClusterRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *UninstallPromClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UninstallPromClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UninstallPromClusterRequest) SetAliyunLang(v string) *UninstallPromClusterRequest {
	s.AliyunLang = &v
	return s
}

func (s *UninstallPromClusterRequest) SetClusterId(v string) *UninstallPromClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UninstallPromClusterRequest) SetRegionId(v string) *UninstallPromClusterRequest {
	s.RegionId = &v
	return s
}

func (s *UninstallPromClusterRequest) Validate() error {
	return dara.Validate(s)
}
