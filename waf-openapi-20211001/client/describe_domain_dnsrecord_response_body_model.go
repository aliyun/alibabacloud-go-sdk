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
	// The DNS status. Valid values:
	//
	// - **cnameMatched**: Normal.
	//
	// - **vipMatched**: A record.
	//
	// - **wafVip**: The VIP of another WAF is used.
	//
	// - **unRecord**: No DNS resolution is configured.
	//
	// - **unUsed**: Traffic does not pass through WAF.
	//
	// - **checkTimeout**: The detection timed out.
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
