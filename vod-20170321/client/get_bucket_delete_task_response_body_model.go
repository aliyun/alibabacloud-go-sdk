// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketDeleteTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetBucketDeleteTaskResponseBody
	GetRequestId() *string
	SetTaskInfo(v *GetBucketDeleteTaskResponseBodyTaskInfo) *GetBucketDeleteTaskResponseBody
	GetTaskInfo() *GetBucketDeleteTaskResponseBodyTaskInfo
}

type GetBucketDeleteTaskResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskInfo  *GetBucketDeleteTaskResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s GetBucketDeleteTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketDeleteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketDeleteTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBucketDeleteTaskResponseBody) GetTaskInfo() *GetBucketDeleteTaskResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *GetBucketDeleteTaskResponseBody) SetRequestId(v string) *GetBucketDeleteTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBucketDeleteTaskResponseBody) SetTaskInfo(v *GetBucketDeleteTaskResponseBodyTaskInfo) *GetBucketDeleteTaskResponseBody {
	s.TaskInfo = v
	return s
}

func (s *GetBucketDeleteTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBucketDeleteTaskResponseBodyTaskInfo struct {
	AttachedMediaRemain *int64  `json:"AttachedMediaRemain,omitempty" xml:"AttachedMediaRemain,omitempty"`
	CreationTime        *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DeleteFiles         *bool   `json:"DeleteFiles,omitempty" xml:"DeleteFiles,omitempty"`
	ImageRemain         *int64  `json:"ImageRemain,omitempty" xml:"ImageRemain,omitempty"`
	ModificationTime    *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageLocation     *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	StorageSize         *int64  `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	VideoRemain         *int64  `json:"VideoRemain,omitempty" xml:"VideoRemain,omitempty"`
}

func (s GetBucketDeleteTaskResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBucketDeleteTaskResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) GetAttachedMediaRemain() *int64 {
	return s.AttachedMediaRemain
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) GetDeleteFiles() *bool {
	return s.DeleteFiles
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) GetImageRemain() *int64 {
	return s.ImageRemain
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) GetVideoRemain() *int64 {
	return s.VideoRemain
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) SetAttachedMediaRemain(v int64) *GetBucketDeleteTaskResponseBodyTaskInfo {
	s.AttachedMediaRemain = &v
	return s
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) SetCreationTime(v string) *GetBucketDeleteTaskResponseBodyTaskInfo {
	s.CreationTime = &v
	return s
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) SetDeleteFiles(v bool) *GetBucketDeleteTaskResponseBodyTaskInfo {
	s.DeleteFiles = &v
	return s
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) SetImageRemain(v int64) *GetBucketDeleteTaskResponseBodyTaskInfo {
	s.ImageRemain = &v
	return s
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) SetModificationTime(v string) *GetBucketDeleteTaskResponseBodyTaskInfo {
	s.ModificationTime = &v
	return s
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) SetStatus(v string) *GetBucketDeleteTaskResponseBodyTaskInfo {
	s.Status = &v
	return s
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) SetStorageLocation(v string) *GetBucketDeleteTaskResponseBodyTaskInfo {
	s.StorageLocation = &v
	return s
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) SetStorageSize(v int64) *GetBucketDeleteTaskResponseBodyTaskInfo {
	s.StorageSize = &v
	return s
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) SetVideoRemain(v int64) *GetBucketDeleteTaskResponseBodyTaskInfo {
	s.VideoRemain = &v
	return s
}

func (s *GetBucketDeleteTaskResponseBodyTaskInfo) Validate() error {
	return dara.Validate(s)
}
