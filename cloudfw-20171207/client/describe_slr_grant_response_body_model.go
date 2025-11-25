// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlrGrantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsGranted(v bool) *DescribeSlrGrantResponseBody
	GetIsGranted() *bool
	SetRequestId(v string) *DescribeSlrGrantResponseBody
	GetRequestId() *string
	SetUserType(v string) *DescribeSlrGrantResponseBody
	GetUserType() *string
}

type DescribeSlrGrantResponseBody struct {
	// example:
	//
	// true
	IsGranted *bool `json:"IsGranted,omitempty" xml:"IsGranted,omitempty"`
	// example:
	//
	// 1BD3D277-AE2F-5609-893F-FF7A72A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// sub
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeSlrGrantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlrGrantResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlrGrantResponseBody) GetIsGranted() *bool {
	return s.IsGranted
}

func (s *DescribeSlrGrantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlrGrantResponseBody) GetUserType() *string {
	return s.UserType
}

func (s *DescribeSlrGrantResponseBody) SetIsGranted(v bool) *DescribeSlrGrantResponseBody {
	s.IsGranted = &v
	return s
}

func (s *DescribeSlrGrantResponseBody) SetRequestId(v string) *DescribeSlrGrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlrGrantResponseBody) SetUserType(v string) *DescribeSlrGrantResponseBody {
	s.UserType = &v
	return s
}

func (s *DescribeSlrGrantResponseBody) Validate() error {
	return dara.Validate(s)
}
