// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRecallManagementServicesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRecallManagementServicesResponseBody
	GetNextToken() *string
	SetRecallManagementServices(v []*ListRecallManagementServicesResponseBodyRecallManagementServices) *ListRecallManagementServicesResponseBody
	GetRecallManagementServices() []*ListRecallManagementServicesResponseBodyRecallManagementServices
	SetRequestId(v string) *ListRecallManagementServicesResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListRecallManagementServicesResponseBody
	GetTotalCount() *string
}

type ListRecallManagementServicesResponseBody struct {
	// example:
	//
	// 0
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// ""
	NextToken                *string                                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RecallManagementServices []*ListRecallManagementServicesResponseBodyRecallManagementServices `json:"RecallManagementServices,omitempty" xml:"RecallManagementServices,omitempty" type:"Repeated"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecallManagementServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecallManagementServicesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRecallManagementServicesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecallManagementServicesResponseBody) GetRecallManagementServices() []*ListRecallManagementServicesResponseBodyRecallManagementServices {
	return s.RecallManagementServices
}

func (s *ListRecallManagementServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecallManagementServicesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListRecallManagementServicesResponseBody) SetMaxResults(v int32) *ListRecallManagementServicesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRecallManagementServicesResponseBody) SetNextToken(v string) *ListRecallManagementServicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecallManagementServicesResponseBody) SetRecallManagementServices(v []*ListRecallManagementServicesResponseBodyRecallManagementServices) *ListRecallManagementServicesResponseBody {
	s.RecallManagementServices = v
	return s
}

func (s *ListRecallManagementServicesResponseBody) SetRequestId(v string) *ListRecallManagementServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecallManagementServicesResponseBody) SetTotalCount(v string) *ListRecallManagementServicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRecallManagementServicesResponseBody) Validate() error {
	if s.RecallManagementServices != nil {
		for _, item := range s.RecallManagementServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRecallManagementServicesResponseBodyRecallManagementServices struct {
	// example:
	//
	// 1
	CurrentRecallManagementServiceVersionId *string `json:"CurrentRecallManagementServiceVersionId,omitempty" xml:"CurrentRecallManagementServiceVersionId,omitempty"`
	// example:
	//
	// version-1
	CurrentRecallManagementServiceVersionName *string `json:"CurrentRecallManagementServiceVersionName,omitempty" xml:"CurrentRecallManagementServiceVersionName,omitempty"`
	// example:
	//
	// this is a test recall
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// hot_group_recall
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 3
	RecallManagementServiceId *string `json:"RecallManagementServiceId,omitempty" xml:"RecallManagementServiceId,omitempty"`
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListRecallManagementServicesResponseBodyRecallManagementServices) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementServicesResponseBodyRecallManagementServices) GoString() string {
	return s.String()
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) GetCurrentRecallManagementServiceVersionId() *string {
	return s.CurrentRecallManagementServiceVersionId
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) GetCurrentRecallManagementServiceVersionName() *string {
	return s.CurrentRecallManagementServiceVersionName
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) GetDescription() *string {
	return s.Description
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) GetName() *string {
	return s.Name
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) GetRecallManagementServiceId() *string {
	return s.RecallManagementServiceId
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) GetStatus() *string {
	return s.Status
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) SetCurrentRecallManagementServiceVersionId(v string) *ListRecallManagementServicesResponseBodyRecallManagementServices {
	s.CurrentRecallManagementServiceVersionId = &v
	return s
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) SetCurrentRecallManagementServiceVersionName(v string) *ListRecallManagementServicesResponseBodyRecallManagementServices {
	s.CurrentRecallManagementServiceVersionName = &v
	return s
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) SetDescription(v string) *ListRecallManagementServicesResponseBodyRecallManagementServices {
	s.Description = &v
	return s
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) SetGmtCreateTime(v string) *ListRecallManagementServicesResponseBodyRecallManagementServices {
	s.GmtCreateTime = &v
	return s
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) SetGmtModifiedTime(v string) *ListRecallManagementServicesResponseBodyRecallManagementServices {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) SetName(v string) *ListRecallManagementServicesResponseBodyRecallManagementServices {
	s.Name = &v
	return s
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) SetRecallManagementServiceId(v string) *ListRecallManagementServicesResponseBodyRecallManagementServices {
	s.RecallManagementServiceId = &v
	return s
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) SetStatus(v string) *ListRecallManagementServicesResponseBodyRecallManagementServices {
	s.Status = &v
	return s
}

func (s *ListRecallManagementServicesResponseBodyRecallManagementServices) Validate() error {
	return dara.Validate(s)
}
