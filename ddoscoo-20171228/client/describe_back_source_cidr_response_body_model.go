// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackSourceCidrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCidrList(v []*string) *DescribeBackSourceCidrResponseBody
	GetCidrList() []*string
	SetRequestId(v string) *DescribeBackSourceCidrResponseBody
	GetRequestId() *string
}

type DescribeBackSourceCidrResponseBody struct {
	CidrList []*string `json:"CidrList,omitempty" xml:"CidrList,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackSourceCidrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackSourceCidrResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrResponseBody) GetCidrList() []*string {
	return s.CidrList
}

func (s *DescribeBackSourceCidrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackSourceCidrResponseBody) SetCidrList(v []*string) *DescribeBackSourceCidrResponseBody {
	s.CidrList = v
	return s
}

func (s *DescribeBackSourceCidrResponseBody) SetRequestId(v string) *DescribeBackSourceCidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackSourceCidrResponseBody) Validate() error {
	return dara.Validate(s)
}
