// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*InstanceInfo) *ListInstancesOutput
	GetInstances() []*InstanceInfo
	SetRequestId(v string) *ListInstancesOutput
	GetRequestId() *string
}

type ListInstancesOutput struct {
	Instances []*InstanceInfo `json:"instances" xml:"instances" type:"Repeated"`
	RequestId *string         `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListInstancesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesOutput) GoString() string {
	return s.String()
}

func (s *ListInstancesOutput) GetInstances() []*InstanceInfo {
	return s.Instances
}

func (s *ListInstancesOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesOutput) SetInstances(v []*InstanceInfo) *ListInstancesOutput {
	s.Instances = v
	return s
}

func (s *ListInstancesOutput) SetRequestId(v string) *ListInstancesOutput {
	s.RequestId = &v
	return s
}

func (s *ListInstancesOutput) Validate() error {
	return dara.Validate(s)
}
