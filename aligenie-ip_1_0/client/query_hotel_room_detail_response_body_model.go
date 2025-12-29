// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHotelRoomDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *QueryHotelRoomDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryHotelRoomDetailResponseBody
	GetRequestId() *string
	SetResult(v *QueryHotelRoomDetailResponseBodyResult) *QueryHotelRoomDetailResponseBody
	GetResult() *QueryHotelRoomDetailResponseBodyResult
	SetStatusCode(v int32) *QueryHotelRoomDetailResponseBody
	GetStatusCode() *int32
}

type QueryHotelRoomDetailResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryHotelRoomDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QueryHotelRoomDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryHotelRoomDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryHotelRoomDetailResponseBody) GetResult() *QueryHotelRoomDetailResponseBodyResult {
	return s.Result
}

func (s *QueryHotelRoomDetailResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryHotelRoomDetailResponseBody) SetMessage(v string) *QueryHotelRoomDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBody) SetRequestId(v string) *QueryHotelRoomDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBody) SetResult(v *QueryHotelRoomDetailResponseBodyResult) *QueryHotelRoomDetailResponseBody {
	s.Result = v
	return s
}

func (s *QueryHotelRoomDetailResponseBody) SetStatusCode(v int32) *QueryHotelRoomDetailResponseBody {
	s.StatusCode = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryHotelRoomDetailResponseBodyResult struct {
	AuthAccounts []*QueryHotelRoomDetailResponseBodyResultAuthAccounts `json:"AuthAccounts,omitempty" xml:"AuthAccounts,omitempty" type:"Repeated"`
	// example:
	//
	// rcu
	ConnectType        *string                                              `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	CreatorAccountName *string                                              `json:"CreatorAccountName,omitempty" xml:"CreatorAccountName,omitempty"`
	DeviceInfos        []*QueryHotelRoomDetailResponseBodyResultDeviceInfos `json:"DeviceInfos,omitempty" xml:"DeviceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// a7***83
	HotelId         *string                                                `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelName       *string                                                `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	OtherService    *QueryHotelRoomDetailResponseBodyResultOtherService    `json:"OtherService,omitempty" xml:"OtherService,omitempty" type:"Struct"`
	RoomControlInfo *QueryHotelRoomDetailResponseBodyResultRoomControlInfo `json:"RoomControlInfo,omitempty" xml:"RoomControlInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2001
	RoomNo          *string                                                `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	RoomServiceInfo *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo `json:"RoomServiceInfo,omitempty" xml:"RoomServiceInfo,omitempty" type:"Struct"`
}

func (s QueryHotelRoomDetailResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBodyResult) GetAuthAccounts() []*QueryHotelRoomDetailResponseBodyResultAuthAccounts {
	return s.AuthAccounts
}

func (s *QueryHotelRoomDetailResponseBodyResult) GetConnectType() *string {
	return s.ConnectType
}

func (s *QueryHotelRoomDetailResponseBodyResult) GetCreatorAccountName() *string {
	return s.CreatorAccountName
}

func (s *QueryHotelRoomDetailResponseBodyResult) GetDeviceInfos() []*QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	return s.DeviceInfos
}

func (s *QueryHotelRoomDetailResponseBodyResult) GetHotelId() *string {
	return s.HotelId
}

func (s *QueryHotelRoomDetailResponseBodyResult) GetHotelName() *string {
	return s.HotelName
}

func (s *QueryHotelRoomDetailResponseBodyResult) GetOtherService() *QueryHotelRoomDetailResponseBodyResultOtherService {
	return s.OtherService
}

func (s *QueryHotelRoomDetailResponseBodyResult) GetRoomControlInfo() *QueryHotelRoomDetailResponseBodyResultRoomControlInfo {
	return s.RoomControlInfo
}

func (s *QueryHotelRoomDetailResponseBodyResult) GetRoomNo() *string {
	return s.RoomNo
}

