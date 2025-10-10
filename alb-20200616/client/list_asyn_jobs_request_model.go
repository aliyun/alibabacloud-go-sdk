// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsynJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *ListAsynJobsRequest
	GetApiName() *string
	SetBeginTime(v int64) *ListAsynJobsRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *ListAsynJobsRequest
	GetEndTime() *int64
	SetJobIds(v []*string) *ListAsynJobsRequest
	GetJobIds() []*string
	SetMaxResults(v int64) *ListAsynJobsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListAsynJobsRequest
	GetNextToken() *string
	SetResourceIds(v []*string) *ListAsynJobsRequest
	GetResourceIds() []*string
	SetResourceType(v string) *ListAsynJobsRequest
	GetResourceType() *string
}

type ListAsynJobsRequest struct {
	// The name of the operation.
	//
	// example:
	//
	// CreateLoadBalancer
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The timestamp that indicates the start time of the task. Unit: milliseconds.
	//
	// Specify the timestamp in the Unix format to indicate the total amount of time that is from 00:00:00 (UTC+0) on January 1, 1970 to when the status of the asynchronous task is queried.
	//
	// example:
	//
	// 2021-06-03T17:22Z
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The timestamp that indicates the end time of the task. Unit: milliseconds.
	//
	// Specify the timestamp in the Unix format to indicate the total amount of time that is from 00:00:00 (UTC+0) on January 1, 1970 to when the status of the asynchronous task is returned.
	//
	// example:
	//
	// 2021-06-04T17:22Z
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The asynchronous task IDs.
	JobIds []*string `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource IDs.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The type of the associated resource. Valid values:
	//
	// 	- **loadbalancer**: an Application Load Balancer (ALB) instance
	//
	// 	- **listener**: a listener
	//
	// 	- **rule**: a forwarding rule
	//
	// 	- **acl**: an access control list (ACL)
	//
	// 	- **securitypolicy**: a security policy
	//
	// 	- **servergroup**: a server group
	//
	// example:
	//
	// acl
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListAsynJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAsynJobsRequest) GoString() string {
	return s.String()
}

func (s *ListAsynJobsRequest) GetApiName() *string {
	return s.ApiName
}

func (s *ListAsynJobsRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListAsynJobsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAsynJobsRequest) GetJobIds() []*string {
	return s.JobIds
}

func (s *ListAsynJobsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListAsynJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAsynJobsRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *ListAsynJobsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAsynJobsRequest) SetApiName(v string) *ListAsynJobsRequest {
	s.ApiName = &v
	return s
}

func (s *ListAsynJobsRequest) SetBeginTime(v int64) *ListAsynJobsRequest {
	s.BeginTime = &v
	return s
}

func (s *ListAsynJobsRequest) SetEndTime(v int64) *ListAsynJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAsynJobsRequest) SetJobIds(v []*string) *ListAsynJobsRequest {
	s.JobIds = v
	return s
}

func (s *ListAsynJobsRequest) SetMaxResults(v int64) *ListAsynJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAsynJobsRequest) SetNextToken(v string) *ListAsynJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAsynJobsRequest) SetResourceIds(v []*string) *ListAsynJobsRequest {
	s.ResourceIds = v
	return s
}

func (s *ListAsynJobsRequest) SetResourceType(v string) *ListAsynJobsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListAsynJobsRequest) Validate() error {
	return dara.Validate(s)
}
