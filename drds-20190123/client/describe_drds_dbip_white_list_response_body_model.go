// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDBIpWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpWhiteList(v *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) *DescribeDrdsDBIpWhiteListResponseBody
	GetIpWhiteList() *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList
	SetRequestId(v string) *DescribeDrdsDBIpWhiteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsDBIpWhiteListResponseBody
	GetSuccess() *bool
}

type DescribeDrdsDBIpWhiteListResponseBody struct {
	// The IP address whitelist.
	IpWhiteList *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList `json:"IpWhiteList,omitempty" xml:"IpWhiteList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 017453B9-0001-4745-87BF-DD612D850ED0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDBIpWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) GetIpWhiteList() *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList {
	return s.IpWhiteList
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) SetIpWhiteList(v *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) *DescribeDrdsDBIpWhiteListResponseBody {
	s.IpWhiteList = v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) SetRequestId(v string) *DescribeDrdsDBIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) SetSuccess(v bool) *DescribeDrdsDBIpWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList struct {
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) GetIp() []*string {
	return s.Ip
}

func (s *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) SetIp(v []*string) *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList {
	s.Ip = v
	return s
}

func (s *DescribeDrdsDBIpWhiteListResponseBodyIpWhiteList) Validate() error {
	return dara.Validate(s)
}
