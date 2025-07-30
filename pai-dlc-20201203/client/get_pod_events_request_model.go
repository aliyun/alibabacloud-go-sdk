// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPodEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetPodEventsRequest
	GetEndTime() *string
	SetMaxEventsNum(v int32) *GetPodEventsRequest
	GetMaxEventsNum() *int32
	SetPodUid(v string) *GetPodEventsRequest
	GetPodUid() *string
	SetStartTime(v string) *GetPodEventsRequest
	GetStartTime() *string
}

type GetPodEventsRequest struct {
	// The end time (UTC).
	//
	// example:
	//
	// 2020-11-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of events that can be returned.
	//
	// example:
	//
	// 100
	MaxEventsNum *int32 `json:"MaxEventsNum,omitempty" xml:"MaxEventsNum,omitempty"`
	// The node UID. Call [GetJob](https://help.aliyun.com/document_detail/459677.html) to get the node UID.
	//
	// example:
	//
	// dlc-20210126170216-*****-chief-0
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// The start time (UTC).
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetPodEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPodEventsRequest) GoString() string {
	return s.String()
}

func (s *GetPodEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetPodEventsRequest) GetMaxEventsNum() *int32 {
	return s.MaxEventsNum
}

func (s *GetPodEventsRequest) GetPodUid() *string {
	return s.PodUid
}

func (s *GetPodEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetPodEventsRequest) SetEndTime(v string) *GetPodEventsRequest {
	s.EndTime = &v
	return s
}

func (s *GetPodEventsRequest) SetMaxEventsNum(v int32) *GetPodEventsRequest {
	s.MaxEventsNum = &v
	return s
}

func (s *GetPodEventsRequest) SetPodUid(v string) *GetPodEventsRequest {
	s.PodUid = &v
	return s
}

func (s *GetPodEventsRequest) SetStartTime(v string) *GetPodEventsRequest {
	s.StartTime = &v
	return s
}

func (s *GetPodEventsRequest) Validate() error {
	return dara.Validate(s)
}
