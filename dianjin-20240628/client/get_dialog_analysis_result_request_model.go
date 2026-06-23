// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDialogAnalysisResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsc(v bool) *GetDialogAnalysisResultRequest
	GetAsc() *bool
	SetEndTime(v string) *GetDialogAnalysisResultRequest
	GetEndTime() *string
	SetSessionIds(v []*string) *GetDialogAnalysisResultRequest
	GetSessionIds() []*string
	SetStartTime(v string) *GetDialogAnalysisResultRequest
	GetStartTime() *string
	SetUseUrl(v bool) *GetDialogAnalysisResultRequest
	GetUseUrl() *bool
}

type GetDialogAnalysisResultRequest struct {
	// Whether to sort in ascending order. Default is true, which sorts by session creation time in ascending order. If false, sorts in descending order.
	//
	// example:
	//
	// true
	Asc *bool `json:"asc,omitempty" xml:"asc,omitempty"`
	// The end time, which must be in yyyy-MM-dd HH:mm:ss format. If sessionIds are provided, the system queries session analysis results based on these IDs.
	//
	// example:
	//
	// 2024-09-23 09:20:02
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Session ID list. When useUrl is true, the response includes OSS URLs. You can specify up to 1000 sessions. If you specify more than 1000, only the first 1000 are processed. When useUrl is false, the response includes full analysis results. You can specify up to 10 sessions. If you specify more than 10, only the first 10 are processed. This parameter is optional. If sessionIds is empty, the API retrieves results for sessions created between startTime and endTime. If sessionIds is not empty, the API retrieves results for the specified sessions. You cannot leave both sessionIds and the time range empty.
	SessionIds []*string `json:"sessionIds,omitempty" xml:"sessionIds,omitempty" type:"Repeated"`
	// Start time in yyyy-MM-dd HH:mm:ss format. If sessionIds is not empty, you can query the session analysis results using the specified session IDs.
	//
	// example:
	//
	// 2024-09-14 09:11:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// Whether to return an OSS URL instead of full analysis results. If true, the response includes an OSS URL that expires in one hour. Default is true. Supports up to 1000 sessions. If you specify more than 1000, only the first 1000 are processed. If false, the response includes full analysis results. Supports up to 10 sessions. If you specify more than 10, only the first 10 are processed.
	//
	// example:
	//
	// true
	UseUrl *bool `json:"useUrl,omitempty" xml:"useUrl,omitempty"`
}

func (s GetDialogAnalysisResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDialogAnalysisResultRequest) GoString() string {
	return s.String()
}

func (s *GetDialogAnalysisResultRequest) GetAsc() *bool {
	return s.Asc
}

func (s *GetDialogAnalysisResultRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetDialogAnalysisResultRequest) GetSessionIds() []*string {
	return s.SessionIds
}

func (s *GetDialogAnalysisResultRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetDialogAnalysisResultRequest) GetUseUrl() *bool {
	return s.UseUrl
}

func (s *GetDialogAnalysisResultRequest) SetAsc(v bool) *GetDialogAnalysisResultRequest {
	s.Asc = &v
	return s
}

func (s *GetDialogAnalysisResultRequest) SetEndTime(v string) *GetDialogAnalysisResultRequest {
	s.EndTime = &v
	return s
}

func (s *GetDialogAnalysisResultRequest) SetSessionIds(v []*string) *GetDialogAnalysisResultRequest {
	s.SessionIds = v
	return s
}

func (s *GetDialogAnalysisResultRequest) SetStartTime(v string) *GetDialogAnalysisResultRequest {
	s.StartTime = &v
	return s
}

func (s *GetDialogAnalysisResultRequest) SetUseUrl(v bool) *GetDialogAnalysisResultRequest {
	s.UseUrl = &v
	return s
}

func (s *GetDialogAnalysisResultRequest) Validate() error {
	return dara.Validate(s)
}
