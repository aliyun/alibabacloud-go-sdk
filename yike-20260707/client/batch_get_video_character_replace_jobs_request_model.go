// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetVideoCharacterReplaceJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *BatchGetVideoCharacterReplaceJobsRequest
	GetJobIds() *string
}

type BatchGetVideoCharacterReplaceJobsRequest struct {
	// The list of job IDs. Separate multiple IDs with commas.
	//
	// example:
	//
	// ag_83baeea38b534ee1937b8a4acd08****,ag_f7eeed5d5e294dfba4bb347177e1****
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
}

func (s BatchGetVideoCharacterReplaceJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetVideoCharacterReplaceJobsRequest) GoString() string {
	return s.String()
}

func (s *BatchGetVideoCharacterReplaceJobsRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *BatchGetVideoCharacterReplaceJobsRequest) SetJobIds(v string) *BatchGetVideoCharacterReplaceJobsRequest {
	s.JobIds = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsRequest) Validate() error {
	return dara.Validate(s)
}
