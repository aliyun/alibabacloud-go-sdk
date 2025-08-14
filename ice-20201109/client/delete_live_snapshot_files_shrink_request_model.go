// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSnapshotFilesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimestampListShrink(v string) *DeleteLiveSnapshotFilesShrinkRequest
	GetCreateTimestampListShrink() *string
	SetDeleteOriginalFile(v bool) *DeleteLiveSnapshotFilesShrinkRequest
	GetDeleteOriginalFile() *bool
	SetJobId(v string) *DeleteLiveSnapshotFilesShrinkRequest
	GetJobId() *string
}

type DeleteLiveSnapshotFilesShrinkRequest struct {
	// The list of timestamps when the jobs were created. The values are UNIX timestamps representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. A maximum of 200 jobs can be deleted at a time.
	//
	// This parameter is required.
	CreateTimestampListShrink *string `json:"CreateTimestampList,omitempty" xml:"CreateTimestampList,omitempty"`
	// Specifies whether to delete the original files at the same time. Default value: false.
	//
	// example:
	//
	// true
	DeleteOriginalFile *bool `json:"DeleteOriginalFile,omitempty" xml:"DeleteOriginalFile,omitempty"`
	// The ID of the snapshot job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteLiveSnapshotFilesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSnapshotFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotFilesShrinkRequest) GetCreateTimestampListShrink() *string {
	return s.CreateTimestampListShrink
}

func (s *DeleteLiveSnapshotFilesShrinkRequest) GetDeleteOriginalFile() *bool {
	return s.DeleteOriginalFile
}

func (s *DeleteLiveSnapshotFilesShrinkRequest) GetJobId() *string {
	return s.JobId
}

func (s *DeleteLiveSnapshotFilesShrinkRequest) SetCreateTimestampListShrink(v string) *DeleteLiveSnapshotFilesShrinkRequest {
	s.CreateTimestampListShrink = &v
	return s
}

func (s *DeleteLiveSnapshotFilesShrinkRequest) SetDeleteOriginalFile(v bool) *DeleteLiveSnapshotFilesShrinkRequest {
	s.DeleteOriginalFile = &v
	return s
}

func (s *DeleteLiveSnapshotFilesShrinkRequest) SetJobId(v string) *DeleteLiveSnapshotFilesShrinkRequest {
	s.JobId = &v
	return s
}

func (s *DeleteLiveSnapshotFilesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
