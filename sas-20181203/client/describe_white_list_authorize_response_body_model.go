// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListAuthorizeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableAuthorizeNum(v int32) *DescribeWhiteListAuthorizeResponseBody
	GetAvailableAuthorizeNum() *int32
	SetRequestId(v string) *DescribeWhiteListAuthorizeResponseBody
	GetRequestId() *string
}

type DescribeWhiteListAuthorizeResponseBody struct {
	// The available quota.
	//
	// example:
	//
	// 3
	AvailableAuthorizeNum *int32 `json:"AvailableAuthorizeNum,omitempty" xml:"AvailableAuthorizeNum,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWhiteListAuthorizeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListAuthorizeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListAuthorizeResponseBody) GetAvailableAuthorizeNum() *int32 {
	return s.AvailableAuthorizeNum
}

func (s *DescribeWhiteListAuthorizeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWhiteListAuthorizeResponseBody) SetAvailableAuthorizeNum(v int32) *DescribeWhiteListAuthorizeResponseBody {
	s.AvailableAuthorizeNum = &v
	return s
}

func (s *DescribeWhiteListAuthorizeResponseBody) SetRequestId(v string) *DescribeWhiteListAuthorizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhiteListAuthorizeResponseBody) Validate() error {
	return dara.Validate(s)
}
