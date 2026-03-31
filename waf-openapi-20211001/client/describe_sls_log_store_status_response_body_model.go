// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsLogStoreStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExistStatus(v bool) *DescribeSlsLogStoreStatusResponseBody
	GetExistStatus() *bool
	SetRequestId(v string) *DescribeSlsLogStoreStatusResponseBody
	GetRequestId() *string
}

type DescribeSlsLogStoreStatusResponseBody struct {
	// Indicates whether a Logstore is created for WAF. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ExistStatus *bool `json:"ExistStatus,omitempty" xml:"ExistStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 184F538F-C115-5C89-A4EF-C79CD2E29AC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSlsLogStoreStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsLogStoreStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogStoreStatusResponseBody) GetExistStatus() *bool {
	return s.ExistStatus
}

func (s *DescribeSlsLogStoreStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlsLogStoreStatusResponseBody) SetExistStatus(v bool) *DescribeSlsLogStoreStatusResponseBody {
	s.ExistStatus = &v
	return s
}

func (s *DescribeSlsLogStoreStatusResponseBody) SetRequestId(v string) *DescribeSlsLogStoreStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsLogStoreStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
