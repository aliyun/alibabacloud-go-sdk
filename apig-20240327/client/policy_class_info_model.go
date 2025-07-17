// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPolicyClassInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *PolicyClassInfo
	GetAlias() *string
	SetAttachableResourceTypes(v []*string) *PolicyClassInfo
	GetAttachableResourceTypes() []*string
	SetClassId(v string) *PolicyClassInfo
	GetClassId() *string
	SetConfigExample(v string) *PolicyClassInfo
	GetConfigExample() *string
	SetDeprecated(v bool) *PolicyClassInfo
	GetDeprecated() *bool
	SetDescription(v string) *PolicyClassInfo
	GetDescription() *string
	SetDirection(v string) *PolicyClassInfo
	GetDirection() *string
	SetEnableLog(v bool) *PolicyClassInfo
	GetEnableLog() *bool
	SetExecutePriority(v string) *PolicyClassInfo
	GetExecutePriority() *string
	SetExecuteStage(v string) *PolicyClassInfo
	GetExecuteStage() *string
	SetName(v string) *PolicyClassInfo
	GetName() *string
	SetType(v string) *PolicyClassInfo
	GetType() *string
	SetVersion(v string) *PolicyClassInfo
	GetVersion() *string
}

type PolicyClassInfo struct {
	Alias                   *string   `json:"alias,omitempty" xml:"alias,omitempty"`
	AttachableResourceTypes []*string `json:"attachableResourceTypes,omitempty" xml:"attachableResourceTypes,omitempty" type:"Repeated"`
	ClassId                 *string   `json:"classId,omitempty" xml:"classId,omitempty"`
	ConfigExample           *string   `json:"configExample,omitempty" xml:"configExample,omitempty"`
	Deprecated              *bool     `json:"deprecated,omitempty" xml:"deprecated,omitempty"`
	Description             *string   `json:"description,omitempty" xml:"description,omitempty"`
	Direction               *string   `json:"direction,omitempty" xml:"direction,omitempty"`
	EnableLog               *bool     `json:"enableLog,omitempty" xml:"enableLog,omitempty"`
	ExecutePriority         *string   `json:"executePriority,omitempty" xml:"executePriority,omitempty"`
	ExecuteStage            *string   `json:"executeStage,omitempty" xml:"executeStage,omitempty"`
	Name                    *string   `json:"name,omitempty" xml:"name,omitempty"`
	Type                    *string   `json:"type,omitempty" xml:"type,omitempty"`
	Version                 *string   `json:"version,omitempty" xml:"version,omitempty"`
}

func (s PolicyClassInfo) String() string {
	return dara.Prettify(s)
}

func (s PolicyClassInfo) GoString() string {
	return s.String()
}

func (s *PolicyClassInfo) GetAlias() *string {
	return s.Alias
}

func (s *PolicyClassInfo) GetAttachableResourceTypes() []*string {
	return s.AttachableResourceTypes
}

func (s *PolicyClassInfo) GetClassId() *string {
	return s.ClassId
}

func (s *PolicyClassInfo) GetConfigExample() *string {
	return s.ConfigExample
}

func (s *PolicyClassInfo) GetDeprecated() *bool {
	return s.Deprecated
}

func (s *PolicyClassInfo) GetDescription() *string {
	return s.Description
}

func (s *PolicyClassInfo) GetDirection() *string {
	return s.Direction
}

func (s *PolicyClassInfo) GetEnableLog() *bool {
	return s.EnableLog
}

func (s *PolicyClassInfo) GetExecutePriority() *string {
	return s.ExecutePriority
}

func (s *PolicyClassInfo) GetExecuteStage() *string {
	return s.ExecuteStage
}

func (s *PolicyClassInfo) GetName() *string {
	return s.Name
}

func (s *PolicyClassInfo) GetType() *string {
	return s.Type
}

func (s *PolicyClassInfo) GetVersion() *string {
	return s.Version
}

func (s *PolicyClassInfo) SetAlias(v string) *PolicyClassInfo {
	s.Alias = &v
	return s
}

func (s *PolicyClassInfo) SetAttachableResourceTypes(v []*string) *PolicyClassInfo {
	s.AttachableResourceTypes = v
	return s
}

func (s *PolicyClassInfo) SetClassId(v string) *PolicyClassInfo {
	s.ClassId = &v
	return s
}

func (s *PolicyClassInfo) SetConfigExample(v string) *PolicyClassInfo {
	s.ConfigExample = &v
	return s
}

func (s *PolicyClassInfo) SetDeprecated(v bool) *PolicyClassInfo {
	s.Deprecated = &v
	return s
}

func (s *PolicyClassInfo) SetDescription(v string) *PolicyClassInfo {
	s.Description = &v
	return s
}

func (s *PolicyClassInfo) SetDirection(v string) *PolicyClassInfo {
	s.Direction = &v
	return s
}

func (s *PolicyClassInfo) SetEnableLog(v bool) *PolicyClassInfo {
	s.EnableLog = &v
	return s
}

func (s *PolicyClassInfo) SetExecutePriority(v string) *PolicyClassInfo {
	s.ExecutePriority = &v
	return s
}

func (s *PolicyClassInfo) SetExecuteStage(v string) *PolicyClassInfo {
	s.ExecuteStage = &v
	return s
}

func (s *PolicyClassInfo) SetName(v string) *PolicyClassInfo {
	s.Name = &v
	return s
}

func (s *PolicyClassInfo) SetType(v string) *PolicyClassInfo {
	s.Type = &v
	return s
}

func (s *PolicyClassInfo) SetVersion(v string) *PolicyClassInfo {
	s.Version = &v
	return s
}

func (s *PolicyClassInfo) Validate() error {
	return dara.Validate(s)
}
