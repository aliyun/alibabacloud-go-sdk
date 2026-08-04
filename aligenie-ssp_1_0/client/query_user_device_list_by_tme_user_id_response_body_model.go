// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserDeviceListByTmeUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryUserDeviceListByTmeUserIdResponseBody
	GetCode() *int32
	SetMessage(v string) *QueryUserDeviceListByTmeUserIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryUserDeviceListByTmeUserIdResponseBody
	GetRequestId() *string
	SetResult(v *QueryUserDeviceListByTmeUserIdResponseBodyResult) *QueryUserDeviceListByTmeUserIdResponseBody
	GetResult() *QueryUserDeviceListByTmeUserIdResponseBodyResult
	SetSuccess(v bool) *QueryUserDeviceListByTmeUserIdResponseBody
	GetSuccess() *bool
}

type QueryUserDeviceListByTmeUserIdResponseBody struct {
	// Response code
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 860194F7-9593-50EA-8E53-BCEC0D325A00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Response Result
	Result *QueryUserDeviceListByTmeUserIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Flag indicating whether the invocation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserDeviceListByTmeUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserDeviceListByTmeUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserDeviceListByTmeUserIdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryUserDeviceListByTmeUserIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryUserDeviceListByTmeUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserDeviceListByTmeUserIdResponseBody) GetResult() *QueryUserDeviceListByTmeUserIdResponseBodyResult {
	return s.Result
}

func (s *QueryUserDeviceListByTmeUserIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUserDeviceListByTmeUserIdResponseBody) SetCode(v int32) *QueryUserDeviceListByTmeUserIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBody) SetMessage(v string) *QueryUserDeviceListByTmeUserIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBody) SetRequestId(v string) *QueryUserDeviceListByTmeUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBody) SetResult(v *QueryUserDeviceListByTmeUserIdResponseBodyResult) *QueryUserDeviceListByTmeUserIdResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBody) SetSuccess(v bool) *QueryUserDeviceListByTmeUserIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryUserDeviceListByTmeUserIdResponseBodyResult struct {
	// Tmall Genie User List
	AligenieUserInfoList []*QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList `json:"AligenieUserInfoList,omitempty" xml:"AligenieUserInfoList,omitempty" type:"Repeated"`
	// entity key (pass-through by third party)
	//
	// example:
	//
	// 12****7
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// entity Type (pass-through by third party)
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// "KG": KuGou
	//
	// "KW": Kuwo
	//
	// "QM": QQ Music
	//
	// example:
	//
	// KG
	Sp *string `json:"Sp,omitempty" xml:"Sp,omitempty"`
}

func (s QueryUserDeviceListByTmeUserIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryUserDeviceListByTmeUserIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResult) GetAligenieUserInfoList() []*QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList {
	return s.AligenieUserInfoList
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResult) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResult) GetEncodeType() *string {
	return s.EncodeType
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResult) GetSp() *string {
	return s.Sp
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResult) SetAligenieUserInfoList(v []*QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList) *QueryUserDeviceListByTmeUserIdResponseBodyResult {
	s.AligenieUserInfoList = v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResult) SetEncodeKey(v string) *QueryUserDeviceListByTmeUserIdResponseBodyResult {
	s.EncodeKey = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResult) SetEncodeType(v string) *QueryUserDeviceListByTmeUserIdResponseBodyResult {
	s.EncodeType = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResult) SetSp(v string) *QueryUserDeviceListByTmeUserIdResponseBodyResult {
	s.Sp = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResult) Validate() error {
	if s.AligenieUserInfoList != nil {
		for _, item := range s.AligenieUserInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList struct {
	// User Authorization device List
	AuthorizedDeviceList []*QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList `json:"AuthorizedDeviceList,omitempty" xml:"AuthorizedDeviceList,omitempty" type:"Repeated"`
	// User ID
	//
	// example:
	//
	// R457Av3qg/OXTwVnFt12z6MwNe0HAS699V6n63OaLdu+VmwvhcNfMzBd+la553wWJhj3kBMjgHq2Y2dyCFoDBg==
	OpenUserId *string `json:"OpenUserId,omitempty" xml:"OpenUserId,omitempty"`
	// User nickname
	//
	// example:
	//
	// a***e
	UserNickname *string `json:"UserNickname,omitempty" xml:"UserNickname,omitempty"`
}

func (s QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList) String() string {
	return dara.Prettify(s)
}

func (s QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList) GoString() string {
	return s.String()
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList) GetAuthorizedDeviceList() []*QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList {
	return s.AuthorizedDeviceList
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList) GetOpenUserId() *string {
	return s.OpenUserId
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList) GetUserNickname() *string {
	return s.UserNickname
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList) SetAuthorizedDeviceList(v []*QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList) *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList {
	s.AuthorizedDeviceList = v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList) SetOpenUserId(v string) *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList {
	s.OpenUserId = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList) SetUserNickname(v string) *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList {
	s.UserNickname = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoList) Validate() error {
	if s.AuthorizedDeviceList != nil {
		for _, item := range s.AuthorizedDeviceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList struct {
	// device name
	//
	// example:
	//
	// CC
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// Indicates whether the device is online
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// Device ID
	//
	// example:
	//
	// fjwZiYQdtkaI95fHaLNjYcaOA/mxUPzxxw2J5iBiTBnjUCWKwER4TSHCqkBnNOYvGJ4bRZA9KzBB2naS4r/Am0lSe8ECDAAOcJ9QKLFF6DM=
	OpenDeviceId *string `json:"OpenDeviceId,omitempty" xml:"OpenDeviceId,omitempty"`
	// Device ID exposed to TME
	TmeDeviceId *string `json:"TmeDeviceId,omitempty" xml:"TmeDeviceId,omitempty"`
	// TME product ID
	TmeProductId *string `json:"TmeProductId,omitempty" xml:"TmeProductId,omitempty"`
}

func (s QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList) String() string {
	return dara.Prettify(s)
}

func (s QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList) GoString() string {
	return s.String()
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList) GetDeviceName() *string {
	return s.DeviceName
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList) GetOnline() *bool {
	return s.Online
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList) GetOpenDeviceId() *string {
	return s.OpenDeviceId
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList) GetTmeDeviceId() *string {
	return s.TmeDeviceId
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList) GetTmeProductId() *string {
	return s.TmeProductId
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList) SetDeviceName(v string) *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList {
	s.DeviceName = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList) SetOnline(v bool) *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList {
	s.Online = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList) SetOpenDeviceId(v string) *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList {
	s.OpenDeviceId = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList) SetTmeDeviceId(v string) *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList {
	s.TmeDeviceId = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList) SetTmeProductId(v string) *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList {
	s.TmeProductId = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponseBodyResultAligenieUserInfoListAuthorizedDeviceList) Validate() error {
	return dara.Validate(s)
}
