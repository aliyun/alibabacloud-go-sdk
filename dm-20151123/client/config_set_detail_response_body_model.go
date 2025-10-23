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
	Detail *ConfigSetDetailResponseBodyDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
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
	// example:
	//
	// xxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// xxx
	Id     *string                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	IpPool *ConfigSetDetailResponseBodyDetailIpPool `json:"IpPool,omitempty" xml:"IpPool,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *ConfigSetDetailResponseBodyDetail) GetName() *string {
	return s.Name
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

func (s *ConfigSetDetailResponseBodyDetail) SetName(v string) *ConfigSetDetailResponseBodyDetail {
	s.Name = &v
	return s
}

func (s *ConfigSetDetailResponseBodyDetail) Validate() error {
	if s.IpPool != nil {
		if err := s.IpPool.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigSetDetailResponseBodyDetailIpPool struct {
	// example:
	//
	// xxx
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
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
