// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v *ConfigSetDetailResponseBodyDetail) *ConfigSetDetailResponseBody
	GetDetail() *ConfigSetDetailResponseBodyDetail
	SetRequestId(v string) *ConfigSetDetailResponseBody
	GetRequestId() *string
}

type ConfigSetDetailResponseBody struct {
	// The configuration set information.
	Detail *ConfigSetDetailResponseBodyDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigSetDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigSetDetailResponseBody) GetDetail() *ConfigSetDetailResponseBodyDetail {
	return s.Detail
}

func (s *ConfigSetDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigSetDetailResponseBody) SetDetail(v *ConfigSetDetailResponseBodyDetail) *ConfigSetDetailResponseBody {
	s.Detail = v
	return s
}

func (s *ConfigSetDetailResponseBody) SetRequestId(v string) *ConfigSetDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigSetDetailResponseBody) Validate() error {
	if s.Detail != nil {
		if err := s.Detail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigSetDetailResponseBodyDetail struct {
	// The description.
	//
	// example:
	//
	// xxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The configuration set ID.
	//
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The associated IP pool.
	IpPool                 *ConfigSetDetailResponseBodyDetailIpPool `json:"IpPool,omitempty" xml:"IpPool,omitempty" type:"Struct"`
	IsPublicChannelBackoff *bool                                    `json:"IsPublicChannelBackoff,omitempty" xml:"IsPublicChannelBackoff,omitempty"`
	// The configuration set name.
	//
	// example:
	//
	// xxx
	Name             *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	ValidationOption *ConfigSetDetailResponseBodyDetailValidationOption `json:"ValidationOption,omitempty" xml:"ValidationOption,omitempty" type:"Struct"`
}

func (s ConfigSetDetailResponseBodyDetail) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetDetailResponseBodyDetail) GoString() string {
	return s.String()
}

func (s *ConfigSetDetailResponseBodyDetail) GetDescription() *string {
	return s.Description
}

func (s *ConfigSetDetailResponseBodyDetail) GetId() *string {
	return s.Id
}

func (s *ConfigSetDetailResponseBodyDetail) GetIpPool() *ConfigSetDetailResponseBodyDetailIpPool {
	return s.IpPool
}

func (s *ConfigSetDetailResponseBodyDetail) GetIsPublicChannelBackoff() *bool {
	return s.IsPublicChannelBackoff
}

func (s *ConfigSetDetailResponseBodyDetail) GetName() *string {
	return s.Name
}

func (s *ConfigSetDetailResponseBodyDetail) GetValidationOption() *ConfigSetDetailResponseBodyDetailValidationOption {
	return s.ValidationOption
}

func (s *ConfigSetDetailResponseBodyDetail) SetDescription(v string) *ConfigSetDetailResponseBodyDetail {
	s.Description = &v
	return s
}

func (s *ConfigSetDetailResponseBodyDetail) SetId(v string) *ConfigSetDetailResponseBodyDetail {
	s.Id = &v
	return s
}

func (s *ConfigSetDetailResponseBodyDetail) SetIpPool(v *ConfigSetDetailResponseBodyDetailIpPool) *ConfigSetDetailResponseBodyDetail {
	s.IpPool = v
	return s
}

func (s *ConfigSetDetailResponseBodyDetail) SetIsPublicChannelBackoff(v bool) *ConfigSetDetailResponseBodyDetail {
	s.IsPublicChannelBackoff = &v
	return s
}

func (s *ConfigSetDetailResponseBodyDetail) SetName(v string) *ConfigSetDetailResponseBodyDetail {
	s.Name = &v
	return s
}

func (s *ConfigSetDetailResponseBodyDetail) SetValidationOption(v *ConfigSetDetailResponseBodyDetailValidationOption) *ConfigSetDetailResponseBodyDetail {
	s.ValidationOption = v
	return s
}

func (s *ConfigSetDetailResponseBodyDetail) Validate() error {
	if s.IpPool != nil {
		if err := s.IpPool.Validate(); err != nil {
			return err
		}
	}
	if s.ValidationOption != nil {
		if err := s.ValidationOption.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigSetDetailResponseBodyDetailIpPool struct {
	// The associated IP pool ID.
	//
	// example:
	//
	// xxx
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	// The associated IP pool name.
	//
	// example:
	//
	// xxx
	IpPoolName *string `json:"IpPoolName,omitempty" xml:"IpPoolName,omitempty"`
}

func (s ConfigSetDetailResponseBodyDetailIpPool) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetDetailResponseBodyDetailIpPool) GoString() string {
	return s.String()
}

func (s *ConfigSetDetailResponseBodyDetailIpPool) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *ConfigSetDetailResponseBodyDetailIpPool) GetIpPoolName() *string {
	return s.IpPoolName
}

func (s *ConfigSetDetailResponseBodyDetailIpPool) SetIpPoolId(v string) *ConfigSetDetailResponseBodyDetailIpPool {
	s.IpPoolId = &v
	return s
}

func (s *ConfigSetDetailResponseBodyDetailIpPool) SetIpPoolName(v string) *ConfigSetDetailResponseBodyDetailIpPool {
	s.IpPoolName = &v
	return s
}

func (s *ConfigSetDetailResponseBodyDetailIpPool) Validate() error {
	return dara.Validate(s)
}

type ConfigSetDetailResponseBodyDetailValidationOption struct {
	Enabled                *bool     `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ForbiddenStatusList    []*string `json:"ForbiddenStatusList,omitempty" xml:"ForbiddenStatusList,omitempty" type:"Repeated"`
	ForbiddenSubStatusList []*string `json:"ForbiddenSubStatusList,omitempty" xml:"ForbiddenSubStatusList,omitempty" type:"Repeated"`
}

func (s ConfigSetDetailResponseBodyDetailValidationOption) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetDetailResponseBodyDetailValidationOption) GoString() string {
	return s.String()
}

func (s *ConfigSetDetailResponseBodyDetailValidationOption) GetEnabled() *bool {
	return s.Enabled
}

func (s *ConfigSetDetailResponseBodyDetailValidationOption) GetForbiddenStatusList() []*string {
	return s.ForbiddenStatusList
}

func (s *ConfigSetDetailResponseBodyDetailValidationOption) GetForbiddenSubStatusList() []*string {
	return s.ForbiddenSubStatusList
}

func (s *ConfigSetDetailResponseBodyDetailValidationOption) SetEnabled(v bool) *ConfigSetDetailResponseBodyDetailValidationOption {
	s.Enabled = &v
	return s
}

func (s *ConfigSetDetailResponseBodyDetailValidationOption) SetForbiddenStatusList(v []*string) *ConfigSetDetailResponseBodyDetailValidationOption {
	s.ForbiddenStatusList = v
	return s
}

func (s *ConfigSetDetailResponseBodyDetailValidationOption) SetForbiddenSubStatusList(v []*string) *ConfigSetDetailResponseBodyDetailValidationOption {
	s.ForbiddenSubStatusList = v
	return s
}

func (s *ConfigSetDetailResponseBodyDetailValidationOption) Validate() error {
	return dara.Validate(s)
}
