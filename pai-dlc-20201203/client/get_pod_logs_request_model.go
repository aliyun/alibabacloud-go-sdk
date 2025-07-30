// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPodLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadToFile(v bool) *GetPodLogsRequest
	GetDownloadToFile() *bool
	SetEndTime(v string) *GetPodLogsRequest
	GetEndTime() *string
	SetMaxLines(v int32) *GetPodLogsRequest
	GetMaxLines() *int32
	SetPodUid(v string) *GetPodLogsRequest
	GetPodUid() *string
	SetStartTime(v string) *GetPodLogsRequest
	GetStartTime() *string
}

type GetPodLogsRequest struct {
	// Specifies whether to download the log file. Default value: false. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// true
	DownloadToFile *bool `json:"DownloadToFile,omitempty" xml:"DownloadToFile,omitempty"`
	// The end time of the query. Default value: current time.
	//
	// example:
	//
	// 2020-11-08T17:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of log entries. Default value: 2000.
	//
	// example:
	//
	// 100
	MaxLines *int32 `json:"MaxLines,omitempty" xml:"MaxLines,omitempty"`
	// The node UID. For more information about how to obtain a node UID, see [GetJob](https://help.aliyun.com/document_detail/459677.html).
	//
	// example:
	//
	// fe846462-af2c-4521-bd6f-96787a57****
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// The start time of the query. Default value: 7 days ago.
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetPodLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPodLogsRequest) GoString() string {
	return s.String()
}

func (s *GetPodLogsRequest) GetDownloadToFile() *bool {
	return s.DownloadToFile
}

func (s *GetPodLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetPodLogsRequest) GetMaxLines() *int32 {
	return s.MaxLines
}

func (s *GetPodLogsRequest) GetPodUid() *string {
	return s.PodUid
}

func (s *GetPodLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetPodLogsRequest) SetDownloadToFile(v bool) *GetPodLogsRequest {
	s.DownloadToFile = &v
	return s
}

func (s *GetPodLogsRequest) SetEndTime(v string) *GetPodLogsRequest {
	s.EndTime = &v
	return s
}

func (s *GetPodLogsRequest) SetMaxLines(v int32) *GetPodLogsRequest {
	s.MaxLines = &v
	return s
}

func (s *GetPodLogsRequest) SetPodUid(v string) *GetPodLogsRequest {
	s.PodUid = &v
	return s
}

func (s *GetPodLogsRequest) SetStartTime(v string) *GetPodLogsRequest {
	s.StartTime = &v
	return s
}

func (s *GetPodLogsRequest) Validate() error {
	return dara.Validate(s)
}
