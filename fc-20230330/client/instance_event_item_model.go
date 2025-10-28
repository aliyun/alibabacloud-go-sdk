// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceEventItem interface {
	dara.Model
	String() string
	GoString() string
	SetChildren(v []*InstanceEventItem) *InstanceEventItem
	GetChildren() []*InstanceEventItem
	SetLevel(v string) *InstanceEventItem
	GetLevel() *string
	SetMessage(v string) *InstanceEventItem
	GetMessage() *string
	SetTime(v string) *InstanceEventItem
	GetTime() *string
	SetType(v string) *InstanceEventItem
	GetType() *string
}

type InstanceEventItem struct {
	Children []*InstanceEventItem `json:"children" xml:"children" type:"Repeated"`
	Level    *string              `json:"level,omitempty" xml:"level,omitempty"`
	Message  *string              `json:"message,omitempty" xml:"message,omitempty"`
	Time     *string              `json:"time,omitempty" xml:"time,omitempty"`
	Type     *string              `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InstanceEventItem) String() string {
	return dara.Prettify(s)
}

func (s InstanceEventItem) GoString() string {
	return s.String()
}

func (s *InstanceEventItem) GetChildren() []*InstanceEventItem {
	return s.Children
}

func (s *InstanceEventItem) GetLevel() *string {
	return s.Level
}

func (s *InstanceEventItem) GetMessage() *string {
	return s.Message
}

func (s *InstanceEventItem) GetTime() *string {
	return s.Time
}

func (s *InstanceEventItem) GetType() *string {
	return s.Type
}

func (s *InstanceEventItem) SetChildren(v []*InstanceEventItem) *InstanceEventItem {
	s.Children = v
	return s
}

func (s *InstanceEventItem) SetLevel(v string) *InstanceEventItem {
	s.Level = &v
	return s
}

func (s *InstanceEventItem) SetMessage(v string) *InstanceEventItem {
	s.Message = &v
	return s
}

func (s *InstanceEventItem) SetTime(v string) *InstanceEventItem {
	s.Time = &v
	return s
}

func (s *InstanceEventItem) SetType(v string) *InstanceEventItem {
	s.Type = &v
	return s
}

func (s *InstanceEventItem) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
