// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *DescribeGatewayResponseBody
	GetChargeType() *string
	SetCreateTime(v string) *DescribeGatewayResponseBody
	GetCreateTime() *string
	SetExternalClusterId(v string) *DescribeGatewayResponseBody
	GetExternalClusterId() *string
	SetGatewayId(v string) *DescribeGatewayResponseBody
	GetGatewayId() *string
	SetGatewayName(v string) *DescribeGatewayResponseBody
	GetGatewayName() *string
	SetInstanceType(v string) *DescribeGatewayResponseBody
	GetInstanceType() *string
	SetInternetDomain(v string) *DescribeGatewayResponseBody
	GetInternetDomain() *string
	SetInternetEnabled(v bool) *DescribeGatewayResponseBody
	GetInternetEnabled() *bool
	SetInternetStatus(v string) *DescribeGatewayResponseBody
	GetInternetStatus() *string
	SetIntranetDomain(v string) *DescribeGatewayResponseBody
	GetIntranetDomain() *string
	SetIntranetEnabled(v bool) *DescribeGatewayResponseBody
	GetIntranetEnabled() *bool
	SetIsDefault(v bool) *DescribeGatewayResponseBody
	GetIsDefault() *bool
	SetLabels(v []*DescribeGatewayResponseBodyLabels) *DescribeGatewayResponseBody
	GetLabels() []*DescribeGatewayResponseBodyLabels
	SetReplicas(v int32) *DescribeGatewayResponseBody
	GetReplicas() *int32
	SetRequestId(v string) *DescribeGatewayResponseBody
	GetRequestId() *string
	SetSSLRedirectionEnabled(v bool) *DescribeGatewayResponseBody
	GetSSLRedirectionEnabled() *bool
	SetStatus(v string) *DescribeGatewayResponseBody
	GetStatus() *string
	SetUpdateTime(v string) *DescribeGatewayResponseBody
	GetUpdateTime() *string
}

