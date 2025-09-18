// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceList(v []*QueryDeviceListResponseBodyDeviceList) *QueryDeviceListResponseBody
	GetDeviceList() []*QueryDeviceListResponseBodyDeviceList
	SetEncodeKey(v string) *QueryDeviceListResponseBody
	GetEncodeKey() *string
	SetEncodeType(v string) *QueryDeviceListResponseBody
	GetEncodeType() *string
	SetRequestId(v string) *QueryDeviceListResponseBody
	GetRequestId() *string
}

type QueryDeviceListResponseBody struct {
	DeviceList []*QueryDeviceListResponseBodyDeviceList `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Repeated"`
	// example:
	//
	// 125****0946
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDeviceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceListResponseBody) GetDeviceList() []*QueryDeviceListResponseBodyDeviceList {
	return s.DeviceList
}

func (s *QueryDeviceListResponseBody) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *QueryDeviceListResponseBody) GetEncodeType() *string {
	return s.EncodeType
}

func (s *QueryDeviceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDeviceListResponseBody) SetDeviceList(v []*QueryDeviceListResponseBodyDeviceList) *QueryDeviceListResponseBody {
	s.DeviceList = v
	return s
}

func (s *QueryDeviceListResponseBody) SetEncodeKey(v string) *QueryDeviceListResponseBody {
	s.EncodeKey = &v
	return s
}

func (s *QueryDeviceListResponseBody) SetEncodeType(v string) *QueryDeviceListResponseBody {
	s.EncodeType = &v
	return s
}

func (s *QueryDeviceListResponseBody) SetRequestId(v string) *QueryDeviceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryDeviceListResponseBodyDeviceList struct {
	// example:
	//
	// https://XXXXXX
	DeviceIconUrl *string `json:"DeviceIconUrl,omitempty" xml:"DeviceIconUrl,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// example:
	//
	// jMR2********ojVJXk=
	DeviceOpenId   *string                                                `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	DeviceUnionIds []*QueryDeviceListResponseBodyDeviceListDeviceUnionIds `json:"DeviceUnionIds,omitempty" xml:"DeviceUnionIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Online *string `json:"Online,omitempty" xml:"Online,omitempty"`
}

func (s QueryDeviceListResponseBodyDeviceList) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceListResponseBodyDeviceList) GoString() string {
	return s.String()
}

func (s *QueryDeviceListResponseBodyDeviceList) GetDeviceIconUrl() *string {
	return s.DeviceIconUrl
}

func (s *QueryDeviceListResponseBodyDeviceList) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryDeviceListResponseBodyDeviceList) GetDeviceOpenId() *string {
	return s.DeviceOpenId
}

func (s *QueryDeviceListResponseBodyDeviceList) GetDeviceUnionIds() []*QueryDeviceListResponseBodyDeviceListDeviceUnionIds {
	return s.DeviceUnionIds
}

func (s *QueryDeviceListResponseBodyDeviceList) GetOnline() *string {
	return s.Online
}

func (s *QueryDeviceListResponseBodyDeviceList) SetDeviceIconUrl(v string) *QueryDeviceListResponseBodyDeviceList {
	s.DeviceIconUrl = &v
	return s
}

func (s *QueryDeviceListResponseBodyDeviceList) SetDeviceName(v string) *QueryDeviceListResponseBodyDeviceList {
	s.DeviceName = &v
	return s
}

func (s *QueryDeviceListResponseBodyDeviceList) SetDeviceOpenId(v string) *QueryDeviceListResponseBodyDeviceList {
	s.DeviceOpenId = &v
	return s
}

func (s *QueryDeviceListResponseBodyDeviceList) SetDeviceUnionIds(v []*QueryDeviceListResponseBodyDeviceListDeviceUnionIds) *QueryDeviceListResponseBodyDeviceList {
	s.DeviceUnionIds = v
	return s
}

func (s *QueryDeviceListResponseBodyDeviceList) SetOnline(v string) *QueryDeviceListResponseBodyDeviceList {
	s.Online = &v
	return s
}

func (s *QueryDeviceListResponseBodyDeviceList) Validate() error {
	return dara.Validate(s)
}

type QueryDeviceListResponseBodyDeviceListDeviceUnionIds struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	UnionId        *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
}

func (s QueryDeviceListResponseBodyDeviceListDeviceUnionIds) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceListResponseBodyDeviceListDeviceUnionIds) GoString() string {
	return s.String()
}

func (s *QueryDeviceListResponseBodyDeviceListDeviceUnionIds) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *QueryDeviceListResponseBodyDeviceListDeviceUnionIds) GetUnionId() *string {
	return s.UnionId
}

func (s *QueryDeviceListResponseBodyDeviceListDeviceUnionIds) SetOrganizationId(v string) *QueryDeviceListResponseBodyDeviceListDeviceUnionIds {
	s.OrganizationId = &v
	return s
}

func (s *QueryDeviceListResponseBodyDeviceListDeviceUnionIds) SetUnionId(v string) *QueryDeviceListResponseBodyDeviceListDeviceUnionIds {
	s.UnionId = &v
	return s
}

func (s *QueryDeviceListResponseBodyDeviceListDeviceUnionIds) Validate() error {
	return dara.Validate(s)
}
