// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObserveGroupRelatedEntityType interface {
	dara.Model
	String() string
	GoString() string
	SetCnName(v string) *ObserveGroupRelatedEntityType
	GetCnName() *string
	SetDirection(v string) *ObserveGroupRelatedEntityType
	GetDirection() *string
	SetDomain(v string) *ObserveGroupRelatedEntityType
	GetDomain() *string
	SetEnName(v string) *ObserveGroupRelatedEntityType
	GetEnName() *string
	SetEntityGroup(v string) *ObserveGroupRelatedEntityType
	GetEntityGroup() *string
	SetRelation(v string) *ObserveGroupRelatedEntityType
	GetRelation() *string
	SetType(v string) *ObserveGroupRelatedEntityType
	GetType() *string
	SetVersion(v string) *ObserveGroupRelatedEntityType
	GetVersion() *string
}

type ObserveGroupRelatedEntityType struct {
	// The Chinese name of the entity.
	//
	// example:
	//
	// RocketMQ Topic
	CnName *string `json:"cnName,omitempty" xml:"cnName,omitempty"`
	// The association direction, which is currently fixed as outgoing, indicating that associated entities are searched along the direct outgoing edges of the primary entity.
	//
	// example:
	//
	// outgoing
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// The entity domain.
	//
	// example:
	//
	// acs
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The English name of the entity.
	//
	// example:
	//
	// RocketMQ Topic
	EnName *string `json:"enName,omitempty" xml:"enName,omitempty"`
	// The CommonSchema group to which the entity belongs.
	//
	// example:
	//
	// acs-rocketmq
	EntityGroup *string `json:"entityGroup,omitempty" xml:"entityGroup,omitempty"`
	// The direct outgoing edge relationship, such as use, run_on, calls, or contains.
	//
	// example:
	//
	// contains
	Relation *string `json:"relation,omitempty" xml:"relation,omitempty"`
	// The associated entity type.
	//
	// example:
	//
	// acs.rocketmq.topic
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The effective CommonSchema version.
	//
	// example:
	//
	// 0.1.14
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ObserveGroupRelatedEntityType) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupRelatedEntityType) GoString() string {
	return s.String()
}

func (s *ObserveGroupRelatedEntityType) GetCnName() *string {
	return s.CnName
}

func (s *ObserveGroupRelatedEntityType) GetDirection() *string {
	return s.Direction
}

func (s *ObserveGroupRelatedEntityType) GetDomain() *string {
	return s.Domain
}

func (s *ObserveGroupRelatedEntityType) GetEnName() *string {
	return s.EnName
}

func (s *ObserveGroupRelatedEntityType) GetEntityGroup() *string {
	return s.EntityGroup
}

func (s *ObserveGroupRelatedEntityType) GetRelation() *string {
	return s.Relation
}

func (s *ObserveGroupRelatedEntityType) GetType() *string {
	return s.Type
}

func (s *ObserveGroupRelatedEntityType) GetVersion() *string {
	return s.Version
}

func (s *ObserveGroupRelatedEntityType) SetCnName(v string) *ObserveGroupRelatedEntityType {
	s.CnName = &v
	return s
}

func (s *ObserveGroupRelatedEntityType) SetDirection(v string) *ObserveGroupRelatedEntityType {
	s.Direction = &v
	return s
}

func (s *ObserveGroupRelatedEntityType) SetDomain(v string) *ObserveGroupRelatedEntityType {
	s.Domain = &v
	return s
}

func (s *ObserveGroupRelatedEntityType) SetEnName(v string) *ObserveGroupRelatedEntityType {
	s.EnName = &v
	return s
}

func (s *ObserveGroupRelatedEntityType) SetEntityGroup(v string) *ObserveGroupRelatedEntityType {
	s.EntityGroup = &v
	return s
}

func (s *ObserveGroupRelatedEntityType) SetRelation(v string) *ObserveGroupRelatedEntityType {
	s.Relation = &v
	return s
}

func (s *ObserveGroupRelatedEntityType) SetType(v string) *ObserveGroupRelatedEntityType {
	s.Type = &v
	return s
}

func (s *ObserveGroupRelatedEntityType) SetVersion(v string) *ObserveGroupRelatedEntityType {
	s.Version = &v
	return s
}

func (s *ObserveGroupRelatedEntityType) Validate() error {
	return dara.Validate(s)
}
