// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupResourceGroupEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *LookupResourceGroupEventsRequest
	GetEndTime() *string
	SetEventCategory(v string) *LookupResourceGroupEventsRequest
	GetEventCategory() *string
	SetLookupAttributes(v []*LookupResourceGroupEventsRequestLookupAttributes) *LookupResourceGroupEventsRequest
	GetLookupAttributes() []*LookupResourceGroupEventsRequestLookupAttributes
	SetMaxResults(v string) *LookupResourceGroupEventsRequest
	GetMaxResults() *string
	SetNextToken(v string) *LookupResourceGroupEventsRequest
	GetNextToken() *string
	SetResourceGroupDisplayName(v string) *LookupResourceGroupEventsRequest
	GetResourceGroupDisplayName() *string
	SetResourceGroupId(v string) *LookupResourceGroupEventsRequest
	GetResourceGroupId() *string
	SetStartTime(v string) *LookupResourceGroupEventsRequest
	GetStartTime() *string
}

type LookupResourceGroupEventsRequest struct {
	// The end of the time range to query. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-11-30 23:43:16
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event type.
	//
	// This parameter is required.
	//
	// example:
	//
	// MemberShipChange
	EventCategory *string `json:"EventCategory,omitempty" xml:"EventCategory,omitempty"`
	// The attributes used for advanced search.
	LookupAttributes []*LookupResourceGroupEventsRequestLookupAttributes `json:"LookupAttributes,omitempty" xml:"LookupAttributes,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource group name.
	//
	// example:
	//
	// ProjectA
	ResourceGroupDisplayName *string `json:"ResourceGroupDisplayName,omitempty" xml:"ResourceGroupDisplayName,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-9gLOoK****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-11-30 23:43:16
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s LookupResourceGroupEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s LookupResourceGroupEventsRequest) GoString() string {
	return s.String()
}

func (s *LookupResourceGroupEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *LookupResourceGroupEventsRequest) GetEventCategory() *string {
	return s.EventCategory
}

func (s *LookupResourceGroupEventsRequest) GetLookupAttributes() []*LookupResourceGroupEventsRequestLookupAttributes {
	return s.LookupAttributes
}

func (s *LookupResourceGroupEventsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *LookupResourceGroupEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *LookupResourceGroupEventsRequest) GetResourceGroupDisplayName() *string {
	return s.ResourceGroupDisplayName
}

func (s *LookupResourceGroupEventsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *LookupResourceGroupEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *LookupResourceGroupEventsRequest) SetEndTime(v string) *LookupResourceGroupEventsRequest {
	s.EndTime = &v
	return s
}

func (s *LookupResourceGroupEventsRequest) SetEventCategory(v string) *LookupResourceGroupEventsRequest {
	s.EventCategory = &v
	return s
}

func (s *LookupResourceGroupEventsRequest) SetLookupAttributes(v []*LookupResourceGroupEventsRequestLookupAttributes) *LookupResourceGroupEventsRequest {
	s.LookupAttributes = v
	return s
}

func (s *LookupResourceGroupEventsRequest) SetMaxResults(v string) *LookupResourceGroupEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *LookupResourceGroupEventsRequest) SetNextToken(v string) *LookupResourceGroupEventsRequest {
	s.NextToken = &v
	return s
}

func (s *LookupResourceGroupEventsRequest) SetResourceGroupDisplayName(v string) *LookupResourceGroupEventsRequest {
	s.ResourceGroupDisplayName = &v
	return s
}

func (s *LookupResourceGroupEventsRequest) SetResourceGroupId(v string) *LookupResourceGroupEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *LookupResourceGroupEventsRequest) SetStartTime(v string) *LookupResourceGroupEventsRequest {
	s.StartTime = &v
	return s
}

func (s *LookupResourceGroupEventsRequest) Validate() error {
	if s.LookupAttributes != nil {
		for _, item := range s.LookupAttributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type LookupResourceGroupEventsRequestLookupAttributes struct {
	// The key of the attribute.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the attribute.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s LookupResourceGroupEventsRequestLookupAttributes) String() string {
	return dara.Prettify(s)
}

func (s LookupResourceGroupEventsRequestLookupAttributes) GoString() string {
	return s.String()
}

func (s *LookupResourceGroupEventsRequestLookupAttributes) GetKey() *string {
	return s.Key
}

func (s *LookupResourceGroupEventsRequestLookupAttributes) GetValue() *string {
	return s.Value
}

func (s *LookupResourceGroupEventsRequestLookupAttributes) SetKey(v string) *LookupResourceGroupEventsRequestLookupAttributes {
	s.Key = &v
	return s
}

func (s *LookupResourceGroupEventsRequestLookupAttributes) SetValue(v string) *LookupResourceGroupEventsRequestLookupAttributes {
	s.Value = &v
	return s
}

func (s *LookupResourceGroupEventsRequestLookupAttributes) Validate() error {
	return dara.Validate(s)
}
