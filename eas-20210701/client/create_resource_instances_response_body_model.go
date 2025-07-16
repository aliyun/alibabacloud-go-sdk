// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *CreateResourceInstancesResponseBody
	GetInstanceIds() []*string
	SetMessage(v string) *CreateResourceInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateResourceInstancesResponseBody
	GetRequestId() *string
}

type CreateResourceInstancesResponseBody struct {
	// The instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// Create 5 new ecs instance(s) in resource [eas-r-asdasdasd] successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceInstancesResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateResourceInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateResourceInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceInstancesResponseBody) SetInstanceIds(v []*string) *CreateResourceInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateResourceInstancesResponseBody) SetMessage(v string) *CreateResourceInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *CreateResourceInstancesResponseBody) SetRequestId(v string) *CreateResourceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
