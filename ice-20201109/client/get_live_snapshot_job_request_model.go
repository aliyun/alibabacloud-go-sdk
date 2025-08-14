// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveSnapshotJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetLiveSnapshotJobRequest
	GetJobId() *string
}

type GetLiveSnapshotJobRequest struct {
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetLiveSnapshotJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLiveSnapshotJobRequest) GoString() string {
	return s.String()
}

func (s *GetLiveSnapshotJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetLiveSnapshotJobRequest) SetJobId(v string) *GetLiveSnapshotJobRequest {
	s.JobId = &v
	return s
}

func (s *GetLiveSnapshotJobRequest) Validate() error {
	return dara.Validate(s)
}
