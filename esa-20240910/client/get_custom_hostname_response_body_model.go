// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomHostnameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomHostnameModel(v *GetCustomHostnameResponseBodyCustomHostnameModel) *GetCustomHostnameResponseBody
	GetCustomHostnameModel() *GetCustomHostnameResponseBodyCustomHostnameModel
	SetRequestId(v string) *GetCustomHostnameResponseBody
	GetRequestId() *string
}

type GetCustomHostnameResponseBody struct {
	CustomHostnameModel *GetCustomHostnameResponseBodyCustomHostnameModel `json:"CustomHostnameModel,omitempty" xml:"CustomHostnameModel,omitempty" type:"Struct"`
	// 本次请求的唯一标识
	//
	// example:
	//
	// 7C414690-9D7B-5D66-9CD9-AD0B3F25ED49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCustomHostnameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomHostnameResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomHostnameResponseBody) GetCustomHostnameModel() *GetCustomHostnameResponseBodyCustomHostnameModel {
	return s.CustomHostnameModel
}

func (s *GetCustomHostnameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomHostnameResponseBody) SetCustomHostnameModel(v *GetCustomHostnameResponseBodyCustomHostnameModel) *GetCustomHostnameResponseBody {
	s.CustomHostnameModel = v
	return s
}

func (s *GetCustomHostnameResponseBody) SetRequestId(v string) *GetCustomHostnameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomHostnameResponseBody) Validate() error {
	if s.CustomHostnameModel != nil {
		if err := s.CustomHostnameModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomHostnameResponseBodyCustomHostnameModel struct {
	// example:
	//
	// 30000478
	CasId *int64 `json:"CasId,omitempty" xml:"CasId,omitempty"`
	// 免费证书申请错误码
	//
	// example:
	//
	// 2
	CertApplyCode *int64 `json:"CertApplyCode,omitempty" xml:"CertApplyCode,omitempty"`
	// 免费证书申请错误说明
	//
	// example:
	//
	// canceled
	CertApplyMessage *string `json:"CertApplyMessage,omitempty" xml:"CertApplyMessage,omitempty"`
	// 证书校验HTTP名称
	//
	// example:
	//
	// http://custom.site.com/.well-known/acme-challenge/jLmMHlEaZ3jb352Qo3ciaSoAC8KZ5Hk0F-4_1xLQtgc
	CertHttpKey *string `json:"CertHttpKey,omitempty" xml:"CertHttpKey,omitempty"`
	// 证书校验HTTP内容
	//
	// example:
	//
	// jLmMHlEaZ3jb352Qo3ciaSoAC8KZ5Hk0F-4_1xLQtgc.GridYdfJJB5PgFEL-t89XfaFvMPB4f2-I9fwLpKl6e0
	CertHttpValue *string `json:"CertHttpValue,omitempty" xml:"CertHttpValue,omitempty"`
	// example:
	//
	// baba2c9e90e840b3b55698cedf02b308
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// 证书过期时间
	//
	// example:
	//
	// 2026-04-19T11:15:20Z
	CertNotAfter *string `json:"CertNotAfter,omitempty" xml:"CertNotAfter,omitempty"`
	// 证书状态
	//
	// example:
	//
	// OK
	CertStatus *string `json:"CertStatus,omitempty" xml:"CertStatus,omitempty"`
	// 证书校验TXT名称
	//
	// example:
	//
	// _acme-challenge.custom.site.com
	CertTxtKey *string `json:"CertTxtKey,omitempty" xml:"CertTxtKey,omitempty"`
	// 证书校验TXT内容
	//
	// example:
	//
	// lcKYad3UQXgrZLvMm_6TBUYKK4xTkGmninV0Mzx4gjM
	CertTxtValue *string `json:"CertTxtValue,omitempty" xml:"CertTxtValue,omitempty"`
	// 证书类型
	//
	// example:
	//
	// free
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// 上传的证书公钥
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// example:
	//
	// existing_custom_hostname
	ConflictWith *string `json:"ConflictWith,omitempty" xml:"ConflictWith,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 2026-04-19T11:15:20Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 用户自定义的主机名
	//
	// example:
	//
	// custom.site.com
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// 1234567890123
	HostnameId *int64 `json:"HostnameId,omitempty" xml:"HostnameId,omitempty"`
	// example:
	//
	// missing_icp
	OfflineReason *string `json:"OfflineReason,omitempty" xml:"OfflineReason,omitempty"`
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// 绑定的源站记录ID
	//
	// example:
	//
	// 4042843419650112
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// 绑定的源站记录名
	//
	// example:
	//
	// origin.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// 与主机名关联的站点ID
	//
	// example:
	//
	// 890601022130656
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// 关联站点名称
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// SSL开关的状态
	//
	// example:
	//
	// on
	SslFlag *string `json:"SslFlag,omitempty" xml:"SslFlag,omitempty"`
	// 自定义主机名状态
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间
	//
	// example:
	//
	// 2026-04-19T11:15:20Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 归属校验TXT内容
	//
	// example:
	//
	// verify_16ab7f4d389d4dff6655f995c6a997bd
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
	// 归属校验TXT名称
	//
	// example:
	//
	// _esa_custom_hostname.custom.site.com
	VerifyHost *string `json:"VerifyHost,omitempty" xml:"VerifyHost,omitempty"`
}

func (s GetCustomHostnameResponseBodyCustomHostnameModel) String() string {
	return dara.Prettify(s)
}

func (s GetCustomHostnameResponseBodyCustomHostnameModel) GoString() string {
	return s.String()
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetCasId() *int64 {
	return s.CasId
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetCertApplyCode() *int64 {
	return s.CertApplyCode
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetCertApplyMessage() *string {
	return s.CertApplyMessage
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetCertHttpKey() *string {
	return s.CertHttpKey
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetCertHttpValue() *string {
	return s.CertHttpValue
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetCertId() *string {
	return s.CertId
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetCertNotAfter() *string {
	return s.CertNotAfter
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetCertStatus() *string {
	return s.CertStatus
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetCertTxtKey() *string {
	return s.CertTxtKey
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetCertTxtValue() *string {
	return s.CertTxtValue
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetCertType() *string {
	return s.CertType
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetCertificate() *string {
	return s.Certificate
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetConflictWith() *string {
	return s.ConflictWith
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetHostname() *string {
	return s.Hostname
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetHostnameId() *int64 {
	return s.HostnameId
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetOfflineReason() *string {
	return s.OfflineReason
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetRecordId() *int64 {
	return s.RecordId
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetRecordName() *string {
	return s.RecordName
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetSiteName() *string {
	return s.SiteName
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetSslFlag() *string {
	return s.SslFlag
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetStatus() *string {
	return s.Status
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) GetVerifyHost() *string {
	return s.VerifyHost
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetCasId(v int64) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.CasId = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetCertApplyCode(v int64) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.CertApplyCode = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetCertApplyMessage(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.CertApplyMessage = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetCertHttpKey(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.CertHttpKey = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetCertHttpValue(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.CertHttpValue = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetCertId(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.CertId = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetCertNotAfter(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.CertNotAfter = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetCertStatus(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.CertStatus = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetCertTxtKey(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.CertTxtKey = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetCertTxtValue(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.CertTxtValue = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetCertType(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.CertType = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetCertificate(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.Certificate = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetConflictWith(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.ConflictWith = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetCreateTime(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.CreateTime = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetHostname(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.Hostname = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetHostnameId(v int64) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.HostnameId = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetOfflineReason(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.OfflineReason = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetPrivateKey(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.PrivateKey = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetRecordId(v int64) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.RecordId = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetRecordName(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.RecordName = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetSiteId(v int64) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.SiteId = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetSiteName(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.SiteName = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetSslFlag(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.SslFlag = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetStatus(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.Status = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetUpdateTime(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.UpdateTime = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetVerifyCode(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.VerifyCode = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) SetVerifyHost(v string) *GetCustomHostnameResponseBodyCustomHostnameModel {
	s.VerifyHost = &v
	return s
}

func (s *GetCustomHostnameResponseBodyCustomHostnameModel) Validate() error {
	return dara.Validate(s)
}
