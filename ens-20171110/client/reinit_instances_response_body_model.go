// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinitInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceResponses(v []*InstanceOperateResponse) *ReinitInstancesResponseBody
	GetInstanceResponses() []*InstanceOperateResponse
	SetRequestId(v string) *ReinitInstancesResponseBody
	GetRequestId() *string
}

type ReinitInstancesResponseBody struct {
	InstanceResponses []*InstanceOperateResponse `json:"InstanceResponses,omitempty" xml:"InstanceResponses,omitempty" type:"Repeated"`
	RequestId         *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReinitInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReinitInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ReinitInstancesResponseBody) GetInstanceResponses() []*InstanceOperateResponse {
	return s.InstanceResponses
}

func (s *ReinitInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReinitInstancesResponseBody) SetInstanceResponses(v []*InstanceOperateResponse) *ReinitInstancesResponseBody {
	s.InstanceResponses = v
	return s
}

func (s *ReinitInstancesResponseBody) SetRequestId(v string) *ReinitInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReinitInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
