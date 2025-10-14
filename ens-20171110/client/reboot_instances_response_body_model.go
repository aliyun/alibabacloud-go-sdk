// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceResponses(v []*InstanceOperateResponse) *RebootInstancesResponseBody
	GetInstanceResponses() []*InstanceOperateResponse
	SetRequestId(v string) *RebootInstancesResponseBody
	GetRequestId() *string
}

type RebootInstancesResponseBody struct {
	InstanceResponses []*InstanceOperateResponse `json:"InstanceResponses,omitempty" xml:"InstanceResponses,omitempty" type:"Repeated"`
	RequestId         *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RebootInstancesResponseBody) GetInstanceResponses() []*InstanceOperateResponse {
	return s.InstanceResponses
}

func (s *RebootInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootInstancesResponseBody) SetInstanceResponses(v []*InstanceOperateResponse) *RebootInstancesResponseBody {
	s.InstanceResponses = v
	return s
}

func (s *RebootInstancesResponseBody) SetRequestId(v string) *RebootInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootInstancesResponseBody) Validate() error {
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
