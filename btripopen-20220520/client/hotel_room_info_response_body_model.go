// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelRoomInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelRoomInfoResponseBody
	GetCode() *string
	SetMessage(v string) *HotelRoomInfoResponseBody
	GetMessage() *string
	SetModule(v []*HotelRoomInfoResponseBodyModule) *HotelRoomInfoResponseBody
	GetModule() []*HotelRoomInfoResponseBodyModule
	SetRequestId(v string) *HotelRoomInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelRoomInfoResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelRoomInfoResponseBody
	GetTraceId() *string
}

type HotelRoomInfoResponseBody struct {
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// operation success.
	Message *string                            `json:"message,omitempty" xml:"message,omitempty"`
	Module  []*HotelRoomInfoResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelRoomInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelRoomInfoResponseBody) GoString() string {
	return s.String()
}

func (s *HotelRoomInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelRoomInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelRoomInfoResponseBody) GetModule() []*HotelRoomInfoResponseBodyModule {
	return s.Module
}

func (s *HotelRoomInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelRoomInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelRoomInfoResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelRoomInfoResponseBody) SetCode(v string) *HotelRoomInfoResponseBody {
	s.Code = &v
	return s
}

func (s *HotelRoomInfoResponseBody) SetMessage(v string) *HotelRoomInfoResponseBody {
	s.Message = &v
	return s
}

func (s *HotelRoomInfoResponseBody) SetModule(v []*HotelRoomInfoResponseBodyModule) *HotelRoomInfoResponseBody {
	s.Module = v
	return s
}

func (s *HotelRoomInfoResponseBody) SetRequestId(v string) *HotelRoomInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelRoomInfoResponseBody) SetSuccess(v bool) *HotelRoomInfoResponseBody {
	s.Success = &v
	return s
}

func (s *HotelRoomInfoResponseBody) SetTraceId(v string) *HotelRoomInfoResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelRoomInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelRoomInfoResponseBodyModule struct {
	BedInfoGroupList []*HotelRoomInfoResponseBodyModuleBedInfoGroupList `json:"bed_info_group_list,omitempty" xml:"bed_info_group_list,omitempty" type:"Repeated"`
	BedInfos         []*HotelRoomInfoResponseBodyModuleBedInfos         `json:"bed_infos,omitempty" xml:"bed_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ExtraBed *int32 `json:"extra_bed,omitempty" xml:"extra_bed,omitempty"`
	// example:
	//
	// demo
	ExtraBedDesc *string `json:"extra_bed_desc,omitempty" xml:"extra_bed_desc,omitempty"`
	// example:
	//
	// 2
	Floor *string `json:"floor,omitempty" xml:"floor,omitempty"`
	// example:
	//
	// 0
	InternetWay *string `json:"internet_way,omitempty" xml:"internet_way,omitempty"`
	// example:
	//
	// 3
	MaxOccupancy     *int32    `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	RoomDesc         *string   `json:"room_desc,omitempty" xml:"room_desc,omitempty"`
	RoomFacilities   *string   `json:"room_facilities,omitempty" xml:"room_facilities,omitempty"`
	RoomFacilityList []*string `json:"room_facility_list,omitempty" xml:"room_facility_list,omitempty" type:"Repeated"`
	// example:
	//
	// 84536009
	RoomId *string `json:"room_id,omitempty" xml:"room_id,omitempty"`
	// example:
	//
	// //img.alicdn.com/imgextra/i2/6000000007493/O1CN010Vmxaz25DqUblX82A_!!6000000007493-2-hotel.png
	RoomImage  *string                                      `json:"room_image,omitempty" xml:"room_image,omitempty"`
	RoomImages []*HotelRoomInfoResponseBodyModuleRoomImages `json:"room_images,omitempty" xml:"room_images,omitempty" type:"Repeated"`
	RoomName   *string                                      `json:"room_name,omitempty" xml:"room_name,omitempty"`
	RoomType   *int32                                       `json:"room_type,omitempty" xml:"room_type,omitempty"`
	Roomarea   *string                                      `json:"roomarea,omitempty" xml:"roomarea,omitempty"`
	// example:
	//
	// 3
	Rooms *int32  `json:"rooms,omitempty" xml:"rooms,omitempty"`
	Smoke *string `json:"smoke,omitempty" xml:"smoke,omitempty"`
	// example:
	//
	// 2
	Window *string `json:"window,omitempty" xml:"window,omitempty"`
	// example:
	//
	// 0
	WindowBad *string `json:"window_bad,omitempty" xml:"window_bad,omitempty"`
	// example:
	//
	// demo
	WindowView *string `json:"window_view,omitempty" xml:"window_view,omitempty"`
}

