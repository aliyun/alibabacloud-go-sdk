// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetServiceEndpointResponseBody
	GetCode() *string
	SetData(v *GetServiceEndpointResponseBodyData) *GetServiceEndpointResponseBody
	GetData() *GetServiceEndpointResponseBodyData
	SetMessage(v string) *GetServiceEndpointResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetServiceEndpointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetServiceEndpointResponseBody
	GetSuccess() *bool
}

type GetServiceEndpointResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetServiceEndpointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// req-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetServiceEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetServiceEndpointResponseBody) GetData() *GetServiceEndpointResponseBodyData {
	return s.Data
}

func (s *GetServiceEndpointResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceEndpointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServiceEndpointResponseBody) SetCode(v string) *GetServiceEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceEndpointResponseBody) SetData(v *GetServiceEndpointResponseBodyData) *GetServiceEndpointResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceEndpointResponseBody) SetMessage(v string) *GetServiceEndpointResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceEndpointResponseBody) SetRequestId(v string) *GetServiceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceEndpointResponseBody) SetSuccess(v bool) *GetServiceEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceEndpointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceEndpointResponseBodyData struct {
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	Component      *string `json:"Component,omitempty" xml:"Component,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainType     *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EndpointId     *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointName   *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NetworkType    *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime     *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetServiceEndpointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointResponseBodyData) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *GetServiceEndpointResponseBodyData) GetComponent() *string {
	return s.Component
}

func (s *GetServiceEndpointResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetServiceEndpointResponseBodyData) GetDomain() *string {
	return s.Domain
}

func (s *GetServiceEndpointResponseBodyData) GetDomainType() *string {
	return s.DomainType
}

func (s *GetServiceEndpointResponseBodyData) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetServiceEndpointResponseBodyData) GetEndpointName() *string {
	return s.EndpointName
}

func (s *GetServiceEndpointResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetServiceEndpointResponseBodyData) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetServiceEndpointResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceEndpointResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetServiceEndpointResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetServiceEndpointResponseBodyData) SetCertIdentifier(v string) *GetServiceEndpointResponseBodyData {
	s.CertIdentifier = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetComponent(v string) *GetServiceEndpointResponseBodyData {
	s.Component = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetCreateTime(v string) *GetServiceEndpointResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetDomain(v string) *GetServiceEndpointResponseBodyData {
	s.Domain = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetDomainType(v string) *GetServiceEndpointResponseBodyData {
	s.DomainType = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetEndpointId(v string) *GetServiceEndpointResponseBodyData {
	s.EndpointId = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetEndpointName(v string) *GetServiceEndpointResponseBodyData {
	s.EndpointName = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetInstanceId(v string) *GetServiceEndpointResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetNetworkType(v string) *GetServiceEndpointResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetRegionId(v string) *GetServiceEndpointResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetStatus(v string) *GetServiceEndpointResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetUpdateTime(v string) *GetServiceEndpointResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) Validate() error {
	return dara.Validate(s)
}
