// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEipInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *CreateEipInstanceRequest
	GetBandwidth() *int64
	SetClientToken(v string) *CreateEipInstanceRequest
	GetClientToken() *string
	SetDescription(v string) *CreateEipInstanceRequest
	GetDescription() *string
	SetEnsRegionId(v string) *CreateEipInstanceRequest
	GetEnsRegionId() *string
	SetInstanceChargeType(v string) *CreateEipInstanceRequest
	GetInstanceChargeType() *string
	SetInstanceId(v string) *CreateEipInstanceRequest
	GetInstanceId() *string
	SetInternetChargeType(v string) *CreateEipInstanceRequest
	GetInternetChargeType() *string
	SetIpAddress(v string) *CreateEipInstanceRequest
	GetIpAddress() *string
	SetIsp(v string) *CreateEipInstanceRequest
	GetIsp() *string
	SetName(v string) *CreateEipInstanceRequest
	GetName() *string
	SetTag(v []*CreateEipInstanceRequestTag) *CreateEipInstanceRequest
	GetTag() []*CreateEipInstanceRequestTag
}

type CreateEipInstanceRequest struct {
	// The peak bandwidth of the EIP. Default value: 5. Valid values: 5 to 10000. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// Ensures the idempotence of the request to prevent duplicate operations caused by multiple retries.
	//
	// - Generate a parameter value from your client and make sure that the value is unique among different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters in length.
	//
	// - If you use a ClientToken that has been used before and the other request parameters remain unchanged, the client receives the same result as the first submission, without affecting the server-side state.
	//
	// - You can retry when the API times out or returns the PROCESSING error code. Idempotence is effective in this case. If the response returns an HTTP 200 status code, the retry returns the same result as the previous call without affecting the server-side state. If the response returns an HTTP 4xx status code and the error code is not PROCESSING, idempotence is not effective.
	//
	// - The ClientToken is valid for 10 minutes.
	//
	// example:
	//
	// 26C28756-2586-17AF-B802-0DC50D8FDEBB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description.
	//
	// example:
	//
	// yourDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Edge Node Service (ENS) edge zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-suzhou-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// Billable methods of the EIP. Currently, only **PostPaid*	- (pay-as-you-go billing method) is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID of the EIP.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Billable methods of the EIP. Currently, only **95BandwidthByMonth*	- (monthly 95th percentile bandwidth billing) is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95BandwidthByMonth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The IP address.
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The Internet service provider (ISP). Valid values:
	//
	// - **cmcc**: China Mobile.
	//
	// - **unicom**: China Unicom.
	//
	// - **telecom**: China Telecom.
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The name of the EIP instance.
	//
	// example:
	//
	// EIP1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource tags.
	Tag []*CreateEipInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateEipInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEipInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateEipInstanceRequest) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateEipInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEipInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEipInstanceRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateEipInstanceRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateEipInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEipInstanceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateEipInstanceRequest) GetIpAddress() *string {
	return s.IpAddress
}

func (s *CreateEipInstanceRequest) GetIsp() *string {
	return s.Isp
}

func (s *CreateEipInstanceRequest) GetName() *string {
	return s.Name
}

func (s *CreateEipInstanceRequest) GetTag() []*CreateEipInstanceRequestTag {
	return s.Tag
}

func (s *CreateEipInstanceRequest) SetBandwidth(v int64) *CreateEipInstanceRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateEipInstanceRequest) SetClientToken(v string) *CreateEipInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEipInstanceRequest) SetDescription(v string) *CreateEipInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateEipInstanceRequest) SetEnsRegionId(v string) *CreateEipInstanceRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateEipInstanceRequest) SetInstanceChargeType(v string) *CreateEipInstanceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateEipInstanceRequest) SetInstanceId(v string) *CreateEipInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateEipInstanceRequest) SetInternetChargeType(v string) *CreateEipInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateEipInstanceRequest) SetIpAddress(v string) *CreateEipInstanceRequest {
	s.IpAddress = &v
	return s
}

func (s *CreateEipInstanceRequest) SetIsp(v string) *CreateEipInstanceRequest {
	s.Isp = &v
	return s
}

func (s *CreateEipInstanceRequest) SetName(v string) *CreateEipInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateEipInstanceRequest) SetTag(v []*CreateEipInstanceRequestTag) *CreateEipInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateEipInstanceRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEipInstanceRequestTag struct {
	// The tag key of the instance. Valid values of N: **1*	- to **20**.
	//
	// - Cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// - Can be up to 64 characters in length.
	//
	// - Cannot be an empty string.
	//
	// example:
	//
	// team
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`.
	//
	// example:
	//
	// tagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEipInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEipInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEipInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEipInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEipInstanceRequestTag) SetKey(v string) *CreateEipInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEipInstanceRequestTag) SetValue(v string) *CreateEipInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEipInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
