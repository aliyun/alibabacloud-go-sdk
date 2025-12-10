// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobHistoryList(v *ListJobHistoryResp) *ListJobHistoryResponseBody
	GetJobHistoryList() *ListJobHistoryResp
}

type ListJobHistoryResponseBody struct {
	// The running history of the migration task.
	JobHistoryList *ListJobHistoryResp `json:"JobHistoryList,omitempty" xml:"JobHistoryList,omitempty"`
}

func (s ListJobHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobHistoryResponseBody) GetJobHistoryList() *ListJobHistoryResp {
	return s.JobHistoryList
}

func (s *ListJobHistoryResponseBody) SetJobHistoryList(v *ListJobHistoryResp) *ListJobHistoryResponseBody {
	s.JobHistoryList = v
	return s
}

func (s *ListJobHistoryResponseBody) Validate() error {
	if s.JobHistoryList != nil {
		if err := s.JobHistoryList.Validate(); err != nil {
			return err
		}
	}
	return nil
}
