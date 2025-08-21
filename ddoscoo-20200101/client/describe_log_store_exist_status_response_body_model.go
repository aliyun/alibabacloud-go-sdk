// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogStoreExistStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExistStatus(v bool) *DescribeLogStoreExistStatusResponseBody
	GetExistStatus() *bool
	SetRequestId(v string) *DescribeLogStoreExistStatusResponseBody
	GetRequestId() *string
}

type DescribeLogStoreExistStatusResponseBody struct {
	// Indicates whether a Logstore is created for Anti-DDoS Pro or Anti-DDoS Premium. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	ExistStatus *bool `json:"ExistStatus,omitempty" xml:"ExistStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLogStoreExistStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreExistStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreExistStatusResponseBody) GetExistStatus() *bool {
	return s.ExistStatus
}

func (s *DescribeLogStoreExistStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogStoreExistStatusResponseBody) SetExistStatus(v bool) *DescribeLogStoreExistStatusResponseBody {
	s.ExistStatus = &v
	return s
}

func (s *DescribeLogStoreExistStatusResponseBody) SetRequestId(v string) *DescribeLogStoreExistStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogStoreExistStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
