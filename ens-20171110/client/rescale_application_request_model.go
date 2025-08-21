// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRescaleApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RescaleApplicationRequest
	GetAppId() *string
	SetRescaleLevel(v string) *RescaleApplicationRequest
	GetRescaleLevel() *string
	SetRescaleType(v string) *RescaleApplicationRequest
	GetRescaleType() *string
	SetResourceSelector(v string) *RescaleApplicationRequest
	GetResourceSelector() *string
	SetTimeout(v int32) *RescaleApplicationRequest
	GetTimeout() *int32
	SetToAppVersion(v string) *RescaleApplicationRequest
	GetToAppVersion() *string
}

type RescaleApplicationRequest struct {
	// The ID of the application. You can query the application ID by calling the ListApplications operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 474bdef0-d149-4695-abfb-52912d9143f0
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The level of resource scaling. The value is of the enumeration type. Valid values:
	//
	// 	- AreaIspCode (default): scales resources based on the Internet service provider (ISP).
	//
	// 	- RegionId: scales resources based on the edge node.
	//
	// 	- InstanceId: scales resources based on the instance ID. Resource scale-out specifies resource hosting and scale-in specifies resource release.
	//
	// Default value: AreaIspCode.
	//
	// example:
	//
	// RegionId
	RescaleLevel *string `json:"RescaleLevel,omitempty" xml:"RescaleLevel,omitempty"`
	// The scaling operation. The value must be of the enumerated data type. Valid values:
	//
	// 	- Add: adds new resources.
	//
	// 	- Del: releases resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// Add
	RescaleType *string `json:"RescaleType,omitempty" xml:"RescaleType,omitempty"`
	// The required resources. The value must be a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"regionCode\\": \\"cn-wuxi-telecom_unicom_cmcc-3\\",    \\"ispCode\\": \\"telecom\\",    \\"count\\": 2	},{    \\"regionCode\\": \\"cn-shanghai-cmcc\\",    \\"count\\": 4	}]
	ResourceSelector *string `json:"ResourceSelector,omitempty" xml:"ResourceSelector,omitempty"`
	// The timeout period for asynchronous scaling. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 1800
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The version number of the application deployment package. By default, the stable version number is used. This parameter takes effect only when you perform resource scale-out.
	//
	// example:
	//
	// v2
	ToAppVersion *string `json:"ToAppVersion,omitempty" xml:"ToAppVersion,omitempty"`
}

func (s RescaleApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s RescaleApplicationRequest) GoString() string {
	return s.String()
}

func (s *RescaleApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *RescaleApplicationRequest) GetRescaleLevel() *string {
	return s.RescaleLevel
}

func (s *RescaleApplicationRequest) GetRescaleType() *string {
	return s.RescaleType
}

func (s *RescaleApplicationRequest) GetResourceSelector() *string {
	return s.ResourceSelector
}

func (s *RescaleApplicationRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *RescaleApplicationRequest) GetToAppVersion() *string {
	return s.ToAppVersion
}

func (s *RescaleApplicationRequest) SetAppId(v string) *RescaleApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RescaleApplicationRequest) SetRescaleLevel(v string) *RescaleApplicationRequest {
	s.RescaleLevel = &v
	return s
}

func (s *RescaleApplicationRequest) SetRescaleType(v string) *RescaleApplicationRequest {
	s.RescaleType = &v
	return s
}

func (s *RescaleApplicationRequest) SetResourceSelector(v string) *RescaleApplicationRequest {
	s.ResourceSelector = &v
	return s
}

func (s *RescaleApplicationRequest) SetTimeout(v int32) *RescaleApplicationRequest {
	s.Timeout = &v
	return s
}

func (s *RescaleApplicationRequest) SetToAppVersion(v string) *RescaleApplicationRequest {
	s.ToAppVersion = &v
	return s
}

func (s *RescaleApplicationRequest) Validate() error {
	return dara.Validate(s)
}
