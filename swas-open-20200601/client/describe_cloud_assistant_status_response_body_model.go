// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudAssistantStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAssistantStatus(v []*DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) *DescribeCloudAssistantStatusResponseBody
	GetCloudAssistantStatus() []*DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus
	SetPageNumber(v int32) *DescribeCloudAssistantStatusResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCloudAssistantStatusResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCloudAssistantStatusResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCloudAssistantStatusResponseBody
	GetTotalCount() *int32
}

type DescribeCloudAssistantStatusResponseBody struct {
	// Indicates whether the Cloud Assistant client is installed on the server.
	CloudAssistantStatus []*DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus `json:"CloudAssistantStatus,omitempty" xml:"CloudAssistantStatus,omitempty" type:"Repeated"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudAssistantStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusResponseBody) GetCloudAssistantStatus() []*DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus {
	return s.CloudAssistantStatus
}

func (s *DescribeCloudAssistantStatusResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCloudAssistantStatusResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudAssistantStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudAssistantStatusResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCloudAssistantStatusResponseBody) SetCloudAssistantStatus(v []*DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) *DescribeCloudAssistantStatusResponseBody {
	s.CloudAssistantStatus = v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetPageNumber(v int32) *DescribeCloudAssistantStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetPageSize(v int32) *DescribeCloudAssistantStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetRequestId(v string) *DescribeCloudAssistantStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetTotalCount(v int32) *DescribeCloudAssistantStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) Validate() error {
	if s.CloudAssistantStatus != nil {
		for _, item := range s.CloudAssistantStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus struct {
	// The ID of the simple application server.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the Cloud Assistant client is installed on the server.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) GetStatus() *bool {
	return s.Status
}

func (s *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) SetInstanceId(v string) *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) SetStatus(v bool) *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus {
	s.Status = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) Validate() error {
	return dara.Validate(s)
}
