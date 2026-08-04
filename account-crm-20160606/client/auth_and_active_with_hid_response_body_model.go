// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthAndActiveWithHidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AuthAndActiveWithHidResponseBody
	GetCode() *string
	SetData(v *AuthAndActiveWithHidResponseBodyData) *AuthAndActiveWithHidResponseBody
	GetData() *AuthAndActiveWithHidResponseBodyData
	SetMsg(v string) *AuthAndActiveWithHidResponseBody
	GetMsg() *string
	SetRequestId(v string) *AuthAndActiveWithHidResponseBody
	GetRequestId() *string
}

type AuthAndActiveWithHidResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AuthAndActiveWithHidResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                               `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthAndActiveWithHidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthAndActiveWithHidResponseBody) GoString() string {
	return s.String()
}

func (s *AuthAndActiveWithHidResponseBody) GetCode() *string {
	return s.Code
}

func (s *AuthAndActiveWithHidResponseBody) GetData() *AuthAndActiveWithHidResponseBodyData {
	return s.Data
}

func (s *AuthAndActiveWithHidResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *AuthAndActiveWithHidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthAndActiveWithHidResponseBody) SetCode(v string) *AuthAndActiveWithHidResponseBody {
	s.Code = &v
	return s
}

func (s *AuthAndActiveWithHidResponseBody) SetData(v *AuthAndActiveWithHidResponseBodyData) *AuthAndActiveWithHidResponseBody {
	s.Data = v
	return s
}

func (s *AuthAndActiveWithHidResponseBody) SetMsg(v string) *AuthAndActiveWithHidResponseBody {
	s.Msg = &v
	return s
}

func (s *AuthAndActiveWithHidResponseBody) SetRequestId(v string) *AuthAndActiveWithHidResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthAndActiveWithHidResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AuthAndActiveWithHidResponseBodyData struct {
	AccountModel *AuthAndActiveWithHidResponseBodyDataAccountModel `json:"AccountModel,omitempty" xml:"AccountModel,omitempty" type:"Struct"`
	SessionModel *AuthAndActiveWithHidResponseBodyDataSessionModel `json:"SessionModel,omitempty" xml:"SessionModel,omitempty" type:"Struct"`
}

func (s AuthAndActiveWithHidResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AuthAndActiveWithHidResponseBodyData) GoString() string {
	return s.String()
}

func (s *AuthAndActiveWithHidResponseBodyData) GetAccountModel() *AuthAndActiveWithHidResponseBodyDataAccountModel {
	return s.AccountModel
}

func (s *AuthAndActiveWithHidResponseBodyData) GetSessionModel() *AuthAndActiveWithHidResponseBodyDataSessionModel {
	return s.SessionModel
}

func (s *AuthAndActiveWithHidResponseBodyData) SetAccountModel(v *AuthAndActiveWithHidResponseBodyDataAccountModel) *AuthAndActiveWithHidResponseBodyData {
	s.AccountModel = v
	return s
}

func (s *AuthAndActiveWithHidResponseBodyData) SetSessionModel(v *AuthAndActiveWithHidResponseBodyDataSessionModel) *AuthAndActiveWithHidResponseBodyData {
	s.SessionModel = v
	return s
}

func (s *AuthAndActiveWithHidResponseBodyData) Validate() error {
	if s.AccountModel != nil {
		if err := s.AccountModel.Validate(); err != nil {
			return err
		}
	}
	if s.SessionModel != nil {
		if err := s.SessionModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AuthAndActiveWithHidResponseBodyDataAccountModel struct {
	AliyunId   *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Email      *string `json:"Email,omitempty" xml:"Email,omitempty"`
	HavanaId   *int64  `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	Mobile     *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	PK         *string `json:"PK,omitempty" xml:"PK,omitempty"`
}

func (s AuthAndActiveWithHidResponseBodyDataAccountModel) String() string {
	return dara.Prettify(s)
}

func (s AuthAndActiveWithHidResponseBodyDataAccountModel) GoString() string {
	return s.String()
}

func (s *AuthAndActiveWithHidResponseBodyDataAccountModel) GetAliyunId() *string {
	return s.AliyunId
}

func (s *AuthAndActiveWithHidResponseBodyDataAccountModel) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *AuthAndActiveWithHidResponseBodyDataAccountModel) GetEmail() *string {
	return s.Email
}

func (s *AuthAndActiveWithHidResponseBodyDataAccountModel) GetHavanaId() *int64 {
	return s.HavanaId
}

func (s *AuthAndActiveWithHidResponseBodyDataAccountModel) GetMobile() *string {
	return s.Mobile
}

func (s *AuthAndActiveWithHidResponseBodyDataAccountModel) GetPK() *string {
	return s.PK
}

func (s *AuthAndActiveWithHidResponseBodyDataAccountModel) SetAliyunId(v string) *AuthAndActiveWithHidResponseBodyDataAccountModel {
	s.AliyunId = &v
	return s
}

func (s *AuthAndActiveWithHidResponseBodyDataAccountModel) SetCreateTime(v int64) *AuthAndActiveWithHidResponseBodyDataAccountModel {
	s.CreateTime = &v
	return s
}

func (s *AuthAndActiveWithHidResponseBodyDataAccountModel) SetEmail(v string) *AuthAndActiveWithHidResponseBodyDataAccountModel {
	s.Email = &v
	return s
}

func (s *AuthAndActiveWithHidResponseBodyDataAccountModel) SetHavanaId(v int64) *AuthAndActiveWithHidResponseBodyDataAccountModel {
	s.HavanaId = &v
	return s
}

func (s *AuthAndActiveWithHidResponseBodyDataAccountModel) SetMobile(v string) *AuthAndActiveWithHidResponseBodyDataAccountModel {
	s.Mobile = &v
	return s
}

func (s *AuthAndActiveWithHidResponseBodyDataAccountModel) SetPK(v string) *AuthAndActiveWithHidResponseBodyDataAccountModel {
	s.PK = &v
	return s
}

func (s *AuthAndActiveWithHidResponseBodyDataAccountModel) Validate() error {
	return dara.Validate(s)
}

type AuthAndActiveWithHidResponseBodyDataSessionModel struct {
	AliyunPK    *string `json:"AliyunPK,omitempty" xml:"AliyunPK,omitempty"`
	LoginTicket *string `json:"LoginTicket,omitempty" xml:"LoginTicket,omitempty"`
}

func (s AuthAndActiveWithHidResponseBodyDataSessionModel) String() string {
	return dara.Prettify(s)
}

func (s AuthAndActiveWithHidResponseBodyDataSessionModel) GoString() string {
	return s.String()
}

func (s *AuthAndActiveWithHidResponseBodyDataSessionModel) GetAliyunPK() *string {
	return s.AliyunPK
}

func (s *AuthAndActiveWithHidResponseBodyDataSessionModel) GetLoginTicket() *string {
	return s.LoginTicket
}

func (s *AuthAndActiveWithHidResponseBodyDataSessionModel) SetAliyunPK(v string) *AuthAndActiveWithHidResponseBodyDataSessionModel {
	s.AliyunPK = &v
	return s
}

func (s *AuthAndActiveWithHidResponseBodyDataSessionModel) SetLoginTicket(v string) *AuthAndActiveWithHidResponseBodyDataSessionModel {
	s.LoginTicket = &v
	return s
}

func (s *AuthAndActiveWithHidResponseBodyDataSessionModel) Validate() error {
	return dara.Validate(s)
}
