// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGateways(v []*ListGatewayResponseBodyGateways) *ListGatewayResponseBody
	GetGateways() []*ListGatewayResponseBodyGateways
	SetPageNumber(v int32) *ListGatewayResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListGatewayResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListGatewayResponseBody
	GetTotalCount() *int64
}

type ListGatewayResponseBody struct {
	// The private gateways.
	Gateways []*ListGatewayResponseBodyGateways `json:"Gateways,omitempty" xml:"Gateways,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of private gateways returned.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBody) GetGateways() []*ListGatewayResponseBodyGateways {
	return s.Gateways
}

func (s *ListGatewayResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListGatewayResponseBody) SetGateways(v []*ListGatewayResponseBodyGateways) *ListGatewayResponseBody {
	s.Gateways = v
	return s
}

func (s *ListGatewayResponseBody) SetPageNumber(v int32) *ListGatewayResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayResponseBody) SetPageSize(v int32) *ListGatewayResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListGatewayResponseBody) SetRequestId(v string) *ListGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayResponseBody) SetTotalCount(v int64) *ListGatewayResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListGatewayResponseBodyGateways struct {
	// The billing method. Valid values:
	//
	// 	- PrePaid: subscription.
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the private gateway was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-05-19T14:19:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The private gateway alias.
	//
	// example:
	//
	// mygateway1
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// The type of instances used for the private gateway.
	//
	// example:
	//
	// 2c4g
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The public endpoint.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******-1801786532******.cn-wulanchabu.pai-eas.aliyuncs.com
	InternetDomain *string `json:"InternetDomain,omitempty" xml:"InternetDomain,omitempty"`
	// Indicates whether Internet access is enabled.
	//
	// example:
	//
	// true
	InternetEnabled *bool `json:"InternetEnabled,omitempty" xml:"InternetEnabled,omitempty"`
	// The internal endpoint.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******-1801786532******-vpc.cn-wulanchabu.pai-eas.aliyuncs.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// Indicates whether it is the default private gateway.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The number of nodes in the private gateway.
	//
	// example:
	//
	// 2
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// Specifies whether to enable HTTP to HTTPS redirection.
	//
	// example:
	//
	// true
	SSLRedirectionEnabled *bool `json:"SSLRedirectionEnabled,omitempty" xml:"SSLRedirectionEnabled,omitempty"`
	// The state of the private gateway.
	//
	// Valid values:
	//
	// 	- Creating
	//
	// 	- Stopped
	//
	// 	- Failed
	//
	// 	- Running
	//
	// 	- Deleted
	//
	// 	- Deleting
	//
	// 	- Waiting
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the private gateway was updated. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-02-24T11:52:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListGatewayResponseBodyGateways) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBodyGateways) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyGateways) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListGatewayResponseBodyGateways) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListGatewayResponseBodyGateways) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListGatewayResponseBodyGateways) GetGatewayName() *string {
	return s.GatewayName
}

func (s *ListGatewayResponseBodyGateways) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListGatewayResponseBodyGateways) GetInternetDomain() *string {
	return s.InternetDomain
}

func (s *ListGatewayResponseBodyGateways) GetInternetEnabled() *bool {
	return s.InternetEnabled
}

func (s *ListGatewayResponseBodyGateways) GetIntranetDomain() *string {
	return s.IntranetDomain
}

func (s *ListGatewayResponseBodyGateways) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListGatewayResponseBodyGateways) GetReplicas() *int32 {
	return s.Replicas
}

func (s *ListGatewayResponseBodyGateways) GetSSLRedirectionEnabled() *bool {
	return s.SSLRedirectionEnabled
}

func (s *ListGatewayResponseBodyGateways) GetStatus() *string {
	return s.Status
}

func (s *ListGatewayResponseBodyGateways) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListGatewayResponseBodyGateways) SetChargeType(v string) *ListGatewayResponseBodyGateways {
	s.ChargeType = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetCreateTime(v string) *ListGatewayResponseBodyGateways {
	s.CreateTime = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetGatewayId(v string) *ListGatewayResponseBodyGateways {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetGatewayName(v string) *ListGatewayResponseBodyGateways {
	s.GatewayName = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetInstanceType(v string) *ListGatewayResponseBodyGateways {
	s.InstanceType = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetInternetDomain(v string) *ListGatewayResponseBodyGateways {
	s.InternetDomain = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetInternetEnabled(v bool) *ListGatewayResponseBodyGateways {
	s.InternetEnabled = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetIntranetDomain(v string) *ListGatewayResponseBodyGateways {
	s.IntranetDomain = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetIsDefault(v bool) *ListGatewayResponseBodyGateways {
	s.IsDefault = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetReplicas(v int32) *ListGatewayResponseBodyGateways {
	s.Replicas = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetSSLRedirectionEnabled(v bool) *ListGatewayResponseBodyGateways {
	s.SSLRedirectionEnabled = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetStatus(v string) *ListGatewayResponseBodyGateways {
	s.Status = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetUpdateTime(v string) *ListGatewayResponseBodyGateways {
	s.UpdateTime = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) Validate() error {
	return dara.Validate(s)
}
