// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataCenter(v string) *ListLiveMessageAppsRequest
	GetDataCenter() *string
	SetNextPageToken(v int64) *ListLiveMessageAppsRequest
	GetNextPageToken() *int64
	SetSortType(v int32) *ListLiveMessageAppsRequest
	GetSortType() *int32
}

type ListLiveMessageAppsRequest struct {
	// The data center. The value must be the same as the data center specified in [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html). Valid values: cn-shanghai (China (Shanghai)) and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The starting position of the query page. If this parameter is left empty or set to -1, the first page is returned by default.
	//
	// example:
	//
	// -1
	NextPageToken *int64 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The sort type by creation time. Valid values:
	//
	// - 1: ascending order.
	//
	// - 2: descending order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SortType *int32 `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s ListLiveMessageAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageAppsRequest) GoString() string {
	return s.String()
}

func (s *ListLiveMessageAppsRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ListLiveMessageAppsRequest) GetNextPageToken() *int64 {
	return s.NextPageToken
}

func (s *ListLiveMessageAppsRequest) GetSortType() *int32 {
	return s.SortType
}

func (s *ListLiveMessageAppsRequest) SetDataCenter(v string) *ListLiveMessageAppsRequest {
	s.DataCenter = &v
	return s
}

func (s *ListLiveMessageAppsRequest) SetNextPageToken(v int64) *ListLiveMessageAppsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListLiveMessageAppsRequest) SetSortType(v int32) *ListLiveMessageAppsRequest {
	s.SortType = &v
	return s
}

func (s *ListLiveMessageAppsRequest) Validate() error {
	return dara.Validate(s)
}