func (s HotelRoomInfoResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelRoomInfoResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelRoomInfoResponseBodyModule) GetBedInfoGroupList() []*HotelRoomInfoResponseBodyModuleBedInfoGroupList {
	return s.BedInfoGroupList
}

func (s *HotelRoomInfoResponseBodyModule) GetBedInfos() []*HotelRoomInfoResponseBodyModuleBedInfos {
	return s.BedInfos
}

func (s *HotelRoomInfoResponseBodyModule) GetExtraBed() *int32 {
	return s.ExtraBed
}

func (s *HotelRoomInfoResponseBodyModule) GetExtraBedDesc() *string {
	return s.ExtraBedDesc
}

func (s *HotelRoomInfoResponseBodyModule) GetFloor() *string {
	return s.Floor
}

func (s *HotelRoomInfoResponseBodyModule) GetInternetWay() *string {
	return s.InternetWay
}

func (s *HotelRoomInfoResponseBodyModule) GetMaxOccupancy() *int32 {
	return s.MaxOccupancy
}

func (s *HotelRoomInfoResponseBodyModule) GetRoomDesc() *string {
	return s.RoomDesc
}

func (s *HotelRoomInfoResponseBodyModule) GetRoomFacilities() *string {
	return s.RoomFacilities
}

func (s *HotelRoomInfoResponseBodyModule) GetRoomFacilityList() []*string {
	return s.RoomFacilityList
}

func (s *HotelRoomInfoResponseBodyModule) GetRoomId() *string {
	return s.RoomId
}

func (s *HotelRoomInfoResponseBodyModule) GetRoomImage() *string {
	return s.RoomImage
}

func (s *HotelRoomInfoResponseBodyModule) GetRoomImages() []*HotelRoomInfoResponseBodyModuleRoomImages {
	return s.RoomImages
}

func (s *HotelRoomInfoResponseBodyModule) GetRoomName() *string {
	return s.RoomName
}

func (s *HotelRoomInfoResponseBodyModule) GetRoomType() *int32 {
	return s.RoomType
}

func (s *HotelRoomInfoResponseBodyModule) GetRoomarea() *string {
	return s.Roomarea
}

func (s *HotelRoomInfoResponseBodyModule) GetRooms() *int32 {
	return s.Rooms
}

func (s *HotelRoomInfoResponseBodyModule) GetSmoke() *string {
	return s.Smoke
}

func (s *HotelRoomInfoResponseBodyModule) GetWindow() *string {
	return s.Window
}

func (s *HotelRoomInfoResponseBodyModule) GetWindowBad() *string {
	return s.WindowBad
}

func (s *HotelRoomInfoResponseBodyModule) GetWindowView() *string {
	return s.WindowView
}

