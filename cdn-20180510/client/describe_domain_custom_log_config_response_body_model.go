// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCustomLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DescribeDomainCustomLogConfigResponseBody
	GetConfigId() *string
	SetRemark(v string) *DescribeDomainCustomLogConfigResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescribeDomainCustomLogConfigResponseBody
	GetRequestId() *string
	SetSample(v string) *DescribeDomainCustomLogConfigResponseBody
	GetSample() *string
	SetTag(v string) *DescribeDomainCustomLogConfigResponseBody
	GetTag() *string
}

type DescribeDomainCustomLogConfigResponseBody struct {
	// The ID of the log configuration.
	//
	// example:
	//
	// 123
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The format of the log configuration.
	//
	// example:
	//
	// $time_iso8601_$request_method_$
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 94E3559F-7B6A-4A5E-AFFD-44E2A208A249
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sample log configuration.
	//
	// example:
	//
	// [9/Jun/2015:01:58:09+0800]188.165.15.75-1542\\"-\\"\\"GET http://www.aliyun.com/index.html\\
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// The tag information about the log configuration.
	//
	// example:
	//
	// book
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeDomainCustomLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCustomLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainCustomLogConfigResponseBody) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeDomainCustomLogConfigResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescribeDomainCustomLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainCustomLogConfigResponseBody) GetSample() *string {
	return s.Sample
}

func (s *DescribeDomainCustomLogConfigResponseBody) GetTag() *string {
	return s.Tag
}

func (s *DescribeDomainCustomLogConfigResponseBody) SetConfigId(v string) *DescribeDomainCustomLogConfigResponseBody {
	s.ConfigId = &v
	return s
}

func (s *DescribeDomainCustomLogConfigResponseBody) SetRemark(v string) *DescribeDomainCustomLogConfigResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeDomainCustomLogConfigResponseBody) SetRequestId(v string) *DescribeDomainCustomLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainCustomLogConfigResponseBody) SetSample(v string) *DescribeDomainCustomLogConfigResponseBody {
	s.Sample = &v
	return s
}

func (s *DescribeDomainCustomLogConfigResponseBody) SetTag(v string) *DescribeDomainCustomLogConfigResponseBody {
	s.Tag = &v
	return s
}

func (s *DescribeDomainCustomLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
