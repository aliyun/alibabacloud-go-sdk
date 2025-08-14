// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenInterRegionTrafficQosQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQosQueueId(v string) *CreateCenInterRegionTrafficQosQueueResponseBody
	GetQosQueueId() *string
	SetRequestId(v string) *CreateCenInterRegionTrafficQosQueueResponseBody
	GetRequestId() *string
}

type CreateCenInterRegionTrafficQosQueueResponseBody struct {
	// The ID of the queue.
	//
	// example:
	//
	// qos-queue-irqhi8k5fdyuu5****
	QosQueueId *string `json:"QosQueueId,omitempty" xml:"QosQueueId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 845F66F6-5C27-53A1-9428-B859086237B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCenInterRegionTrafficQosQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCenInterRegionTrafficQosQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenInterRegionTrafficQosQueueResponseBody) GetQosQueueId() *string {
	return s.QosQueueId
}

func (s *CreateCenInterRegionTrafficQosQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCenInterRegionTrafficQosQueueResponseBody) SetQosQueueId(v string) *CreateCenInterRegionTrafficQosQueueResponseBody {
	s.QosQueueId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueResponseBody) SetRequestId(v string) *CreateCenInterRegionTrafficQosQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
