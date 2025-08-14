// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSnapshotJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetSnapshotJobRequest
	GetJobId() *string
}

type GetSnapshotJobRequest struct {
	// The ID of the snapshot job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetSnapshotJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotJobRequest) GoString() string {
	return s.String()
}

func (s *GetSnapshotJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetSnapshotJobRequest) SetJobId(v string) *GetSnapshotJobRequest {
	s.JobId = &v
	return s
}

func (s *GetSnapshotJobRequest) Validate() error {
	return dara.Validate(s)
}
