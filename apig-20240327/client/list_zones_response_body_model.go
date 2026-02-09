// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListZonesResponseBody
	GetCode() *string
	SetData(v *ListZonesResponseBodyData) *ListZonesResponseBody
	GetData() *ListZonesResponseBodyData
	SetMessage(v string) *ListZonesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListZonesResponseBody
	GetRequestId() *string
}

type ListZonesResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *ListZonesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// E8079207-B651-592A-A565-23E9EE5673B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListZonesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListZonesResponseBody) GetData() *ListZonesResponseBodyData {
	return s.Data
}

func (s *ListZonesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListZonesResponseBody) SetCode(v string) *ListZonesResponseBody {
	s.Code = &v
	return s
}

func (s *ListZonesResponseBody) SetData(v *ListZonesResponseBodyData) *ListZonesResponseBody {
	s.Data = v
	return s
}

func (s *ListZonesResponseBody) SetMessage(v string) *ListZonesResponseBody {
	s.Message = &v
	return s
}

func (s *ListZonesResponseBody) SetRequestId(v string) *ListZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListZonesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListZonesResponseBodyData struct {
	// The list of queried zones.
	Items []*ListZonesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListZonesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListZonesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListZonesResponseBodyData) GetItems() []*ListZonesResponseBodyDataItems {
	return s.Items
}

func (s *ListZonesResponseBodyData) SetItems(v []*ListZonesResponseBodyDataItems) *ListZonesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListZonesResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListZonesResponseBodyDataItems struct {
	// Whether QAT (Quality Assurance Testing) is supported in this zone
	//
	// example:
	//
	// true
	SupportQat *string `json:"supportQat,omitempty" xml:"supportQat,omitempty"`
	// The zone identifier
	//
	// example:
	//
	// cn-shenzhen-c
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListZonesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListZonesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListZonesResponseBodyDataItems) GetSupportQat() *string {
	return s.SupportQat
}

func (s *ListZonesResponseBodyDataItems) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListZonesResponseBodyDataItems) SetSupportQat(v string) *ListZonesResponseBodyDataItems {
	s.SupportQat = &v
	return s
}

func (s *ListZonesResponseBodyDataItems) SetZoneId(v string) *ListZonesResponseBodyDataItems {
	s.ZoneId = &v
	return s
}

func (s *ListZonesResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
