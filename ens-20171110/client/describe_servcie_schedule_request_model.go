// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServcieScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeServcieScheduleRequest
	GetAppId() *string
	SetPodConfigName(v string) *DescribeServcieScheduleRequest
	GetPodConfigName() *string
	SetUuid(v string) *DescribeServcieScheduleRequest
	GetUuid() *string
}

type DescribeServcieScheduleRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 474bdef0-d149-4695-abfb-52912d91****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter does not take effect.
	//
	// example:
	//
	// android
	PodConfigName *string `json:"PodConfigName,omitempty" xml:"PodConfigName,omitempty"`
	// The unique ID of the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// hdm_f022bf160dc69e2d8eb421e508eb****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeServcieScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServcieScheduleRequest) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeServcieScheduleRequest) GetPodConfigName() *string {
	return s.PodConfigName
}

func (s *DescribeServcieScheduleRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeServcieScheduleRequest) SetAppId(v string) *DescribeServcieScheduleRequest {
	s.AppId = &v
	return s
}

func (s *DescribeServcieScheduleRequest) SetPodConfigName(v string) *DescribeServcieScheduleRequest {
	s.PodConfigName = &v
	return s
}

func (s *DescribeServcieScheduleRequest) SetUuid(v string) *DescribeServcieScheduleRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeServcieScheduleRequest) Validate() error {
	return dara.Validate(s)
}
