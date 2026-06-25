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
	SetCallerUid(v string) *Group
	GetCallerUid() *string
	SetClusterId(v string) *Group
	GetClusterId() *string
	SetCreateTime(v string) *Group
	GetCreateTime() *string
	SetInternetEndpoint(v string) *Group
	GetInternetEndpoint() *string
	SetIntranetEndpoint(v string) *Group
	GetIntranetEndpoint() *string
	SetLabels(v []*GroupLabels) *Group
	GetLabels() []*GroupLabels
	SetName(v string) *Group
	GetName() *string
	SetNetwork(v *GroupNetwork) *Group
	GetNetwork() *GroupNetwork
	SetParentUid(v string) *Group
	GetParentUid() *string
	SetQueueService(v string) *Group
	GetQueueService() *string
	SetTrafficMode(v string) *Group
	GetTrafficMode() *string
	SetUpdateTime(v string) *Group
	GetUpdateTime() *string
}

type Group struct {
	// The access token for the traffic entry of the service group.
	//
	// example:
	//
	// MzJiMDI5MDliODc0MTlkYmI0ZDhlYmExYjczYTIyZTE3Zm********
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	CallerUid   *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// The region in which the service group resides.
	//
	// example:
	//
	// cn-shanghai
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The time when the service group was created. The time is in UTC.
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
	IntranetEndpoint *string        `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	Labels           []*GroupLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The name of the service group.
	//
	// example:
	//
	// foo
	Name      *string       `json:"Name,omitempty" xml:"Name,omitempty"`
	Network   *GroupNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	ParentUid *string       `json:"ParentUid,omitempty" xml:"ParentUid,omitempty"`
	// The queue services contained in the service group.
	//
	// example:
	//
	// qservice
	QueueService *string `json:"QueueService,omitempty" xml:"QueueService,omitempty"`
	// The traffic mode.
	//
	// example:
	//
	// auto
	TrafficMode *string `json:"TrafficMode,omitempty" xml:"TrafficMode,omitempty"`
	// The time when the service group was last updated. The time is in UTC.
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

func (s *Group) GetCallerUid() *string {
	return s.CallerUid
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

func (s *Group) GetLabels() []*GroupLabels {
	return s.Labels
}

func (s *Group) GetName() *string {
	return s.Name
}

func (s *Group) GetNetwork() *GroupNetwork {
	return s.Network
}

func (s *Group) GetParentUid() *string {
	return s.ParentUid
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

func (s *Group) SetCallerUid(v string) *Group {
	s.CallerUid = &v
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

func (s *Group) SetLabels(v []*GroupLabels) *Group {
	s.Labels = v
	return s
}

func (s *Group) SetName(v string) *Group {
	s.Name = &v
	return s
}

func (s *Group) SetNetwork(v *GroupNetwork) *Group {
	s.Network = v
	return s
}

func (s *Group) SetParentUid(v string) *Group {
	s.ParentUid = &v
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
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GroupLabels struct {
	LabelKey   *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	LabelValue *string `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
}

func (s GroupLabels) String() string {
	return dara.Prettify(s)
}

func (s GroupLabels) GoString() string {
	return s.String()
}

func (s *GroupLabels) GetLabelKey() *string {
	return s.LabelKey
}

func (s *GroupLabels) GetLabelValue() *string {
	return s.LabelValue
}

func (s *GroupLabels) SetLabelKey(v string) *GroupLabels {
	s.LabelKey = &v
	return s
}

func (s *GroupLabels) SetLabelValue(v string) *GroupLabels {
	s.LabelValue = &v
	return s
}

func (s *GroupLabels) Validate() error {
	return dara.Validate(s)
}

type GroupNetwork struct {
	GatewayId       *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId       *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GroupNetwork) String() string {
	return dara.Prettify(s)
}

func (s GroupNetwork) GoString() string {
	return s.String()
}

func (s *GroupNetwork) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GroupNetwork) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GroupNetwork) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GroupNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *GroupNetwork) SetGatewayId(v string) *GroupNetwork {
	s.GatewayId = &v
	return s
}

func (s *GroupNetwork) SetSecurityGroupId(v string) *GroupNetwork {
	s.SecurityGroupId = &v
	return s
}

func (s *GroupNetwork) SetVSwitchId(v string) *GroupNetwork {
	s.VSwitchId = &v
	return s
}

func (s *GroupNetwork) SetVpcId(v string) *GroupNetwork {
	s.VpcId = &v
	return s
}

func (s *GroupNetwork) Validate() error {
	return dara.Validate(s)
}
