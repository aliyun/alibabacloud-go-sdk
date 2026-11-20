// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListXTelephonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListXTelephonesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListXTelephonesResponseBody
	GetCode() *string
	SetData(v *ListXTelephonesResponseBodyData) *ListXTelephonesResponseBody
	GetData() *ListXTelephonesResponseBodyData
	SetMessage(v string) *ListXTelephonesResponseBody
	GetMessage() *string
	SetSuccess(v bool) *ListXTelephonesResponseBody
	GetSuccess() *bool
}

type ListXTelephonesResponseBody struct {
	// The details of the access denied error.
	//
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code. Valid values:
	//
	// - 0000: Success.
	//
	// - Other values: Failure.
	//
	// example:
	//
	// 0000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data collection.
	Data *ListXTelephonesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: Success.
	//
	// - false: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListXTelephonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListXTelephonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListXTelephonesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListXTelephonesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListXTelephonesResponseBody) GetData() *ListXTelephonesResponseBodyData {
	return s.Data
}

func (s *ListXTelephonesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListXTelephonesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListXTelephonesResponseBody) SetAccessDeniedDetail(v string) *ListXTelephonesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListXTelephonesResponseBody) SetCode(v string) *ListXTelephonesResponseBody {
	s.Code = &v
	return s
}

func (s *ListXTelephonesResponseBody) SetData(v *ListXTelephonesResponseBodyData) *ListXTelephonesResponseBody {
	s.Data = v
	return s
}

func (s *ListXTelephonesResponseBody) SetMessage(v string) *ListXTelephonesResponseBody {
	s.Message = &v
	return s
}

func (s *ListXTelephonesResponseBody) SetSuccess(v bool) *ListXTelephonesResponseBody {
	s.Success = &v
	return s
}

func (s *ListXTelephonesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListXTelephonesResponseBodyData struct {
	// The data list.
	List []*ListXTelephonesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries that meet the query conditions.
	//
	// example:
	//
	// 50
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListXTelephonesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListXTelephonesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListXTelephonesResponseBodyData) GetList() []*ListXTelephonesResponseBodyDataList {
	return s.List
}

func (s *ListXTelephonesResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListXTelephonesResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListXTelephonesResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListXTelephonesResponseBodyData) SetList(v []*ListXTelephonesResponseBodyDataList) *ListXTelephonesResponseBodyData {
	s.List = v
	return s
}

func (s *ListXTelephonesResponseBodyData) SetPageNo(v int64) *ListXTelephonesResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListXTelephonesResponseBodyData) SetPageSize(v int64) *ListXTelephonesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListXTelephonesResponseBodyData) SetTotal(v int64) *ListXTelephonesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListXTelephonesResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListXTelephonesResponseBodyDataList struct {
	// The reason for binding failure.
	//
	// example:
	//
	// Binding failed: user ID card blacklisted
	AuthMsg *string `json:"AuthMsg,omitempty" xml:"AuthMsg,omitempty"`
	// The binding time. The format is YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2024-08-29 17:23:58
	BindTime *string `json:"BindTime,omitempty" xml:"BindTime,omitempty"`
	// The purchase time. The format is YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2024-08-29 17:23:58
	BuyTime *string `json:"BuyTime,omitempty" xml:"BuyTime,omitempty"`
	// The customer number pool key.
	//
	// example:
	//
	// FC533e6eeb81f4400c87ef3745a21a1a
	CustomerPoolKey *string `json:"CustomerPoolKey,omitempty" xml:"CustomerPoolKey,omitempty"`
	// The number pool name.
	//
	// example:
	//
	// Test number pool
	CustomerPoolName *string `json:"CustomerPoolName,omitempty" xml:"CustomerPoolName,omitempty"`
	// The release time. The format is YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2024-08-29 17:23:58
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The SMS activation status. Valid values:
	//
	// - 0: Not activated.
	//
	// - 1: Activated.
	//
	// example:
	//
	// 0
	SmsStatus *string `json:"SmsStatus,omitempty" xml:"SmsStatus,omitempty"`
	// The X number.
	//
	// example:
	//
	// 17816876546
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	// The number status. Valid values:
	//
	// - 0: Idle.
	//
	// - 1: Allocation completed and pending purchase.
	//
	// - 2: Purchase completed and pending verification.
	//
	// - 3: Real-name verification in progress.
	//
	// - 4: Real-name verification succeeded.
	//
	// - 5: Verification failed.
	//
	// - 6: Unbinding in progress.
	//
	// - 7: Unbinding failed.
	//
	// - 8: Released.
	//
	// - 99: Released due to timeout.
	//
	// example:
	//
	// 0
	TelephoneStatus *string `json:"TelephoneStatus,omitempty" xml:"TelephoneStatus,omitempty"`
	// The unbinding time. The format is YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2024-08-29 17:23:58
	UnbindTime *string `json:"UnbindTime,omitempty" xml:"UnbindTime,omitempty"`
}

func (s ListXTelephonesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListXTelephonesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListXTelephonesResponseBodyDataList) GetAuthMsg() *string {
	return s.AuthMsg
}

func (s *ListXTelephonesResponseBodyDataList) GetBindTime() *string {
	return s.BindTime
}

func (s *ListXTelephonesResponseBodyDataList) GetBuyTime() *string {
	return s.BuyTime
}

func (s *ListXTelephonesResponseBodyDataList) GetCustomerPoolKey() *string {
	return s.CustomerPoolKey
}

func (s *ListXTelephonesResponseBodyDataList) GetCustomerPoolName() *string {
	return s.CustomerPoolName
}

func (s *ListXTelephonesResponseBodyDataList) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *ListXTelephonesResponseBodyDataList) GetSmsStatus() *string {
	return s.SmsStatus
}

func (s *ListXTelephonesResponseBodyDataList) GetTelephone() *string {
	return s.Telephone
}

func (s *ListXTelephonesResponseBodyDataList) GetTelephoneStatus() *string {
	return s.TelephoneStatus
}

func (s *ListXTelephonesResponseBodyDataList) GetUnbindTime() *string {
	return s.UnbindTime
}

func (s *ListXTelephonesResponseBodyDataList) SetAuthMsg(v string) *ListXTelephonesResponseBodyDataList {
	s.AuthMsg = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetBindTime(v string) *ListXTelephonesResponseBodyDataList {
	s.BindTime = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetBuyTime(v string) *ListXTelephonesResponseBodyDataList {
	s.BuyTime = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetCustomerPoolKey(v string) *ListXTelephonesResponseBodyDataList {
	s.CustomerPoolKey = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetCustomerPoolName(v string) *ListXTelephonesResponseBodyDataList {
	s.CustomerPoolName = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetReleaseTime(v string) *ListXTelephonesResponseBodyDataList {
	s.ReleaseTime = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetSmsStatus(v string) *ListXTelephonesResponseBodyDataList {
	s.SmsStatus = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetTelephone(v string) *ListXTelephonesResponseBodyDataList {
	s.Telephone = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetTelephoneStatus(v string) *ListXTelephonesResponseBodyDataList {
	s.TelephoneStatus = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) SetUnbindTime(v string) *ListXTelephonesResponseBodyDataList {
	s.UnbindTime = &v
	return s
}

func (s *ListXTelephonesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
