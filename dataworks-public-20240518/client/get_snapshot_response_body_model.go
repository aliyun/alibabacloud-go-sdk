// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSnapshotResponseBody
	GetRequestId() *string
	SetSnapshot(v *GetSnapshotResponseBodySnapshot) *GetSnapshotResponseBody
	GetSnapshot() *GetSnapshotResponseBodySnapshot
}

type GetSnapshotResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0bc14115-1234-5678-ABCD-159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The snapshot.
	Snapshot *GetSnapshotResponseBodySnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Struct"`
}

func (s GetSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *GetSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSnapshotResponseBody) GetSnapshot() *GetSnapshotResponseBodySnapshot {
	return s.Snapshot
}

func (s *GetSnapshotResponseBody) SetRequestId(v string) *GetSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSnapshotResponseBody) SetSnapshot(v *GetSnapshotResponseBodySnapshot) *GetSnapshotResponseBody {
	s.Snapshot = v
	return s
}

func (s *GetSnapshotResponseBody) Validate() error {
	if s.Snapshot != nil {
		if err := s.Snapshot.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSnapshotResponseBodySnapshot struct {
	// The submit comment.
	//
	// example:
	//
	// snapshot comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The structured snapshot content. This parameter is mutually exclusive with ContentUrl.
	Content *GetSnapshotResponseBodySnapshotContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The OSS pre-signed download URL. This parameter is mutually exclusive with Content.
	//
	// example:
	//
	// https://oss.example.com/presigned
	ContentUrl *string `json:"ContentUrl,omitempty" xml:"ContentUrl,omitempty"`
	// The snapshot creation time in millisecond timestamp.
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

func (s GetSnapshotResponseBodySnapshot) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotResponseBodySnapshot) GoString() string {
	return s.String()
}

func (s *GetSnapshotResponseBodySnapshot) GetComment() *string {
	return s.Comment
}

func (s *GetSnapshotResponseBodySnapshot) GetContent() *GetSnapshotResponseBodySnapshotContent {
	return s.Content
}

func (s *GetSnapshotResponseBodySnapshot) GetContentUrl() *string {
	return s.ContentUrl
}

func (s *GetSnapshotResponseBodySnapshot) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetSnapshotResponseBodySnapshot) GetCreator() *string {
	return s.Creator
}

func (s *GetSnapshotResponseBodySnapshot) GetId() *string {
	return s.Id
}

func (s *GetSnapshotResponseBodySnapshot) GetNamespace() *string {
	return s.Namespace
}

func (s *GetSnapshotResponseBodySnapshot) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetSnapshotResponseBodySnapshot) GetType() *string {
	return s.Type
}

func (s *GetSnapshotResponseBodySnapshot) GetVersion() *int32 {
	return s.Version
}

func (s *GetSnapshotResponseBodySnapshot) SetComment(v string) *GetSnapshotResponseBodySnapshot {
	s.Comment = &v
	return s
}

func (s *GetSnapshotResponseBodySnapshot) SetContent(v *GetSnapshotResponseBodySnapshotContent) *GetSnapshotResponseBodySnapshot {
	s.Content = v
	return s
}

func (s *GetSnapshotResponseBodySnapshot) SetContentUrl(v string) *GetSnapshotResponseBodySnapshot {
	s.ContentUrl = &v
	return s
}

func (s *GetSnapshotResponseBodySnapshot) SetCreateTime(v int64) *GetSnapshotResponseBodySnapshot {
	s.CreateTime = &v
	return s
}

func (s *GetSnapshotResponseBodySnapshot) SetCreator(v string) *GetSnapshotResponseBodySnapshot {
	s.Creator = &v
	return s
}

func (s *GetSnapshotResponseBodySnapshot) SetId(v string) *GetSnapshotResponseBodySnapshot {
	s.Id = &v
	return s
}

func (s *GetSnapshotResponseBodySnapshot) SetNamespace(v string) *GetSnapshotResponseBodySnapshot {
	s.Namespace = &v
	return s
}

func (s *GetSnapshotResponseBodySnapshot) SetObjectId(v string) *GetSnapshotResponseBodySnapshot {
	s.ObjectId = &v
	return s
}

func (s *GetSnapshotResponseBodySnapshot) SetType(v string) *GetSnapshotResponseBodySnapshot {
	s.Type = &v
	return s
}

func (s *GetSnapshotResponseBodySnapshot) SetVersion(v int32) *GetSnapshotResponseBodySnapshot {
	s.Version = &v
	return s
}

func (s *GetSnapshotResponseBodySnapshot) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSnapshotResponseBodySnapshotContent struct {
	// The node script content.
	//
	// example:
	//
	// SELECT 1
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// FlowSpec JSON
	//
	// example:
	//
	// {"version":"1.1.0"}
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The stage code.
	//
	// example:
	//
	// DEV
	StageCode *string `json:"StageCode,omitempty" xml:"StageCode,omitempty"`
}

func (s GetSnapshotResponseBodySnapshotContent) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotResponseBodySnapshotContent) GoString() string {
	return s.String()
}

func (s *GetSnapshotResponseBodySnapshotContent) GetContent() *string {
	return s.Content
}

func (s *GetSnapshotResponseBodySnapshotContent) GetSpec() *string {
	return s.Spec
}

func (s *GetSnapshotResponseBodySnapshotContent) GetStageCode() *string {
	return s.StageCode
}

func (s *GetSnapshotResponseBodySnapshotContent) SetContent(v string) *GetSnapshotResponseBodySnapshotContent {
	s.Content = &v
	return s
}

func (s *GetSnapshotResponseBodySnapshotContent) SetSpec(v string) *GetSnapshotResponseBodySnapshotContent {
	s.Spec = &v
	return s
}

func (s *GetSnapshotResponseBodySnapshotContent) SetStageCode(v string) *GetSnapshotResponseBodySnapshotContent {
	s.StageCode = &v
	return s
}

func (s *GetSnapshotResponseBodySnapshotContent) Validate() error {
	return dara.Validate(s)
}
