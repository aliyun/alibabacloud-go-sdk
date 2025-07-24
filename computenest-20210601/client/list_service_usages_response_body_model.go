// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceUsagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServiceUsagesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceUsagesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceUsagesResponseBody
	GetRequestId() *string
	SetServiceUsages(v []*ListServiceUsagesResponseBodyServiceUsages) *ListServiceUsagesResponseBody
	GetServiceUsages() []*ListServiceUsagesResponseBodyServiceUsages
	SetTotalCount(v int32) *ListServiceUsagesResponseBody
	GetTotalCount() *int32
}

type ListServiceUsagesResponseBody struct {
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAYChudnQUoBH+mGWFpb6oP0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 18AD0960-A9FE-1AC8-ADF8-22131Fxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service applications.
	ServiceUsages []*ListServiceUsagesResponseBodyServiceUsages `json:"ServiceUsages,omitempty" xml:"ServiceUsages,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceUsagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceUsagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceUsagesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceUsagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceUsagesResponseBody) GetServiceUsages() []*ListServiceUsagesResponseBodyServiceUsages {
	return s.ServiceUsages
}

func (s *ListServiceUsagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListServiceUsagesResponseBody) SetMaxResults(v int32) *ListServiceUsagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetNextToken(v string) *ListServiceUsagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetRequestId(v string) *ListServiceUsagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetServiceUsages(v []*ListServiceUsagesResponseBodyServiceUsages) *ListServiceUsagesResponseBody {
	s.ServiceUsages = v
	return s
}

func (s *ListServiceUsagesResponseBody) SetTotalCount(v int32) *ListServiceUsagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceUsagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListServiceUsagesResponseBodyServiceUsages struct {
	// The review comment.
	//
	// example:
	//
	// Approved
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2022-05-25T02:02:02Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-c9f36ec6d19b4exxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// LobelChat
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The state of the service application. Valid values:
	//
	// 	- Submitted: The application is submitted for review.
	//
	// 	- Approved: The application is approved.
	//
	// 	- Rejected: The application is rejected.
	//
	// 	- Canceled: The application is canceled.
	//
	// example:
	//
	// Submitted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// TestSupplier
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The time when the application was updated.
	//
	// example:
	//
	// 2022-05-25T02:02:02Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 127383705958xxxx
	UserAliUid *int64 `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
	// The information about the applicants.
	UserInformation map[string]*string `json:"UserInformation,omitempty" xml:"UserInformation,omitempty"`
}

func (s ListServiceUsagesResponseBodyServiceUsages) String() string {
	return dara.Prettify(s)
}

func (s ListServiceUsagesResponseBodyServiceUsages) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponseBodyServiceUsages) GetComments() *string {
	return s.Comments
}

func (s *ListServiceUsagesResponseBodyServiceUsages) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListServiceUsagesResponseBodyServiceUsages) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServiceUsagesResponseBodyServiceUsages) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListServiceUsagesResponseBodyServiceUsages) GetStatus() *string {
	return s.Status
}

func (s *ListServiceUsagesResponseBodyServiceUsages) GetSupplierName() *string {
	return s.SupplierName
}

func (s *ListServiceUsagesResponseBodyServiceUsages) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListServiceUsagesResponseBodyServiceUsages) GetUserAliUid() *int64 {
	return s.UserAliUid
}

func (s *ListServiceUsagesResponseBodyServiceUsages) GetUserInformation() map[string]*string {
	return s.UserInformation
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetComments(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.Comments = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetCreateTime(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.CreateTime = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetServiceId(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.ServiceId = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetServiceName(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.ServiceName = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetStatus(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.Status = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetSupplierName(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.SupplierName = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUpdateTime(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.UpdateTime = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUserAliUid(v int64) *ListServiceUsagesResponseBodyServiceUsages {
	s.UserAliUid = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUserInformation(v map[string]*string) *ListServiceUsagesResponseBodyServiceUsages {
	s.UserInformation = v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) Validate() error {
	return dara.Validate(s)
}
