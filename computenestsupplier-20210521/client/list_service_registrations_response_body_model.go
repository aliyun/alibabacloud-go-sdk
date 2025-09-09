// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceRegistrationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServiceRegistrationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceRegistrationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceRegistrationsResponseBody
	GetRequestId() *string
	SetServiceRegistrations(v []*ListServiceRegistrationsResponseBodyServiceRegistrations) *ListServiceRegistrationsResponseBody
	GetServiceRegistrations() []*ListServiceRegistrationsResponseBodyServiceRegistrations
	SetTotalCount(v int32) *ListServiceRegistrationsResponseBody
	GetTotalCount() *int32
}

type ListServiceRegistrationsResponseBody struct {
	// Number of items per page in a paginated query. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Service registration information.
	ServiceRegistrations []*ListServiceRegistrationsResponseBodyServiceRegistrations `json:"ServiceRegistrations,omitempty" xml:"ServiceRegistrations,omitempty" type:"Repeated"`
	// Total number of records that meet the criteria.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceRegistrationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceRegistrationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceRegistrationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceRegistrationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceRegistrationsResponseBody) GetServiceRegistrations() []*ListServiceRegistrationsResponseBodyServiceRegistrations {
	return s.ServiceRegistrations
}

func (s *ListServiceRegistrationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListServiceRegistrationsResponseBody) SetMaxResults(v int32) *ListServiceRegistrationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceRegistrationsResponseBody) SetNextToken(v string) *ListServiceRegistrationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceRegistrationsResponseBody) SetRequestId(v string) *ListServiceRegistrationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceRegistrationsResponseBody) SetServiceRegistrations(v []*ListServiceRegistrationsResponseBodyServiceRegistrations) *ListServiceRegistrationsResponseBody {
	s.ServiceRegistrations = v
	return s
}

func (s *ListServiceRegistrationsResponseBody) SetTotalCount(v int32) *ListServiceRegistrationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceRegistrationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListServiceRegistrationsResponseBodyServiceRegistrations struct {
	// Comment.
	//
	// example:
	//
	// some info is missing
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Finish time.
	//
	// example:
	//
	// 2021-05-23T00:00:00Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// Registration ID.
	//
	// example:
	//
	// sr-001
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// Service ID.
	//
	// example:
	//
	// service-f4c0026a254bxxxxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Registration status. Allowed values:
	//
	// - Submitted
	//
	// - Approved
	//
	// - Rejected
	//
	// - Canceled
	//
	// - Executed
	//
	// - Executed: Executed.
	//
	// example:
	//
	// Rejected
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Submit time.
	//
	// example:
	//
	// 2021-05-22T00:00:00Z
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s ListServiceRegistrationsResponseBodyServiceRegistrations) String() string {
	return dara.Prettify(s)
}

func (s ListServiceRegistrationsResponseBodyServiceRegistrations) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) GetComment() *string {
	return s.Comment
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) GetRegistrationId() *string {
	return s.RegistrationId
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) GetStatus() *string {
	return s.Status
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetComment(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.Comment = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetFinishTime(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.FinishTime = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetRegistrationId(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.RegistrationId = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetServiceId(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.ServiceId = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetStatus(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.Status = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) SetSubmitTime(v string) *ListServiceRegistrationsResponseBodyServiceRegistrations {
	s.SubmitTime = &v
	return s
}

func (s *ListServiceRegistrationsResponseBodyServiceRegistrations) Validate() error {
	return dara.Validate(s)
}
