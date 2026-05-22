// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKvAccountStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeKvAccountStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeKvAccountStatusResponseBody
	GetStatus() *string
}

type DescribeKvAccountStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether Edge KV is activated for the Alibaba Cloud account.
	//
	// 	- **online**
	//
	// 	- **offline**
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeKvAccountStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKvAccountStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKvAccountStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKvAccountStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeKvAccountStatusResponseBody) SetRequestId(v string) *DescribeKvAccountStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKvAccountStatusResponseBody) SetStatus(v string) *DescribeKvAccountStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeKvAccountStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
