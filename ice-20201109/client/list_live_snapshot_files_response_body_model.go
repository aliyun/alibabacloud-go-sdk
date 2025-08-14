// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveSnapshotFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileList(v []*ListLiveSnapshotFilesResponseBodyFileList) *ListLiveSnapshotFilesResponseBody
	GetFileList() []*ListLiveSnapshotFilesResponseBodyFileList
	SetNextStartTime(v string) *ListLiveSnapshotFilesResponseBody
	GetNextStartTime() *string
	SetRequestId(v string) *ListLiveSnapshotFilesResponseBody
	GetRequestId() *string
}

type ListLiveSnapshotFilesResponseBody struct {
	// The list of files.
	FileList []*ListLiveSnapshotFilesResponseBodyFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// The start time of the next page. If no value is returned, the pagination ends.
	//
	// example:
	//
	// 2022-02-02T22:22:22Z
	NextStartTime *string `json:"NextStartTime,omitempty" xml:"NextStartTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLiveSnapshotFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveSnapshotFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveSnapshotFilesResponseBody) GetFileList() []*ListLiveSnapshotFilesResponseBodyFileList {
	return s.FileList
}

func (s *ListLiveSnapshotFilesResponseBody) GetNextStartTime() *string {
	return s.NextStartTime
}

func (s *ListLiveSnapshotFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveSnapshotFilesResponseBody) SetFileList(v []*ListLiveSnapshotFilesResponseBodyFileList) *ListLiveSnapshotFilesResponseBody {
	s.FileList = v
	return s
}

func (s *ListLiveSnapshotFilesResponseBody) SetNextStartTime(v string) *ListLiveSnapshotFilesResponseBody {
	s.NextStartTime = &v
	return s
}

func (s *ListLiveSnapshotFilesResponseBody) SetRequestId(v string) *ListLiveSnapshotFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveSnapshotFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLiveSnapshotFilesResponseBodyFileList struct {
	// The time when the template was created.
	//
	// example:
	//
	// 2022-02-02T22:22:22Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creation timestamp that is used as an input parameter for a delete API operation.
	//
	// example:
	//
	// 1619503516000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// Specifies whether to overlay snapshots.
	//
	// example:
	//
	// true
	IsOverlay *bool `json:"IsOverlay,omitempty" xml:"IsOverlay,omitempty"`
	// The OSS bucket.
	//
	// example:
	//
	// testbucket
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The Object Storage Service (OSS) domain name.
	//
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// The location in which the OSS object is stored.
	OssObject *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
}

func (s ListLiveSnapshotFilesResponseBodyFileList) String() string {
	return dara.Prettify(s)
}

func (s ListLiveSnapshotFilesResponseBodyFileList) GoString() string {
	return s.String()
}

func (s *ListLiveSnapshotFilesResponseBodyFileList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLiveSnapshotFilesResponseBodyFileList) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListLiveSnapshotFilesResponseBodyFileList) GetIsOverlay() *bool {
	return s.IsOverlay
}

func (s *ListLiveSnapshotFilesResponseBodyFileList) GetOssBucket() *string {
	return s.OssBucket
}

func (s *ListLiveSnapshotFilesResponseBodyFileList) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *ListLiveSnapshotFilesResponseBodyFileList) GetOssObject() *string {
	return s.OssObject
}

func (s *ListLiveSnapshotFilesResponseBodyFileList) SetCreateTime(v string) *ListLiveSnapshotFilesResponseBodyFileList {
	s.CreateTime = &v
	return s
}

func (s *ListLiveSnapshotFilesResponseBodyFileList) SetCreateTimestamp(v int64) *ListLiveSnapshotFilesResponseBodyFileList {
	s.CreateTimestamp = &v
	return s
}

func (s *ListLiveSnapshotFilesResponseBodyFileList) SetIsOverlay(v bool) *ListLiveSnapshotFilesResponseBodyFileList {
	s.IsOverlay = &v
	return s
}

func (s *ListLiveSnapshotFilesResponseBodyFileList) SetOssBucket(v string) *ListLiveSnapshotFilesResponseBodyFileList {
	s.OssBucket = &v
	return s
}

func (s *ListLiveSnapshotFilesResponseBodyFileList) SetOssEndpoint(v string) *ListLiveSnapshotFilesResponseBodyFileList {
	s.OssEndpoint = &v
	return s
}

func (s *ListLiveSnapshotFilesResponseBodyFileList) SetOssObject(v string) *ListLiveSnapshotFilesResponseBodyFileList {
	s.OssObject = &v
	return s
}

func (s *ListLiveSnapshotFilesResponseBodyFileList) Validate() error {
	return dara.Validate(s)
}
