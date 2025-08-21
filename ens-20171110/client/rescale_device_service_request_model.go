// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRescaleDeviceServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RescaleDeviceServiceRequest
	GetAppId() *string
	SetImageId(v string) *RescaleDeviceServiceRequest
	GetImageId() *string
	SetIpType(v int32) *RescaleDeviceServiceRequest
	GetIpType() *int32
	SetRescaleLevel(v string) *RescaleDeviceServiceRequest
	GetRescaleLevel() *string
	SetRescaleType(v string) *RescaleDeviceServiceRequest
	GetRescaleType() *string
	SetResourceInfo(v string) *RescaleDeviceServiceRequest
	GetResourceInfo() *string
	SetResourceSelector(v string) *RescaleDeviceServiceRequest
	GetResourceSelector() *string
	SetResourceSpec(v string) *RescaleDeviceServiceRequest
	GetResourceSpec() *string
	SetServiceId(v string) *RescaleDeviceServiceRequest
	GetServiceId() *string
	SetTimeout(v int64) *RescaleDeviceServiceRequest
	GetTimeout() *int64
}

type RescaleDeviceServiceRequest struct {
	// The ID of the application.
	//
	// example:
	//
	// 7aedc50b-b1cb-4a7c-9e3d-4cf3c9ee55a4
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// m-5rynw9g1ow1e928lb83bqmbnf
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the network. The value must be of the enumerated data type. Valid values:
	//
	// 	- **1*	- (default): Internet
	//
	// 	- **2**: internal network
	//
	// example:
	//
	// 1
	IpType *int32 `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The region level of the scale-out. Set the value to RegionId. RegionId specifies that resource scale-out is performed based on the ID of the edge node.
	//
	// This parameter is required.
	//
	// example:
	//
	// RegionId
	RescaleLevel *string `json:"RescaleLevel,omitempty" xml:"RescaleLevel,omitempty"`
	// The scaling operation. Set the value to Add to add new resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// Add
	RescaleType *string `json:"RescaleType,omitempty" xml:"RescaleType,omitempty"`
	// The information about the resource specification template. The value must be a JSON string.
	//
	// example:
	//
	// {\\"imageId\\":\\"m-5s4z4c10avgwvwtn33gl2vgob\\",\\"ipType\\":2,\\"specName\\":\\"ens.pfb-c3m7.medium\\"}
	ResourceInfo *string `json:"ResourceInfo,omitempty" xml:"ResourceInfo,omitempty"`
	// The required resources. The value must be a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"regionCode\\": \\"cn-wuxi-telecom_unicom_cmcc-3\\",    \\"ispCode\\": \\"telecom\\",    \\"count\\": 2	},{    \\"regionCode\\": \\"cn-shanghai-cmcc\\",    \\"count\\": 4	}]
	ResourceSelector *string `json:"ResourceSelector,omitempty" xml:"ResourceSelector,omitempty"`
	// The resource specification.
	//
	// example:
	//
	// ens.a6e.large
	ResourceSpec *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// service-01c6dd6e93f040698566
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The timeout period for asynchronous scale-out. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 1800
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s RescaleDeviceServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s RescaleDeviceServiceRequest) GoString() string {
	return s.String()
}

func (s *RescaleDeviceServiceRequest) GetAppId() *string {
	return s.AppId
}

func (s *RescaleDeviceServiceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *RescaleDeviceServiceRequest) GetIpType() *int32 {
	return s.IpType
}

func (s *RescaleDeviceServiceRequest) GetRescaleLevel() *string {
	return s.RescaleLevel
}

func (s *RescaleDeviceServiceRequest) GetRescaleType() *string {
	return s.RescaleType
}

func (s *RescaleDeviceServiceRequest) GetResourceInfo() *string {
	return s.ResourceInfo
}

func (s *RescaleDeviceServiceRequest) GetResourceSelector() *string {
	return s.ResourceSelector
}

func (s *RescaleDeviceServiceRequest) GetResourceSpec() *string {
	return s.ResourceSpec
}

func (s *RescaleDeviceServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *RescaleDeviceServiceRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *RescaleDeviceServiceRequest) SetAppId(v string) *RescaleDeviceServiceRequest {
	s.AppId = &v
	return s
}

func (s *RescaleDeviceServiceRequest) SetImageId(v string) *RescaleDeviceServiceRequest {
	s.ImageId = &v
	return s
}

func (s *RescaleDeviceServiceRequest) SetIpType(v int32) *RescaleDeviceServiceRequest {
	s.IpType = &v
	return s
}

func (s *RescaleDeviceServiceRequest) SetRescaleLevel(v string) *RescaleDeviceServiceRequest {
	s.RescaleLevel = &v
	return s
}

func (s *RescaleDeviceServiceRequest) SetRescaleType(v string) *RescaleDeviceServiceRequest {
	s.RescaleType = &v
	return s
}

func (s *RescaleDeviceServiceRequest) SetResourceInfo(v string) *RescaleDeviceServiceRequest {
	s.ResourceInfo = &v
	return s
}

func (s *RescaleDeviceServiceRequest) SetResourceSelector(v string) *RescaleDeviceServiceRequest {
	s.ResourceSelector = &v
	return s
}

func (s *RescaleDeviceServiceRequest) SetResourceSpec(v string) *RescaleDeviceServiceRequest {
	s.ResourceSpec = &v
	return s
}

func (s *RescaleDeviceServiceRequest) SetServiceId(v string) *RescaleDeviceServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *RescaleDeviceServiceRequest) SetTimeout(v int64) *RescaleDeviceServiceRequest {
	s.Timeout = &v
	return s
}

func (s *RescaleDeviceServiceRequest) Validate() error {
	return dara.Validate(s)
}
