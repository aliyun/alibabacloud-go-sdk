// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudRecordStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCloudRecordStatusResponseBody
	GetRequestId() *string
	SetStatus(v int32) *DescribeCloudRecordStatusResponseBody
	GetStatus() *int32
}

type DescribeCloudRecordStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCloudRecordStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudRecordStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudRecordStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudRecordStatusResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCloudRecordStatusResponseBody) SetRequestId(v string) *DescribeCloudRecordStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudRecordStatusResponseBody) SetStatus(v int32) *DescribeCloudRecordStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeCloudRecordStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
