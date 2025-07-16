// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMirrorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRatio(v string) *DescribeServiceMirrorResponseBody
	GetRatio() *string
	SetRequestId(v string) *DescribeServiceMirrorResponseBody
	GetRequestId() *string
	SetServiceName(v string) *DescribeServiceMirrorResponseBody
	GetServiceName() *string
	SetTarget(v string) *DescribeServiceMirrorResponseBody
	GetTarget() *string
}

type DescribeServiceMirrorResponseBody struct {
	// The percentage of traffic that you want to mirror. Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service name.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The destination services to which you want to mirror traffic.
	//
	// example:
	//
	// foo2,foo3
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s DescribeServiceMirrorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMirrorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMirrorResponseBody) GetRatio() *string {
	return s.Ratio
}

func (s *DescribeServiceMirrorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceMirrorResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeServiceMirrorResponseBody) GetTarget() *string {
	return s.Target
}

func (s *DescribeServiceMirrorResponseBody) SetRatio(v string) *DescribeServiceMirrorResponseBody {
	s.Ratio = &v
	return s
}

func (s *DescribeServiceMirrorResponseBody) SetRequestId(v string) *DescribeServiceMirrorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMirrorResponseBody) SetServiceName(v string) *DescribeServiceMirrorResponseBody {
	s.ServiceName = &v
	return s
}

func (s *DescribeServiceMirrorResponseBody) SetTarget(v string) *DescribeServiceMirrorResponseBody {
	s.Target = &v
	return s
}

func (s *DescribeServiceMirrorResponseBody) Validate() error {
	return dara.Validate(s)
}