func (s *QueryHotelRoomDetailResponseBodyResult) GetRoomServiceInfo() *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo {
	return s.RoomServiceInfo
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetAuthAccounts(v []*QueryHotelRoomDetailResponseBodyResultAuthAccounts) *QueryHotelRoomDetailResponseBodyResult {
	s.AuthAccounts = v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetConnectType(v string) *QueryHotelRoomDetailResponseBodyResult {
	s.ConnectType = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetCreatorAccountName(v string) *QueryHotelRoomDetailResponseBodyResult {
	s.CreatorAccountName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetDeviceInfos(v []*QueryHotelRoomDetailResponseBodyResultDeviceInfos) *QueryHotelRoomDetailResponseBodyResult {
	s.DeviceInfos = v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetHotelId(v string) *QueryHotelRoomDetailResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetHotelName(v string) *QueryHotelRoomDetailResponseBodyResult {
	s.HotelName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetOtherService(v *QueryHotelRoomDetailResponseBodyResultOtherService) *QueryHotelRoomDetailResponseBodyResult {
	s.OtherService = v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetRoomControlInfo(v *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) *QueryHotelRoomDetailResponseBodyResult {
	s.RoomControlInfo = v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetRoomNo(v string) *QueryHotelRoomDetailResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) SetRoomServiceInfo(v *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) *QueryHotelRoomDetailResponseBodyResult {
	s.RoomServiceInfo = v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResult) Validate() error {
	if s.AuthAccounts != nil {
		for _, item := range s.AuthAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DeviceInfos != nil {
		for _, item := range s.DeviceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OtherService != nil {
		if err := s.OtherService.Validate(); err != nil {
			return err
		}
	}
	if s.RoomControlInfo != nil {
		if err := s.RoomControlInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RoomServiceInfo != nil {
		if err := s.RoomServiceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryHotelRoomDetailResponseBodyResultAuthAccounts struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 2023-01-01 12:00:00
	AuthTime *string `json:"AuthTime,omitempty" xml:"AuthTime,omitempty"`
}

func (s QueryHotelRoomDetailResponseBodyResultAuthAccounts) String() string {
	return dara.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBodyResultAuthAccounts) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBodyResultAuthAccounts) GetAccountName() *string {
	return s.AccountName
}

func (s *QueryHotelRoomDetailResponseBodyResultAuthAccounts) GetAuthTime() *string {
	return s.AuthTime
}

func (s *QueryHotelRoomDetailResponseBodyResultAuthAccounts) SetAccountName(v string) *QueryHotelRoomDetailResponseBodyResultAuthAccounts {
	s.AccountName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultAuthAccounts) SetAuthTime(v string) *QueryHotelRoomDetailResponseBodyResultAuthAccounts {
	s.AuthTime = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultAuthAccounts) Validate() error {
	return dara.Validate(s)
}

type QueryHotelRoomDetailResponseBodyResultDeviceInfos struct {
	// example:
	//
	// 2023-01-01 12:00:00
	ActiveTime *string `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// example:
	//
	// 6.1.8-RS-20230425.1806
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" xml:"FirmwareVersion,omitempty"`
	// example:
	//
	// fa:03:23:58:c3:00
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 1
	OnlineStatus *int32 `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// example:
	//
	// sag42dlz4qf
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// example:
	//
	// 41c95c18a0a643bcb58edf438877def5
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryHotelRoomDetailResponseBodyResultDeviceInfos) String() string {
	return dara.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBodyResultDeviceInfos) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) GetActiveTime() *string {
	return s.ActiveTime
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) GetFirmwareVersion() *string {
	return s.FirmwareVersion
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) GetMac() *string {
	return s.Mac
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) GetOnlineStatus() *int32 {
	return s.OnlineStatus
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) GetSn() *string {
	return s.Sn
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) GetUuid() *string {
	return s.Uuid
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) SetActiveTime(v string) *QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	s.ActiveTime = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) SetDeviceName(v string) *QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	s.DeviceName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) SetFirmwareVersion(v string) *QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	s.FirmwareVersion = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) SetMac(v string) *QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	s.Mac = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) SetOnlineStatus(v int32) *QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	s.OnlineStatus = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) SetSn(v string) *QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	s.Sn = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) SetUuid(v string) *QueryHotelRoomDetailResponseBodyResultDeviceInfos {
	s.Uuid = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultDeviceInfos) Validate() error {
	return dara.Validate(s)
}

type QueryHotelRoomDetailResponseBodyResultOtherService struct {
	// example:
	//
	// false
	OpenCall *bool `json:"OpenCall,omitempty" xml:"OpenCall,omitempty"`
	// example:
	//
	// 0
	UnhandleTickets *int32 `json:"UnhandleTickets,omitempty" xml:"UnhandleTickets,omitempty"`
}

func (s QueryHotelRoomDetailResponseBodyResultOtherService) String() string {
	return dara.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBodyResultOtherService) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBodyResultOtherService) GetOpenCall() *bool {
	return s.OpenCall
}

func (s *QueryHotelRoomDetailResponseBodyResultOtherService) GetUnhandleTickets() *int32 {
	return s.UnhandleTickets
}

func (s *QueryHotelRoomDetailResponseBodyResultOtherService) SetOpenCall(v bool) *QueryHotelRoomDetailResponseBodyResultOtherService {
	s.OpenCall = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultOtherService) SetUnhandleTickets(v int32) *QueryHotelRoomDetailResponseBodyResultOtherService {
	s.UnhandleTickets = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultOtherService) Validate() error {
	return dara.Validate(s)
}

type QueryHotelRoomDetailResponseBodyResultRoomControlInfo struct {
	// example:
	//
	// 78
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// app
	AppName     *string                                                             `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DeviceInfos []*QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos `json:"DeviceInfos,omitempty" xml:"DeviceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// http://www.xxx.com
	RcuUrl *string `json:"RcuUrl,omitempty" xml:"RcuUrl,omitempty"`
	// example:
	//
	// 1170
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s QueryHotelRoomDetailResponseBodyResultRoomControlInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBodyResultRoomControlInfo) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) GetAppId() *int64 {
	return s.AppId
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) GetAppName() *string {
	return s.AppName
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) GetDeviceInfos() []*QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	return s.DeviceInfos
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) GetRcuUrl() *string {
	return s.RcuUrl
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) GetTemplateName() *string {
	return s.TemplateName
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) SetAppId(v int64) *QueryHotelRoomDetailResponseBodyResultRoomControlInfo {
	s.AppId = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) SetAppName(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfo {
	s.AppName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) SetDeviceInfos(v []*QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) *QueryHotelRoomDetailResponseBodyResultRoomControlInfo {
	s.DeviceInfos = v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) SetRcuUrl(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfo {
	s.RcuUrl = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) SetTemplateId(v int64) *QueryHotelRoomDetailResponseBodyResultRoomControlInfo {
	s.TemplateId = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) SetTemplateName(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfo {
	s.TemplateName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfo) Validate() error {
	if s.DeviceInfos != nil {
		for _, item := range s.DeviceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos struct {
	// example:
	//
	// light
	CategoryEnName *string `json:"CategoryEnName,omitempty" xml:"CategoryEnName,omitempty"`
	// example:
	//
	// 3
	CategoryId   *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// rcu
	DeviceConnectType *string `json:"DeviceConnectType,omitempty" xml:"DeviceConnectType,omitempty"`
	// example:
	//
	// 4
	DeviceCount *int32 `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	// example:
	//
	// readLight
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// example:
	//
	// room
	LocationEnName *string `json:"LocationEnName,omitempty" xml:"LocationEnName,omitempty"`
	// example:
	//
	// 1
	LocationId   *int64  `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	// example:
	//
	// a1ueWGP6W2L
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) String() string {
	return dara.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) GetCategoryEnName() *string {
	return s.CategoryEnName
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) GetCategoryName() *string {
	return s.CategoryName
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) GetDeviceConnectType() *string {
	return s.DeviceConnectType
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) GetDeviceCount() *int32 {
	return s.DeviceCount
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) GetDeviceId() *string {
	return s.DeviceId
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) GetLocationEnName() *string {
	return s.LocationEnName
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) GetLocationId() *int64 {
	return s.LocationId
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) GetLocationName() *string {
	return s.LocationName
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) GetProductKey() *string {
	return s.ProductKey
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetCategoryEnName(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.CategoryEnName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetCategoryId(v int64) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.CategoryId = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetCategoryName(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.CategoryName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetDeviceConnectType(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.DeviceConnectType = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetDeviceCount(v int32) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.DeviceCount = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetDeviceId(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.DeviceId = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetDeviceName(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.DeviceName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetLocationEnName(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.LocationEnName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetLocationId(v int64) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.LocationId = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetLocationName(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.LocationName = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) SetProductKey(v string) *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos {
	s.ProductKey = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomControlInfoDeviceInfos) Validate() error {
	return dara.Validate(s)
}

type QueryHotelRoomDetailResponseBodyResultRoomServiceInfo struct {
	// example:
	//
	// 0
	BookServiceCnt *int32 `json:"BookServiceCnt,omitempty" xml:"BookServiceCnt,omitempty"`
	// example:
	//
	// 10
	GoodsServiceCnt *int32 `json:"GoodsServiceCnt,omitempty" xml:"GoodsServiceCnt,omitempty"`
	// example:
	//
	// 10
	RepairServiceCnt *int32 `json:"RepairServiceCnt,omitempty" xml:"RepairServiceCnt,omitempty"`
	// example:
	//
	// 12
	RoomServiceCnt *int32 `json:"RoomServiceCnt,omitempty" xml:"RoomServiceCnt,omitempty"`
}

func (s QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) GetBookServiceCnt() *int32 {
	return s.BookServiceCnt
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) GetGoodsServiceCnt() *int32 {
	return s.GoodsServiceCnt
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) GetRepairServiceCnt() *int32 {
	return s.RepairServiceCnt
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) GetRoomServiceCnt() *int32 {
	return s.RoomServiceCnt
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) SetBookServiceCnt(v int32) *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo {
	s.BookServiceCnt = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) SetGoodsServiceCnt(v int32) *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo {
	s.GoodsServiceCnt = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) SetRepairServiceCnt(v int32) *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo {
	s.RepairServiceCnt = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) SetRoomServiceCnt(v int32) *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo {
	s.RoomServiceCnt = &v
	return s
}

func (s *QueryHotelRoomDetailResponseBodyResultRoomServiceInfo) Validate() error {
	return dara.Validate(s)
}
