// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForNetworkZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListApplicationsForNetworkZoneResponseBodyApplications) *ListApplicationsForNetworkZoneResponseBody
	GetApplications() []*ListApplicationsForNetworkZoneResponseBodyApplications
	SetMaxResults(v int32) *ListApplicationsForNetworkZoneResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationsForNetworkZoneResponseBody
	GetNextToken() *string
	SetPreviousToken(v string) *ListApplicationsForNetworkZoneResponseBody
	GetPreviousToken() *string
	SetRequestId(v string) *ListApplicationsForNetworkZoneResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationsForNetworkZoneResponseBody
	GetTotalCount() *int64
}

type ListApplicationsForNetworkZoneResponseBody struct {
	// The list of applications.
	Applications []*ListApplicationsForNetworkZoneResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token returned for the next query.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The token returned for the previous query.
	//
	// example:
	//
	// PTxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationsForNetworkZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForNetworkZoneResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsForNetworkZoneResponseBody) GetApplications() []*ListApplicationsForNetworkZoneResponseBodyApplications {
	return s.Applications
}

func (s *ListApplicationsForNetworkZoneResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationsForNetworkZoneResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsForNetworkZoneResponseBody) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListApplicationsForNetworkZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsForNetworkZoneResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationsForNetworkZoneResponseBody) SetApplications(v []*ListApplicationsForNetworkZoneResponseBodyApplications) *ListApplicationsForNetworkZoneResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsForNetworkZoneResponseBody) SetMaxResults(v int32) *ListApplicationsForNetworkZoneResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsForNetworkZoneResponseBody) SetNextToken(v string) *ListApplicationsForNetworkZoneResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsForNetworkZoneResponseBody) SetPreviousToken(v string) *ListApplicationsForNetworkZoneResponseBody {
	s.PreviousToken = &v
	return s
}

func (s *ListApplicationsForNetworkZoneResponseBody) SetRequestId(v string) *ListApplicationsForNetworkZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsForNetworkZoneResponseBody) SetTotalCount(v int64) *ListApplicationsForNetworkZoneResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationsForNetworkZoneResponseBody) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsForNetworkZoneResponseBodyApplications struct {
	// The ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// cloudSSO
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListApplicationsForNetworkZoneResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForNetworkZoneResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsForNetworkZoneResponseBodyApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationsForNetworkZoneResponseBodyApplications) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListApplicationsForNetworkZoneResponseBodyApplications) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationsForNetworkZoneResponseBodyApplications) SetApplicationId(v string) *ListApplicationsForNetworkZoneResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsForNetworkZoneResponseBodyApplications) SetApplicationName(v string) *ListApplicationsForNetworkZoneResponseBodyApplications {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationsForNetworkZoneResponseBodyApplications) SetInstanceId(v string) *ListApplicationsForNetworkZoneResponseBodyApplications {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsForNetworkZoneResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}
