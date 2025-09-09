// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecycleBinStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRecycleBinStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeRecycleBinStatusResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DescribeRecycleBinStatusResponseBody
	GetSuccess() *bool
}

type DescribeRecycleBinStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ACB5258F-25AF-4D7C-8FAA-B6FE60******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the table recycle bin. Valid values:
	//
	// 	- disable: The table recycle bin is enabled.
	//
	// 	- enable: The table recycle bin is disabled.
	//
	// example:
	//
	// disable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRecycleBinStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecycleBinStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecycleBinStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeRecycleBinStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRecycleBinStatusResponseBody) SetRequestId(v string) *DescribeRecycleBinStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecycleBinStatusResponseBody) SetStatus(v string) *DescribeRecycleBinStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeRecycleBinStatusResponseBody) SetSuccess(v bool) *DescribeRecycleBinStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRecycleBinStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
