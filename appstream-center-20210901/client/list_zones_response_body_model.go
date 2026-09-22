// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListZonesModel(v *ListZonesResponseBodyListZonesModel) *ListZonesResponseBody
	GetListZonesModel() *ListZonesResponseBodyListZonesModel
	SetRequestId(v string) *ListZonesResponseBody
	GetRequestId() *string
}

type ListZonesResponseBody struct {
	// The zone query result.
	ListZonesModel *ListZonesResponseBodyListZonesModel `json:"ListZonesModel,omitempty" xml:"ListZonesModel,omitempty" type:"Struct"`
	// The request ID. You can use this ID to locate and troubleshoot issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListZonesResponseBody) GetListZonesModel() *ListZonesResponseBodyListZonesModel {
	return s.ListZonesModel
}

func (s *ListZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListZonesResponseBody) SetListZonesModel(v *ListZonesResponseBodyListZonesModel) *ListZonesResponseBody {
	s.ListZonesModel = v
	return s
}

func (s *ListZonesResponseBody) SetRequestId(v string) *ListZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListZonesResponseBody) Validate() error {
	if s.ListZonesModel != nil {
		if err := s.ListZonesModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListZonesResponseBodyListZonesModel struct {
	// The list of available zone IDs for the specified product type and operating system type in the current region. When creating a resource that requires a vSwitch, select a vSwitch in one of these zones.
	Zones []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListZonesResponseBodyListZonesModel) String() string {
	return dara.Prettify(s)
}

func (s ListZonesResponseBodyListZonesModel) GoString() string {
	return s.String()
}

func (s *ListZonesResponseBodyListZonesModel) GetZones() []*string {
	return s.Zones
}

func (s *ListZonesResponseBodyListZonesModel) SetZones(v []*string) *ListZonesResponseBodyListZonesModel {
	s.Zones = v
	return s
}

func (s *ListZonesResponseBodyListZonesModel) Validate() error {
	return dara.Validate(s)
}
