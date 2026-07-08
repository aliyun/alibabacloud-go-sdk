// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcePortResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeResourcePortResponseBody
	GetRequestId() *string
	SetResourcePorts(v []*string) *DescribeResourcePortResponseBody
	GetResourcePorts() []*string
}

type DescribeResourcePortResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 66A98669-CC6E-4F3E-80A6-3014697B11AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// All HTTP and HTTPS listener ports that are added to WAF.
	ResourcePorts []*string `json:"ResourcePorts,omitempty" xml:"ResourcePorts,omitempty" type:"Repeated"`
}

func (s DescribeResourcePortResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePortResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcePortResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourcePortResponseBody) GetResourcePorts() []*string {
	return s.ResourcePorts
}

func (s *DescribeResourcePortResponseBody) SetRequestId(v string) *DescribeResourcePortResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePortResponseBody) SetResourcePorts(v []*string) *DescribeResourcePortResponseBody {
	s.ResourcePorts = v
	return s
}

func (s *DescribeResourcePortResponseBody) Validate() error {
	return dara.Validate(s)
}
