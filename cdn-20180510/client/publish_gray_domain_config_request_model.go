// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishGrayDomainConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomCountryId(v int32) *PublishGrayDomainConfigRequest
	GetCustomCountryId() *int32
	SetCustomPercent(v int32) *PublishGrayDomainConfigRequest
	GetCustomPercent() *int32
	SetCustomProvinceId(v int32) *PublishGrayDomainConfigRequest
	GetCustomProvinceId() *int32
	SetDomainName(v string) *PublishGrayDomainConfigRequest
	GetDomainName() *string
	SetPublishMode(v string) *PublishGrayDomainConfigRequest
	GetPublishMode() *string
}

type PublishGrayDomainConfigRequest struct {
	CustomCountryId *int32 `json:"CustomCountryId,omitempty" xml:"CustomCountryId,omitempty"`
	// example:
	//
	// 15
	CustomPercent    *int32 `json:"CustomPercent,omitempty" xml:"CustomPercent,omitempty"`
	CustomProvinceId *int32 `json:"CustomProvinceId,omitempty" xml:"CustomProvinceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// publishByCustom
	PublishMode *string `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
}

func (s PublishGrayDomainConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishGrayDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *PublishGrayDomainConfigRequest) GetCustomCountryId() *int32 {
	return s.CustomCountryId
}

func (s *PublishGrayDomainConfigRequest) GetCustomPercent() *int32 {
	return s.CustomPercent
}

func (s *PublishGrayDomainConfigRequest) GetCustomProvinceId() *int32 {
	return s.CustomProvinceId
}

func (s *PublishGrayDomainConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *PublishGrayDomainConfigRequest) GetPublishMode() *string {
	return s.PublishMode
}

func (s *PublishGrayDomainConfigRequest) SetCustomCountryId(v int32) *PublishGrayDomainConfigRequest {
	s.CustomCountryId = &v
	return s
}

func (s *PublishGrayDomainConfigRequest) SetCustomPercent(v int32) *PublishGrayDomainConfigRequest {
	s.CustomPercent = &v
	return s
}

func (s *PublishGrayDomainConfigRequest) SetCustomProvinceId(v int32) *PublishGrayDomainConfigRequest {
	s.CustomProvinceId = &v
	return s
}

func (s *PublishGrayDomainConfigRequest) SetDomainName(v string) *PublishGrayDomainConfigRequest {
	s.DomainName = &v
	return s
}

func (s *PublishGrayDomainConfigRequest) SetPublishMode(v string) *PublishGrayDomainConfigRequest {
	s.PublishMode = &v
	return s
}

func (s *PublishGrayDomainConfigRequest) Validate() error {
	return dara.Validate(s)
}
