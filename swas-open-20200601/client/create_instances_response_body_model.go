// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *CreateInstancesResponseBody
	GetInstanceIds() []*string
	SetRequestId(v string) *CreateInstancesResponseBody
	GetRequestId() *string
}

type CreateInstancesResponseBody struct {
	// The instance IDs of the Simple Application Server instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E1FEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstancesResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstancesResponseBody) SetInstanceIds(v []*string) *CreateInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateInstancesResponseBody) SetRequestId(v string) *CreateInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
