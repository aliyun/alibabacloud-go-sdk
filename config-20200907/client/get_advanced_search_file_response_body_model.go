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
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DownloadUrl                   *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ResourceInventoryGenerateTime *int64  `json:"ResourceInventoryGenerateTime,omitempty" xml:"ResourceInventoryGenerateTime,omitempty"`
	Status                        *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
