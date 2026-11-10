// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *QueryPushRecordsRequest
	GetAppKey() *int64
	SetEndTime(v string) *QueryPushRecordsRequest
	GetEndTime() *string
	SetKeyword(v string) *QueryPushRecordsRequest
	GetKeyword() *string
	SetNextToken(v string) *QueryPushRecordsRequest
	GetNextToken() *string
	SetPage(v int32) *QueryPushRecordsRequest
	GetPage() *int32
	SetPageSize(v int32) *QueryPushRecordsRequest
	GetPageSize() *int32
	SetPushType(v string) *QueryPushRecordsRequest
	GetPushType() *string
	SetSource(v string) *QueryPushRecordsRequest
	GetSource() *string
	SetStartTime(v string) *QueryPushRecordsRequest
	GetStartTime() *string
	SetTarget(v string) *QueryPushRecordsRequest
	GetTarget() *string
}

type QueryPushRecordsRequest struct {
	// The AppKey information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 333526247
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The end time of the query. The time is in UTC and follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-29T06:24:58Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The query keyword. The query is performed based on the MessageId, Title, and Body fields. The system performs Chinese word segmentation on the title and body of the push, and uses the segmented words for the query instead of substring matching.
	//
	// example:
	//
	// Statistics data test notification 805
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated and no longer used.
	//
	// example:
	//
	// FFPpkmhCPm*****************xjk=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number of the push records. Default value: 1. Maximum value: 10000.
	//
	// example:
	//
	// 8
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The push type. Valid values:
	//
	// - **MESSAGE**: message
	//
	// - **NOTICE**: notification
	//
	// - **LIVE_ACTIVITY**: live activity
	//
	// example:
	//
	// NOTICE
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// The push source. Valid values:
	//
	// - **API**: push initiated by OpenAPI, including pushes initiated from the OpenAPI platform console
	//
	// - **CONSOLE**: push initiated from the Mobile Push console
	//
	// - **OpenAPIExplorer**: push initiated from the Alibaba Cloud OpenAPI Explorer platform
	//
	// example:
	//
	// API
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The start time of the query. The time is in UTC and follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-15T02:05:24Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The push target. Valid values:
	//
	// - **DEVICE**: push by device
	//
	// - **ACCOUNT**: push by account
	//
	// - **ALIAS**: push by alias
	//
	// - **TAG**: push by tag
	//
	// - **ALL**: push to all devices
	//
	// - **TBD**: initialize continuous push, where the push target is specified by the subsequent ContinuouslyPush operation
	//
	// example:
	//
	// DEVICE
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s QueryPushRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPushRecordsRequest) GoString() string {
	return s.String()
}

func (s *QueryPushRecordsRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *QueryPushRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryPushRecordsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryPushRecordsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryPushRecordsRequest) GetPage() *int32 {
	return s.Page
}

func (s *QueryPushRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPushRecordsRequest) GetPushType() *string {
	return s.PushType
}

func (s *QueryPushRecordsRequest) GetSource() *string {
	return s.Source
}

func (s *QueryPushRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryPushRecordsRequest) GetTarget() *string {
	return s.Target
}

func (s *QueryPushRecordsRequest) SetAppKey(v int64) *QueryPushRecordsRequest {
	s.AppKey = &v
	return s
}

func (s *QueryPushRecordsRequest) SetEndTime(v string) *QueryPushRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPushRecordsRequest) SetKeyword(v string) *QueryPushRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *QueryPushRecordsRequest) SetNextToken(v string) *QueryPushRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryPushRecordsRequest) SetPage(v int32) *QueryPushRecordsRequest {
	s.Page = &v
	return s
}

func (s *QueryPushRecordsRequest) SetPageSize(v int32) *QueryPushRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPushRecordsRequest) SetPushType(v string) *QueryPushRecordsRequest {
	s.PushType = &v
	return s
}

func (s *QueryPushRecordsRequest) SetSource(v string) *QueryPushRecordsRequest {
	s.Source = &v
	return s
}

func (s *QueryPushRecordsRequest) SetStartTime(v string) *QueryPushRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *QueryPushRecordsRequest) SetTarget(v string) *QueryPushRecordsRequest {
	s.Target = &v
	return s
}

func (s *QueryPushRecordsRequest) Validate() error {
	return dara.Validate(s)
}
