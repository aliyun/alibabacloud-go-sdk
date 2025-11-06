// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDnsHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDnsHostList(v []*QueryDnsHostResponseBodyDnsHostList) *QueryDnsHostResponseBody
	GetDnsHostList() []*QueryDnsHostResponseBodyDnsHostList
	SetRequestId(v string) *QueryDnsHostResponseBody
	GetRequestId() *string
}

type QueryDnsHostResponseBody struct {
	DnsHostList []*QueryDnsHostResponseBodyDnsHostList `json:"DnsHostList,omitempty" xml:"DnsHostList,omitempty" type:"Repeated"`
	// example:
	//
	// 18A313DD-3AF3-40AA-84F9-56BA45DC511F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDnsHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDnsHostResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDnsHostResponseBody) GetDnsHostList() []*QueryDnsHostResponseBodyDnsHostList {
	return s.DnsHostList
}

func (s *QueryDnsHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDnsHostResponseBody) SetDnsHostList(v []*QueryDnsHostResponseBodyDnsHostList) *QueryDnsHostResponseBody {
	s.DnsHostList = v
	return s
}

func (s *QueryDnsHostResponseBody) SetRequestId(v string) *QueryDnsHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDnsHostResponseBody) Validate() error {
	if s.DnsHostList != nil {
		for _, item := range s.DnsHostList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDnsHostResponseBodyDnsHostList struct {
	// example:
	//
	// ns3
	DnsName *string   `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	IpList  []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
}

func (s QueryDnsHostResponseBodyDnsHostList) String() string {
	return dara.Prettify(s)
}

func (s QueryDnsHostResponseBodyDnsHostList) GoString() string {
	return s.String()
}

func (s *QueryDnsHostResponseBodyDnsHostList) GetDnsName() *string {
	return s.DnsName
}

func (s *QueryDnsHostResponseBodyDnsHostList) GetIpList() []*string {
	return s.IpList
}

func (s *QueryDnsHostResponseBodyDnsHostList) SetDnsName(v string) *QueryDnsHostResponseBodyDnsHostList {
	s.DnsName = &v
	return s
}

func (s *QueryDnsHostResponseBodyDnsHostList) SetIpList(v []*string) *QueryDnsHostResponseBodyDnsHostList {
	s.IpList = v
	return s
}

func (s *QueryDnsHostResponseBodyDnsHostList) Validate() error {
	return dara.Validate(s)
}
