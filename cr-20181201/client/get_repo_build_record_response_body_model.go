// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoBuildRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRecordId(v string) *GetRepoBuildRecordResponseBody
	GetBuildRecordId() *string
	SetCode(v string) *GetRepoBuildRecordResponseBody
	GetCode() *string
	SetEndTime(v int64) *GetRepoBuildRecordResponseBody
	GetEndTime() *int64
	SetImage(v *GetRepoBuildRecordResponseBodyImage) *GetRepoBuildRecordResponseBody
	GetImage() *GetRepoBuildRecordResponseBodyImage
	SetIsSuccess(v bool) *GetRepoBuildRecordResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetRepoBuildRecordResponseBody
	GetRequestId() *string
	SetStartTime(v int64) *GetRepoBuildRecordResponseBody
	GetStartTime() *int64
	SetStatus(v string) *GetRepoBuildRecordResponseBody
	GetStatus() *string
}

type GetRepoBuildRecordResponseBody struct {
	// The ID of the image building record.
	//
	// example:
	//
	// 79174CBA-8556-443A-8976-22C922D7****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the image building was completed.
	//
	// example:
	//
	// 1568718698000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The information about the image.
	Image *GetRepoBuildRecordResponseBodyImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// a78ec6fb-16ea-4649-93b7-f52afba7d9de1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the image building started.
	//
	// example:
	//
	// 1568718468000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the instance.
	//
	// example:
	//
	// true
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRepoBuildRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRepoBuildRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoBuildRecordResponseBody) GetBuildRecordId() *string {
	return s.BuildRecordId
}

func (s *GetRepoBuildRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRepoBuildRecordResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetRepoBuildRecordResponseBody) GetImage() *GetRepoBuildRecordResponseBodyImage {
	return s.Image
}

func (s *GetRepoBuildRecordResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetRepoBuildRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRepoBuildRecordResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetRepoBuildRecordResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetRepoBuildRecordResponseBody) SetBuildRecordId(v string) *GetRepoBuildRecordResponseBody {
	s.BuildRecordId = &v
	return s
}

func (s *GetRepoBuildRecordResponseBody) SetCode(v string) *GetRepoBuildRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoBuildRecordResponseBody) SetEndTime(v int64) *GetRepoBuildRecordResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetRepoBuildRecordResponseBody) SetImage(v *GetRepoBuildRecordResponseBodyImage) *GetRepoBuildRecordResponseBody {
	s.Image = v
	return s
}

func (s *GetRepoBuildRecordResponseBody) SetIsSuccess(v bool) *GetRepoBuildRecordResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoBuildRecordResponseBody) SetRequestId(v string) *GetRepoBuildRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepoBuildRecordResponseBody) SetStartTime(v int64) *GetRepoBuildRecordResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetRepoBuildRecordResponseBody) SetStatus(v string) *GetRepoBuildRecordResponseBody {
	s.Status = &v
	return s
}

func (s *GetRepoBuildRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRepoBuildRecordResponseBodyImage struct {
	// The tag of the image.
	//
	// example:
	//
	// master
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the image repository belongs.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s GetRepoBuildRecordResponseBodyImage) String() string {
	return dara.Prettify(s)
}

func (s GetRepoBuildRecordResponseBodyImage) GoString() string {
	return s.String()
}

func (s *GetRepoBuildRecordResponseBodyImage) GetImageTag() *string {
	return s.ImageTag
}

func (s *GetRepoBuildRecordResponseBodyImage) GetRepoName() *string {
	return s.RepoName
}

func (s *GetRepoBuildRecordResponseBodyImage) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *GetRepoBuildRecordResponseBodyImage) SetImageTag(v string) *GetRepoBuildRecordResponseBodyImage {
	s.ImageTag = &v
	return s
}

func (s *GetRepoBuildRecordResponseBodyImage) SetRepoName(v string) *GetRepoBuildRecordResponseBodyImage {
	s.RepoName = &v
	return s
}

func (s *GetRepoBuildRecordResponseBodyImage) SetRepoNamespaceName(v string) *GetRepoBuildRecordResponseBodyImage {
	s.RepoNamespaceName = &v
	return s
}

func (s *GetRepoBuildRecordResponseBodyImage) Validate() error {
	return dara.Validate(s)
}
