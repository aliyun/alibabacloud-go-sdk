// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomHostnamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostnames(v []*ListCustomHostnamesResponseBodyHostnames) *ListCustomHostnamesResponseBody
	GetHostnames() []*ListCustomHostnamesResponseBodyHostnames
	SetPageNumber(v int32) *ListCustomHostnamesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomHostnamesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCustomHostnamesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCustomHostnamesResponseBody
	GetTotalCount() *int32
}

type ListCustomHostnamesResponseBody struct {
	Hostnames []*ListCustomHostnamesResponseBodyHostnames `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomHostnamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomHostnamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomHostnamesResponseBody) GetHostnames() []*ListCustomHostnamesResponseBodyHostnames {
	return s.Hostnames
}

func (s *ListCustomHostnamesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomHostnamesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomHostnamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomHostnamesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomHostnamesResponseBody) SetHostnames(v []*ListCustomHostnamesResponseBodyHostnames) *ListCustomHostnamesResponseBody {
	s.Hostnames = v
	return s
}

func (s *ListCustomHostnamesResponseBody) SetPageNumber(v int32) *ListCustomHostnamesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomHostnamesResponseBody) SetPageSize(v int32) *ListCustomHostnamesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomHostnamesResponseBody) SetRequestId(v string) *ListCustomHostnamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomHostnamesResponseBody) SetTotalCount(v int32) *ListCustomHostnamesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomHostnamesResponseBody) Validate() error {
	if s.Hostnames != nil {
		for _, item := range s.Hostnames {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomHostnamesResponseBodyHostnames struct {
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
	// 证书过期时间
	//
	// example:
	//
	// 2026-04-19T11:15:20Z
	CertNotAfter *string `json:"CertNotAfter,omitempty" xml:"CertNotAfter,omitempty"`
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
	// 绑定的源站记录ID
	//
	// example:
	//
	// 3386032455886912
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

func (s ListCustomHostnamesResponseBodyHostnames) String() string {
	return dara.Prettify(s)
}

func (s ListCustomHostnamesResponseBodyHostnames) GoString() string {
	return s.String()
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetCertApplyCode() *int64 {
	return s.CertApplyCode
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetCertApplyMessage() *string {
	return s.CertApplyMessage
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetCertHttpKey() *string {
	return s.CertHttpKey
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetCertHttpValue() *string {
	return s.CertHttpValue
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetCertNotAfter() *string {
	return s.CertNotAfter
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetCertStatus() *string {
	return s.CertStatus
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetCertTxtKey() *string {
	return s.CertTxtKey
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetCertTxtValue() *string {
	return s.CertTxtValue
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetCertType() *string {
	return s.CertType
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetConflictWith() *string {
	return s.ConflictWith
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetHostname() *string {
	return s.Hostname
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetHostnameId() *int64 {
	return s.HostnameId
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetOfflineReason() *string {
	return s.OfflineReason
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetRecordId() *int64 {
	return s.RecordId
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetRecordName() *string {
	return s.RecordName
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetSiteName() *string {
	return s.SiteName
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetSslFlag() *string {
	return s.SslFlag
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetStatus() *string {
	return s.Status
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *ListCustomHostnamesResponseBodyHostnames) GetVerifyHost() *string {
	return s.VerifyHost
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetCertApplyCode(v int64) *ListCustomHostnamesResponseBodyHostnames {
	s.CertApplyCode = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetCertApplyMessage(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.CertApplyMessage = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetCertHttpKey(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.CertHttpKey = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetCertHttpValue(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.CertHttpValue = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetCertNotAfter(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.CertNotAfter = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetCertStatus(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.CertStatus = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetCertTxtKey(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.CertTxtKey = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetCertTxtValue(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.CertTxtValue = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetCertType(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.CertType = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetConflictWith(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.ConflictWith = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetCreateTime(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.CreateTime = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetHostname(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.Hostname = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetHostnameId(v int64) *ListCustomHostnamesResponseBodyHostnames {
	s.HostnameId = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetOfflineReason(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.OfflineReason = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetRecordId(v int64) *ListCustomHostnamesResponseBodyHostnames {
	s.RecordId = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetRecordName(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.RecordName = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetSiteId(v int64) *ListCustomHostnamesResponseBodyHostnames {
	s.SiteId = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetSiteName(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.SiteName = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetSslFlag(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.SslFlag = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetStatus(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.Status = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetUpdateTime(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.UpdateTime = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetVerifyCode(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.VerifyCode = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) SetVerifyHost(v string) *ListCustomHostnamesResponseBodyHostnames {
	s.VerifyHost = &v
	return s
}

func (s *ListCustomHostnamesResponseBodyHostnames) Validate() error {
	return dara.Validate(s)
}
