// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWhitelistSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertNo(v string) *CreateWhitelistSettingRequest
	GetCertNo() *string
	SetCertifyId(v string) *CreateWhitelistSettingRequest
	GetCertifyId() *string
	SetLang(v string) *CreateWhitelistSettingRequest
	GetLang() *string
	SetRemark(v string) *CreateWhitelistSettingRequest
	GetRemark() *string
	SetSceneId(v int64) *CreateWhitelistSettingRequest
	GetSceneId() *int64
	SetServiceCode(v string) *CreateWhitelistSettingRequest
	GetServiceCode() *string
	SetSourceIp(v string) *CreateWhitelistSettingRequest
	GetSourceIp() *string
	SetValidDay(v int32) *CreateWhitelistSettingRequest
	GetValidDay() *int32
	SetWhitelistType(v string) *CreateWhitelistSettingRequest
	GetWhitelistType() *string
}

type CreateWhitelistSettingRequest struct {
	// The ID card number to add to the whitelist.
	//
	// example:
	//
	// 44018219950810XXXX
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// The certificate ID. The user associated with this authentication is added to the whitelist.
	//
	// example:
	//
	// sha6d0405f42926084e396e76a037d00
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// The user language.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The remarks for the whitelist.
	//
	// example:
	//
	// 测试白名单。
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The scene ID.
	//
	// example:
	//
	// 1000014526
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The service type. Valid values:
	//
	// - **antcloudauth**: Financial-grade ID Verification.
	//
	// - **cloudauthst*	- (discontinued): ID Verification Enhanced Edition.
	//
	// This parameter is required.
	//
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The source IP address of the visitor. You do not need to specify this parameter. The system automatically obtains the value.
	//
	// example:
	//
	// 27.115.63.58
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The number of valid days after the whitelist is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	ValidDay *int32 `json:"ValidDay,omitempty" xml:"ValidDay,omitempty"`
	// example:
	//
	// SUPER
	WhitelistType *string `json:"WhitelistType,omitempty" xml:"WhitelistType,omitempty"`
}

func (s CreateWhitelistSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWhitelistSettingRequest) GoString() string {
	return s.String()
}

func (s *CreateWhitelistSettingRequest) GetCertNo() *string {
	return s.CertNo
}

func (s *CreateWhitelistSettingRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *CreateWhitelistSettingRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateWhitelistSettingRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateWhitelistSettingRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *CreateWhitelistSettingRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *CreateWhitelistSettingRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateWhitelistSettingRequest) GetValidDay() *int32 {
	return s.ValidDay
}

func (s *CreateWhitelistSettingRequest) GetWhitelistType() *string {
	return s.WhitelistType
}

func (s *CreateWhitelistSettingRequest) SetCertNo(v string) *CreateWhitelistSettingRequest {
	s.CertNo = &v
	return s
}

func (s *CreateWhitelistSettingRequest) SetCertifyId(v string) *CreateWhitelistSettingRequest {
	s.CertifyId = &v
	return s
}

func (s *CreateWhitelistSettingRequest) SetLang(v string) *CreateWhitelistSettingRequest {
	s.Lang = &v
	return s
}

func (s *CreateWhitelistSettingRequest) SetRemark(v string) *CreateWhitelistSettingRequest {
	s.Remark = &v
	return s
}

func (s *CreateWhitelistSettingRequest) SetSceneId(v int64) *CreateWhitelistSettingRequest {
	s.SceneId = &v
	return s
}

func (s *CreateWhitelistSettingRequest) SetServiceCode(v string) *CreateWhitelistSettingRequest {
	s.ServiceCode = &v
	return s
}

func (s *CreateWhitelistSettingRequest) SetSourceIp(v string) *CreateWhitelistSettingRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateWhitelistSettingRequest) SetValidDay(v int32) *CreateWhitelistSettingRequest {
	s.ValidDay = &v
	return s
}

func (s *CreateWhitelistSettingRequest) SetWhitelistType(v string) *CreateWhitelistSettingRequest {
	s.WhitelistType = &v
	return s
}

func (s *CreateWhitelistSettingRequest) Validate() error {
	return dara.Validate(s)
}
