// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceInfo interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *InstanceInfo
	GetImageUrl() *string
	SetInstanceId(v string) *InstanceInfo
	GetInstanceId() *string
	SetStatus(v string) *InstanceInfo
	GetStatus() *string
	SetVersionId(v string) *InstanceInfo
	GetVersionId() *string
}

type InstanceInfo struct {
	ImageUrl   *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	VersionId  *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s InstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s InstanceInfo) GoString() string {
	return s.String()
}

func (s *InstanceInfo) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *InstanceInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InstanceInfo) GetStatus() *string {
	return s.Status
}

func (s *InstanceInfo) GetVersionId() *string {
	return s.VersionId
}

func (s *InstanceInfo) SetImageUrl(v string) *InstanceInfo {
	s.ImageUrl = &v
	return s
}

func (s *InstanceInfo) SetInstanceId(v string) *InstanceInfo {
	s.InstanceId = &v
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
