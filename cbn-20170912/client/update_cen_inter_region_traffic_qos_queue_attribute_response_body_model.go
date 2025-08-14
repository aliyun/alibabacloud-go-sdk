// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCenInterRegionTrafficQosQueueAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCenInterRegionTrafficQosQueueAttributeResponseBody
	GetRequestId() *string
}

type UpdateCenInterRegionTrafficQosQueueAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6DF9A765-BCD2-5C7E-8C32-C35C8A361A39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCenInterRegionTrafficQosQueueAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCenInterRegionTrafficQosQueueAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeResponseBody) SetRequestId(v string) *UpdateCenInterRegionTrafficQosQueueAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
