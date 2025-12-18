// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTimeMs(v int64) *InstanceInfo
	GetCreatedTimeMs() *int64
	SetDestroyedTimeMs(v int64) *InstanceInfo
	GetDestroyedTimeMs() *int64
	SetInstanceId(v string) *InstanceInfo
	GetInstanceId() *string
	SetQualifier(v string) *InstanceInfo
	GetQualifier() *string
	SetResourceType(v string) *InstanceInfo
	GetResourceType() *string
	SetStatus(v string) *InstanceInfo
	GetStatus() *string
	SetVersionId(v string) *InstanceInfo
	GetVersionId() *string
}

type InstanceInfo struct {
	CreatedTimeMs   *int64 `json:"createdTimeMs,omitempty" xml:"createdTimeMs,omitempty"`
	DestroyedTimeMs *int64 `json:"destroyedTimeMs,omitempty" xml:"destroyedTimeMs,omitempty"`
	// example:
	//
	// 1ef6b6ff-7f7b-485e-ab49-501ac681****
	InstanceId   *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Qualifier    *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	VersionId    *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s InstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s InstanceInfo) GoString() string {
	return s.String()
}

func (s *InstanceInfo) GetCreatedTimeMs() *int64 {
	return s.CreatedTimeMs
}

func (s *InstanceInfo) GetDestroyedTimeMs() *int64 {
	return s.DestroyedTimeMs
}

func (s *InstanceInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InstanceInfo) GetQualifier() *string {
	return s.Qualifier
}

func (s *InstanceInfo) GetResourceType() *string {
	return s.ResourceType
}

func (s *InstanceInfo) GetStatus() *string {
	return s.Status
}

func (s *InstanceInfo) GetVersionId() *string {
	return s.VersionId
}

func (s *InstanceInfo) SetCreatedTimeMs(v int64) *InstanceInfo {
	s.CreatedTimeMs = &v
	return s
}

func (s *InstanceInfo) SetDestroyedTimeMs(v int64) *InstanceInfo {
	s.DestroyedTimeMs = &v
	return s
}

func (s *InstanceInfo) SetInstanceId(v string) *InstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *InstanceInfo) SetQualifier(v string) *InstanceInfo {
	s.Qualifier = &v
	return s
}

func (s *InstanceInfo) SetResourceType(v string) *InstanceInfo {
	s.ResourceType = &v
	return s
}

func (s *InstanceInfo) SetStatus(v string) *InstanceInfo {
	s.Status = &v
	return s
}

func (s *InstanceInfo) SetVersionId(v string) *InstanceInfo {
	s.VersionId = &v
	return s
}

func (s *InstanceInfo) Validate() error {
	return dara.Validate(s)
}
