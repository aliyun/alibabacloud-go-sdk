// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v []*ListChatConfigurationsRequestFilters) *ListChatConfigurationsRequest
	GetFilters() []*ListChatConfigurationsRequestFilters
	SetMaxResults(v int32) *ListChatConfigurationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListChatConfigurationsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListChatConfigurationsRequest
	GetRegionId() *string
}

type ListChatConfigurationsRequest struct {
	// example:
	//
	// [{"Name": "status", "Value": "inactive"}, {"Name": "type", "Value": "private"}]
	Filters []*ListChatConfigurationsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AQEBARoBBgEAAgIBBQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListChatConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListChatConfigurationsRequest) GetFilters() []*ListChatConfigurationsRequestFilters {
	return s.Filters
}

func (s *ListChatConfigurationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListChatConfigurationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListChatConfigurationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListChatConfigurationsRequest) SetFilters(v []*ListChatConfigurationsRequestFilters) *ListChatConfigurationsRequest {
	s.Filters = v
	return s
}

func (s *ListChatConfigurationsRequest) SetMaxResults(v int32) *ListChatConfigurationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListChatConfigurationsRequest) SetNextToken(v string) *ListChatConfigurationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListChatConfigurationsRequest) SetRegionId(v string) *ListChatConfigurationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListChatConfigurationsRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListChatConfigurationsRequestFilters struct {
	// example:
	//
	// status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// inactive
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListChatConfigurationsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListChatConfigurationsRequestFilters) GoString() string {
	return s.String()
}

func (s *ListChatConfigurationsRequestFilters) GetName() *string {
	return s.Name
}

func (s *ListChatConfigurationsRequestFilters) GetValue() *string {
	return s.Value
}

func (s *ListChatConfigurationsRequestFilters) SetName(v string) *ListChatConfigurationsRequestFilters {
	s.Name = &v
	return s
}

func (s *ListChatConfigurationsRequestFilters) SetValue(v string) *ListChatConfigurationsRequestFilters {
	s.Value = &v
	return s
}

func (s *ListChatConfigurationsRequestFilters) Validate() error {
	return dara.Validate(s)
}
