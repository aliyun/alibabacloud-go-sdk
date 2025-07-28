// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourceAccessedPortsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHttp(v []*int32) *DescribeCloudResourceAccessedPortsResponseBody
	GetHttp() []*int32
	SetHttps(v []*int32) *DescribeCloudResourceAccessedPortsResponseBody
	GetHttps() []*int32
	SetRequestId(v string) *DescribeCloudResourceAccessedPortsResponseBody
	GetRequestId() *string
}

type DescribeCloudResourceAccessedPortsResponseBody struct {
	// The HTTP ports.
	Http []*int32 `json:"Http,omitempty" xml:"Http,omitempty" type:"Repeated"`
	// The HTTPS ports.
	Https []*int32 `json:"Https,omitempty" xml:"Https,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C1823E96-EF4B-5BD2-9E02-1D18****3ED8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudResourceAccessedPortsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceAccessedPortsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceAccessedPortsResponseBody) GetHttp() []*int32 {
	return s.Http
}

func (s *DescribeCloudResourceAccessedPortsResponseBody) GetHttps() []*int32 {
	return s.Https
}

func (s *DescribeCloudResourceAccessedPortsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudResourceAccessedPortsResponseBody) SetHttp(v []*int32) *DescribeCloudResourceAccessedPortsResponseBody {
	s.Http = v
	return s
}

func (s *DescribeCloudResourceAccessedPortsResponseBody) SetHttps(v []*int32) *DescribeCloudResourceAccessedPortsResponseBody {
	s.Https = v
	return s
}

func (s *DescribeCloudResourceAccessedPortsResponseBody) SetRequestId(v string) *DescribeCloudResourceAccessedPortsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudResourceAccessedPortsResponseBody) Validate() error {
	return dara.Validate(s)
}
