// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAccessModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainModeList(v []*DescribeDomainAccessModeResponseBodyDomainModeList) *DescribeDomainAccessModeResponseBody
	GetDomainModeList() []*DescribeDomainAccessModeResponseBodyDomainModeList
	SetRequestId(v string) *DescribeDomainAccessModeResponseBody
	GetRequestId() *string
}

type DescribeDomainAccessModeResponseBody struct {
	DomainModeList []*DescribeDomainAccessModeResponseBodyDomainModeList `json:"DomainModeList,omitempty" xml:"DomainModeList,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainAccessModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAccessModeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainAccessModeResponseBody) GetDomainModeList() []*DescribeDomainAccessModeResponseBodyDomainModeList {
	return s.DomainModeList
}

func (s *DescribeDomainAccessModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainAccessModeResponseBody) SetDomainModeList(v []*DescribeDomainAccessModeResponseBodyDomainModeList) *DescribeDomainAccessModeResponseBody {
	s.DomainModeList = v
	return s
}

func (s *DescribeDomainAccessModeResponseBody) SetRequestId(v string) *DescribeDomainAccessModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainAccessModeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainAccessModeResponseBodyDomainModeList struct {
	// example:
	//
	// 1
	AccessMode *int32 `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainAccessModeResponseBodyDomainModeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAccessModeResponseBodyDomainModeList) GoString() string {
	return s.String()
}

func (s *DescribeDomainAccessModeResponseBodyDomainModeList) GetAccessMode() *int32 {
	return s.AccessMode
}

func (s *DescribeDomainAccessModeResponseBodyDomainModeList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainAccessModeResponseBodyDomainModeList) SetAccessMode(v int32) *DescribeDomainAccessModeResponseBodyDomainModeList {
	s.AccessMode = &v
	return s
}

func (s *DescribeDomainAccessModeResponseBodyDomainModeList) SetDomain(v string) *DescribeDomainAccessModeResponseBodyDomainModeList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAccessModeResponseBodyDomainModeList) Validate() error {
	return dara.Validate(s)
}
