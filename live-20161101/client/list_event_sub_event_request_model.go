// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventSubEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListEventSubEventRequest
	GetAppId() *string
	SetEndTime(v int64) *ListEventSubEventRequest
	GetEndTime() *int64
	SetPageNo(v int64) *ListEventSubEventRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListEventSubEventRequest
	GetPageSize() *int64
	SetStartTime(v int64) *ListEventSubEventRequest
	GetStartTime() *int64
	SetSubscribeId(v string) *ListEventSubEventRequest
	GetSubscribeId() *string
}

type ListEventSubEventRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9qb1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1698201013
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of callback records to return on each page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1698195600
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The subscription ID. You can obtain the ID from the response to the [CreateEventSub](https://help.aliyun.com/document_detail/2848209.html) operation.
	//
	// example:
	//
	// ad53276431c****
	SubscribeId *string `json:"SubscribeId,omitempty" xml:"SubscribeId,omitempty"`
}

func (s ListEventSubEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventSubEventRequest) GoString() string {
	return s.String()
}

func (s *ListEventSubEventRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListEventSubEventRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListEventSubEventRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListEventSubEventRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListEventSubEventRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListEventSubEventRequest) GetSubscribeId() *string {
	return s.SubscribeId
}

func (s *ListEventSubEventRequest) SetAppId(v string) *ListEventSubEventRequest {
	s.AppId = &v
	return s
}

func (s *ListEventSubEventRequest) SetEndTime(v int64) *ListEventSubEventRequest {
	s.EndTime = &v
	return s
}

func (s *ListEventSubEventRequest) SetPageNo(v int64) *ListEventSubEventRequest {
	s.PageNo = &v
	return s
}

func (s *ListEventSubEventRequest) SetPageSize(v int64) *ListEventSubEventRequest {
	s.PageSize = &v
	return s
}

func (s *ListEventSubEventRequest) SetStartTime(v int64) *ListEventSubEventRequest {
	s.StartTime = &v
	return s
}

func (s *ListEventSubEventRequest) SetSubscribeId(v string) *ListEventSubEventRequest {
	s.SubscribeId = &v
	return s
}

func (s *ListEventSubEventRequest) Validate() error {
	return dara.Validate(s)
}
