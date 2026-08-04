// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterInternalAccountForBucResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RegisterInternalAccountForBucResponseBody
	GetCode() *string
	SetData(v *RegisterInternalAccountForBucResponseBodyData) *RegisterInternalAccountForBucResponseBody
	GetData() *RegisterInternalAccountForBucResponseBodyData
	SetLocalizedMessage(v string) *RegisterInternalAccountForBucResponseBody
	GetLocalizedMessage() *string
	SetMessage(v string) *RegisterInternalAccountForBucResponseBody
	GetMessage() *string
	SetMsg(v string) *RegisterInternalAccountForBucResponseBody
	GetMsg() *string
	SetRequestId(v string) *RegisterInternalAccountForBucResponseBody
	GetRequestId() *string
}

type RegisterInternalAccountForBucResponseBody struct {
	Code             *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *RegisterInternalAccountForBucResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	LocalizedMessage *string                                        `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	Message          *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Msg              *string                                        `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId        *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterInternalAccountForBucResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterInternalAccountForBucResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterInternalAccountForBucResponseBody) GetCode() *string {
	return s.Code
}

func (s *RegisterInternalAccountForBucResponseBody) GetData() *RegisterInternalAccountForBucResponseBodyData {
	return s.Data
}

func (s *RegisterInternalAccountForBucResponseBody) GetLocalizedMessage() *string {
	return s.LocalizedMessage
}

func (s *RegisterInternalAccountForBucResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RegisterInternalAccountForBucResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *RegisterInternalAccountForBucResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterInternalAccountForBucResponseBody) SetCode(v string) *RegisterInternalAccountForBucResponseBody {
	s.Code = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBody) SetData(v *RegisterInternalAccountForBucResponseBodyData) *RegisterInternalAccountForBucResponseBody {
	s.Data = v
	return s
}

func (s *RegisterInternalAccountForBucResponseBody) SetLocalizedMessage(v string) *RegisterInternalAccountForBucResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBody) SetMessage(v string) *RegisterInternalAccountForBucResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBody) SetMsg(v string) *RegisterInternalAccountForBucResponseBody {
	s.Msg = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBody) SetRequestId(v string) *RegisterInternalAccountForBucResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RegisterInternalAccountForBucResponseBodyData struct {
	AccountStatus    *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	AccountStructure *string `json:"AccountStructure,omitempty" xml:"AccountStructure,omitempty"`
	ExtendInfo       *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	HavanaId         *string `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	LastLoginTime    *string `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty"`
	OwnerBid         *string `json:"OwnerBid,omitempty" xml:"OwnerBid,omitempty"`
	ParentPk         *string `json:"ParentPk,omitempty" xml:"ParentPk,omitempty"`
	PartnerPk        *string `json:"PartnerPk,omitempty" xml:"PartnerPk,omitempty"`
	Pk               *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Site             *string `json:"Site,omitempty" xml:"Site,omitempty"`
}

func (s RegisterInternalAccountForBucResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RegisterInternalAccountForBucResponseBodyData) GoString() string {
	return s.String()
}

func (s *RegisterInternalAccountForBucResponseBodyData) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *RegisterInternalAccountForBucResponseBodyData) GetAccountStructure() *string {
	return s.AccountStructure
}

func (s *RegisterInternalAccountForBucResponseBodyData) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *RegisterInternalAccountForBucResponseBodyData) GetHavanaId() *string {
	return s.HavanaId
}

func (s *RegisterInternalAccountForBucResponseBodyData) GetLastLoginTime() *string {
	return s.LastLoginTime
}

func (s *RegisterInternalAccountForBucResponseBodyData) GetOwnerBid() *string {
	return s.OwnerBid
}

func (s *RegisterInternalAccountForBucResponseBodyData) GetParentPk() *string {
	return s.ParentPk
}

func (s *RegisterInternalAccountForBucResponseBodyData) GetPartnerPk() *string {
	return s.PartnerPk
}

func (s *RegisterInternalAccountForBucResponseBodyData) GetPk() *string {
	return s.Pk
}

func (s *RegisterInternalAccountForBucResponseBodyData) GetSite() *string {
	return s.Site
}

func (s *RegisterInternalAccountForBucResponseBodyData) SetAccountStatus(v string) *RegisterInternalAccountForBucResponseBodyData {
	s.AccountStatus = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBodyData) SetAccountStructure(v string) *RegisterInternalAccountForBucResponseBodyData {
	s.AccountStructure = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBodyData) SetExtendInfo(v string) *RegisterInternalAccountForBucResponseBodyData {
	s.ExtendInfo = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBodyData) SetHavanaId(v string) *RegisterInternalAccountForBucResponseBodyData {
	s.HavanaId = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBodyData) SetLastLoginTime(v string) *RegisterInternalAccountForBucResponseBodyData {
	s.LastLoginTime = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBodyData) SetOwnerBid(v string) *RegisterInternalAccountForBucResponseBodyData {
	s.OwnerBid = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBodyData) SetParentPk(v string) *RegisterInternalAccountForBucResponseBodyData {
	s.ParentPk = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBodyData) SetPartnerPk(v string) *RegisterInternalAccountForBucResponseBodyData {
	s.PartnerPk = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBodyData) SetPk(v string) *RegisterInternalAccountForBucResponseBodyData {
	s.Pk = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBodyData) SetSite(v string) *RegisterInternalAccountForBucResponseBodyData {
	s.Site = &v
	return s
}

func (s *RegisterInternalAccountForBucResponseBodyData) Validate() error {
	return dara.Validate(s)
}
