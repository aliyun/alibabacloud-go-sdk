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
	SetInternetChargeType(v string) *CreateEipInstanceRequest
	GetInternetChargeType() *string
	SetIsp(v string) *CreateEipInstanceRequest
	GetIsp() *string
	SetName(v string) *CreateEipInstanceRequest
	GetName() *string
	SetTag(v []*CreateEipInstanceRequestTag) *CreateEipInstanceRequest
	GetTag() []*CreateEipInstanceRequestTag
}

type CreateEipInstanceRequest struct {
	// The maximum bandwidth of the EIP. Default value: 5. Valid values: 5 to 10000. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request. This prevents repeated operations caused by multiple retries.
	//
	// 	- You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length.
	//
	// 	- If you use a ClientToken that has been used and other request parameters remain unchanged in a repeated request, the client will receive the same result as the first request. This does not affect the status of your server.
	//
	// 	- You can initiate a retry when the operation times out or the error code is PROCESSING. The idempotence is valid. If HTTP status code 200 is returned, the client receives the same result as the last request. However, your server status is not affected. If HTTP status code 4xx is returned and error code is not PROCESSING, the idempotence is invalid.
	//
	// 	- A client token is valid for 10 minutes.
	//
	// example:
	//
	// 26C28756-2586-17AF-B802-0DC50D8FDEBB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the EIP.
	//
	// example:
	//
	// yourDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Edge Node Service (ENS) node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-suzhou-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The billing method of the EIP. Set the value to **PostPaid**.
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the EIP. Set the value to **95BandwidthByMonth**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95BandwidthByMonth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The Internet service provider. Valid values:
	//
	// 	- **cmcc**: China Mobile.
	//
	// 	- **unicom**: China Unicom.
	//
	// 	- **telecom**: China Telecom.
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The name of the EIP.
	//
	// example:
	//
	// EIP1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tags.
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

func (s *CreateEipInstanceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
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

func (s *CreateEipInstanceRequest) SetInternetChargeType(v string) *CreateEipInstanceRequest {
	s.InternetChargeType = &v
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
	// The key of tag N of the instance. Valid values of N: **1*	- to **20**.
	//
	// 	- The key cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// 	- The key must be up to 64 characters in length.
	//
	// 	- The tag key cannot be an empty string.
	//
	// example:
	//
	// team
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the resource. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length. It cannot start with acs: or contain http:// or https://.
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
