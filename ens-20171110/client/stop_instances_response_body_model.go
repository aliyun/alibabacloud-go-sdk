// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceResponses(v []*InstanceOperateResponse) *StopInstancesResponseBody
	GetInstanceResponses() []*InstanceOperateResponse
	SetRequestId(v string) *StopInstancesResponseBody
	GetRequestId() *string
}

type StopInstancesResponseBody struct {
	InstanceResponses []*InstanceOperateResponse `json:"InstanceResponses,omitempty" xml:"InstanceResponses,omitempty" type:"Repeated"`
	RequestId         *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstancesResponseBody) GetInstanceResponses() []*InstanceOperateResponse {
	return s.InstanceResponses
}

func (s *StopInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopInstancesResponseBody) SetInstanceResponses(v []*InstanceOperateResponse) *StopInstancesResponseBody {
	s.InstanceResponses = v
	return s
}

func (s *StopInstancesResponseBody) SetRequestId(v string) *StopInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstancesResponseBody) Validate() error {
	if s.InstanceResponses != nil {
		for _, item := range s.InstanceResponses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
