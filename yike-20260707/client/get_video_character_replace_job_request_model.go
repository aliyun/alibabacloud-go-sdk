// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoCharacterReplaceJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetVideoCharacterReplaceJobRequest
	GetJobId() *string
}

type GetVideoCharacterReplaceJobRequest struct {
	// The task ID, which is returned by SubmitVideoCharacterReplaceJob.
	//
	// example:
	//
	// ag_83baeea38b534ee1937b8a4acd08****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetVideoCharacterReplaceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCharacterReplaceJobRequest) GoString() string {
	return s.String()
}

func (s *GetVideoCharacterReplaceJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetVideoCharacterReplaceJobRequest) SetJobId(v string) *GetVideoCharacterReplaceJobRequest {
	s.JobId = &v
	return s
}

func (s *GetVideoCharacterReplaceJobRequest) Validate() error {
	return dara.Validate(s)
}
