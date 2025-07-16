// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreloadDetailByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePreloadDetailByIdResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePreloadDetailByIdResponseBody
	GetTotalCount() *int64
	SetUrlDetails(v []*DescribePreloadDetailByIdResponseBodyUrlDetails) *DescribePreloadDetailByIdResponseBody
	GetUrlDetails() []*DescribePreloadDetailByIdResponseBodyUrlDetails
}

type DescribePreloadDetailByIdResponseBody struct {
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// E9D3257A-1B7C-414C-90C1-8D07AC47BCAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of queried tasks.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The details of the task, including the task ID, start time, end time, domain name, success rate, status, returned error code, and completion details of all URL resources.
	UrlDetails []*DescribePreloadDetailByIdResponseBodyUrlDetails `json:"UrlDetails,omitempty" xml:"UrlDetails,omitempty" type:"Repeated"`
}

func (s DescribePreloadDetailByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePreloadDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePreloadDetailByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePreloadDetailByIdResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePreloadDetailByIdResponseBody) GetUrlDetails() []*DescribePreloadDetailByIdResponseBodyUrlDetails {
	return s.UrlDetails
}

func (s *DescribePreloadDetailByIdResponseBody) SetRequestId(v string) *DescribePreloadDetailByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePreloadDetailByIdResponseBody) SetTotalCount(v int64) *DescribePreloadDetailByIdResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePreloadDetailByIdResponseBody) SetUrlDetails(v []*DescribePreloadDetailByIdResponseBodyUrlDetails) *DescribePreloadDetailByIdResponseBody {
	s.UrlDetails = v
	return s
}

func (s *DescribePreloadDetailByIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePreloadDetailByIdResponseBodyUrlDetails struct {
	// The time when the task was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-11-23T02:26:56Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The domain name for prefetching resources.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The time when the task ended. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-11-23T02:27:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The progress of the prefetch task, which indicates the number of points of presence (POPs) on which the prefetch task is completed.
	//
	// example:
	//
	// 100%
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// The turned error code. A value of `0` indicates that the task succeeded.
	//
	// example:
	//
	// 0
	RetCode *string `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **Complete**
	//
	// 	- **Refreshing**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task that you want to query.
	//
	// You can call the PushObjectCache operation to query task IDs. Then, you can use the task IDs to query task status.
	//
	// You can query one task ID at a time.
	//
	// example:
	//
	// 14286878547
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The completion details of all URL resources in the task.
	Urls []*DescribePreloadDetailByIdResponseBodyUrlDetailsUrls `json:"Urls,omitempty" xml:"Urls,omitempty" type:"Repeated"`
}

func (s DescribePreloadDetailByIdResponseBodyUrlDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribePreloadDetailByIdResponseBodyUrlDetails) GoString() string {
	return s.String()
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) GetDomain() *string {
	return s.Domain
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) GetProcess() *string {
	return s.Process
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) GetRetCode() *string {
	return s.RetCode
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) GetStatus() *string {
	return s.Status
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) GetUrls() []*DescribePreloadDetailByIdResponseBodyUrlDetailsUrls {
	return s.Urls
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) SetCreationTime(v string) *DescribePreloadDetailByIdResponseBodyUrlDetails {
	s.CreationTime = &v
	return s
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) SetDomain(v string) *DescribePreloadDetailByIdResponseBodyUrlDetails {
	s.Domain = &v
	return s
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) SetEndTime(v string) *DescribePreloadDetailByIdResponseBodyUrlDetails {
	s.EndTime = &v
	return s
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) SetProcess(v string) *DescribePreloadDetailByIdResponseBodyUrlDetails {
	s.Process = &v
	return s
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) SetRetCode(v string) *DescribePreloadDetailByIdResponseBodyUrlDetails {
	s.RetCode = &v
	return s
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) SetStatus(v string) *DescribePreloadDetailByIdResponseBodyUrlDetails {
	s.Status = &v
	return s
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) SetTaskId(v string) *DescribePreloadDetailByIdResponseBodyUrlDetails {
	s.TaskId = &v
	return s
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) SetUrls(v []*DescribePreloadDetailByIdResponseBodyUrlDetailsUrls) *DescribePreloadDetailByIdResponseBodyUrlDetails {
	s.Urls = v
	return s
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetails) Validate() error {
	return dara.Validate(s)
}

type DescribePreloadDetailByIdResponseBodyUrlDetailsUrls struct {
	// The details of resource prefetch.
	//
	// 	- If the resource is prefetched on all POPs, "Successfully preloaded" is returned.
	//
	// 	- If the resource fails to be prefetched on some POPs, the failure details separated by vertical bars (|) are returned.
	//
	// example:
	//
	// Successfully preloaded
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The success percentage, which indicates the number of POPs on which the resource is prefetched.
	//
	// example:
	//
	// 47%
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The URL of the prefetched resource.
	//
	// example:
	//
	// /abc.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribePreloadDetailByIdResponseBodyUrlDetailsUrls) String() string {
	return dara.Prettify(s)
}

func (s DescribePreloadDetailByIdResponseBodyUrlDetailsUrls) GoString() string {
	return s.String()
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetailsUrls) GetDescription() *string {
	return s.Description
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetailsUrls) GetSuccess() *string {
	return s.Success
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetailsUrls) GetUrl() *string {
	return s.Url
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetailsUrls) SetDescription(v string) *DescribePreloadDetailByIdResponseBodyUrlDetailsUrls {
	s.Description = &v
	return s
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetailsUrls) SetSuccess(v string) *DescribePreloadDetailByIdResponseBodyUrlDetailsUrls {
	s.Success = &v
	return s
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetailsUrls) SetUrl(v string) *DescribePreloadDetailByIdResponseBodyUrlDetailsUrls {
	s.Url = &v
	return s
}

func (s *DescribePreloadDetailByIdResponseBodyUrlDetailsUrls) Validate() error {
	return dara.Validate(s)
}
