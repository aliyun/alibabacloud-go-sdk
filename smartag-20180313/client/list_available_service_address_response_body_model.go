// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableServiceAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAvailableServiceAddressResponseBody
	GetCode() *string
	SetData(v []*ListAvailableServiceAddressResponseBodyData) *ListAvailableServiceAddressResponseBody
	GetData() []*ListAvailableServiceAddressResponseBodyData
	SetMessage(v string) *ListAvailableServiceAddressResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAvailableServiceAddressResponseBody
	GetRequestId() *string
}

type ListAvailableServiceAddressResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListAvailableServiceAddressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 324223F3-93D3-4CE4-B26F-66C0C3809922
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAvailableServiceAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableServiceAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableServiceAddressResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAvailableServiceAddressResponseBody) GetData() []*ListAvailableServiceAddressResponseBodyData {
	return s.Data
}

func (s *ListAvailableServiceAddressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAvailableServiceAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAvailableServiceAddressResponseBody) SetCode(v string) *ListAvailableServiceAddressResponseBody {
	s.Code = &v
	return s
}

func (s *ListAvailableServiceAddressResponseBody) SetData(v []*ListAvailableServiceAddressResponseBodyData) *ListAvailableServiceAddressResponseBody {
	s.Data = v
	return s
}

func (s *ListAvailableServiceAddressResponseBody) SetMessage(v string) *ListAvailableServiceAddressResponseBody {
	s.Message = &v
	return s
}

func (s *ListAvailableServiceAddressResponseBody) SetRequestId(v string) *ListAvailableServiceAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableServiceAddressResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAvailableServiceAddressResponseBodyData struct {
	// The service address.
	//
	// example:
	//
	// 192.168.1.1
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The type of service address.
	//
	// example:
	//
	// ProbeTask
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
}

func (s ListAvailableServiceAddressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableServiceAddressResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAvailableServiceAddressResponseBodyData) GetAddress() *string {
	return s.Address
}

func (s *ListAvailableServiceAddressResponseBodyData) GetAddressType() *string {
	return s.AddressType
}

func (s *ListAvailableServiceAddressResponseBodyData) SetAddress(v string) *ListAvailableServiceAddressResponseBodyData {
	s.Address = &v
	return s
}

func (s *ListAvailableServiceAddressResponseBodyData) SetAddressType(v string) *ListAvailableServiceAddressResponseBodyData {
	s.AddressType = &v
	return s
}

func (s *ListAvailableServiceAddressResponseBodyData) Validate() error {
	return dara.Validate(s)
}
