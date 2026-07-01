// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSnapshotFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteFileResultList(v []*DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList) *DeleteLiveSnapshotFilesResponseBody
	GetDeleteFileResultList() []*DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList
	SetRequestId(v string) *DeleteLiveSnapshotFilesResponseBody
	GetRequestId() *string
}

type DeleteLiveSnapshotFilesResponseBody struct {
	// An array of deletion results.
	DeleteFileResultList []*DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList `json:"DeleteFileResultList,omitempty" xml:"DeleteFileResultList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// ****2876-6263-4B75-8F2C-CD0F7FCF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveSnapshotFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSnapshotFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotFilesResponseBody) GetDeleteFileResultList() []*DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList {
	return s.DeleteFileResultList
}

func (s *DeleteLiveSnapshotFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveSnapshotFilesResponseBody) SetDeleteFileResultList(v []*DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList) *DeleteLiveSnapshotFilesResponseBody {
	s.DeleteFileResultList = v
	return s
}

func (s *DeleteLiveSnapshotFilesResponseBody) SetRequestId(v string) *DeleteLiveSnapshotFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveSnapshotFilesResponseBody) Validate() error {
	if s.DeleteFileResultList != nil {
		for _, item := range s.DeleteFileResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList struct {
	// The creation timestamp of the file.
	//
	// example:
	//
	// 1660638613798
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The deletion result. A value of `OK` indicates the operation succeeded. Other values indicate that it failed.
	//
	// example:
	//
	// OK
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList) GetResult() *string {
	return s.Result
}

func (s *DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList) SetCreateTimestamp(v int64) *DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList {
	s.CreateTimestamp = &v
	return s
}

func (s *DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList) SetResult(v string) *DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList {
	s.Result = &v
	return s
}

func (s *DeleteLiveSnapshotFilesResponseBodyDeleteFileResultList) Validate() error {
	return dara.Validate(s)
}
