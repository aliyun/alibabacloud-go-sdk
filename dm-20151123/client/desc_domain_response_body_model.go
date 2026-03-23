// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCnameAuthStatus(v string) *DescDomainResponseBody
	GetCnameAuthStatus() *string
	SetCnameConfirmStatus(v string) *DescDomainResponseBody
	GetCnameConfirmStatus() *string
	SetCnameRecord(v string) *DescDomainResponseBody
	GetCnameRecord() *string
	SetCreateTime(v string) *DescDomainResponseBody
	GetCreateTime() *string
	SetDefaultDomain(v string) *DescDomainResponseBody
	GetDefaultDomain() *string
	SetDkimAuthStatus(v string) *DescDomainResponseBody
	GetDkimAuthStatus() *string
	SetDkimPublicKey(v string) *DescDomainResponseBody
	GetDkimPublicKey() *string
	SetDkimRR(v string) *DescDomainResponseBody
	GetDkimRR() *string
	SetDkimRsaLength(v int32) *DescDomainResponseBody
	GetDkimRsaLength() *int32
	SetDmarcAuthStatus(v int32) *DescDomainResponseBody
	GetDmarcAuthStatus() *int32
	SetDmarcHostRecord(v string) *DescDomainResponseBody
	GetDmarcHostRecord() *string
	SetDmarcRecord(v string) *DescDomainResponseBody
	GetDmarcRecord() *string
	SetDnsDmarc(v string) *DescDomainResponseBody
	GetDnsDmarc() *string
	SetDnsMx(v string) *DescDomainResponseBody
	GetDnsMx() *string
	SetDnsSpf(v string) *DescDomainResponseBody
	GetDnsSpf() *string
	SetDnsTxt(v string) *DescDomainResponseBody
	GetDnsTxt() *string
	SetDomainId(v string) *DescDomainResponseBody
	GetDomainId() *string
	SetDomainName(v string) *DescDomainResponseBody
	GetDomainName() *string
	SetDomainStatus(v string) *DescDomainResponseBody
	GetDomainStatus() *string
	SetDomainType(v string) *DescDomainResponseBody
	GetDomainType() *string
	SetHostRecord(v string) *DescDomainResponseBody
	GetHostRecord() *string
	SetIcpStatus(v string) *DescDomainResponseBody
	GetIcpStatus() *string
	SetMxAuthStatus(v string) *DescDomainResponseBody
	GetMxAuthStatus() *string
	SetMxRecord(v string) *DescDomainResponseBody
	GetMxRecord() *string
	SetRequestId(v string) *DescDomainResponseBody
	GetRequestId() *string
	SetSpfAuthStatus(v string) *DescDomainResponseBody
	GetSpfAuthStatus() *string
	SetSpfRecord(v string) *DescDomainResponseBody
	GetSpfRecord() *string
	SetSpfRecordV2(v string) *DescDomainResponseBody
	GetSpfRecordV2() *string
	SetTlDomainName(v string) *DescDomainResponseBody
	GetTlDomainName() *string
	SetTracefRecord(v string) *DescDomainResponseBody
	GetTracefRecord() *string
}

