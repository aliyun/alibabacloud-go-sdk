// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLatestSubmitDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLatestSubmitDetailResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetLatestSubmitDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetLatestSubmitDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLatestSubmitDetailResponseBody
	GetRequestId() *string
	SetSubmitDetailResult(v *GetLatestSubmitDetailResponseBodySubmitDetailResult) *GetLatestSubmitDetailResponseBody
	GetSubmitDetailResult() *GetLatestSubmitDetailResponseBodySubmitDetailResult
	SetSuccess(v bool) *GetLatestSubmitDetailResponseBody
	GetSuccess() *bool
}

type GetLatestSubmitDetailResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The pending submit details.
	SubmitDetailResult *GetLatestSubmitDetailResponseBodySubmitDetailResult `json:"SubmitDetailResult,omitempty" xml:"SubmitDetailResult,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLatestSubmitDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLatestSubmitDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetLatestSubmitDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLatestSubmitDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetLatestSubmitDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLatestSubmitDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLatestSubmitDetailResponseBody) GetSubmitDetailResult() *GetLatestSubmitDetailResponseBodySubmitDetailResult {
	return s.SubmitDetailResult
}

func (s *GetLatestSubmitDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLatestSubmitDetailResponseBody) SetCode(v string) *GetLatestSubmitDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBody) SetHttpStatusCode(v int32) *GetLatestSubmitDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBody) SetMessage(v string) *GetLatestSubmitDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBody) SetRequestId(v string) *GetLatestSubmitDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBody) SetSubmitDetailResult(v *GetLatestSubmitDetailResponseBodySubmitDetailResult) *GetLatestSubmitDetailResponseBody {
	s.SubmitDetailResult = v
	return s
}

func (s *GetLatestSubmitDetailResponseBody) SetSuccess(v bool) *GetLatestSubmitDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBody) Validate() error {
	if s.SubmitDetailResult != nil {
		if err := s.SubmitDetailResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLatestSubmitDetailResponseBodySubmitDetailResult struct {
	// The submit record ID.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The publish status. Valid values:
	//
	// - NOT_VERIFIED: not verified.
	//
	// - APPROVING: approval in progress.
	//
	// - APPROVE_FAILED: approval failed.
	//
	// - APPROVE_SUCCESS: approval succeeded.
	//
	// - APPROVE_REVOKED: approval revoked.
	//
	// - CHECKING: publish verification in progress.
	//
	// - CHECK_FAILED: publish verification failed.
	//
	// - CHECK_SUCCESS: publish verification succeeded.
	//
	// - PUBLISHING: publishing in progress.
	//
	// - FAILED: publish failed.
	//
	// - SUCCESS: publish succeeded.
	//
	// example:
	//
	// SUCCESS
	PublishStatus *string `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// The pending submit record object.
	ReleaseObject *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject `json:"ReleaseObject,omitempty" xml:"ReleaseObject,omitempty" type:"Struct"`
	// The submit status. Valid values:
	//
	// - TO_BE_PUBLISHED: pending.
	//
	// - PUBLISHED: published.
	//
	// - PUBLISHING: publishing.
	//
	// example:
	//
	// TO_BE_PUBLISHED
	SubmitStatus *string `json:"SubmitStatus,omitempty" xml:"SubmitStatus,omitempty"`
	// The record tag.
	//
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetLatestSubmitDetailResponseBodySubmitDetailResult) String() string {
	return dara.Prettify(s)
}

func (s GetLatestSubmitDetailResponseBodySubmitDetailResult) GoString() string {
	return s.String()
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResult) GetId() *int64 {
	return s.Id
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResult) GetPublishStatus() *string {
	return s.PublishStatus
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResult) GetReleaseObject() *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject {
	return s.ReleaseObject
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResult) GetSubmitStatus() *string {
	return s.SubmitStatus
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResult) GetTag() *string {
	return s.Tag
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResult) SetId(v int64) *GetLatestSubmitDetailResponseBodySubmitDetailResult {
	s.Id = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResult) SetPublishStatus(v string) *GetLatestSubmitDetailResponseBodySubmitDetailResult {
	s.PublishStatus = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResult) SetReleaseObject(v *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) *GetLatestSubmitDetailResponseBodySubmitDetailResult {
	s.ReleaseObject = v
	return s
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResult) SetSubmitStatus(v string) *GetLatestSubmitDetailResponseBodySubmitDetailResult {
	s.SubmitStatus = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResult) SetTag(v string) *GetLatestSubmitDetailResponseBodySubmitDetailResult {
	s.Tag = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResult) Validate() error {
	if s.ReleaseObject != nil {
		if err := s.ReleaseObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject struct {
	// The change type. Valid values: CREATE, UPDATE, and DELETE.
	//
	// example:
	//
	// DELETE
	ChangeType *string `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// The node ID.
	//
	// example:
	//
	// n_1234
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The object version.
	//
	// example:
	//
	// 1
	ObjectVersion *string `json:"ObjectVersion,omitempty" xml:"ObjectVersion,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The submit comment.
	//
	// example:
	//
	// 提交说明
	SubmitComment *string `json:"SubmitComment,omitempty" xml:"SubmitComment,omitempty"`
	// The submit object.
	SubmitObject *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject `json:"SubmitObject,omitempty" xml:"SubmitObject,omitempty" type:"Struct"`
}

func (s GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) String() string {
	return dara.Prettify(s)
}

func (s GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) GoString() string {
	return s.String()
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) GetChangeType() *string {
	return s.ChangeType
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) GetNodeId() *string {
	return s.NodeId
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) GetObjectVersion() *string {
	return s.ObjectVersion
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) GetSubmitComment() *string {
	return s.SubmitComment
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) GetSubmitObject() *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject {
	return s.SubmitObject
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) SetChangeType(v string) *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject {
	s.ChangeType = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) SetNodeId(v string) *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject {
	s.NodeId = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) SetObjectVersion(v string) *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject {
	s.ObjectVersion = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) SetProjectId(v int64) *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject {
	s.ProjectId = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) SetSubmitComment(v string) *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject {
	s.SubmitComment = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) SetSubmitObject(v *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject) *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject {
	s.SubmitObject = v
	return s
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObject) Validate() error {
	if s.SubmitObject != nil {
		if err := s.SubmitObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject struct {
	// The object ID.
	//
	// example:
	//
	// 1234
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The object name.
	//
	// example:
	//
	// abc
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The object type.
	//
	// example:
	//
	// MAX_COMPUTE_SQL
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject) String() string {
	return dara.Prettify(s)
}

func (s GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject) GoString() string {
	return s.String()
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject) GetObjectName() *string {
	return s.ObjectName
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject) SetObjectId(v string) *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject {
	s.ObjectId = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject) SetObjectName(v string) *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject {
	s.ObjectName = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject) SetObjectType(v string) *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject {
	s.ObjectType = &v
	return s
}

func (s *GetLatestSubmitDetailResponseBodySubmitDetailResultReleaseObjectSubmitObject) Validate() error {
	return dara.Validate(s)
}
