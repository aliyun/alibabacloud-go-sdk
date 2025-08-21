// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackSourceCidrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCidrs(v []*string) *DescribeBackSourceCidrResponseBody
	GetCidrs() []*string
	SetRequestId(v string) *DescribeBackSourceCidrResponseBody
	GetRequestId() *string
}

type DescribeBackSourceCidrResponseBody struct {
	// An array that consists of the back-to-origin CIDR blocks of the instance.
	Cidrs []*string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackSourceCidrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackSourceCidrResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrResponseBody) GetCidrs() []*string {
	return s.Cidrs
}

func (s *DescribeBackSourceCidrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackSourceCidrResponseBody) SetCidrs(v []*string) *DescribeBackSourceCidrResponseBody {
	s.Cidrs = v
	return s
}

func (s *DescribeBackSourceCidrResponseBody) SetRequestId(v string) *DescribeBackSourceCidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackSourceCidrResponseBody) Validate() error {
	return dara.Validate(s)
}
