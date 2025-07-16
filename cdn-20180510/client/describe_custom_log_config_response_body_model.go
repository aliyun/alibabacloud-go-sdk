// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemark(v string) *DescribeCustomLogConfigResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescribeCustomLogConfigResponseBody
	GetRequestId() *string
	SetSample(v string) *DescribeCustomLogConfigResponseBody
	GetSample() *string
	SetTag(v string) *DescribeCustomLogConfigResponseBody
	GetTag() *string
}

type DescribeCustomLogConfigResponseBody struct {
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
	// The sample.
	//
	// example:
	//
	// "[9/Jun/2015:01:58:09 +0800] 188.165.15.75 - 1542 \\"-\\" \\"GEThttp: //www.aliyun.com/index.html\\" 200
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// The tag information about the log configuration.
	//
	// example:
	//
	// img1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeCustomLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomLogConfigResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCustomLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomLogConfigResponseBody) GetSample() *string {
	return s.Sample
}

func (s *DescribeCustomLogConfigResponseBody) GetTag() *string {
	return s.Tag
}

func (s *DescribeCustomLogConfigResponseBody) SetRemark(v string) *DescribeCustomLogConfigResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeCustomLogConfigResponseBody) SetRequestId(v string) *DescribeCustomLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomLogConfigResponseBody) SetSample(v string) *DescribeCustomLogConfigResponseBody {
	s.Sample = &v
	return s
}

func (s *DescribeCustomLogConfigResponseBody) SetTag(v string) *DescribeCustomLogConfigResponseBody {
	s.Tag = &v
	return s
}

func (s *DescribeCustomLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
