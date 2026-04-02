// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDirectNotifyChannel interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifiers(v []*string) *DirectNotifyChannel
	GetIdentifiers() []*string
	SetType(v string) *DirectNotifyChannel
	GetType() *string
}

type DirectNotifyChannel struct {
	// 通知对象标识列表
	//
	// This parameter is required.
	Identifiers []*string `json:"identifiers,omitempty" xml:"identifiers,omitempty" type:"Repeated"`
	// 通知渠道类型
	//
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DirectNotifyChannel) String() string {
	return dara.Prettify(s)
}

func (s DirectNotifyChannel) GoString() string {
	return s.String()
}

func (s *DirectNotifyChannel) GetIdentifiers() []*string {
	return s.Identifiers
}

func (s *DirectNotifyChannel) GetType() *string {
	return s.Type
}

func (s *DirectNotifyChannel) SetIdentifiers(v []*string) *DirectNotifyChannel {
	s.Identifiers = v
	return s
}

func (s *DirectNotifyChannel) SetType(v string) *DirectNotifyChannel {
	s.Type = &v
	return s
}

func (s *DirectNotifyChannel) Validate() error {
	return dara.Validate(s)
}
