// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWuyingServerSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModifyWuyingServerSpecModel(v *ModifyWuyingServerSpecResponseBodyModifyWuyingServerSpecModel) *ModifyWuyingServerSpecResponseBody
	GetModifyWuyingServerSpecModel() *ModifyWuyingServerSpecResponseBodyModifyWuyingServerSpecModel
	SetRequestId(v string) *ModifyWuyingServerSpecResponseBody
	GetRequestId() *string
}

type ModifyWuyingServerSpecResponseBody struct {
	// The processing result of the specification change.
	ModifyWuyingServerSpecModel *ModifyWuyingServerSpecResponseBodyModifyWuyingServerSpecModel `json:"ModifyWuyingServerSpecModel,omitempty" xml:"ModifyWuyingServerSpecModel,omitempty" type:"Struct"`
	// The ID of the request. It is used for troubleshooting.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWuyingServerSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWuyingServerSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWuyingServerSpecResponseBody) GetModifyWuyingServerSpecModel() *ModifyWuyingServerSpecResponseBodyModifyWuyingServerSpecModel {
	return s.ModifyWuyingServerSpecModel
}

func (s *ModifyWuyingServerSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWuyingServerSpecResponseBody) SetModifyWuyingServerSpecModel(v *ModifyWuyingServerSpecResponseBodyModifyWuyingServerSpecModel) *ModifyWuyingServerSpecResponseBody {
	s.ModifyWuyingServerSpecModel = v
	return s
}

func (s *ModifyWuyingServerSpecResponseBody) SetRequestId(v string) *ModifyWuyingServerSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWuyingServerSpecResponseBody) Validate() error {
	if s.ModifyWuyingServerSpecModel != nil {
		if err := s.ModifyWuyingServerSpecModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyWuyingServerSpecResponseBodyModifyWuyingServerSpecModel struct {
	// The ID of the order generated for this specification change. Returning an order ID does not mean that the specification change is complete.
	//
	// example:
	//
	// 23429322113****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyWuyingServerSpecResponseBodyModifyWuyingServerSpecModel) String() string {
	return dara.Prettify(s)
}

func (s ModifyWuyingServerSpecResponseBodyModifyWuyingServerSpecModel) GoString() string {
	return s.String()
}

func (s *ModifyWuyingServerSpecResponseBodyModifyWuyingServerSpecModel) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyWuyingServerSpecResponseBodyModifyWuyingServerSpecModel) SetOrderId(v string) *ModifyWuyingServerSpecResponseBodyModifyWuyingServerSpecModel {
	s.OrderId = &v
	return s
}

func (s *ModifyWuyingServerSpecResponseBodyModifyWuyingServerSpecModel) Validate() error {
	return dara.Validate(s)
}
