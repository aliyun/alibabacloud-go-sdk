// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDescribeCdnIpInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpInfoList(v []*BatchDescribeCdnIpInfoResponseBodyIpInfoList) *BatchDescribeCdnIpInfoResponseBody
	GetIpInfoList() []*BatchDescribeCdnIpInfoResponseBodyIpInfoList
	SetRequestId(v string) *BatchDescribeCdnIpInfoResponseBody
	GetRequestId() *string
}

type BatchDescribeCdnIpInfoResponseBody struct {
	// The results about IP addresses returned.
	IpInfoList []*BatchDescribeCdnIpInfoResponseBodyIpInfoList `json:"IpInfoList,omitempty" xml:"IpInfoList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 55ADD936-763F-5E1A-BF54-2EA3F6E94A52
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDescribeCdnIpInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDescribeCdnIpInfoResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDescribeCdnIpInfoResponseBody) GetIpInfoList() []*BatchDescribeCdnIpInfoResponseBodyIpInfoList {
	return s.IpInfoList
}

func (s *BatchDescribeCdnIpInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDescribeCdnIpInfoResponseBody) SetIpInfoList(v []*BatchDescribeCdnIpInfoResponseBodyIpInfoList) *BatchDescribeCdnIpInfoResponseBody {
	s.IpInfoList = v
	return s
}

func (s *BatchDescribeCdnIpInfoResponseBody) SetRequestId(v string) *BatchDescribeCdnIpInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDescribeCdnIpInfoResponseBody) Validate() error {
	if s.IpInfoList != nil {
		for _, item := range s.IpInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchDescribeCdnIpInfoResponseBodyIpInfoList struct {
	// Indicates whether the IP address belongs to an Alibaba Cloud CDN point of presence (POP).
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CdnIp *string `json:"CdnIp,omitempty" xml:"CdnIp,omitempty"`
	// The city to which the IP address belongs.
	//
	// example:
	//
	// Beijing
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The country to which the IP address belongs.
	//
	// example:
	//
	// China
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 111.XXX.XXX.230
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The ISP to which the IP address belongs.
	//
	// example:
	//
	// Move
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// The province to which the IP address belongs.
	//
	// example:
	//
	// Beijing
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s BatchDescribeCdnIpInfoResponseBodyIpInfoList) String() string {
	return dara.Prettify(s)
}

func (s BatchDescribeCdnIpInfoResponseBodyIpInfoList) GoString() string {
	return s.String()
}

func (s *BatchDescribeCdnIpInfoResponseBodyIpInfoList) GetCdnIp() *string {
	return s.CdnIp
}

func (s *BatchDescribeCdnIpInfoResponseBodyIpInfoList) GetCity() *string {
	return s.City
}

func (s *BatchDescribeCdnIpInfoResponseBodyIpInfoList) GetCountry() *string {
	return s.Country
}

func (s *BatchDescribeCdnIpInfoResponseBodyIpInfoList) GetIpAddress() *string {
	return s.IpAddress
}

func (s *BatchDescribeCdnIpInfoResponseBodyIpInfoList) GetIspName() *string {
	return s.IspName
}

func (s *BatchDescribeCdnIpInfoResponseBodyIpInfoList) GetProvince() *string {
	return s.Province
}

func (s *BatchDescribeCdnIpInfoResponseBodyIpInfoList) SetCdnIp(v string) *BatchDescribeCdnIpInfoResponseBodyIpInfoList {
	s.CdnIp = &v
	return s
}

func (s *BatchDescribeCdnIpInfoResponseBodyIpInfoList) SetCity(v string) *BatchDescribeCdnIpInfoResponseBodyIpInfoList {
	s.City = &v
	return s
}

func (s *BatchDescribeCdnIpInfoResponseBodyIpInfoList) SetCountry(v string) *BatchDescribeCdnIpInfoResponseBodyIpInfoList {
	s.Country = &v
	return s
}

func (s *BatchDescribeCdnIpInfoResponseBodyIpInfoList) SetIpAddress(v string) *BatchDescribeCdnIpInfoResponseBodyIpInfoList {
	s.IpAddress = &v
	return s
}

func (s *BatchDescribeCdnIpInfoResponseBodyIpInfoList) SetIspName(v string) *BatchDescribeCdnIpInfoResponseBodyIpInfoList {
	s.IspName = &v
	return s
}

func (s *BatchDescribeCdnIpInfoResponseBodyIpInfoList) SetProvince(v string) *BatchDescribeCdnIpInfoResponseBodyIpInfoList {
	s.Province = &v
	return s
}

func (s *BatchDescribeCdnIpInfoResponseBodyIpInfoList) Validate() error {
	return dara.Validate(s)
}
