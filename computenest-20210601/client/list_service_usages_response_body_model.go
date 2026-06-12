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
	// The number of entries returned on each page. The maximum value is 100. The default value is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It is returned when the results are not complete. To retrieve the next page of results, include this token in the NextToken parameter of the next request.
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
	// The service usage details.
	ServiceUsages []*ListServiceUsagesResponseBodyServiceUsages `json:"ServiceUsages,omitempty" xml:"ServiceUsages,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 1
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
	if s.ServiceUsages != nil {
		for _, item := range s.ServiceUsages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceUsagesResponseBodyServiceUsages struct {
	// The remarks of the approval.
	//
	// example:
	//
	// Review approved.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The time when the request was created.
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
	// Deployment link permission request
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service status. Valid values:
	//
	// - Submitted: The request is submitted and pending approval.
	//
	// - Approved: The request is approved.
	//
	// - Rejected: The request is rejected.
	//
	// - Canceled: The request is canceled.
	//
	// example:
	//
	// Submitted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the supplier.
	//
	// example:
	//
	// Acceptance test
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The time when the request was last updated.
	//
	// example:
	//
	// 2022-05-25T02:02:02Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The Alibaba Cloud account ID of the user.
	//
	// example:
	//
	// 127383705958xxxx
	UserAliUid *int64 `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
	// The user information.
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
