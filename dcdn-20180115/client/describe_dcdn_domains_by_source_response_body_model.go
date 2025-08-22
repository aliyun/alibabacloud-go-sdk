// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainsBySourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainInfo(v []*DescribeDcdnDomainsBySourceResponseBodyDomainInfo) *DescribeDcdnDomainsBySourceResponseBody
	GetDomainInfo() []*DescribeDcdnDomainsBySourceResponseBodyDomainInfo
	SetRequestId(v string) *DescribeDcdnDomainsBySourceResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainsBySourceResponseBody struct {
	// The information about each origin server and the corresponding domain names.
	//
	// This parameter is required.
	DomainInfo []*DescribeDcdnDomainsBySourceResponseBodyDomainInfo `json:"DomainInfo,omitempty" xml:"DomainInfo,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainsBySourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainsBySourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainsBySourceResponseBody) GetDomainInfo() []*DescribeDcdnDomainsBySourceResponseBodyDomainInfo {
	return s.DomainInfo
}

func (s *DescribeDcdnDomainsBySourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainsBySourceResponseBody) SetDomainInfo(v []*DescribeDcdnDomainsBySourceResponseBodyDomainInfo) *DescribeDcdnDomainsBySourceResponseBody {
	s.DomainInfo = v
	return s
}

func (s *DescribeDcdnDomainsBySourceResponseBody) SetRequestId(v string) *DescribeDcdnDomainsBySourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainsBySourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainsBySourceResponseBodyDomainInfo struct {
	// The information about the domain names.
	DomainList []*DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// The origin server.
	//
	// example:
	//
	// example.com
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeDcdnDomainsBySourceResponseBodyDomainInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainsBySourceResponseBodyDomainInfo) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfo) GetDomainList() []*DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList {
	return s.DomainList
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfo) GetSource() *string {
	return s.Source
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfo) SetDomainList(v []*DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) *DescribeDcdnDomainsBySourceResponseBodyDomainInfo {
	s.DomainList = v
	return s
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfo) SetSource(v string) *DescribeDcdnDomainsBySourceResponseBodyDomainInfo {
	s.Source = &v
	return s
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList struct {
	// The creation time.
	//
	// example:
	//
	// 2021-08-21T03:05:20+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The CNAME record assigned to the domain name.
	//
	// example:
	//
	// example.org.alikunlun.com
	DomainCname *string `json:"DomainCname,omitempty" xml:"DomainCname,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.org
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The workload type of the accelerated domain name. Valid value:
	//
	// 	- **ipa**: layer 4 acceleration
	//
	// 	- **dynamic**: layer 7 acceleration
	//
	// example:
	//
	// dynamic
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The status of the domain name. Valid value:
	//
	// 	- **applying**: The domain name is under review.
	//
	// 	- **configuring**: The domain name is being configured.
	//
	// 	- **online**: The domain name is working as expected.
	//
	// 	- **stopping**: The domain name is being stopped.
	//
	// 	- **offline**: The domain name is disabled.
	//
	// 	- **disabling**: The domain name is being removed.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the domain name was updated.
	//
	// example:
	//
	// 2022-12-01T03:26:55+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) GetDomainCname() *string {
	return s.DomainCname
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) SetCreateTime(v string) *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList {
	s.CreateTime = &v
	return s
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) SetDomainCname(v string) *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList {
	s.DomainCname = &v
	return s
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) SetDomainName(v string) *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) SetDomainType(v string) *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList {
	s.DomainType = &v
	return s
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) SetStatus(v string) *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList {
	s.Status = &v
	return s
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) SetUpdateTime(v string) *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDcdnDomainsBySourceResponseBodyDomainInfoDomainList) Validate() error {
	return dara.Validate(s)
}
