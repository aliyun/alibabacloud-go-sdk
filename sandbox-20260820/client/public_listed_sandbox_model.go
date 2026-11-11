// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicListedSandbox interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPause(v bool) *PublicListedSandbox
	GetAutoPause() *bool
	SetAutoResume(v bool) *PublicListedSandbox
	GetAutoResume() *bool
	SetCpu(v float32) *PublicListedSandbox
	GetCpu() *float32
	SetDiskSize(v int32) *PublicListedSandbox
	GetDiskSize() *int32
	SetEndAt(v string) *PublicListedSandbox
	GetEndAt() *string
	SetMemorySize(v int32) *PublicListedSandbox
	GetMemorySize() *int32
	SetResourceGroupID(v string) *PublicListedSandbox
	GetResourceGroupID() *string
	SetResourceTags(v []*PublicListedSandboxResourceTags) *PublicListedSandbox
	GetResourceTags() []*PublicListedSandboxResourceTags
	SetSandboxID(v string) *PublicListedSandbox
	GetSandboxID() *string
	SetStartedAt(v string) *PublicListedSandbox
	GetStartedAt() *string
	SetState(v string) *PublicListedSandbox
	GetState() *string
	SetTeamID(v string) *PublicListedSandbox
	GetTeamID() *string
	SetTeamName(v string) *PublicListedSandbox
	GetTeamName() *string
	SetTemplateID(v string) *PublicListedSandbox
	GetTemplateID() *string
	SetTemplateName(v string) *PublicListedSandbox
	GetTemplateName() *string
	SetUserID(v string) *PublicListedSandbox
	GetUserID() *string
}

