// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsynJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v []*ListAsynJobsResponseBodyJobs) *ListAsynJobsResponseBody
	GetJobs() []*ListAsynJobsResponseBodyJobs
	SetMaxResults(v int64) *ListAsynJobsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListAsynJobsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAsynJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAsynJobsResponseBody
	GetTotalCount() *int64
}

type ListAsynJobsResponseBody struct {
	// The tasks.
	Jobs []*ListAsynJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAsynJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAsynJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsynJobsResponseBody) GetJobs() []*ListAsynJobsResponseBodyJobs {
	return s.Jobs
}

func (s *ListAsynJobsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListAsynJobsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAsynJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAsynJobsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAsynJobsResponseBody) SetJobs(v []*ListAsynJobsResponseBodyJobs) *ListAsynJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListAsynJobsResponseBody) SetMaxResults(v int64) *ListAsynJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAsynJobsResponseBody) SetNextToken(v string) *ListAsynJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAsynJobsResponseBody) SetRequestId(v string) *ListAsynJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsynJobsResponseBody) SetTotalCount(v int64) *ListAsynJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAsynJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAsynJobsResponseBodyJobs struct {
	// The name of the operation.
	//
	// example:
	//
	// CreateLoadBalancer
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The timestamp that indicates the start time of the task. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 2134663231234
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// If the value of **Status*	- is Failed, an error code is returned.
	//
	// example:
	//
	// 506
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// If the value of **Status*	- is Failed, an error message is returned.
	//
	// example:
	//
	// AllocateEipAddress Failed
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B5****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The timestamp that indicates the end time of the task. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 2144663233315
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The type of the operation. Valid values:
	//
	// 	- **Create**
	//
	// 	- **Update**
	//
	// 	- **Delete**
	//
	// example:
	//
	// Create
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The associated resource ID.
	//
	// example:
	//
	// alb-o8mszt95oamfjy****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the associated resource. Valid values:
	//
	// 	- **loadbalancer**: an ALB instance
	//
	// 	- **listener**: a listener
	//
	// 	- **rule**: a forwarding rule
	//
	// 	- **acl**: an ACL
	//
	// 	- **securitypolicy**: a security policy
	//
	// 	- **servergroup**: a server group
	//
	// example:
	//
	// acl
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **Succeeded**
	//
	// 	- **Failed**
	//
	// 	- **Processing**
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAsynJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListAsynJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListAsynJobsResponseBodyJobs) GetApiName() *string {
	return s.ApiName
}

func (s *ListAsynJobsResponseBodyJobs) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAsynJobsResponseBodyJobs) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAsynJobsResponseBodyJobs) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAsynJobsResponseBodyJobs) GetId() *string {
	return s.Id
}

func (s *ListAsynJobsResponseBodyJobs) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListAsynJobsResponseBodyJobs) GetOperateType() *string {
	return s.OperateType
}

func (s *ListAsynJobsResponseBodyJobs) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListAsynJobsResponseBodyJobs) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAsynJobsResponseBodyJobs) GetStatus() *string {
	return s.Status
}

func (s *ListAsynJobsResponseBodyJobs) SetApiName(v string) *ListAsynJobsResponseBodyJobs {
	s.ApiName = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetCreateTime(v int64) *ListAsynJobsResponseBodyJobs {
	s.CreateTime = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetErrorCode(v string) *ListAsynJobsResponseBodyJobs {
	s.ErrorCode = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetErrorMessage(v string) *ListAsynJobsResponseBodyJobs {
	s.ErrorMessage = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetId(v string) *ListAsynJobsResponseBodyJobs {
	s.Id = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetModifyTime(v int64) *ListAsynJobsResponseBodyJobs {
	s.ModifyTime = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetOperateType(v string) *ListAsynJobsResponseBodyJobs {
	s.OperateType = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetResourceId(v string) *ListAsynJobsResponseBodyJobs {
	s.ResourceId = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetResourceType(v string) *ListAsynJobsResponseBodyJobs {
	s.ResourceType = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetStatus(v string) *ListAsynJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) Validate() error {
	return dara.Validate(s)
}
