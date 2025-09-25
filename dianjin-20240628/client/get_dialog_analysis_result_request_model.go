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
	// example:
	//
	// true
	Asc *bool `json:"asc,omitempty" xml:"asc,omitempty"`
	// example:
	//
	// 2024-09-23 09:20:02
	EndTime    *string   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	SessionIds []*string `json:"sessionIds,omitempty" xml:"sessionIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-14 09:11:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
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
