// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWuyingServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateWuyingServerResponseBodyData) *CreateWuyingServerResponseBody
	GetData() *CreateWuyingServerResponseBodyData
	SetRequestId(v string) *CreateWuyingServerResponseBody
	GetRequestId() *string
}

type CreateWuyingServerResponseBody struct {
	Data *CreateWuyingServerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWuyingServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWuyingServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWuyingServerResponseBody) GetData() *CreateWuyingServerResponseBodyData {
	return s.Data
}

func (s *CreateWuyingServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWuyingServerResponseBody) SetData(v *CreateWuyingServerResponseBodyData) *CreateWuyingServerResponseBody {
	s.Data = v
	return s
}

func (s *CreateWuyingServerResponseBody) SetRequestId(v string) *CreateWuyingServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWuyingServerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWuyingServerResponseBodyData struct {
	// example:
	//
	// 23429322113****
	OrderId            *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	WuyingServerIdList []*string `json:"WuyingServerIdList,omitempty" xml:"WuyingServerIdList,omitempty" type:"Repeated"`
}

func (s CreateWuyingServerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWuyingServerResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWuyingServerResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateWuyingServerResponseBodyData) GetWuyingServerIdList() []*string {
	return s.WuyingServerIdList
}

func (s *CreateWuyingServerResponseBodyData) SetOrderId(v string) *CreateWuyingServerResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateWuyingServerResponseBodyData) SetWuyingServerIdList(v []*string) *CreateWuyingServerResponseBodyData {
	s.WuyingServerIdList = v
	return s
}

func (s *CreateWuyingServerResponseBodyData) Validate() error {
	return dara.Validate(s)
}
