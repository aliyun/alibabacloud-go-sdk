// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCdnDeliverTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliverId(v string) *CreateCdnDeliverTaskResponseBody
	GetDeliverId() *string
	SetRequestId(v string) *CreateCdnDeliverTaskResponseBody
	GetRequestId() *string
}

type CreateCdnDeliverTaskResponseBody struct {
	// The ID of the tracking task.
	//
	// example:
	//
	// 1025
	DeliverId *string `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCdnDeliverTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCdnDeliverTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCdnDeliverTaskResponseBody) GetDeliverId() *string {
	return s.DeliverId
}

func (s *CreateCdnDeliverTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCdnDeliverTaskResponseBody) SetDeliverId(v string) *CreateCdnDeliverTaskResponseBody {
	s.DeliverId = &v
	return s
}

func (s *CreateCdnDeliverTaskResponseBody) SetRequestId(v string) *CreateCdnDeliverTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCdnDeliverTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