type PublicListedSandbox struct {
	// example:
	//
	// true
	AutoPause *bool `json:"autoPause,omitempty" xml:"autoPause,omitempty"`
	// example:
	//
	// true
	AutoResume *bool `json:"autoResume,omitempty" xml:"autoResume,omitempty"`
	// example:
	//
	// 2
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// 2048
	DiskSize *int32 `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	// example:
	//
	// 2026-10-30T03:09:21.313Z
	EndAt *string `json:"endAt,omitempty" xml:"endAt,omitempty"`
	// example:
	//
	// 10240
	MemorySize *int32 `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// example:
	//
	// rg-acfmwxqyrgwabcd
	ResourceGroupID *string                            `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	ResourceTags    []*PublicListedSandboxResourceTags `json:"resourceTags" xml:"resourceTags" type:"Repeated"`
	// example:
	//
	// sbx-4a4356a4-b731-42c1-ba8e-650f725be098
	SandboxID *string `json:"sandboxID,omitempty" xml:"sandboxID,omitempty"`
	// example:
	//
	// 2026-10-30T02:09:21.313Z
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// example:
	//
	// running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// 13b721e6-8cc8-5df2-af13-80316f7508af
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// example:
	//
	// my-team
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
	// example:
	//
	// us7dxqaezw5uu7aa2cm5
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
	// example:
	//
	// my-base-template
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// example:
	//
	// 123456789
	UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s PublicListedSandbox) String() string {
	return dara.Prettify(s)
}

func (s PublicListedSandbox) GoString() string {
	return s.String()
}

func (s *PublicListedSandbox) GetAutoPause() *bool {
	return s.AutoPause
}

func (s *PublicListedSandbox) GetAutoResume() *bool {
	return s.AutoResume
}

func (s *PublicListedSandbox) GetCpu() *float32 {
	return s.Cpu
}

func (s *PublicListedSandbox) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *PublicListedSandbox) GetEndAt() *string {
	return s.EndAt
}

func (s *PublicListedSandbox) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *PublicListedSandbox) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *PublicListedSandbox) GetResourceTags() []*PublicListedSandboxResourceTags {
	return s.ResourceTags
}

func (s *PublicListedSandbox) GetSandboxID() *string {
	return s.SandboxID
}

func (s *PublicListedSandbox) GetStartedAt() *string {
	return s.StartedAt
}

func (s *PublicListedSandbox) GetState() *string {
	return s.State
}

func (s *PublicListedSandbox) GetTeamID() *string {
	return s.TeamID
}

func (s *PublicListedSandbox) GetTeamName() *string {
	return s.TeamName
}

func (s *PublicListedSandbox) GetTemplateID() *string {
	return s.TemplateID
}

func (s *PublicListedSandbox) GetTemplateName() *string {
	return s.TemplateName
}

func (s *PublicListedSandbox) GetUserID() *string {
	return s.UserID
}

func (s *PublicListedSandbox) SetAutoPause(v bool) *PublicListedSandbox {
	s.AutoPause = &v
	return s
}

func (s *PublicListedSandbox) SetAutoResume(v bool) *PublicListedSandbox {
	s.AutoResume = &v
	return s
}

func (s *PublicListedSandbox) SetCpu(v float32) *PublicListedSandbox {
	s.Cpu = &v
	return s
}

func (s *PublicListedSandbox) SetDiskSize(v int32) *PublicListedSandbox {
	s.DiskSize = &v
	return s
}

func (s *PublicListedSandbox) SetEndAt(v string) *PublicListedSandbox {
	s.EndAt = &v
	return s
}

func (s *PublicListedSandbox) SetMemorySize(v int32) *PublicListedSandbox {
	s.MemorySize = &v
	return s
}

func (s *PublicListedSandbox) SetResourceGroupID(v string) *PublicListedSandbox {
	s.ResourceGroupID = &v
	return s
}

func (s *PublicListedSandbox) SetResourceTags(v []*PublicListedSandboxResourceTags) *PublicListedSandbox {
	s.ResourceTags = v
	return s
}

func (s *PublicListedSandbox) SetSandboxID(v string) *PublicListedSandbox {
	s.SandboxID = &v
	return s
}

func (s *PublicListedSandbox) SetStartedAt(v string) *PublicListedSandbox {
	s.StartedAt = &v
	return s
}

func (s *PublicListedSandbox) SetState(v string) *PublicListedSandbox {
	s.State = &v
	return s
}

func (s *PublicListedSandbox) SetTeamID(v string) *PublicListedSandbox {
	s.TeamID = &v
	return s
}

func (s *PublicListedSandbox) SetTeamName(v string) *PublicListedSandbox {
	s.TeamName = &v
	return s
}

func (s *PublicListedSandbox) SetTemplateID(v string) *PublicListedSandbox {
	s.TemplateID = &v
	return s
}

func (s *PublicListedSandbox) SetTemplateName(v string) *PublicListedSandbox {
	s.TemplateName = &v
	return s
}

func (s *PublicListedSandbox) SetUserID(v string) *PublicListedSandbox {
	s.UserID = &v
	return s
}

func (s *PublicListedSandbox) Validate() error {
	if s.ResourceTags != nil {
		for _, item := range s.ResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PublicListedSandboxResourceTags struct {
	// example:
	//
	// project
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// example:
	//
	// demo
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s PublicListedSandboxResourceTags) String() string {
	return dara.Prettify(s)
}

func (s PublicListedSandboxResourceTags) GoString() string {
	return s.String()
}

func (s *PublicListedSandboxResourceTags) GetTagKey() *string {
	return s.TagKey
}

func (s *PublicListedSandboxResourceTags) GetTagValue() *string {
	return s.TagValue
}

func (s *PublicListedSandboxResourceTags) SetTagKey(v string) *PublicListedSandboxResourceTags {
	s.TagKey = &v
	return s
}

func (s *PublicListedSandboxResourceTags) SetTagValue(v string) *PublicListedSandboxResourceTags {
	s.TagValue = &v
	return s
}

func (s *PublicListedSandboxResourceTags) Validate() error {
	return dara.Validate(s)
}
