// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCanTrySasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanTry(v int32) *DescribeCanTrySasResponseBody
	GetCanTry() *int32
	SetRequestId(v string) *DescribeCanTrySasResponseBody
	GetRequestId() *string
}

type DescribeCanTrySasResponseBody struct {
	// Indicates whether you have the permissions on the trial use of Security Center. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 0
	CanTry *int32 `json:"CanTry,omitempty" xml:"CanTry,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E90DE229-9FC6-58F6-BF4B-03AD6179****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCanTrySasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCanTrySasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCanTrySasResponseBody) GetCanTry() *int32 {
	return s.CanTry
}

func (s *DescribeCanTrySasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCanTrySasResponseBody) SetCanTry(v int32) *DescribeCanTrySasResponseBody {
	s.CanTry = &v
	return s
}

func (s *DescribeCanTrySasResponseBody) SetRequestId(v string) *DescribeCanTrySasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCanTrySasResponseBody) Validate() error {
	return dara.Validate(s)
}
