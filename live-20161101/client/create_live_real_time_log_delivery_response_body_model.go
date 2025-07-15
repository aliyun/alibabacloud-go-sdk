// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveRealTimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLiveRealTimeLogDeliveryResponseBody
	GetRequestId() *string
}

type CreateLiveRealTimeLogDeliveryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F32C57AA-7BF8-49AE-A2CC-9F42390F5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLiveRealTimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveRealTimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveRealTimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLiveRealTimeLogDeliveryResponseBody) SetRequestId(v string) *CreateLiveRealTimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveRealTimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
