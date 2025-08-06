// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodRealTimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVodRealTimeLogDeliveryResponseBody
	GetRequestId() *string
}

type CreateVodRealTimeLogDeliveryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVodRealTimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVodRealTimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVodRealTimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVodRealTimeLogDeliveryResponseBody) SetRequestId(v string) *CreateVodRealTimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVodRealTimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
