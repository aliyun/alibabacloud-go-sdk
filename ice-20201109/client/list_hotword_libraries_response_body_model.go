// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotwordLibrariesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHotwordLibraryList(v []*ListHotwordLibrariesResponseBodyHotwordLibraryList) *ListHotwordLibrariesResponseBody
	GetHotwordLibraryList() []*ListHotwordLibrariesResponseBodyHotwordLibraryList
	SetMaxResults(v int32) *ListHotwordLibrariesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListHotwordLibrariesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListHotwordLibrariesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHotwordLibrariesResponseBody
	GetTotalCount() *int32
}

type ListHotwordLibrariesResponseBody struct {
	// The hotword libraries.
	HotwordLibraryList []*ListHotwordLibrariesResponseBodyHotwordLibraryList `json:"HotwordLibraryList,omitempty" xml:"HotwordLibraryList,omitempty" type:"Repeated"`
	// The maximum number of hotword libraries that can be returned.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token that can be used in the next request to retrieve a new page of results. If it is empty, all results are returned.
	//
	// example:
	//
	// CBB6BC61D08
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ****9262E3DA-07FA-4862-FCBB6BC61D08*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of hotword libraries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHotwordLibrariesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotwordLibrariesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotwordLibrariesResponseBody) GetHotwordLibraryList() []*ListHotwordLibrariesResponseBodyHotwordLibraryList {
	return s.HotwordLibraryList
}

func (s *ListHotwordLibrariesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHotwordLibrariesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHotwordLibrariesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotwordLibrariesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHotwordLibrariesResponseBody) SetHotwordLibraryList(v []*ListHotwordLibrariesResponseBodyHotwordLibraryList) *ListHotwordLibrariesResponseBody {
	s.HotwordLibraryList = v
	return s
}

func (s *ListHotwordLibrariesResponseBody) SetMaxResults(v int32) *ListHotwordLibrariesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHotwordLibrariesResponseBody) SetNextToken(v string) *ListHotwordLibrariesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHotwordLibrariesResponseBody) SetRequestId(v string) *ListHotwordLibrariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotwordLibrariesResponseBody) SetTotalCount(v int32) *ListHotwordLibrariesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHotwordLibrariesResponseBody) Validate() error {
	if s.HotwordLibraryList != nil {
		for _, item := range s.HotwordLibraryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotwordLibrariesResponseBodyHotwordLibraryList struct {
	// The time when the hotword library was created.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the hotword library. It can be up to 200 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the hotword library.
	//
	// example:
	//
	// a93b91141c0f422fa114af203f8b****
	HotwordLibraryId *string `json:"HotwordLibraryId,omitempty" xml:"HotwordLibraryId,omitempty"`
	// The time when the hotword library was last modified.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the hotword library.
	//
	// example:
	//
	// my_hotwords
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The usage scenario of the hotword library. Valid values:
	//
	// 	- ASR: Automatic Speech Recognition
	//
	// 	- StructuredMediaAssets: structured media analysis
	//
	// 	- VideoTranslation: Video translation This field cannot be modified after the hotword library is created.
	//
	// example:
	//
	// ASR
	UsageScenario *string `json:"UsageScenario,omitempty" xml:"UsageScenario,omitempty"`
}

func (s ListHotwordLibrariesResponseBodyHotwordLibraryList) String() string {
	return dara.Prettify(s)
}

func (s ListHotwordLibrariesResponseBodyHotwordLibraryList) GoString() string {
	return s.String()
}

func (s *ListHotwordLibrariesResponseBodyHotwordLibraryList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListHotwordLibrariesResponseBodyHotwordLibraryList) GetDescription() *string {
	return s.Description
}

func (s *ListHotwordLibrariesResponseBodyHotwordLibraryList) GetHotwordLibraryId() *string {
	return s.HotwordLibraryId
}

func (s *ListHotwordLibrariesResponseBodyHotwordLibraryList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListHotwordLibrariesResponseBodyHotwordLibraryList) GetName() *string {
	return s.Name
}

func (s *ListHotwordLibrariesResponseBodyHotwordLibraryList) GetUsageScenario() *string {
	return s.UsageScenario
}

func (s *ListHotwordLibrariesResponseBodyHotwordLibraryList) SetCreationTime(v string) *ListHotwordLibrariesResponseBodyHotwordLibraryList {
	s.CreationTime = &v
	return s
}

func (s *ListHotwordLibrariesResponseBodyHotwordLibraryList) SetDescription(v string) *ListHotwordLibrariesResponseBodyHotwordLibraryList {
	s.Description = &v
	return s
}

func (s *ListHotwordLibrariesResponseBodyHotwordLibraryList) SetHotwordLibraryId(v string) *ListHotwordLibrariesResponseBodyHotwordLibraryList {
	s.HotwordLibraryId = &v
	return s
}

func (s *ListHotwordLibrariesResponseBodyHotwordLibraryList) SetModifiedTime(v string) *ListHotwordLibrariesResponseBodyHotwordLibraryList {
	s.ModifiedTime = &v
	return s
}

func (s *ListHotwordLibrariesResponseBodyHotwordLibraryList) SetName(v string) *ListHotwordLibrariesResponseBodyHotwordLibraryList {
	s.Name = &v
	return s
}

func (s *ListHotwordLibrariesResponseBodyHotwordLibraryList) SetUsageScenario(v string) *ListHotwordLibrariesResponseBodyHotwordLibraryList {
	s.UsageScenario = &v
	return s
}

func (s *ListHotwordLibrariesResponseBodyHotwordLibraryList) Validate() error {
	return dara.Validate(s)
}
