// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdpConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetIdpConfigResponseBodyData) *GetIdpConfigResponseBody
	GetData() *GetIdpConfigResponseBodyData
	SetRequestId(v string) *GetIdpConfigResponseBody
	GetRequestId() *string
}

type GetIdpConfigResponseBody struct {
	Data *GetIdpConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIdpConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdpConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdpConfigResponseBody) GetData() *GetIdpConfigResponseBodyData {
	return s.Data
}

func (s *GetIdpConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdpConfigResponseBody) SetData(v *GetIdpConfigResponseBodyData) *GetIdpConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetIdpConfigResponseBody) SetRequestId(v string) *GetIdpConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdpConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetIdpConfigResponseBodyData struct {
	// AccessKey ID
	//
	// example:
	//
	// LTAI5tJVztnh6Nn***
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// AccessKey Secret
	//
	// example:
	//
	// E75ktr5jENiR3ssjC***
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// https://172.10.10.2:4321/getGroup?name=%s&pass=%s
	GetGroupUrl *string `json:"GetGroupUrl,omitempty" xml:"GetGroupUrl,omitempty"`
	// example:
	//
	// 1465
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// <?xml version="1.0" encoding="utf-8"?>***
	IdpMetadata *string `json:"IdpMetadata,omitempty" xml:"IdpMetadata,omitempty"`
	// example:
	//
	// totp
	MfaConfigType *string `json:"MfaConfigType,omitempty" xml:"MfaConfigType,omitempty"`
	// example:
	//
	// password
	MobileLoginType *string `json:"MobileLoginType,omitempty" xml:"MobileLoginType,omitempty"`
	// example:
	//
	// totp
	MobileMfaConfigType *string `json:"MobileMfaConfigType,omitempty" xml:"MobileMfaConfigType,omitempty"`
	// example:
	//
	// 1482,1355
	MultiIdpInfo *string `json:"MultiIdpInfo,omitempty" xml:"MultiIdpInfo,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// password
	PcLoginType *string `json:"PcLoginType,omitempty" xml:"PcLoginType,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// CSAS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2024-02-26T02:02:42Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// QVhaU0RDR0JIWVV***
	VerifyAesKey *string `json:"VerifyAesKey,omitempty" xml:"VerifyAesKey,omitempty"`
	// example:
	//
	// 7JAr3fYtnl***
	VerifyToken *string `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
	// example:
	//
	// http://172.10.10.1:1234/otp_verify
	VerifyUrl *string `json:"VerifyUrl,omitempty" xml:"VerifyUrl,omitempty"`
}

func (s GetIdpConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIdpConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIdpConfigResponseBodyData) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetIdpConfigResponseBodyData) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *GetIdpConfigResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetIdpConfigResponseBodyData) GetGetGroupUrl() *string {
	return s.GetGroupUrl
}

func (s *GetIdpConfigResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetIdpConfigResponseBodyData) GetIdpMetadata() *string {
	return s.IdpMetadata
}

func (s *GetIdpConfigResponseBodyData) GetMfaConfigType() *string {
	return s.MfaConfigType
}

func (s *GetIdpConfigResponseBodyData) GetMobileLoginType() *string {
	return s.MobileLoginType
}

func (s *GetIdpConfigResponseBodyData) GetMobileMfaConfigType() *string {
	return s.MobileMfaConfigType
}

func (s *GetIdpConfigResponseBodyData) GetMultiIdpInfo() *string {
	return s.MultiIdpInfo
}

func (s *GetIdpConfigResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetIdpConfigResponseBodyData) GetPcLoginType() *string {
	return s.PcLoginType
}

func (s *GetIdpConfigResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetIdpConfigResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetIdpConfigResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetIdpConfigResponseBodyData) GetVerifyAesKey() *string {
	return s.VerifyAesKey
}

func (s *GetIdpConfigResponseBodyData) GetVerifyToken() *string {
	return s.VerifyToken
}

func (s *GetIdpConfigResponseBodyData) GetVerifyUrl() *string {
	return s.VerifyUrl
}

func (s *GetIdpConfigResponseBodyData) SetAccessKey(v string) *GetIdpConfigResponseBodyData {
	s.AccessKey = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetAccessKeySecret(v string) *GetIdpConfigResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetDescription(v string) *GetIdpConfigResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetGetGroupUrl(v string) *GetIdpConfigResponseBodyData {
	s.GetGroupUrl = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetId(v string) *GetIdpConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetIdpMetadata(v string) *GetIdpConfigResponseBodyData {
	s.IdpMetadata = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetMfaConfigType(v string) *GetIdpConfigResponseBodyData {
	s.MfaConfigType = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetMobileLoginType(v string) *GetIdpConfigResponseBodyData {
	s.MobileLoginType = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetMobileMfaConfigType(v string) *GetIdpConfigResponseBodyData {
	s.MobileMfaConfigType = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetMultiIdpInfo(v string) *GetIdpConfigResponseBodyData {
	s.MultiIdpInfo = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetName(v string) *GetIdpConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetPcLoginType(v string) *GetIdpConfigResponseBodyData {
	s.PcLoginType = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetStatus(v string) *GetIdpConfigResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetType(v string) *GetIdpConfigResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetUpdateTime(v string) *GetIdpConfigResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetVerifyAesKey(v string) *GetIdpConfigResponseBodyData {
	s.VerifyAesKey = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetVerifyToken(v string) *GetIdpConfigResponseBodyData {
	s.VerifyToken = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetVerifyUrl(v string) *GetIdpConfigResponseBodyData {
	s.VerifyUrl = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