type DescDomainResponseBody struct {
	// The CNAME authentication flag. 0: Succeeded. 1: Failed.
	//
	// example:
	//
	// 1
	CnameAuthStatus *string `json:"CnameAuthStatus,omitempty" xml:"CnameAuthStatus,omitempty"`
	// Indicates whether the CNAME host record was modified. A value of 1 means the record was modified. Reverting to the original value is also considered a modification. A value of 0 means the record was not modified.
	//
	// example:
	//
	// 0
	CnameConfirmStatus *string `json:"CnameConfirmStatus,omitempty" xml:"CnameConfirmStatus,omitempty"`
	// The custom part of the CNAME host record.
	//
	// example:
	//
	// dmtrace
	CnameRecord *string `json:"CnameRecord,omitempty" xml:"CnameRecord,omitempty"`
	// The time when the domain name was created.
	//
	// example:
	//
	// 2025-03-19T12:49Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the domain name is the default domain name.
	//
	// Value: 0 (No). This field is deprecated.
	//
	// example:
	//
	// 0
	DefaultDomain *string `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	// The DKIM authentication flag. Indicates if the DKIM record in your DNS settings passed verification. 0: Passed. 1: Not passed.
	//
	// example:
	//
	// 0
	DkimAuthStatus *string `json:"DkimAuthStatus,omitempty" xml:"DkimAuthStatus,omitempty"`
	// The DKIM public key. This is the value of the DKIM record to configure in your DNS settings.
	//
	// example:
	//
	// v=DKIM1; k=rsa; p=MIGfMA0GCSqGSI...
	DkimPublicKey *string `json:"DkimPublicKey,omitempty" xml:"DkimPublicKey,omitempty"`
	// The DKIM host record. This is the key of the DKIM record to configure in your DNS settings.
	//
	// example:
	//
	// aliyun-cn-hangzhou._domainkey.hangzhou26
	DkimRR        *string `json:"DkimRR,omitempty" xml:"DkimRR,omitempty"`
	DkimRsaLength *int32  `json:"DkimRsaLength,omitempty" xml:"DkimRsaLength,omitempty"`
	// The DMARC authentication flag. Indicates if the DMARC record in your DNS settings passed verification. 0: Passed. 1: Not passed.
	//
	// example:
	//
	// 1
	DmarcAuthStatus *int32 `json:"DmarcAuthStatus,omitempty" xml:"DmarcAuthStatus,omitempty"`
	// The DMARC host record value.
	//
	// example:
	//
	// _dmarc.xxx
	DmarcHostRecord *string `json:"DmarcHostRecord,omitempty" xml:"DmarcHostRecord,omitempty"`
	// The DMARC record value.
	//
	// example:
	//
	// v=DMARC1;p=none;rua=mailto:dmarc_report@service.aliyun.com
	DmarcRecord *string `json:"DmarcRecord,omitempty" xml:"DmarcRecord,omitempty"`
	// The DMARC record value parsed from the public domain name.
	//
	// example:
	//
	// v=DMARC1;p=none;rua=mailto:dmarc_report@service.aliyun.com
	DnsDmarc *string `json:"DnsDmarc,omitempty" xml:"DnsDmarc,omitempty"`
	// The MX record value parsed from the public domain name.
	//
	// example:
	//
	// mx01.dm.aliyun.com
	DnsMx *string `json:"DnsMx,omitempty" xml:"DnsMx,omitempty"`
	// The SPF record value parsed from the public domain name.
	//
	// example:
	//
	// v=xxxx
	DnsSpf *string `json:"DnsSpf,omitempty" xml:"DnsSpf,omitempty"`
	// The ownership record value parsed from the public domain name.
	//
	// example:
	//
	// 0c40d5f125af4e42892a
	DnsTxt *string `json:"DnsTxt,omitempty" xml:"DnsTxt,omitempty"`
	// The domain name ID.
	//
	// example:
	//
	// 158910
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// test.example.net
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The domain status. This indicates whether the domain name passed authentication. Valid values:
	//
	// - **0**: Active. The domain name passed authentication.
	//
	// - **1**: Inactive. The domain name failed authentication.
	//
	// example:
	//
	// 1
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The ownership record provided by the Direct Mail console.
	//
	// example:
	//
	// 0c40d5f125af4e42892a
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The host record.
	//
	// example:
	//
	// xxx
	HostRecord *string `json:"HostRecord,omitempty" xml:"HostRecord,omitempty"`
	// The ICP filing status. **1*	- indicates that the domain name has an ICP filing. **0*	- indicates that the domain name does not have an ICP filing.
	//
	// example:
	//
	// 1
	IcpStatus *string `json:"IcpStatus,omitempty" xml:"IcpStatus,omitempty"`
	// The MX authentication flag. 0: Succeeded. 1: Failed.
	//
	// example:
	//
	// 1
	MxAuthStatus *string `json:"MxAuthStatus,omitempty" xml:"MxAuthStatus,omitempty"`
	// The MX record value provided by the Direct Mail console.
	//
	// example:
	//
	// mx01.dm.aliyun.com
	MxRecord *string `json:"MxRecord,omitempty" xml:"MxRecord,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 51B74264-46B4-43C8-A9A0-6B8E8BC04F34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SPF authentication flag. 0: Succeeded. 1: Failed.
	//
	// example:
	//
	// 1
	SpfAuthStatus *string `json:"SpfAuthStatus,omitempty" xml:"SpfAuthStatus,omitempty"`
	// The SPF record value provided by the Direct Mail console.
	//
	// example:
	//
	// include:spf1.dm.aliyun.com
	SpfRecord *string `json:"SpfRecord,omitempty" xml:"SpfRecord,omitempty"`
	// The SPF record. This field replaces the \\`spfRecord\\` field. You can directly display the value of this field without needing to calculate it from the response.
	//
	// example:
	//
	// v=spf1 include:spf1.dm.aliyun.com -all
	SpfRecordV2 *string `json:"SpfRecordV2,omitempty" xml:"SpfRecordV2,omitempty"`
	// The primary domain name.
	//
	// example:
	//
	// example.com
	TlDomainName *string `json:"TlDomainName,omitempty" xml:"TlDomainName,omitempty"`
	// The CNAME record value provided by the Direct Mail console.
	//
	// example:
	//
	// tracedm.aliyuncs.com
	TracefRecord *string `json:"TracefRecord,omitempty" xml:"TracefRecord,omitempty"`
}

