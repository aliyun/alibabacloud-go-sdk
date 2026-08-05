// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
}

type DescribeRegionsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// AE3554F3-313D-5712-919C-A334615C5DE5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}
