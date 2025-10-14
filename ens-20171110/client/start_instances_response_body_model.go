// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceResponses(v []*InstanceOperateResponse) *StartInstancesResponseBody
	GetInstanceResponses() []*InstanceOperateResponse
	SetRequestId(v string) *StartInstancesResponseBody
	GetRequestId() *string
}

type StartInstancesResponseBody struct {
	InstanceResponses []*InstanceOperateResponse `json:"InstanceResponses,omitempty" xml:"InstanceResponses,omitempty" type:"Repeated"`
	RequestId         *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstancesResponseBody) GetInstanceResponses() []*InstanceOperateResponse {
	return s.InstanceResponses
}

func (s *StartInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartInstancesResponseBody) SetInstanceResponses(v []*InstanceOperateResponse) *StartInstancesResponseBody {
	s.InstanceResponses = v
	return s
}

func (s *StartInstancesResponseBody) SetRequestId(v string) *StartInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartInstancesResponseBody) Validate() error {
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
