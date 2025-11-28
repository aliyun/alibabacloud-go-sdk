// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBroadcastTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *BroadcastTemplate
	GetCreateTime() *string
	SetId(v string) *BroadcastTemplate
	GetId() *string
	SetModifiedTime(v string) *BroadcastTemplate
	GetModifiedTime() *string
	SetName(v string) *BroadcastTemplate
	GetName() *string
	SetVariables(v []*TemplateVariable) *BroadcastTemplate
	GetVariables() []*TemplateVariable
}

type BroadcastTemplate struct {
	CreateTime   *string             `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Id           *string             `json:"id,omitempty" xml:"id,omitempty"`
	ModifiedTime *string             `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Name         *string             `json:"name,omitempty" xml:"name,omitempty"`
	Variables    []*TemplateVariable `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s BroadcastTemplate) String() string {
	return dara.Prettify(s)
}

func (s BroadcastTemplate) GoString() string {
	return s.String()
}

func (s *BroadcastTemplate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BroadcastTemplate) GetId() *string {
	return s.Id
}

func (s *BroadcastTemplate) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *BroadcastTemplate) GetName() *string {
	return s.Name
}

func (s *BroadcastTemplate) GetVariables() []*TemplateVariable {
	return s.Variables
}

func (s *BroadcastTemplate) SetCreateTime(v string) *BroadcastTemplate {
	s.CreateTime = &v
	return s
}

func (s *BroadcastTemplate) SetId(v string) *BroadcastTemplate {
	s.Id = &v
	return s
}

func (s *BroadcastTemplate) SetModifiedTime(v string) *BroadcastTemplate {
	s.ModifiedTime = &v
	return s
}

func (s *BroadcastTemplate) SetName(v string) *BroadcastTemplate {
	s.Name = &v
	return s
}

func (s *BroadcastTemplate) SetVariables(v []*TemplateVariable) *BroadcastTemplate {
	s.Variables = v
	return s
}

func (s *BroadcastTemplate) Validate() error {
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
