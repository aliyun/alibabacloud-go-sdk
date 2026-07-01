// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaBasicInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthTimeout(v int64) *ListMediaBasicInfosRequest
	GetAuthTimeout() *int64
	SetBusinessType(v string) *ListMediaBasicInfosRequest
	GetBusinessType() *string
	SetEndTime(v string) *ListMediaBasicInfosRequest
	GetEndTime() *string
	SetIncludeFileBasicInfo(v bool) *ListMediaBasicInfosRequest
	GetIncludeFileBasicInfo() *bool
	SetMaxResults(v int32) *ListMediaBasicInfosRequest
	GetMaxResults() *int32
	SetMediaId(v string) *ListMediaBasicInfosRequest
	GetMediaId() *string
	SetMediaType(v string) *ListMediaBasicInfosRequest
	GetMediaType() *string
	SetNextToken(v string) *ListMediaBasicInfosRequest
	GetNextToken() *string
	SetSortBy(v string) *ListMediaBasicInfosRequest
	GetSortBy() *string
	SetSource(v string) *ListMediaBasicInfosRequest
	GetSource() *string
	SetStartTime(v string) *ListMediaBasicInfosRequest
	GetStartTime() *string
	SetStatus(v string) *ListMediaBasicInfosRequest
	GetStatus() *string
}

type ListMediaBasicInfosRequest struct {
	// The authentication expiration time, in seconds.
	//
	// - Minimum value: 1.
	//
	// - Maximum value: 86400.
	//
	// - Default value: 3600.
	//
	// example:
	//
	// 30
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The business type. Valid values:
	//
	// - `subtitles`: subtitles
	//
	// - `watermark`: watermark
	//
	// - `opening`: opening
	//
	// - `ending`: ending
	//
	// - `general`: General
	//
	// example:
	//
	// opening
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The end time of the query range, based on the creation time (`utcCreated`).
	//
	// - The query returns results created at or before this time (inclusive).
	//
	// - The time must be in UTC and comply with the ISO 8601 standard. The format is `YYYY-MM-DD\\"T\\"HH:mm:ss\\"Z\\"`. For example, `2017-01-11T12:00:00Z` corresponds to 20:00:00 on January 11, 2017 (UTC+8).
	//
	// example:
	//
	// 2020-12-20T13:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Set to `true` to include basic file information in the response.
	//
	// example:
	//
	// true
	IncludeFileBasicInfo *bool `json:"IncludeFileBasicInfo,omitempty" xml:"IncludeFileBasicInfo,omitempty"`
	// The maximum number of results to return per page.
	//
	// Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The media ID.
	//
	// example:
	//
	// ****019b82e24b37a1c2958dec38****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The media type. Valid values:
	//
	// - `image`: image
	//
	// - `video`: video
	//
	// - `audio`: audio
	//
	// - `text`: text
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The pagination token used to retrieve the next page of results. An empty value indicates that all results have been returned.
	//
	// example:
	//
	// pSa1SQ0wCe5pzVrQ6mWZEw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order for results based on the creation time (`utcCreated`). The default is descending. Valid values:
	//
	// - `desc`: Descending order
	//
	// - `asc`: Ascending order
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The source. Valid values:
	//
	// - `oss`: OSS
	//
	// - `vod`: video on demand
	//
	// - `live`: live streaming
	//
	// - `general`: Other sources (default).
	//
	// example:
	//
	// oss
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The start time of the query range, based on the creation time (`utcCreated`).
	//
	// - The query returns results created after this time (exclusive).
	//
	// - The time must be in UTC and comply with the ISO 8601 standard. The format is `YYYY-MM-DD\\"T\\"HH:mm:ss\\"Z\\"`. For example, `2017-01-11T12:00:00Z` corresponds to 20:00:00 on January 11, 2017 (UTC+8).
	//
	// example:
	//
	// 2020-12-20T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The media status. Valid values:
	//
	// - `Init`: The source file is not ready.
	//
	// - `Preparing`: The source file is being prepared. For example, it is being uploaded or composited.
	//
	// - `PrepareFail`: The source file failed to prepare. This may occur, for example, if the system fails to retrieve information about the source file.
	//
	// - `Normal`: The source file is ready.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMediaBasicInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMediaBasicInfosRequest) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosRequest) GetAuthTimeout() *int64 {
	return s.AuthTimeout
}

func (s *ListMediaBasicInfosRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ListMediaBasicInfosRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListMediaBasicInfosRequest) GetIncludeFileBasicInfo() *bool {
	return s.IncludeFileBasicInfo
}

func (s *ListMediaBasicInfosRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMediaBasicInfosRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *ListMediaBasicInfosRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ListMediaBasicInfosRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMediaBasicInfosRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListMediaBasicInfosRequest) GetSource() *string {
	return s.Source
}

func (s *ListMediaBasicInfosRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListMediaBasicInfosRequest) GetStatus() *string {
	return s.Status
}

func (s *ListMediaBasicInfosRequest) SetAuthTimeout(v int64) *ListMediaBasicInfosRequest {
	s.AuthTimeout = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetBusinessType(v string) *ListMediaBasicInfosRequest {
	s.BusinessType = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetEndTime(v string) *ListMediaBasicInfosRequest {
	s.EndTime = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetIncludeFileBasicInfo(v bool) *ListMediaBasicInfosRequest {
	s.IncludeFileBasicInfo = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetMaxResults(v int32) *ListMediaBasicInfosRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetMediaId(v string) *ListMediaBasicInfosRequest {
	s.MediaId = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetMediaType(v string) *ListMediaBasicInfosRequest {
	s.MediaType = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetNextToken(v string) *ListMediaBasicInfosRequest {
	s.NextToken = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetSortBy(v string) *ListMediaBasicInfosRequest {
	s.SortBy = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetSource(v string) *ListMediaBasicInfosRequest {
	s.Source = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetStartTime(v string) *ListMediaBasicInfosRequest {
	s.StartTime = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetStatus(v string) *ListMediaBasicInfosRequest {
	s.Status = &v
	return s
}

func (s *ListMediaBasicInfosRequest) Validate() error {
	return dara.Validate(s)
}