func (s DescDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescDomainResponseBody) GetCnameAuthStatus() *string {
	return s.CnameAuthStatus
}

func (s *DescDomainResponseBody) GetCnameConfirmStatus() *string {
	return s.CnameConfirmStatus
}

func (s *DescDomainResponseBody) GetCnameRecord() *string {
	return s.CnameRecord
}

func (s *DescDomainResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescDomainResponseBody) GetDefaultDomain() *string {
	return s.DefaultDomain
}

func (s *DescDomainResponseBody) GetDkimAuthStatus() *string {
	return s.DkimAuthStatus
}

func (s *DescDomainResponseBody) GetDkimPublicKey() *string {
	return s.DkimPublicKey
}

func (s *DescDomainResponseBody) GetDkimRR() *string {
	return s.DkimRR
}

func (s *DescDomainResponseBody) GetDkimRsaLength() *int32 {
	return s.DkimRsaLength
}

func (s *DescDomainResponseBody) GetDmarcAuthStatus() *int32 {
	return s.DmarcAuthStatus
}

func (s *DescDomainResponseBody) GetDmarcHostRecord() *string {
	return s.DmarcHostRecord
}

func (s *DescDomainResponseBody) GetDmarcRecord() *string {
	return s.DmarcRecord
}

func (s *DescDomainResponseBody) GetDnsDmarc() *string {
	return s.DnsDmarc
}

func (s *DescDomainResponseBody) GetDnsMx() *string {
	return s.DnsMx
}

func (s *DescDomainResponseBody) GetDnsSpf() *string {
	return s.DnsSpf
}

func (s *DescDomainResponseBody) GetDnsTxt() *string {
	return s.DnsTxt
}

func (s *DescDomainResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *DescDomainResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescDomainResponseBody) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescDomainResponseBody) GetDomainType() *string {
	return s.DomainType
}

func (s *DescDomainResponseBody) GetHostRecord() *string {
	return s.HostRecord
}

func (s *DescDomainResponseBody) GetIcpStatus() *string {
	return s.IcpStatus
}

func (s *DescDomainResponseBody) GetMxAuthStatus() *string {
	return s.MxAuthStatus
}

func (s *DescDomainResponseBody) GetMxRecord() *string {
	return s.MxRecord
}

func (s *DescDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescDomainResponseBody) GetSpfAuthStatus() *string {
	return s.SpfAuthStatus
}

