// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainDNSRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDNSStatus(v string) *DescribeDomainDNSRecordResponseBody
	GetDNSStatus() *string
	SetRequestId(v string) *DescribeDomainDNSRecordResponseBody
	GetRequestId() *string
}

type DescribeDomainDNSRecordResponseBody struct {
	// The status of the DNS settings. Valid values:
	//
	// 	- **cnameMatched**: The DNS settings are properly configured.
	//
	// 	- **vipMatched**: An A record maps the domain name to the WAF virtual IP address (VIP).
	//
	// 	- **wafVip**: An A record maps the domain name to another WAF VIP.
	//
	// 	- **unRecord**: The domain name does not have a DNS record.
	//
	// 	- **unUsed**: The domain name is not pointed to WAF.
	//
	// 	- **checkTimeout**: The check times out.
	//
	// example:
	//
	// cnameMatched
	DNSStatus *string `json:"DNSStatus,omitempty" xml:"DNSStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D827FCFE-90A7-4330-9326-D33C8B4C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainDNSRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDNSRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainDNSRecordResponseBody) GetDNSStatus() *string {
	return s.DNSStatus
}

func (s *DescribeDomainDNSRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainDNSRecordResponseBody) SetDNSStatus(v string) *DescribeDomainDNSRecordResponseBody {
	s.DNSStatus = &v
	return s
}

func (s *DescribeDomainDNSRecordResponseBody) SetRequestId(v string) *DescribeDomainDNSRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainDNSRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
