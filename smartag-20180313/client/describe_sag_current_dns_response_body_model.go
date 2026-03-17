// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagCurrentDnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMasterDns(v string) *DescribeSagCurrentDnsResponseBody
	GetMasterDns() *string
	SetRequestId(v string) *DescribeSagCurrentDnsResponseBody
	GetRequestId() *string
	SetSlaveDns(v string) *DescribeSagCurrentDnsResponseBody
	GetSlaveDns() *string
}

type DescribeSagCurrentDnsResponseBody struct {
	// The IP address of the primary DNS server.
	//
	// example:
	//
	// 223.XX.XX.5
	MasterDns *string `json:"MasterDns,omitempty" xml:"MasterDns,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0937DEA0-AB4B-42F4-9314-07B97D30282B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IP address of the secondary DNS server.
	//
	// example:
	//
	// 114.XX.XX.114
	SlaveDns *string `json:"SlaveDns,omitempty" xml:"SlaveDns,omitempty"`
}

func (s DescribeSagCurrentDnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagCurrentDnsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagCurrentDnsResponseBody) GetMasterDns() *string {
	return s.MasterDns
}

func (s *DescribeSagCurrentDnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagCurrentDnsResponseBody) GetSlaveDns() *string {
	return s.SlaveDns
}

func (s *DescribeSagCurrentDnsResponseBody) SetMasterDns(v string) *DescribeSagCurrentDnsResponseBody {
	s.MasterDns = &v
	return s
}

func (s *DescribeSagCurrentDnsResponseBody) SetRequestId(v string) *DescribeSagCurrentDnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagCurrentDnsResponseBody) SetSlaveDns(v string) *DescribeSagCurrentDnsResponseBody {
	s.SlaveDns = &v
	return s
}

func (s *DescribeSagCurrentDnsResponseBody) Validate() error {
	return dara.Validate(s)
}
