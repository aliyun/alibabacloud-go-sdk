// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBgpPackByIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeBgpPackByIpResponseBody
	GetCode() *int32
	SetDdosbgpInfo(v *DescribeBgpPackByIpResponseBodyDdosbgpInfo) *DescribeBgpPackByIpResponseBody
	GetDdosbgpInfo() *DescribeBgpPackByIpResponseBodyDdosbgpInfo
	SetRequestId(v string) *DescribeBgpPackByIpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBgpPackByIpResponseBody
	GetSuccess() *bool
}

type DescribeBgpPackByIpResponseBody struct {
	// The HTTP status code of the request.
	//
	// For more information about status codes, see [Common parameters](https://help.aliyun.com/document_detail/118841.html).
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configurations of the instance that is associated with the asset.
	DdosbgpInfo *DescribeBgpPackByIpResponseBodyDdosbgpInfo `json:"DdosbgpInfo,omitempty" xml:"DdosbgpInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E46A08E4-A1CD-5BE9-B580-C4D6E9BC5484
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBgpPackByIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpPackByIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackByIpResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeBgpPackByIpResponseBody) GetDdosbgpInfo() *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	return s.DdosbgpInfo
}

func (s *DescribeBgpPackByIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBgpPackByIpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBgpPackByIpResponseBody) SetCode(v int32) *DescribeBgpPackByIpResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBody) SetDdosbgpInfo(v *DescribeBgpPackByIpResponseBodyDdosbgpInfo) *DescribeBgpPackByIpResponseBody {
	s.DdosbgpInfo = v
	return s
}

func (s *DescribeBgpPackByIpResponseBody) SetRequestId(v string) *DescribeBgpPackByIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBody) SetSuccess(v bool) *DescribeBgpPackByIpResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBgpPackByIpResponseBodyDdosbgpInfo struct {
	// The basic protection threshold of the instance. Unit: Gbit/s.
	//
	// example:
	//
	// 20
	BaseThreshold *int32 `json:"BaseThreshold,omitempty" xml:"BaseThreshold,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	DdosbgpInstanceId *string `json:"DdosbgpInstanceId,omitempty" xml:"DdosbgpInstanceId,omitempty"`
	// The burstable protection threshold of the instance. Unit: Gbit/s.
	//
	// example:
	//
	// 301
	ElasticThreshold *int32 `json:"ElasticThreshold,omitempty" xml:"ElasticThreshold,omitempty"`
	// The expiration time of the instance. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1640448000000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 118.31.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeBgpPackByIpResponseBodyDdosbgpInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpPackByIpResponseBodyDdosbgpInfo) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) GetBaseThreshold() *int32 {
	return s.BaseThreshold
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) GetDdosbgpInstanceId() *string {
	return s.DdosbgpInstanceId
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) GetElasticThreshold() *int32 {
	return s.ElasticThreshold
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) GetIp() *string {
	return s.Ip
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetBaseThreshold(v int32) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.BaseThreshold = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetDdosbgpInstanceId(v string) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.DdosbgpInstanceId = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetElasticThreshold(v int32) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.ElasticThreshold = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetExpireTime(v int64) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.ExpireTime = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) SetIp(v string) *DescribeBgpPackByIpResponseBodyDdosbgpInfo {
	s.Ip = &v
	return s
}

func (s *DescribeBgpPackByIpResponseBodyDdosbgpInfo) Validate() error {
	return dara.Validate(s)
}
