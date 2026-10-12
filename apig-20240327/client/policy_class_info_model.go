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
	// The alias of the policy template.
	//
	// example:
	//
	// Concurrency Control
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// The list of resource types that can be mounted.
	AttachableResourceTypes []*string `json:"attachableResourceTypes,omitempty" xml:"attachableResourceTypes,omitempty" type:"Repeated"`
	// The identifier of the policy template.
	//
	// example:
	//
	// pc-cpqm0s94ckt1displ8ng
	ClassId *string `json:"classId,omitempty" xml:"classId,omitempty"`
	// The configuration example.
	//
	// example:
	//
	// {}
	ConfigExample *string `json:"configExample,omitempty" xml:"configExample,omitempty"`
	// Indicates whether the policy is deprecated.
	//
	// example:
	//
	// false
	Deprecated *bool `json:"deprecated,omitempty" xml:"deprecated,omitempty"`
	// The detailed description of the policy template.
	//
	// example:
	//
	// Controls the number of concurrent threads to ensure system stability. Applicable to scenarios where internal application components or downstream dependencies become unstable, such as slow SQL queries or increased response time from downstream applications
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The direction in which the policy takes effect.
	//
	// example:
	//
	// InBound
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// Indicates whether log recording is enabled.
	//
	// example:
	//
	// true
	EnableLog *bool `json:"enableLog,omitempty" xml:"enableLog,omitempty"`
	// The execution priority.
	//
	// example:
	//
	// 1
	ExecutePriority *string `json:"executePriority,omitempty" xml:"executePriority,omitempty"`
	// The execution stage.
	//
	// example:
	//
	// POST
	ExecuteStage *string `json:"executeStage,omitempty" xml:"executeStage,omitempty"`
	// The name of the policy template.
	//
	// example:
	//
	// ConcurrencyLimit
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The policy type.
	//
	// example:
	//
	// FlowControl
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version of the policy template.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
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