type DescribeGatewayResponseBody struct {
	// The billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the dedicated gateway was created. The time is in UTC.
	//
	// example:
	//
	// 2020-05-19T14:19:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The self-managed cluster ID.
	//
	// example:
	//
	// c935eadf284c14c2da57a2a13ad6******
	ExternalClusterId *string `json:"ExternalClusterId,omitempty" xml:"ExternalClusterId,omitempty"`
	// The dedicated gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The alias of the dedicated gateway.
	//
	// example:
	//
	// mygateway1
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// The instance type used by the dedicated gateway.
	//
	// example:
	//
	// 2c4g
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The public network access domain name.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******-1801786532******.cn-wulanchabu.pai-eas.aliyuncs.com
	InternetDomain *string `json:"InternetDomain,omitempty" xml:"InternetDomain,omitempty"`
	// Indicates whether public network access is enabled.
	//
	// example:
	//
	// true
	InternetEnabled *bool `json:"InternetEnabled,omitempty" xml:"InternetEnabled,omitempty"`
	// The status of public network access enablement or disablement.
	//
	// example:
	//
	// Running
	InternetStatus *string `json:"InternetStatus,omitempty" xml:"InternetStatus,omitempty"`
	// The internal access domain name.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******-1801786532******-vpc.cn-wulanchabu.pai-eas.aliyuncs.com
	IntranetDomain  *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	IntranetEnabled *bool   `json:"IntranetEnabled,omitempty" xml:"IntranetEnabled,omitempty"`
	// Indicates whether the dedicated gateway is the default one.
	//
	// example:
	//
	// true
	IsDefault *bool                                `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Labels    []*DescribeGatewayResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The number of nodes in the dedicated gateway.
	//
	// example:
	//
	// 2
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether HTTP-to-HTTPS redirection is enabled.
	//
	// example:
	//
	// true
	SSLRedirectionEnabled *bool `json:"SSLRedirectionEnabled,omitempty" xml:"SSLRedirectionEnabled,omitempty"`
	// The status of the dedicated gateway.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the dedicated gateway was last updated. The time is in UTC.
	//
	// example:
	//
	// 2021-02-24T11:52:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeGatewayResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGatewayResponseBody) GetExternalClusterId() *string {
	return s.ExternalClusterId
}

func (s *DescribeGatewayResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DescribeGatewayResponseBody) GetGatewayName() *string {
	return s.GatewayName
}

func (s *DescribeGatewayResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeGatewayResponseBody) GetInternetDomain() *string {
	return s.InternetDomain
}

func (s *DescribeGatewayResponseBody) GetInternetEnabled() *bool {
	return s.InternetEnabled
}

func (s *DescribeGatewayResponseBody) GetInternetStatus() *string {
	return s.InternetStatus
}

func (s *DescribeGatewayResponseBody) GetIntranetDomain() *string {
	return s.IntranetDomain
}

func (s *DescribeGatewayResponseBody) GetIntranetEnabled() *bool {
	return s.IntranetEnabled
}

func (s *DescribeGatewayResponseBody) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeGatewayResponseBody) GetLabels() []*DescribeGatewayResponseBodyLabels {
	return s.Labels
}

func (s *DescribeGatewayResponseBody) GetReplicas() *int32 {
	return s.Replicas
}

func (s *DescribeGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGatewayResponseBody) GetSSLRedirectionEnabled() *bool {
	return s.SSLRedirectionEnabled
}

func (s *DescribeGatewayResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeGatewayResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeGatewayResponseBody) SetChargeType(v string) *DescribeGatewayResponseBody {
	s.ChargeType = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetCreateTime(v string) *DescribeGatewayResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetExternalClusterId(v string) *DescribeGatewayResponseBody {
	s.ExternalClusterId = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetGatewayId(v string) *DescribeGatewayResponseBody {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetGatewayName(v string) *DescribeGatewayResponseBody {
	s.GatewayName = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetInstanceType(v string) *DescribeGatewayResponseBody {
	s.InstanceType = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetInternetDomain(v string) *DescribeGatewayResponseBody {
	s.InternetDomain = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetInternetEnabled(v bool) *DescribeGatewayResponseBody {
	s.InternetEnabled = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetInternetStatus(v string) *DescribeGatewayResponseBody {
	s.InternetStatus = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetIntranetDomain(v string) *DescribeGatewayResponseBody {
	s.IntranetDomain = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetIntranetEnabled(v bool) *DescribeGatewayResponseBody {
	s.IntranetEnabled = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetIsDefault(v bool) *DescribeGatewayResponseBody {
	s.IsDefault = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetLabels(v []*DescribeGatewayResponseBodyLabels) *DescribeGatewayResponseBody {
	s.Labels = v
	return s
}

func (s *DescribeGatewayResponseBody) SetReplicas(v int32) *DescribeGatewayResponseBody {
	s.Replicas = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetRequestId(v string) *DescribeGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetSSLRedirectionEnabled(v bool) *DescribeGatewayResponseBody {
	s.SSLRedirectionEnabled = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetStatus(v string) *DescribeGatewayResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetUpdateTime(v string) *DescribeGatewayResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGatewayResponseBody) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGatewayResponseBodyLabels struct {
	// example:
	//
	// key1
	LabelKey *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	// example:
	//
	// value1
	LabelValue *string `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
}

func (s DescribeGatewayResponseBodyLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *DescribeGatewayResponseBodyLabels) GetLabelKey() *string {
	return s.LabelKey
}

func (s *DescribeGatewayResponseBodyLabels) GetLabelValue() *string {
	return s.LabelValue
}

func (s *DescribeGatewayResponseBodyLabels) SetLabelKey(v string) *DescribeGatewayResponseBodyLabels {
	s.LabelKey = &v
	return s
}

func (s *DescribeGatewayResponseBodyLabels) SetLabelValue(v string) *DescribeGatewayResponseBodyLabels {
	s.LabelValue = &v
	return s
}

func (s *DescribeGatewayResponseBodyLabels) Validate() error {
	return dara.Validate(s)
}
