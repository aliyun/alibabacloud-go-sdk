// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroup interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *Group
	GetAccessToken() *string
	SetClusterId(v string) *Group
	GetClusterId() *string
	SetCreateTime(v string) *Group
	GetCreateTime() *string
	SetInternetEndpoint(v string) *Group
	GetInternetEndpoint() *string
	SetIntranetEndpoint(v string) *Group
	GetIntranetEndpoint() *string
	SetName(v string) *Group
	GetName() *string
	SetQueueService(v string) *Group
	GetQueueService() *string
	SetTrafficMode(v string) *Group
	GetTrafficMode() *string
	SetUpdateTime(v string) *Group
	GetUpdateTime() *string
}

type Group struct {
	// The token that is used to access the service group.
	//
	// example:
	//
	// MzJiMDI5MDliODc0MTlkYmI0ZDhlYmExYjczYTIyZTE3Zm********
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The region where the service group resides.
	//
	// example:
	//
	// cn-shanghai
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The time when the service group was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-05-19T14:19:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The public endpoint of the service group.
	//
	// example:
	//
	// http://1110*****.cn-hangzhou.pai-eas.aliyuncs.com/api/predict/test_group
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	// The internal endpoint of the service group.
	//
	// example:
	//
	// http://1110*****.vpc.cn-hangzhou.pai-eas.aliyuncs.com/api/predict/test_group
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// The name of the service group.
	//
	// example:
	//
	// foo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The queue service that is included in the service group.
	//
	// example:
	//
	// qservice
	QueueService *string `json:"QueueService,omitempty" xml:"QueueService,omitempty"`
	// The traffic mode.
	//
	// Valid values:
	//
	// 	- auto: The traffic is automatically allocated based on the number of instances.
	//
	// 	- customized: The traffic is allocated based on the custom weight.
	//
	// example:
	//
	// auto
	TrafficMode *string `json:"TrafficMode,omitempty" xml:"TrafficMode,omitempty"`
	// The time when the service group was updated. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-29T11:13:20Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Group) String() string {
	return dara.Prettify(s)
}

func (s Group) GoString() string {
	return s.String()
}

func (s *Group) GetAccessToken() *string {
	return s.AccessToken
}

func (s *Group) GetClusterId() *string {
	return s.ClusterId
}

func (s *Group) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Group) GetInternetEndpoint() *string {
	return s.InternetEndpoint
}

func (s *Group) GetIntranetEndpoint() *string {
	return s.IntranetEndpoint
}

func (s *Group) GetName() *string {
	return s.Name
}

func (s *Group) GetQueueService() *string {
	return s.QueueService
}

func (s *Group) GetTrafficMode() *string {
	return s.TrafficMode
}

func (s *Group) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *Group) SetAccessToken(v string) *Group {
	s.AccessToken = &v
	return s
}

func (s *Group) SetClusterId(v string) *Group {
	s.ClusterId = &v
	return s
}

func (s *Group) SetCreateTime(v string) *Group {
	s.CreateTime = &v
	return s
}

func (s *Group) SetInternetEndpoint(v string) *Group {
	s.InternetEndpoint = &v
	return s
}

func (s *Group) SetIntranetEndpoint(v string) *Group {
	s.IntranetEndpoint = &v
	return s
}

func (s *Group) SetName(v string) *Group {
	s.Name = &v
	return s
}

func (s *Group) SetQueueService(v string) *Group {
	s.QueueService = &v
	return s
}

func (s *Group) SetTrafficMode(v string) *Group {
	s.TrafficMode = &v
	return s
}

func (s *Group) SetUpdateTime(v string) *Group {
	s.UpdateTime = &v
	return s
}

func (s *Group) Validate() error {
	return dara.Validate(s)
}
