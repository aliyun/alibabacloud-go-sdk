// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUploadTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListUploadTasksRequest
	GetEndTime() *string
	SetSiteId(v int64) *ListUploadTasksRequest
	GetSiteId() *int64
	SetStartTime(v string) *ListUploadTasksRequest
	GetStartTime() *string
	SetType(v string) *ListUploadTasksRequest
	GetType() *string
}

type ListUploadTasksRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SiteId    *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListUploadTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUploadTasksRequest) GoString() string {
	return s.String()
}

func (s *ListUploadTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListUploadTasksRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListUploadTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListUploadTasksRequest) GetType() *string {
	return s.Type
}

func (s *ListUploadTasksRequest) SetEndTime(v string) *ListUploadTasksRequest {
	s.EndTime = &v
	return s
}

func (s *ListUploadTasksRequest) SetSiteId(v int64) *ListUploadTasksRequest {
	s.SiteId = &v
	return s
}

func (s *ListUploadTasksRequest) SetStartTime(v string) *ListUploadTasksRequest {
	s.StartTime = &v
	return s
}

func (s *ListUploadTasksRequest) SetType(v string) *ListUploadTasksRequest {
	s.Type = &v
	return s
}

func (s *ListUploadTasksRequest) Validate() error {
	return dara.Validate(s)
}
