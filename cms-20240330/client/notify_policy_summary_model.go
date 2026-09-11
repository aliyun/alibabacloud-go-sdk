// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyPolicySummary interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *NotifyPolicySummary
	GetCreateTime() *string
	SetDescription(v string) *NotifyPolicySummary
	GetDescription() *string
	SetEnabled(v bool) *NotifyPolicySummary
	GetEnabled() *bool
	SetName(v string) *NotifyPolicySummary
	GetName() *string
	SetNotifyStrategy(v *NotifyStrategyDetail) *NotifyPolicySummary
	GetNotifyStrategy() *NotifyStrategyDetail
	SetUpdateTime(v string) *NotifyPolicySummary
	GetUpdateTime() *string
	SetUserId(v string) *NotifyPolicySummary
	GetUserId() *string
	SetUuid(v string) *NotifyPolicySummary
	GetUuid() *string
	SetVersion(v int32) *NotifyPolicySummary
	GetVersion() *int32
	SetWorkspace(v string) *NotifyPolicySummary
	GetWorkspace() *string
}

type NotifyPolicySummary struct {
	// The creation time, represented as a UNIX millisecond timestamp string.
	//
	// example:
	//
	// 1700000000000
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the notification policy.
	//
	// example:
	//
	// A notification policy for testing
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether the notification policy is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The Policy Name of the notification policy.
	//
	// example:
	//
	// Sample notification policy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The configuration details of the notification policy.
	NotifyStrategy *NotifyStrategyDetail `json:"notifyStrategy,omitempty" xml:"notifyStrategy,omitempty"`
	// The update time, represented as a UNIX millisecond timestamp string.
	//
	// example:
	//
	// 1700000000000
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The Alibaba Cloud account UID.
	//
	// example:
	//
	// 1234567890
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The unique identifier of the notification policy.
	//
	// example:
	//
	// ns-uuid-001
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// The version number of the notification policy.
	//
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
	// The workspace identifier.
	//
	// example:
	//
	// default
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s NotifyPolicySummary) String() string {
	return dara.Prettify(s)
}

func (s NotifyPolicySummary) GoString() string {
	return s.String()
}

func (s *NotifyPolicySummary) GetCreateTime() *string {
	return s.CreateTime
}

func (s *NotifyPolicySummary) GetDescription() *string {
	return s.Description
}

func (s *NotifyPolicySummary) GetEnabled() *bool {
	return s.Enabled
}

func (s *NotifyPolicySummary) GetName() *string {
	return s.Name
}

func (s *NotifyPolicySummary) GetNotifyStrategy() *NotifyStrategyDetail {
	return s.NotifyStrategy
}

func (s *NotifyPolicySummary) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *NotifyPolicySummary) GetUserId() *string {
	return s.UserId
}

func (s *NotifyPolicySummary) GetUuid() *string {
	return s.Uuid
}

func (s *NotifyPolicySummary) GetVersion() *int32 {
	return s.Version
}

func (s *NotifyPolicySummary) GetWorkspace() *string {
	return s.Workspace
}

func (s *NotifyPolicySummary) SetCreateTime(v string) *NotifyPolicySummary {
	s.CreateTime = &v
	return s
}

func (s *NotifyPolicySummary) SetDescription(v string) *NotifyPolicySummary {
	s.Description = &v
	return s
}

func (s *NotifyPolicySummary) SetEnabled(v bool) *NotifyPolicySummary {
	s.Enabled = &v
	return s
}

func (s *NotifyPolicySummary) SetName(v string) *NotifyPolicySummary {
	s.Name = &v
	return s
}

func (s *NotifyPolicySummary) SetNotifyStrategy(v *NotifyStrategyDetail) *NotifyPolicySummary {
	s.NotifyStrategy = v
	return s
}

func (s *NotifyPolicySummary) SetUpdateTime(v string) *NotifyPolicySummary {
	s.UpdateTime = &v
	return s
}

func (s *NotifyPolicySummary) SetUserId(v string) *NotifyPolicySummary {
	s.UserId = &v
	return s
}

func (s *NotifyPolicySummary) SetUuid(v string) *NotifyPolicySummary {
	s.Uuid = &v
	return s
}

func (s *NotifyPolicySummary) SetVersion(v int32) *NotifyPolicySummary {
	s.Version = &v
	return s
}

func (s *NotifyPolicySummary) SetWorkspace(v string) *NotifyPolicySummary {
	s.Workspace = &v
	return s
}

func (s *NotifyPolicySummary) Validate() error {
	if s.NotifyStrategy != nil {
		if err := s.NotifyStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
