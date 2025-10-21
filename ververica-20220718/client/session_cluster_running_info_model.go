// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSessionClusterRunningInfo interface {
	dara.Model
	String() string
	GoString() string
	SetLastUpdateTime(v int64) *SessionClusterRunningInfo
	GetLastUpdateTime() *int64
	SetReferenceDeploymentIds(v []*string) *SessionClusterRunningInfo
	GetReferenceDeploymentIds() []*string
	SetStartedAt(v int64) *SessionClusterRunningInfo
	GetStartedAt() *int64
}

type SessionClusterRunningInfo struct {
	LastUpdateTime         *int64    `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty"`
	ReferenceDeploymentIds []*string `json:"referenceDeploymentIds,omitempty" xml:"referenceDeploymentIds,omitempty" type:"Repeated"`
	StartedAt              *int64    `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
}

func (s SessionClusterRunningInfo) String() string {
	return dara.Prettify(s)
}

func (s SessionClusterRunningInfo) GoString() string {
	return s.String()
}

func (s *SessionClusterRunningInfo) GetLastUpdateTime() *int64 {
	return s.LastUpdateTime
}

func (s *SessionClusterRunningInfo) GetReferenceDeploymentIds() []*string {
	return s.ReferenceDeploymentIds
}

func (s *SessionClusterRunningInfo) GetStartedAt() *int64 {
	return s.StartedAt
}

func (s *SessionClusterRunningInfo) SetLastUpdateTime(v int64) *SessionClusterRunningInfo {
	s.LastUpdateTime = &v
	return s
}

func (s *SessionClusterRunningInfo) SetReferenceDeploymentIds(v []*string) *SessionClusterRunningInfo {
	s.ReferenceDeploymentIds = v
	return s
}

func (s *SessionClusterRunningInfo) SetStartedAt(v int64) *SessionClusterRunningInfo {
	s.StartedAt = &v
	return s
}

func (s *SessionClusterRunningInfo) Validate() error {
	return dara.Validate(s)
}
