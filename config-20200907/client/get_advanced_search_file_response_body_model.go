// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdvancedSearchFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAdvancedSearchFileResponseBody
	GetRequestId() *string
	SetResourceSearch(v *GetAdvancedSearchFileResponseBodyResourceSearch) *GetAdvancedSearchFileResponseBody
	GetResourceSearch() *GetAdvancedSearchFileResponseBodyResourceSearch
}

type GetAdvancedSearchFileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9366FE45-3C83-54FB-8BB1-44176B200706
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource file.
	ResourceSearch *GetAdvancedSearchFileResponseBodyResourceSearch `json:"ResourceSearch,omitempty" xml:"ResourceSearch,omitempty" type:"Struct"`
}

func (s GetAdvancedSearchFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAdvancedSearchFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdvancedSearchFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAdvancedSearchFileResponseBody) GetResourceSearch() *GetAdvancedSearchFileResponseBodyResourceSearch {
	return s.ResourceSearch
}

func (s *GetAdvancedSearchFileResponseBody) SetRequestId(v string) *GetAdvancedSearchFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdvancedSearchFileResponseBody) SetResourceSearch(v *GetAdvancedSearchFileResponseBodyResourceSearch) *GetAdvancedSearchFileResponseBody {
	s.ResourceSearch = v
	return s
}

func (s *GetAdvancedSearchFileResponseBody) Validate() error {
	if s.ResourceSearch != nil {
		if err := s.ResourceSearch.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAdvancedSearchFileResponseBodyResourceSearch struct {
	// The download URL of the resource file.
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The time when the resource file was generated. The value is a timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1688281755480
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

func (s GetAdvancedSearchFileResponseBodyResourceSearch) String() string {
	return dara.Prettify(s)
}

func (s GetAdvancedSearchFileResponseBodyResourceSearch) GoString() string {
	return s.String()
}

func (s *GetAdvancedSearchFileResponseBodyResourceSearch) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetAdvancedSearchFileResponseBodyResourceSearch) GetResourceInventoryGenerateTime() *int64 {
	return s.ResourceInventoryGenerateTime
}

func (s *GetAdvancedSearchFileResponseBodyResourceSearch) GetStatus() *string {
	return s.Status
}

func (s *GetAdvancedSearchFileResponseBodyResourceSearch) SetDownloadUrl(v string) *GetAdvancedSearchFileResponseBodyResourceSearch {
	s.DownloadUrl = &v
	return s
}

func (s *GetAdvancedSearchFileResponseBodyResourceSearch) SetResourceInventoryGenerateTime(v int64) *GetAdvancedSearchFileResponseBodyResourceSearch {
	s.ResourceInventoryGenerateTime = &v
	return s
}

func (s *GetAdvancedSearchFileResponseBodyResourceSearch) SetStatus(v string) *GetAdvancedSearchFileResponseBodyResourceSearch {
	s.Status = &v
	return s
}

func (s *GetAdvancedSearchFileResponseBodyResourceSearch) Validate() error {
	return dara.Validate(s)
}
