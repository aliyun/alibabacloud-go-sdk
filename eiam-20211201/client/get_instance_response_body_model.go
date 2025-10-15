// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstance(v *GetInstanceResponseBodyInstance) *GetInstanceResponseBody
	GetInstance() *GetInstanceResponseBodyInstance
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
}

type GetInstanceResponseBody struct {
	// The details of the instance.
	Instance *GetInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetInstance() *GetInstanceResponseBodyInstance {
	return s.Instance
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) SetInstance(v *GetInstanceResponseBodyInstance) *GetInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceResponseBodyInstance struct {
	// The time when the instance was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1550115455000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The default endpoint of the instance.
	DefaultEndpoint *GetInstanceResponseBodyInstanceDefaultEndpoint `json:"DefaultEndpoint,omitempty" xml:"DefaultEndpoint,omitempty" type:"Struct"`
	// The description of the instance.
	//
	// example:
	//
	// test_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The default domain of the instance.
	DomainConfig *GetInstanceResponseBodyInstanceDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Struct"`
	// The outbound public CIDR blocks of the instance. For example, when you synchronize Active Directory (AD) accounts, the IDaaS EIAM instance accesses your AD service by using the outbound public CIDR blocks.
	EgressAddresses []*string `json:"EgressAddresses,omitempty" xml:"EgressAddresses,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// idaas_abt3pfwojojcq323si6g5xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- creating
	//
	// 	- running
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstance) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetInstanceResponseBodyInstance) GetDefaultEndpoint() *GetInstanceResponseBodyInstanceDefaultEndpoint {
	return s.DefaultEndpoint
}

func (s *GetInstanceResponseBodyInstance) GetDescription() *string {
	return s.Description
}

func (s *GetInstanceResponseBodyInstance) GetDomainConfig() *GetInstanceResponseBodyInstanceDomainConfig {
	return s.DomainConfig
}

func (s *GetInstanceResponseBodyInstance) GetEgressAddresses() []*string {
	return s.EgressAddresses
}

func (s *GetInstanceResponseBodyInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyInstance) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyInstance) SetCreateTime(v int64) *GetInstanceResponseBodyInstance {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDefaultEndpoint(v *GetInstanceResponseBodyInstanceDefaultEndpoint) *GetInstanceResponseBodyInstance {
	s.DefaultEndpoint = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDescription(v string) *GetInstanceResponseBodyInstance {
	s.Description = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDomainConfig(v *GetInstanceResponseBodyInstanceDomainConfig) *GetInstanceResponseBodyInstance {
	s.DomainConfig = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEgressAddresses(v []*string) *GetInstanceResponseBodyInstance {
	s.EgressAddresses = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceId(v string) *GetInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetStatus(v string) *GetInstanceResponseBodyInstance {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) Validate() error {
	if s.DefaultEndpoint != nil {
		if err := s.DefaultEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.DomainConfig != nil {
		if err := s.DomainConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceResponseBodyInstanceDefaultEndpoint struct {
	// The endpoint of the instance.
	//
	// example:
	//
	// example-xxx.aliyunidaas.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The status of the endpoint. Valid values:
	//
	// 	- resolved
	//
	// 	- unresolved
	//
	// example:
	//
	// resolved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceResponseBodyInstanceDefaultEndpoint) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceDefaultEndpoint) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceDefaultEndpoint) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetInstanceResponseBodyInstanceDefaultEndpoint) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyInstanceDefaultEndpoint) SetEndpoint(v string) *GetInstanceResponseBodyInstanceDefaultEndpoint {
	s.Endpoint = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceDefaultEndpoint) SetStatus(v string) *GetInstanceResponseBodyInstanceDefaultEndpoint {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceDefaultEndpoint) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyInstanceDomainConfig struct {
	// The default domain of the instance.
	//
	// example:
	//
	// example-xxx.example.com
	DefaultDomain *string `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	// The init domain of the instance.
	//
	// example:
	//
	// example-xxx.aliyunidaas.com
	InitDomain *string `json:"InitDomain,omitempty" xml:"InitDomain,omitempty"`
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	InitDomainAutoRedirectStatus *string `json:"InitDomainAutoRedirectStatus,omitempty" xml:"InitDomainAutoRedirectStatus,omitempty"`
}

func (s GetInstanceResponseBodyInstanceDomainConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceDomainConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) GetDefaultDomain() *string {
	return s.DefaultDomain
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) GetInitDomain() *string {
	return s.InitDomain
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) GetInitDomainAutoRedirectStatus() *string {
	return s.InitDomainAutoRedirectStatus
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) SetDefaultDomain(v string) *GetInstanceResponseBodyInstanceDomainConfig {
	s.DefaultDomain = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) SetInitDomain(v string) *GetInstanceResponseBodyInstanceDomainConfig {
	s.InitDomain = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) SetInitDomainAutoRedirectStatus(v string) *GetInstanceResponseBodyInstanceDomainConfig {
	s.InitDomainAutoRedirectStatus = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) Validate() error {
	return dara.Validate(s)
}
