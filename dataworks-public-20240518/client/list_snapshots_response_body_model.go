// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListSnapshotsResponseBodyPagingInfo) *ListSnapshotsResponseBody
	GetPagingInfo() *ListSnapshotsResponseBodyPagingInfo
	SetRequestId(v string) *ListSnapshotsResponseBody
	GetRequestId() *string
}

type ListSnapshotsResponseBody struct {
	// The pagination information.
	PagingInfo *ListSnapshotsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 0bc14115-1234-5678-ABCD-159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBody) GetPagingInfo() *ListSnapshotsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSnapshotsResponseBody) SetPagingInfo(v *ListSnapshotsResponseBodyPagingInfo) *ListSnapshotsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListSnapshotsResponseBody) SetRequestId(v string) *ListSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSnapshotsResponseBodyPagingInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of snapshots.
	Snapshots []*ListSnapshotsResponseBodyPagingInfoSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSnapshotsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSnapshotsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSnapshotsResponseBodyPagingInfo) GetSnapshots() []*ListSnapshotsResponseBodyPagingInfoSnapshots {
	return s.Snapshots
}

func (s *ListSnapshotsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSnapshotsResponseBodyPagingInfo) SetPageNumber(v int32) *ListSnapshotsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotsResponseBodyPagingInfo) SetPageSize(v int32) *ListSnapshotsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotsResponseBodyPagingInfo) SetSnapshots(v []*ListSnapshotsResponseBodyPagingInfoSnapshots) *ListSnapshotsResponseBodyPagingInfo {
	s.Snapshots = v
	return s
}

func (s *ListSnapshotsResponseBodyPagingInfo) SetTotalCount(v int32) *ListSnapshotsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListSnapshotsResponseBodyPagingInfo) Validate() error {
	if s.Snapshots != nil {
		for _, item := range s.Snapshots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSnapshotsResponseBodyPagingInfoSnapshots struct {
	// The snapshot comment.
	//
	// example:
	//
	// snapshot comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The snapshot creation time, in millisecond UNIX timestamp.
	//
	// example:
	//
	// 1782370983000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The employee ID of the submitter.
	//
	// example:
	//
	// 209508679263509059
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The unique ID of the snapshot.
	//
	// example:
	//
	// 8b723a9e8bd443af920b77e39aeb4f63
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The namespace.
	//
	// example:
	//
	// 1389623
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The unique ID of the object to which the snapshot belongs.
	//
	// example:
	//
	// 8467231038932407294
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The snapshot type.
	//
	// example:
	//
	// Saved
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 2
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListSnapshotsResponseBodyPagingInfoSnapshots) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponseBodyPagingInfoSnapshots) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) GetComment() *string {
	return s.Comment
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) GetCreator() *string {
	return s.Creator
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) GetId() *string {
	return s.Id
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) GetNamespace() *string {
	return s.Namespace
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) GetType() *string {
	return s.Type
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) GetVersion() *int32 {
	return s.Version
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) SetComment(v string) *ListSnapshotsResponseBodyPagingInfoSnapshots {
	s.Comment = &v
	return s
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) SetCreateTime(v int64) *ListSnapshotsResponseBodyPagingInfoSnapshots {
	s.CreateTime = &v
	return s
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) SetCreator(v string) *ListSnapshotsResponseBodyPagingInfoSnapshots {
	s.Creator = &v
	return s
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) SetId(v string) *ListSnapshotsResponseBodyPagingInfoSnapshots {
	s.Id = &v
	return s
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) SetNamespace(v string) *ListSnapshotsResponseBodyPagingInfoSnapshots {
	s.Namespace = &v
	return s
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) SetObjectId(v string) *ListSnapshotsResponseBodyPagingInfoSnapshots {
	s.ObjectId = &v
	return s
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) SetType(v string) *ListSnapshotsResponseBodyPagingInfoSnapshots {
	s.Type = &v
	return s
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) SetVersion(v int32) *ListSnapshotsResponseBodyPagingInfoSnapshots {
	s.Version = &v
	return s
}

func (s *ListSnapshotsResponseBodyPagingInfoSnapshots) Validate() error {
	return dara.Validate(s)
}
