// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSnapshotFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimestampList(v []*int64) *DeleteLiveSnapshotFilesRequest
	GetCreateTimestampList() []*int64
	SetDeleteOriginalFile(v bool) *DeleteLiveSnapshotFilesRequest
	GetDeleteOriginalFile() *bool
	SetJobId(v string) *DeleteLiveSnapshotFilesRequest
	GetJobId() *string
}

type DeleteLiveSnapshotFilesRequest struct {
	// The list of timestamps when the jobs were created. The values are UNIX timestamps representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. A maximum of 200 jobs can be deleted at a time.
	//
	// This parameter is required.
	CreateTimestampList []*int64 `json:"CreateTimestampList,omitempty" xml:"CreateTimestampList,omitempty" type:"Repeated"`
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

func (s DeleteLiveSnapshotFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSnapshotFilesRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotFilesRequest) GetCreateTimestampList() []*int64 {
	return s.CreateTimestampList
}

func (s *DeleteLiveSnapshotFilesRequest) GetDeleteOriginalFile() *bool {
	return s.DeleteOriginalFile
}

func (s *DeleteLiveSnapshotFilesRequest) GetJobId() *string {
	return s.JobId
}

func (s *DeleteLiveSnapshotFilesRequest) SetCreateTimestampList(v []*int64) *DeleteLiveSnapshotFilesRequest {
	s.CreateTimestampList = v
	return s
}

func (s *DeleteLiveSnapshotFilesRequest) SetDeleteOriginalFile(v bool) *DeleteLiveSnapshotFilesRequest {
	s.DeleteOriginalFile = &v
	return s
}

func (s *DeleteLiveSnapshotFilesRequest) SetJobId(v string) *DeleteLiveSnapshotFilesRequest {
	s.JobId = &v
	return s
}

func (s *DeleteLiveSnapshotFilesRequest) Validate() error {
	return dara.Validate(s)
}
