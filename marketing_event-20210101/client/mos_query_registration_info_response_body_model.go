// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMosQueryRegistrationInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *MosQueryRegistrationInfoResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *MosQueryRegistrationInfoResponseBody
	GetCode() *string
	SetData(v *MosQueryRegistrationInfoResponseBodyData) *MosQueryRegistrationInfoResponseBody
	GetData() *MosQueryRegistrationInfoResponseBodyData
	SetMessage(v string) *MosQueryRegistrationInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *MosQueryRegistrationInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MosQueryRegistrationInfoResponseBody
	GetSuccess() *bool
}

type MosQueryRegistrationInfoResponseBody struct {
	AccessDeniedDetail *string                                   `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *MosQueryRegistrationInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MosQueryRegistrationInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MosQueryRegistrationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *MosQueryRegistrationInfoResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *MosQueryRegistrationInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *MosQueryRegistrationInfoResponseBody) GetData() *MosQueryRegistrationInfoResponseBodyData {
	return s.Data
}

func (s *MosQueryRegistrationInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MosQueryRegistrationInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MosQueryRegistrationInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MosQueryRegistrationInfoResponseBody) SetAccessDeniedDetail(v string) *MosQueryRegistrationInfoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBody) SetCode(v string) *MosQueryRegistrationInfoResponseBody {
	s.Code = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBody) SetData(v *MosQueryRegistrationInfoResponseBodyData) *MosQueryRegistrationInfoResponseBody {
	s.Data = v
	return s
}

func (s *MosQueryRegistrationInfoResponseBody) SetMessage(v string) *MosQueryRegistrationInfoResponseBody {
	s.Message = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBody) SetRequestId(v string) *MosQueryRegistrationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBody) SetSuccess(v bool) *MosQueryRegistrationInfoResponseBody {
	s.Success = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MosQueryRegistrationInfoResponseBodyData struct {
	List  []*MosQueryRegistrationInfoResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Total *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s MosQueryRegistrationInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MosQueryRegistrationInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *MosQueryRegistrationInfoResponseBodyData) GetList() []*MosQueryRegistrationInfoResponseBodyDataList {
	return s.List
}

func (s *MosQueryRegistrationInfoResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *MosQueryRegistrationInfoResponseBodyData) SetList(v []*MosQueryRegistrationInfoResponseBodyDataList) *MosQueryRegistrationInfoResponseBodyData {
	s.List = v
	return s
}

func (s *MosQueryRegistrationInfoResponseBodyData) SetTotal(v int64) *MosQueryRegistrationInfoResponseBodyData {
	s.Total = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBodyData) Validate() error {
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

type MosQueryRegistrationInfoResponseBodyDataList struct {
	ActivityId    *string                                                 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	ChannelName   *string                                                 `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	CheckinStatus *string                                                 `json:"CheckinStatus,omitempty" xml:"CheckinStatus,omitempty"`
	Id            *int64                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	QrCode        *string                                                 `json:"QrCode,omitempty" xml:"QrCode,omitempty"`
	TagInfos      []*MosQueryRegistrationInfoResponseBodyDataListTagInfos `json:"TagInfos,omitempty" xml:"TagInfos,omitempty" type:"Repeated"`
	UserCompany   *string                                                 `json:"UserCompany,omitempty" xml:"UserCompany,omitempty"`
	UserEmail     *string                                                 `json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
	UserName      *string                                                 `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPhone     *string                                                 `json:"UserPhone,omitempty" xml:"UserPhone,omitempty"`
}

func (s MosQueryRegistrationInfoResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s MosQueryRegistrationInfoResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) GetActivityId() *string {
	return s.ActivityId
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) GetChannelName() *string {
	return s.ChannelName
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) GetCheckinStatus() *string {
	return s.CheckinStatus
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) GetQrCode() *string {
	return s.QrCode
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) GetTagInfos() []*MosQueryRegistrationInfoResponseBodyDataListTagInfos {
	return s.TagInfos
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) GetUserCompany() *string {
	return s.UserCompany
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) GetUserEmail() *string {
	return s.UserEmail
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) GetUserName() *string {
	return s.UserName
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) GetUserPhone() *string {
	return s.UserPhone
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) SetActivityId(v string) *MosQueryRegistrationInfoResponseBodyDataList {
	s.ActivityId = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) SetChannelName(v string) *MosQueryRegistrationInfoResponseBodyDataList {
	s.ChannelName = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) SetCheckinStatus(v string) *MosQueryRegistrationInfoResponseBodyDataList {
	s.CheckinStatus = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) SetId(v int64) *MosQueryRegistrationInfoResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) SetQrCode(v string) *MosQueryRegistrationInfoResponseBodyDataList {
	s.QrCode = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) SetTagInfos(v []*MosQueryRegistrationInfoResponseBodyDataListTagInfos) *MosQueryRegistrationInfoResponseBodyDataList {
	s.TagInfos = v
	return s
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) SetUserCompany(v string) *MosQueryRegistrationInfoResponseBodyDataList {
	s.UserCompany = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) SetUserEmail(v string) *MosQueryRegistrationInfoResponseBodyDataList {
	s.UserEmail = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) SetUserName(v string) *MosQueryRegistrationInfoResponseBodyDataList {
	s.UserName = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) SetUserPhone(v string) *MosQueryRegistrationInfoResponseBodyDataList {
	s.UserPhone = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBodyDataList) Validate() error {
	if s.TagInfos != nil {
		for _, item := range s.TagInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MosQueryRegistrationInfoResponseBodyDataListTagInfos struct {
	TagCode *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s MosQueryRegistrationInfoResponseBodyDataListTagInfos) String() string {
	return dara.Prettify(s)
}

func (s MosQueryRegistrationInfoResponseBodyDataListTagInfos) GoString() string {
	return s.String()
}

func (s *MosQueryRegistrationInfoResponseBodyDataListTagInfos) GetTagCode() *string {
	return s.TagCode
}

func (s *MosQueryRegistrationInfoResponseBodyDataListTagInfos) GetTagName() *string {
	return s.TagName
}

func (s *MosQueryRegistrationInfoResponseBodyDataListTagInfos) SetTagCode(v string) *MosQueryRegistrationInfoResponseBodyDataListTagInfos {
	s.TagCode = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBodyDataListTagInfos) SetTagName(v string) *MosQueryRegistrationInfoResponseBodyDataListTagInfos {
	s.TagName = &v
	return s
}

func (s *MosQueryRegistrationInfoResponseBodyDataListTagInfos) Validate() error {
	return dara.Validate(s)
}
