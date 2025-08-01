// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGatewayService interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayTrafficPolicy(v *TrafficPolicy) *GatewayService
	GetGatewayTrafficPolicy() *TrafficPolicy
	SetGatewayUniqueId(v string) *GatewayService
	GetGatewayUniqueId() *string
	SetGroupName(v string) *GatewayService
	GetGroupName() *string
	SetId(v int64) *GatewayService
	GetId() *int64
	SetMetaInfo(v string) *GatewayService
	GetMetaInfo() *string
	SetName(v string) *GatewayService
	GetName() *string
	SetNamespace(v string) *GatewayService
	GetNamespace() *string
	SetSourceType(v string) *GatewayService
	GetSourceType() *string
}

type GatewayService struct {
	GatewayTrafficPolicy *TrafficPolicy `json:"GatewayTrafficPolicy,omitempty" xml:"GatewayTrafficPolicy,omitempty"`
	GatewayUniqueId      *string        `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	GroupName            *string        `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Id                   *int64         `json:"Id,omitempty" xml:"Id,omitempty"`
	MetaInfo             *string        `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	Name                 *string        `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace            *string        `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	SourceType           *string        `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s GatewayService) String() string {
	return dara.Prettify(s)
}

func (s GatewayService) GoString() string {
	return s.String()
}

func (s *GatewayService) GetGatewayTrafficPolicy() *TrafficPolicy {
	return s.GatewayTrafficPolicy
}

func (s *GatewayService) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GatewayService) GetGroupName() *string {
	return s.GroupName
}

func (s *GatewayService) GetId() *int64 {
	return s.Id
}

func (s *GatewayService) GetMetaInfo() *string {
	return s.MetaInfo
}

func (s *GatewayService) GetName() *string {
	return s.Name
}

func (s *GatewayService) GetNamespace() *string {
	return s.Namespace
}

func (s *GatewayService) GetSourceType() *string {
	return s.SourceType
}

func (s *GatewayService) SetGatewayTrafficPolicy(v *TrafficPolicy) *GatewayService {
	s.GatewayTrafficPolicy = v
	return s
}

func (s *GatewayService) SetGatewayUniqueId(v string) *GatewayService {
	s.GatewayUniqueId = &v
	return s
}

func (s *GatewayService) SetGroupName(v string) *GatewayService {
	s.GroupName = &v
	return s
}

func (s *GatewayService) SetId(v int64) *GatewayService {
	s.Id = &v
	return s
}

func (s *GatewayService) SetMetaInfo(v string) *GatewayService {
	s.MetaInfo = &v
	return s
}

func (s *GatewayService) SetName(v string) *GatewayService {
	s.Name = &v
	return s
}

func (s *GatewayService) SetNamespace(v string) *GatewayService {
	s.Namespace = &v
	return s
}

func (s *GatewayService) SetSourceType(v string) *GatewayService {
	s.SourceType = &v
	return s
}

func (s *GatewayService) Validate() error {
	return dara.Validate(s)
}
