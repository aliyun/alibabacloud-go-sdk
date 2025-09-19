// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoDelConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDays(v int32) *DescribeAutoDelConfigResponseBody
	GetDays() *int32
	SetRequestId(v string) *DescribeAutoDelConfigResponseBody
	GetRequestId() *string
}

type DescribeAutoDelConfigResponseBody struct {
	// The number of days during which a detected vulnerability is retained before the vulnerability is automatically deleted.
	//
	// example:
	//
	// 30
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C56F66FD-C4EE-4813-ABDC-4FF94B6C384E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAutoDelConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoDelConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoDelConfigResponseBody) GetDays() *int32 {
	return s.Days
}

func (s *DescribeAutoDelConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoDelConfigResponseBody) SetDays(v int32) *DescribeAutoDelConfigResponseBody {
	s.Days = &v
	return s
}

func (s *DescribeAutoDelConfigResponseBody) SetRequestId(v string) *DescribeAutoDelConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoDelConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
