// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateAdvancedSearchFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAggregateAdvancedSearchFileResponseBody
	GetRequestId() *string
	SetResourceSearch(v *GetAggregateAdvancedSearchFileResponseBodyResourceSearch) *GetAggregateAdvancedSearchFileResponseBody
	GetResourceSearch() *GetAggregateAdvancedSearchFileResponseBodyResourceSearch
}

type GetAggregateAdvancedSearchFileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6DB86284-DB7F-5936-B210-3B53D6D41B03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource file.
	ResourceSearch *GetAggregateAdvancedSearchFileResponseBodyResourceSearch `json:"ResourceSearch,omitempty" xml:"ResourceSearch,omitempty" type:"Struct"`
}

func (s GetAggregateAdvancedSearchFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateAdvancedSearchFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateAdvancedSearchFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateAdvancedSearchFileResponseBody) GetResourceSearch() *GetAggregateAdvancedSearchFileResponseBodyResourceSearch {
	return s.ResourceSearch
}

func (s *GetAggregateAdvancedSearchFileResponseBody) SetRequestId(v string) *GetAggregateAdvancedSearchFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateAdvancedSearchFileResponseBody) SetResourceSearch(v *GetAggregateAdvancedSearchFileResponseBodyResourceSearch) *GetAggregateAdvancedSearchFileResponseBody {
	s.ResourceSearch = v
	return s
}

func (s *GetAggregateAdvancedSearchFileResponseBody) Validate() error {
	if s.ResourceSearch != nil {
		if err := s.ResourceSearch.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateAdvancedSearchFileResponseBodyResourceSearch struct {
	// The download URL of the resource file.
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The time when the resource file was generated. The value is a timestamp.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1691375618130
	ResourceInventoryGenerateTime *int64 `json:"ResourceInventoryGenerateTime,omitempty" xml:"ResourceInventoryGenerateTime,omitempty"`
	// The generation status of the resource file. Valid values:
	//
	// 	- CREATING: The resource file is being generated.
	//
	// 	- COMPLETE: The resource file is generated.
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAggregateAdvancedSearchFileResponseBodyResourceSearch) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateAdvancedSearchFileResponseBodyResourceSearch) GoString() string {
	return s.String()
}

func (s *GetAggregateAdvancedSearchFileResponseBodyResourceSearch) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetAggregateAdvancedSearchFileResponseBodyResourceSearch) GetResourceInventoryGenerateTime() *int64 {
	return s.ResourceInventoryGenerateTime
}

func (s *GetAggregateAdvancedSearchFileResponseBodyResourceSearch) GetStatus() *string {
	return s.Status
}

func (s *GetAggregateAdvancedSearchFileResponseBodyResourceSearch) SetDownloadUrl(v string) *GetAggregateAdvancedSearchFileResponseBodyResourceSearch {
	s.DownloadUrl = &v
	return s
}

func (s *GetAggregateAdvancedSearchFileResponseBodyResourceSearch) SetResourceInventoryGenerateTime(v int64) *GetAggregateAdvancedSearchFileResponseBodyResourceSearch {
	s.ResourceInventoryGenerateTime = &v
	return s
}

func (s *GetAggregateAdvancedSearchFileResponseBodyResourceSearch) SetStatus(v string) *GetAggregateAdvancedSearchFileResponseBodyResourceSearch {
	s.Status = &v
	return s
}

func (s *GetAggregateAdvancedSearchFileResponseBodyResourceSearch) Validate() error {
	return dara.Validate(s)
}
