// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateServiceEndpointResponseBody
	GetCode() *string
	SetData(v *UpdateServiceEndpointResponseBodyData) *UpdateServiceEndpointResponseBody
	GetData() *UpdateServiceEndpointResponseBodyData
	SetMessage(v string) *UpdateServiceEndpointResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceEndpointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateServiceEndpointResponseBody
	GetSuccess() *bool
}

type UpdateServiceEndpointResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UpdateServiceEndpointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateServiceEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateServiceEndpointResponseBody) GetData() *UpdateServiceEndpointResponseBodyData {
	return s.Data
}

func (s *UpdateServiceEndpointResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceEndpointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateServiceEndpointResponseBody) SetCode(v string) *UpdateServiceEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateServiceEndpointResponseBody) SetData(v *UpdateServiceEndpointResponseBodyData) *UpdateServiceEndpointResponseBody {
	s.Data = v
	return s
}

func (s *UpdateServiceEndpointResponseBody) SetMessage(v string) *UpdateServiceEndpointResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceEndpointResponseBody) SetRequestId(v string) *UpdateServiceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceEndpointResponseBody) SetSuccess(v bool) *UpdateServiceEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateServiceEndpointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateServiceEndpointResponseBodyData struct {
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	Component      *string `json:"Component,omitempty" xml:"Component,omitempty"`
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainType     *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EndpointId     *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointName   *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NetworkType    *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s UpdateServiceEndpointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointResponseBodyData) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *UpdateServiceEndpointResponseBodyData) GetComponent() *string {
	return s.Component
}

func (s *UpdateServiceEndpointResponseBodyData) GetDomain() *string {
	return s.Domain
}

func (s *UpdateServiceEndpointResponseBodyData) GetDomainType() *string {
	return s.DomainType
}

func (s *UpdateServiceEndpointResponseBodyData) GetEndpointId() *string {
	return s.EndpointId
}

func (s *UpdateServiceEndpointResponseBodyData) GetEndpointName() *string {
	return s.EndpointName
}

func (s *UpdateServiceEndpointResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateServiceEndpointResponseBodyData) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateServiceEndpointResponseBodyData) SetCertIdentifier(v string) *UpdateServiceEndpointResponseBodyData {
	s.CertIdentifier = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetComponent(v string) *UpdateServiceEndpointResponseBodyData {
	s.Component = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetDomain(v string) *UpdateServiceEndpointResponseBodyData {
	s.Domain = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetDomainType(v string) *UpdateServiceEndpointResponseBodyData {
	s.DomainType = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetEndpointId(v string) *UpdateServiceEndpointResponseBodyData {
	s.EndpointId = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetEndpointName(v string) *UpdateServiceEndpointResponseBodyData {
	s.EndpointName = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetInstanceId(v string) *UpdateServiceEndpointResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetNetworkType(v string) *UpdateServiceEndpointResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) Validate() error {
	return dara.Validate(s)
}
