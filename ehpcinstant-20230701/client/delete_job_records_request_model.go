// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v []*string) *DeleteJobRecordsRequest
	GetJobIds() []*string
}

type DeleteJobRecordsRequest struct {
	// The list of job IDs.
	JobIds []*string `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
}

func (s DeleteJobRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobRecordsRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobRecordsRequest) GetJobIds() []*string {
	return s.JobIds
}

func (s *DeleteJobRecordsRequest) SetJobIds(v []*string) *DeleteJobRecordsRequest {
	s.JobIds = v
	return s
}

func (s *DeleteJobRecordsRequest) Validate() error {
	return dara.Validate(s)
}
