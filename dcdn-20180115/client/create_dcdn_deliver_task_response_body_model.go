// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnDeliverTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliverId(v string) *CreateDcdnDeliverTaskResponseBody
	GetDeliverId() *string
	SetRequestId(v string) *CreateDcdnDeliverTaskResponseBody
	GetRequestId() *string
}

type CreateDcdnDeliverTaskResponseBody struct {
	// The ID of the change tracking task.
	//
	// example:
	//
	// 92
	DeliverId *string `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDcdnDeliverTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnDeliverTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDcdnDeliverTaskResponseBody) GetDeliverId() *string {
	return s.DeliverId
}

func (s *CreateDcdnDeliverTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDcdnDeliverTaskResponseBody) SetDeliverId(v string) *CreateDcdnDeliverTaskResponseBody {
	s.DeliverId = &v
	return s
}

func (s *CreateDcdnDeliverTaskResponseBody) SetRequestId(v string) *CreateDcdnDeliverTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDcdnDeliverTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
