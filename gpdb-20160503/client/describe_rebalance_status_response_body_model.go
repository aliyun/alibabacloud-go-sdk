// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRebalanceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRebalanceStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeRebalanceStatusResponseBody
	GetStatus() *string
}

type DescribeRebalanceStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7F5B5023-94EA-5D5D-AB72-B7B356BA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The rebalance status of the instance. Valid values: Balanced and Imbalanced.
	//
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRebalanceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRebalanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRebalanceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRebalanceStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeRebalanceStatusResponseBody) SetRequestId(v string) *DescribeRebalanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRebalanceStatusResponseBody) SetStatus(v string) *DescribeRebalanceStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeRebalanceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