func (s *HotelRoomInfoResponseBodyModule) SetBedInfoGroupList(v []*HotelRoomInfoResponseBodyModuleBedInfoGroupList) *HotelRoomInfoResponseBodyModule {
	s.BedInfoGroupList = v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetBedInfos(v []*HotelRoomInfoResponseBodyModuleBedInfos) *HotelRoomInfoResponseBodyModule {
	s.BedInfos = v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetExtraBed(v int32) *HotelRoomInfoResponseBodyModule {
	s.ExtraBed = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetExtraBedDesc(v string) *HotelRoomInfoResponseBodyModule {
	s.ExtraBedDesc = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetFloor(v string) *HotelRoomInfoResponseBodyModule {
	s.Floor = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetInternetWay(v string) *HotelRoomInfoResponseBodyModule {
	s.InternetWay = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetMaxOccupancy(v int32) *HotelRoomInfoResponseBodyModule {
	s.MaxOccupancy = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetRoomDesc(v string) *HotelRoomInfoResponseBodyModule {
	s.RoomDesc = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetRoomFacilities(v string) *HotelRoomInfoResponseBodyModule {
	s.RoomFacilities = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetRoomFacilityList(v []*string) *HotelRoomInfoResponseBodyModule {
	s.RoomFacilityList = v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetRoomId(v string) *HotelRoomInfoResponseBodyModule {
	s.RoomId = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetRoomImage(v string) *HotelRoomInfoResponseBodyModule {
	s.RoomImage = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetRoomImages(v []*HotelRoomInfoResponseBodyModuleRoomImages) *HotelRoomInfoResponseBodyModule {
	s.RoomImages = v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetRoomName(v string) *HotelRoomInfoResponseBodyModule {
	s.RoomName = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetRoomType(v int32) *HotelRoomInfoResponseBodyModule {
	s.RoomType = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetRoomarea(v string) *HotelRoomInfoResponseBodyModule {
	s.Roomarea = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetRooms(v int32) *HotelRoomInfoResponseBodyModule {
	s.Rooms = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetSmoke(v string) *HotelRoomInfoResponseBodyModule {
	s.Smoke = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetWindow(v string) *HotelRoomInfoResponseBodyModule {
	s.Window = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetWindowBad(v string) *HotelRoomInfoResponseBodyModule {
	s.WindowBad = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) SetWindowView(v string) *HotelRoomInfoResponseBodyModule {
	s.WindowView = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelRoomInfoResponseBodyModuleBedInfoGroupList struct {
	BedInfos []*HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos `json:"bed_infos,omitempty" xml:"bed_infos,omitempty" type:"Repeated"`
}

func (s HotelRoomInfoResponseBodyModuleBedInfoGroupList) String() string {
	return dara.Prettify(s)
}

func (s HotelRoomInfoResponseBodyModuleBedInfoGroupList) GoString() string {
	return s.String()
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupList) GetBedInfos() []*HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos {
	return s.BedInfos
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupList) SetBedInfos(v []*HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) *HotelRoomInfoResponseBodyModuleBedInfoGroupList {
	s.BedInfos = v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupList) Validate() error {
	return dara.Validate(s)
}

type HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos struct {
	BedDesc *string `json:"bed_desc,omitempty" xml:"bed_desc,omitempty"`
	BedNum  *int32  `json:"bed_num,omitempty" xml:"bed_num,omitempty"`
	BedSize *string `json:"bed_size,omitempty" xml:"bed_size,omitempty"`
	BedType *string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	Length  *string `json:"length,omitempty" xml:"length,omitempty"`
	Width   *string `json:"width,omitempty" xml:"width,omitempty"`
}

func (s HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) String() string {
	return dara.Prettify(s)
}

func (s HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) GoString() string {
	return s.String()
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) GetBedDesc() *string {
	return s.BedDesc
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) GetBedNum() *int32 {
	return s.BedNum
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) GetBedSize() *string {
	return s.BedSize
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) GetBedType() *string {
	return s.BedType
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) GetLength() *string {
	return s.Length
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) GetWidth() *string {
	return s.Width
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) SetBedDesc(v string) *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos {
	s.BedDesc = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) SetBedNum(v int32) *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos {
	s.BedNum = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) SetBedSize(v string) *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos {
	s.BedSize = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) SetBedType(v string) *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos {
	s.BedType = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) SetLength(v string) *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos {
	s.Length = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) SetWidth(v string) *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos {
	s.Width = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleBedInfoGroupListBedInfos) Validate() error {
	return dara.Validate(s)
}

type HotelRoomInfoResponseBodyModuleBedInfos struct {
	BedDesc *string `json:"bed_desc,omitempty" xml:"bed_desc,omitempty"`
	// example:
	//
	// 2
	BedNum *int32 `json:"bed_num,omitempty" xml:"bed_num,omitempty"`
	// example:
	//
	// 1.8*2.0m
	BedSize *string `json:"bed_size,omitempty" xml:"bed_size,omitempty"`
	// example:
	//
	// 1
	BedType *string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	Length  *string `json:"length,omitempty" xml:"length,omitempty"`
	Width   *string `json:"width,omitempty" xml:"width,omitempty"`
}

func (s HotelRoomInfoResponseBodyModuleBedInfos) String() string {
	return dara.Prettify(s)
}

func (s HotelRoomInfoResponseBodyModuleBedInfos) GoString() string {
	return s.String()
}

func (s *HotelRoomInfoResponseBodyModuleBedInfos) GetBedDesc() *string {
	return s.BedDesc
}

func (s *HotelRoomInfoResponseBodyModuleBedInfos) GetBedNum() *int32 {
	return s.BedNum
}

func (s *HotelRoomInfoResponseBodyModuleBedInfos) GetBedSize() *string {
	return s.BedSize
}

func (s *HotelRoomInfoResponseBodyModuleBedInfos) GetBedType() *string {
	return s.BedType
}

func (s *HotelRoomInfoResponseBodyModuleBedInfos) GetLength() *string {
	return s.Length
}

func (s *HotelRoomInfoResponseBodyModuleBedInfos) GetWidth() *string {
	return s.Width
}

func (s *HotelRoomInfoResponseBodyModuleBedInfos) SetBedDesc(v string) *HotelRoomInfoResponseBodyModuleBedInfos {
	s.BedDesc = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleBedInfos) SetBedNum(v int32) *HotelRoomInfoResponseBodyModuleBedInfos {
	s.BedNum = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleBedInfos) SetBedSize(v string) *HotelRoomInfoResponseBodyModuleBedInfos {
	s.BedSize = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleBedInfos) SetBedType(v string) *HotelRoomInfoResponseBodyModuleBedInfos {
	s.BedType = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleBedInfos) SetLength(v string) *HotelRoomInfoResponseBodyModuleBedInfos {
	s.Length = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleBedInfos) SetWidth(v string) *HotelRoomInfoResponseBodyModuleBedInfos {
	s.Width = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleBedInfos) Validate() error {
	return dara.Validate(s)
}

type HotelRoomInfoResponseBodyModuleRoomImages struct {
	BedInfos2 *string `json:"bed_infos2,omitempty" xml:"bed_infos2,omitempty"`
	Tag       *int32  `json:"tag,omitempty" xml:"tag,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i3/6000000000003/O1CN01xkZQR41BtPxK1PQCb_!!6000000000003-0-hotel.jpg
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s HotelRoomInfoResponseBodyModuleRoomImages) String() string {
	return dara.Prettify(s)
}

func (s HotelRoomInfoResponseBodyModuleRoomImages) GoString() string {
	return s.String()
}

func (s *HotelRoomInfoResponseBodyModuleRoomImages) GetBedInfos2() *string {
	return s.BedInfos2
}

func (s *HotelRoomInfoResponseBodyModuleRoomImages) GetTag() *int32 {
	return s.Tag
}

func (s *HotelRoomInfoResponseBodyModuleRoomImages) GetUrl() *string {
	return s.Url
}

func (s *HotelRoomInfoResponseBodyModuleRoomImages) SetBedInfos2(v string) *HotelRoomInfoResponseBodyModuleRoomImages {
	s.BedInfos2 = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleRoomImages) SetTag(v int32) *HotelRoomInfoResponseBodyModuleRoomImages {
	s.Tag = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleRoomImages) SetUrl(v string) *HotelRoomInfoResponseBodyModuleRoomImages {
	s.Url = &v
	return s
}

func (s *HotelRoomInfoResponseBodyModuleRoomImages) Validate() error {
	return dara.Validate(s)
}
