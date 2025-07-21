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
	// CNAME verification flag, 0 for success, 1 for failure.
	//
	// example:
	//
	// 1
	CnameAuthStatus *string `json:"CnameAuthStatus,omitempty" xml:"CnameAuthStatus,omitempty"`
	// Indicates whether the CNAME host record has been modified, 1 for modified (reverting to the original value also counts as modification), 0 for not modified.
	//
	// example:
	//
	// 0
	CnameConfirmStatus *string `json:"CnameConfirmStatus,omitempty" xml:"CnameConfirmStatus,omitempty"`
	// Custom part of the CNAME host record
	//
	// example:
	//
	// dmtrace
	CnameRecord *string `json:"CnameRecord,omitempty" xml:"CnameRecord,omitempty"`
	// Creation time
	//
	// example:
	//
	// 2025-03-19T12:49Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Whether it is the default domain,
	//
	// Value: 0 No (this field is deprecated)
	//
	// example:
	//
	// 0
	DefaultDomain *string `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	// DKIM verification flag, indicating whether the DKIM record set by the user in DNS has passed validation, 0: Passed, 1: Not passed
	//
	// example:
	//
	// 0
	DkimAuthStatus *string `json:"DkimAuthStatus,omitempty" xml:"DkimAuthStatus,omitempty"`
	// DKIM public key value, the value that users need to set for the DKIM record in DNS
	//
	// example:
	//
	// v=DKIM1; k=rsa; p=MIGfMA0GCSqGSI...
	DkimPublicKey *string `json:"DkimPublicKey,omitempty" xml:"DkimPublicKey,omitempty"`
	// DKIM host record, the key that the user needs to set in the DNS for the DKIM record
	//
	// example:
	//
	// aliyun-cn-hangzhou._domainkey.hangzhou26
	DkimRR        *string `json:"DkimRR,omitempty" xml:"DkimRR,omitempty"`
	DkimRsaLength *int32  `json:"DkimRsaLength,omitempty" xml:"DkimRsaLength,omitempty"`
	// DMARC verification flag, indicating whether the DMARC record set by the user in DNS has passed validation, 0: Passed, 1: Not passed
	//
	// example:
	//
	// 1
	DmarcAuthStatus *int32 `json:"DmarcAuthStatus,omitempty" xml:"DmarcAuthStatus,omitempty"`
	// DMARC host record value
	//
	// example:
	//
	// _dmarc.xxx
	DmarcHostRecord *string `json:"DmarcHostRecord,omitempty" xml:"DmarcHostRecord,omitempty"`
	// DMARC record value
	//
	// example:
	//
	// v=DMARC1;p=none;rua=mailto:dmarc_report@service.aliyun.com
	DmarcRecord *string `json:"DmarcRecord,omitempty" xml:"DmarcRecord,omitempty"`
	// DMARC record value resolved through the public domain name
	//
	// example:
	//
	// v=DMARC1;p=none;rua=mailto:dmarc_report@service.aliyun.com
	DnsDmarc *string `json:"DnsDmarc,omitempty" xml:"DnsDmarc,omitempty"`
	// MX record value resolved from the public network domain
	//
	// example:
	//
	// mx01.dm.aliyun.com
	DnsMx *string `json:"DnsMx,omitempty" xml:"DnsMx,omitempty"`
	// SPF record value resolved from the public network domain
	//
	// example:
	//
	// v=xxxx
	DnsSpf *string `json:"DnsSpf,omitempty" xml:"DnsSpf,omitempty"`
	// Ownership record value resolved from the public network domain
	//
	// example:
	//
	// 0c40d5f125af4e42892a
	DnsTxt *string `json:"DnsTxt,omitempty" xml:"DnsTxt,omitempty"`
	// Domain ID
	//
	// example:
	//
	// 158910
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// Domain name
	//
	// example:
	//
	// test.example.net
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Domain status. Indicates whether the verification was successful, with values:
	//
	// - **0**: Available, verified successfully
	//
	// - **1**: Unavailable, verification failed
	//
	// example:
	//
	// 1
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// Ownership record provided by the email push console
	//
	// example:
	//
	// 0c40d5f125af4e42892a
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// Host record
	//
	// example:
	//
	// xxx
	HostRecord *string `json:"HostRecord,omitempty" xml:"HostRecord,omitempty"`
	// Filing status. **1*	- indicates filed, **0*	- indicates not filed.
	//
	// example:
	//
	// 1
	IcpStatus *string `json:"IcpStatus,omitempty" xml:"IcpStatus,omitempty"`
	// MX verification flag, 0 for success, 1 for failure.
	//
	// example:
	//
	// 1
	MxAuthStatus *string `json:"MxAuthStatus,omitempty" xml:"MxAuthStatus,omitempty"`
	// MX record value provided by the email push console
	//
	// example:
	//
	// mx01.dm.aliyun.com
	MxRecord *string `json:"MxRecord,omitempty" xml:"MxRecord,omitempty"`
	// Request ID
	//
	// example:
	//
	// 51B74264-46B4-43C8-A9A0-6B8E8BC04F34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// SPF verification flag, 0 for success, 1 for failure.
	//
	// example:
	//
	// 1
	SpfAuthStatus *string `json:"SpfAuthStatus,omitempty" xml:"SpfAuthStatus,omitempty"`
	// SPF record value provided by the email push console
	//
	// example:
	//
	// include:spf1.dm.aliyun.com
	SpfRecord *string `json:"SpfRecord,omitempty" xml:"SpfRecord,omitempty"`
	// SPF record. Previously, the SPF display content needed to be calculated by the calling end based on the spfRecord in the response. The new field spfRecordV2 replaces spfRecord, and the calling end can directly display this field after obtaining it;
	//
	// example:
	//
	// v=spf1 include:spf1.dm.aliyun.com -all
	SpfRecordV2 *string `json:"SpfRecordV2,omitempty" xml:"SpfRecordV2,omitempty"`
	// Primary domain
	//
	// example:
	//
	// example.com
	TlDomainName *string `json:"TlDomainName,omitempty" xml:"TlDomainName,omitempty"`
	// CNAME record value provided by the email push console
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
