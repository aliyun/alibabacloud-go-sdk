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
	AccessToken      *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	QueueService     *string `json:"QueueService,omitempty" xml:"QueueService,omitempty"`
	TrafficMode      *string `json:"TrafficMode,omitempty" xml:"TrafficMode,omitempty"`
	UpdateTime       *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
