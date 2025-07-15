// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectTrafficQosQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQosId(v string) *CreateExpressConnectTrafficQosQueueResponseBody
	GetQosId() *string
	SetQueueId(v string) *CreateExpressConnectTrafficQosQueueResponseBody
	GetQueueId() *string
	SetRequestId(v string) *CreateExpressConnectTrafficQosQueueResponseBody
	GetRequestId() *string
}

type CreateExpressConnectTrafficQosQueueResponseBody struct {
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-2giu0a6vd5x0mv4700
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the QoS queue.
	//
	// example:
	//
	// qos-queue-9nyx2u7n71s2rcy4n5
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4CF20CC7-D1FC-425B-A15B-DF7C8E2131A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExpressConnectTrafficQosQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectTrafficQosQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectTrafficQosQueueResponseBody) GetQosId() *string {
	return s.QosId
}

func (s *CreateExpressConnectTrafficQosQueueResponseBody) GetQueueId() *string {
	return s.QueueId
}

func (s *CreateExpressConnectTrafficQosQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExpressConnectTrafficQosQueueResponseBody) SetQosId(v string) *CreateExpressConnectTrafficQosQueueResponseBody {
	s.QosId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueResponseBody) SetQueueId(v string) *CreateExpressConnectTrafficQosQueueResponseBody {
	s.QueueId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueResponseBody) SetRequestId(v string) *CreateExpressConnectTrafficQosQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