func (s *DescDomainResponseBody) GetSpfRecord() *string {
	return s.SpfRecord
}

func (s *DescDomainResponseBody) GetSpfRecordV2() *string {
	return s.SpfRecordV2
}

func (s *DescDomainResponseBody) GetTlDomainName() *string {
	return s.TlDomainName
}

func (s *DescDomainResponseBody) GetTracefRecord() *string {
	return s.TracefRecord
}

func (s *DescDomainResponseBody) SetCnameAuthStatus(v string) *DescDomainResponseBody {
	s.CnameAuthStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetCnameConfirmStatus(v string) *DescDomainResponseBody {
	s.CnameConfirmStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetCnameRecord(v string) *DescDomainResponseBody {
	s.CnameRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetCreateTime(v string) *DescDomainResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescDomainResponseBody) SetDefaultDomain(v string) *DescDomainResponseBody {
	s.DefaultDomain = &v
	return s
}

func (s *DescDomainResponseBody) SetDkimAuthStatus(v string) *DescDomainResponseBody {
	s.DkimAuthStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetDkimPublicKey(v string) *DescDomainResponseBody {
	s.DkimPublicKey = &v
	return s
}

func (s *DescDomainResponseBody) SetDkimRR(v string) *DescDomainResponseBody {
	s.DkimRR = &v
	return s
}

func (s *DescDomainResponseBody) SetDkimRsaLength(v int32) *DescDomainResponseBody {
	s.DkimRsaLength = &v
	return s
}

func (s *DescDomainResponseBody) SetDmarcAuthStatus(v int32) *DescDomainResponseBody {
	s.DmarcAuthStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetDmarcHostRecord(v string) *DescDomainResponseBody {
	s.DmarcHostRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetDmarcRecord(v string) *DescDomainResponseBody {
	s.DmarcRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetDnsDmarc(v string) *DescDomainResponseBody {
	s.DnsDmarc = &v
	return s
}

func (s *DescDomainResponseBody) SetDnsMx(v string) *DescDomainResponseBody {
	s.DnsMx = &v
	return s
}

func (s *DescDomainResponseBody) SetDnsSpf(v string) *DescDomainResponseBody {
	s.DnsSpf = &v
	return s
}

func (s *DescDomainResponseBody) SetDnsTxt(v string) *DescDomainResponseBody {
	s.DnsTxt = &v
	return s
}

func (s *DescDomainResponseBody) SetDomainId(v string) *DescDomainResponseBody {
	s.DomainId = &v
	return s
}

func (s *DescDomainResponseBody) SetDomainName(v string) *DescDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescDomainResponseBody) SetDomainStatus(v string) *DescDomainResponseBody {
	s.DomainStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetDomainType(v string) *DescDomainResponseBody {
	s.DomainType = &v
	return s
}

func (s *DescDomainResponseBody) SetHostRecord(v string) *DescDomainResponseBody {
	s.HostRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetIcpStatus(v string) *DescDomainResponseBody {
	s.IcpStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetMxAuthStatus(v string) *DescDomainResponseBody {
	s.MxAuthStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetMxRecord(v string) *DescDomainResponseBody {
	s.MxRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetRequestId(v string) *DescDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescDomainResponseBody) SetSpfAuthStatus(v string) *DescDomainResponseBody {
	s.SpfAuthStatus = &v
	return s
}

func (s *DescDomainResponseBody) SetSpfRecord(v string) *DescDomainResponseBody {
	s.SpfRecord = &v
	return s
}

func (s *DescDomainResponseBody) SetSpfRecordV2(v string) *DescDomainResponseBody {
	s.SpfRecordV2 = &v
	return s
}

func (s *DescDomainResponseBody) SetTlDomainName(v string) *DescDomainResponseBody {
	s.TlDomainName = &v
	return s
}

func (s *DescDomainResponseBody) SetTracefRecord(v string) *DescDomainResponseBody {
	s.TracefRecord = &v
	return s
}

func (s *DescDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
