// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceActiveOpsTask interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceActiveOpsTaskId(v string) *InstanceActiveOpsTask
	GetInstanceActiveOpsTaskId() *string
	SetInstanceActiveOpsTaskStatus(v string) *InstanceActiveOpsTask
	GetInstanceActiveOpsTaskStatus() *string
}

type InstanceActiveOpsTask struct {
	InstanceActiveOpsTaskId     *string `json:"InstanceActiveOpsTaskId,omitempty" xml:"InstanceActiveOpsTaskId,omitempty"`
	InstanceActiveOpsTaskStatus *string `json:"InstanceActiveOpsTaskStatus,omitempty" xml:"InstanceActiveOpsTaskStatus,omitempty"`
}

func (s InstanceActiveOpsTask) String() string {
	return dara.Prettify(s)
}

func (s InstanceActiveOpsTask) GoString() string {
	return s.String()
}

func (s *InstanceActiveOpsTask) GetInstanceActiveOpsTaskId() *string {
	return s.InstanceActiveOpsTaskId
}

func (s *InstanceActiveOpsTask) GetInstanceActiveOpsTaskStatus() *string {
	return s.InstanceActiveOpsTaskStatus
}

func (s *InstanceActiveOpsTask) SetInstanceActiveOpsTaskId(v string) *InstanceActiveOpsTask {
	s.InstanceActiveOpsTaskId = &v
	return s
}

func (s *InstanceActiveOpsTask) SetInstanceActiveOpsTaskStatus(v string) *InstanceActiveOpsTask {
	s.InstanceActiveOpsTaskStatus = &v
	return s
}

func (s *InstanceActiveOpsTask) Validate() error {
	return dara.Validate(s)
}
