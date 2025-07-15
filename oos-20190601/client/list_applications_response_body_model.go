// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody
	GetApplications() []*ListApplicationsResponseBodyApplications
	SetMaxResults(v int32) *ListApplicationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListApplicationsResponseBody
	GetRequestId() *string
}

type ListApplicationsResponseBody struct {
	// The details of the application.
	Applications []*ListApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 12067D53-56A9-561B-ADD6-61429D207117
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) GetApplications() []*ListApplicationsResponseBodyApplications {
	return s.Applications
}

func (s *ListApplicationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsResponseBody) SetApplications(v []*ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsResponseBody) SetMaxResults(v int32) *ListApplicationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsResponseBody) SetNextToken(v string) *ListApplicationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApplicationsResponseBodyApplications struct {
	// The type of the application.
	//
	// example:
	//
	// DingTalk
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2021-09-07T09:09:59Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// MyApplication
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags added to the application.
	//
	// example:
	//
	// {"k1": "v1","k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The time when the application was updated.
	//
	// example:
	//
	// 2021-09-07T09:09:59Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListApplicationsResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplications) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *ListApplicationsResponseBodyApplications) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListApplicationsResponseBodyApplications) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationsResponseBodyApplications) GetName() *string {
	return s.Name
}

func (s *ListApplicationsResponseBodyApplications) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListApplicationsResponseBodyApplications) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListApplicationsResponseBodyApplications) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationType(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationType = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetCreateDate(v string) *ListApplicationsResponseBodyApplications {
	s.CreateDate = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetDescription(v string) *ListApplicationsResponseBodyApplications {
	s.Description = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetName(v string) *ListApplicationsResponseBodyApplications {
	s.Name = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetResourceGroupId(v string) *ListApplicationsResponseBodyApplications {
	s.ResourceGroupId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetTags(v map[string]interface{}) *ListApplicationsResponseBodyApplications {
	s.Tags = v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetUpdateDate(v string) *ListApplicationsResponseBodyApplications {
	s.UpdateDate = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}
