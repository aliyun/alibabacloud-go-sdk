// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *QueryDeviceInfoResponseBodyDeviceInfo) *QueryDeviceInfoResponseBody
	GetDeviceInfo() *QueryDeviceInfoResponseBodyDeviceInfo
	SetRequestId(v string) *QueryDeviceInfoResponseBody
	GetRequestId() *string
}

type QueryDeviceInfoResponseBody struct {
	DeviceInfo *QueryDeviceInfoResponseBodyDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 6EEF262B-EA7D-41DC-89B9-20F3D1E28194
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceInfoResponseBody) GetDeviceInfo() *QueryDeviceInfoResponseBodyDeviceInfo {
	return s.DeviceInfo
}

func (s *QueryDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDeviceInfoResponseBody) SetDeviceInfo(v *QueryDeviceInfoResponseBodyDeviceInfo) *QueryDeviceInfoResponseBody {
	s.DeviceInfo = v
	return s
}

func (s *QueryDeviceInfoResponseBody) SetRequestId(v string) *QueryDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryDeviceInfoResponseBodyDeviceInfo struct {
	// example:
	//
	// test@aliyun.com
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// example:
	//
	// test_alias,test_alias2
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// example:
	//
	// a64ae296f3b04a58a05b30c95****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 5ecc7b4012aaa801b63******5543ccbda6b4930d09629e936e1ac4b762a7df
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	// example:
	//
	// iOS
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 2018-03-27T02:19:40Z
	LastOnlineTime *string `json:"LastOnlineTime,omitempty" xml:"LastOnlineTime,omitempty"`
	Model          *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// false
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// example:
	//
	// 133********
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// true
	PushEnabled *bool `json:"PushEnabled,omitempty" xml:"PushEnabled,omitempty"`
	// example:
	//
	// test_tag,test_tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s QueryDeviceInfoResponseBodyDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceInfoResponseBodyDeviceInfo) GoString() string {
	return s.String()
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) GetAccount() *string {
	return s.Account
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) GetAlias() *string {
	return s.Alias
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) GetBrand() *string {
	return s.Brand
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) GetDeviceId() *string {
	return s.DeviceId
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) GetDeviceToken() *string {
	return s.DeviceToken
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) GetDeviceType() *string {
	return s.DeviceType
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) GetLastOnlineTime() *string {
	return s.LastOnlineTime
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) GetModel() *string {
	return s.Model
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) GetOnline() *bool {
	return s.Online
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) GetPushEnabled() *bool {
	return s.PushEnabled
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) GetTags() *string {
	return s.Tags
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetAccount(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.Account = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetAlias(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.Alias = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetBrand(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.Brand = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetDeviceId(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.DeviceId = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetDeviceToken(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.DeviceToken = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetDeviceType(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.DeviceType = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetLastOnlineTime(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.LastOnlineTime = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetModel(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.Model = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetOnline(v bool) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.Online = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetPhoneNumber(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.PhoneNumber = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetPushEnabled(v bool) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.PushEnabled = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) SetTags(v string) *QueryDeviceInfoResponseBodyDeviceInfo {
	s.Tags = &v
	return s
}

func (s *QueryDeviceInfoResponseBodyDeviceInfo) Validate() error {
	return dara.Validate(s)
}
