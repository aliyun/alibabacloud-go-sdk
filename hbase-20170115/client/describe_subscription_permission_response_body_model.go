// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSubscriptionPermissionResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeSubscriptionPermissionResponseBody
	GetStatus() *string
}

type DescribeSubscriptionPermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSubscriptionPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSubscriptionPermissionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeSubscriptionPermissionResponseBody) SetRequestId(v string) *DescribeSubscriptionPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionPermissionResponseBody) SetStatus(v string) *DescribeSubscriptionPermissionResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSubscriptionPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
