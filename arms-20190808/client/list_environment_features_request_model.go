// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentFeaturesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *ListEnvironmentFeaturesRequest
	GetAliyunLang() *string
	SetEnvironmentId(v string) *ListEnvironmentFeaturesRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *ListEnvironmentFeaturesRequest
	GetRegionId() *string
}

type ListEnvironmentFeaturesRequest struct {
	// The language. Default value: zh.
	//
	// Valid values:
	//
	// 	- en: English.
	//
	// 	- zh: Chinese.
	//
	// example:
	//
	// en
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEnvironmentFeaturesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentFeaturesRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *ListEnvironmentFeaturesRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvironmentFeaturesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvironmentFeaturesRequest) SetAliyunLang(v string) *ListEnvironmentFeaturesRequest {
	s.AliyunLang = &v
	return s
}

func (s *ListEnvironmentFeaturesRequest) SetEnvironmentId(v string) *ListEnvironmentFeaturesRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvironmentFeaturesRequest) SetRegionId(v string) *ListEnvironmentFeaturesRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnvironmentFeaturesRequest) Validate() error {
	return dara.Validate(s)
}
