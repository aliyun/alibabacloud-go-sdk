// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentError(v string) *ListInstancesOutput
	GetCurrentError() *string
	SetInstances(v []*InstanceInfo) *ListInstancesOutput
	GetInstances() []*InstanceInfo
	SetRequestId(v string) *ListInstancesOutput
	GetRequestId() *string
	SetVersionStatus(v map[string]*VersionStatus) *ListInstancesOutput
	GetVersionStatus() map[string]*VersionStatus
}

type ListInstancesOutput struct {
	CurrentError  *string                   `json:"currentError,omitempty" xml:"currentError,omitempty"`
	Instances     []*InstanceInfo           `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	RequestId     *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	VersionStatus map[string]*VersionStatus `json:"versionStatus,omitempty" xml:"versionStatus,omitempty"`
}

func (s ListInstancesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesOutput) GoString() string {
	return s.String()
}

func (s *ListInstancesOutput) GetCurrentError() *string {
	return s.CurrentError
}

func (s *ListInstancesOutput) GetInstances() []*InstanceInfo {
	return s.Instances
}

func (s *ListInstancesOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesOutput) GetVersionStatus() map[string]*VersionStatus {
	return s.VersionStatus
}

func (s *ListInstancesOutput) SetCurrentError(v string) *ListInstancesOutput {
	s.CurrentError = &v
	return s
}

func (s *ListInstancesOutput) SetInstances(v []*InstanceInfo) *ListInstancesOutput {
	s.Instances = v
	return s
}

func (s *ListInstancesOutput) SetRequestId(v string) *ListInstancesOutput {
	s.RequestId = &v
	return s
}

func (s *ListInstancesOutput) SetVersionStatus(v map[string]*VersionStatus) *ListInstancesOutput {
	s.VersionStatus = v
	return s
}

func (s *ListInstancesOutput) Validate() error {
	return dara.Validate(s)
}
