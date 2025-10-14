// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceActiveOpsTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceActiveOpsTask(v *InstanceActiveOpsTask) *CreateInstanceActiveOpsTaskResponseBody
	GetInstanceActiveOpsTask() *InstanceActiveOpsTask
	SetRequestId(v string) *CreateInstanceActiveOpsTaskResponseBody
	GetRequestId() *string
}

type CreateInstanceActiveOpsTaskResponseBody struct {
	InstanceActiveOpsTask *InstanceActiveOpsTask `json:"InstanceActiveOpsTask,omitempty" xml:"InstanceActiveOpsTask,omitempty"`
	RequestId             *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceActiveOpsTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceActiveOpsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceActiveOpsTaskResponseBody) GetInstanceActiveOpsTask() *InstanceActiveOpsTask {
	return s.InstanceActiveOpsTask
}

func (s *CreateInstanceActiveOpsTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceActiveOpsTaskResponseBody) SetInstanceActiveOpsTask(v *InstanceActiveOpsTask) *CreateInstanceActiveOpsTaskResponseBody {
	s.InstanceActiveOpsTask = v
	return s
}

func (s *CreateInstanceActiveOpsTaskResponseBody) SetRequestId(v string) *CreateInstanceActiveOpsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceActiveOpsTaskResponseBody) Validate() error {
	if s.InstanceActiveOpsTask != nil {
		if err := s.InstanceActiveOpsTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}
