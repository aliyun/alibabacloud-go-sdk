// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobHistoryResp interface {
	dara.Model
	String() string
	GoString() string
	SetJobHistory(v []*JobHistory) *ListJobHistoryResp
	GetJobHistory() []*JobHistory
	SetNextMarker(v string) *ListJobHistoryResp
	GetNextMarker() *string
	SetTruncated(v string) *ListJobHistoryResp
	GetTruncated() *string
}

type ListJobHistoryResp struct {
	JobHistory []*JobHistory `json:"JobHistory,omitempty" xml:"JobHistory,omitempty" type:"Repeated"`
	// example:
	//
	// 1#3
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// example:
	//
	// true
	Truncated *string `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s ListJobHistoryResp) String() string {
	return dara.Prettify(s)
}

func (s ListJobHistoryResp) GoString() string {
	return s.String()
}

func (s *ListJobHistoryResp) GetJobHistory() []*JobHistory {
	return s.JobHistory
}

func (s *ListJobHistoryResp) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListJobHistoryResp) GetTruncated() *string {
	return s.Truncated
}

func (s *ListJobHistoryResp) SetJobHistory(v []*JobHistory) *ListJobHistoryResp {
	s.JobHistory = v
	return s
}

func (s *ListJobHistoryResp) SetNextMarker(v string) *ListJobHistoryResp {
	s.NextMarker = &v
	return s
}

func (s *ListJobHistoryResp) SetTruncated(v string) *ListJobHistoryResp {
	s.Truncated = &v
	return s
}

func (s *ListJobHistoryResp) Validate() error {
	if s.JobHistory != nil {
		for _, item := range s.JobHistory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
