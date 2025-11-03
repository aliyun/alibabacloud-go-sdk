// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoBuildRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRecords(v []*ListRepoBuildRecordResponseBodyBuildRecords) *ListRepoBuildRecordResponseBody
	GetBuildRecords() []*ListRepoBuildRecordResponseBodyBuildRecords
	SetCode(v string) *ListRepoBuildRecordResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListRepoBuildRecordResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListRepoBuildRecordResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepoBuildRecordResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRepoBuildRecordResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListRepoBuildRecordResponseBody
	GetTotalCount() *string
}

type ListRepoBuildRecordResponseBody struct {
	// The list of image building records.
	BuildRecords []*ListRepoBuildRecordResponseBodyBuildRecords `json:"BuildRecords,omitempty" xml:"BuildRecords,omitempty" type:"Repeated"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9D23DEDF-E91D-434B-B7D5-9D12C648D166
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRepoBuildRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepoBuildRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordResponseBody) GetBuildRecords() []*ListRepoBuildRecordResponseBodyBuildRecords {
	return s.BuildRecords
}

func (s *ListRepoBuildRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRepoBuildRecordResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListRepoBuildRecordResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepoBuildRecordResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepoBuildRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepoBuildRecordResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListRepoBuildRecordResponseBody) SetBuildRecords(v []*ListRepoBuildRecordResponseBodyBuildRecords) *ListRepoBuildRecordResponseBody {
	s.BuildRecords = v
	return s
}

func (s *ListRepoBuildRecordResponseBody) SetCode(v string) *ListRepoBuildRecordResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoBuildRecordResponseBody) SetIsSuccess(v bool) *ListRepoBuildRecordResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoBuildRecordResponseBody) SetPageNo(v int32) *ListRepoBuildRecordResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepoBuildRecordResponseBody) SetPageSize(v int32) *ListRepoBuildRecordResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepoBuildRecordResponseBody) SetRequestId(v string) *ListRepoBuildRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoBuildRecordResponseBody) SetTotalCount(v string) *ListRepoBuildRecordResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRepoBuildRecordResponseBody) Validate() error {
	if s.BuildRecords != nil {
		for _, item := range s.BuildRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRepoBuildRecordResponseBodyBuildRecords struct {
	// The ID of the image building record.
	//
	// example:
	//
	// 537e08ab-735e-415f-b7c2-160eb87f8****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The status of the image building.
	//
	// example:
	//
	// SUCCESS
	BuildStatus *string `json:"BuildStatus,omitempty" xml:"BuildStatus,omitempty"`
	// The time when the image building ended.
	//
	// example:
	//
	// 1572875610000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The information about the image.
	Image *ListRepoBuildRecordResponseBodyBuildRecordsImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// The time when the image building started.
	//
	// example:
	//
	// 1572872207000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListRepoBuildRecordResponseBodyBuildRecords) String() string {
	return dara.Prettify(s)
}

func (s ListRepoBuildRecordResponseBodyBuildRecords) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) GetBuildRecordId() *string {
	return s.BuildRecordId
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) GetBuildStatus() *string {
	return s.BuildStatus
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) GetEndTime() *string {
	return s.EndTime
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) GetImage() *ListRepoBuildRecordResponseBodyBuildRecordsImage {
	return s.Image
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) GetStartTime() *string {
	return s.StartTime
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) SetBuildRecordId(v string) *ListRepoBuildRecordResponseBodyBuildRecords {
	s.BuildRecordId = &v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) SetBuildStatus(v string) *ListRepoBuildRecordResponseBodyBuildRecords {
	s.BuildStatus = &v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) SetEndTime(v string) *ListRepoBuildRecordResponseBodyBuildRecords {
	s.EndTime = &v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) SetImage(v *ListRepoBuildRecordResponseBodyBuildRecordsImage) *ListRepoBuildRecordResponseBodyBuildRecords {
	s.Image = v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) SetStartTime(v string) *ListRepoBuildRecordResponseBodyBuildRecords {
	s.StartTime = &v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) Validate() error {
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRepoBuildRecordResponseBodyBuildRecordsImage struct {
	// The tag of the image.
	//
	// example:
	//
	// v0.1
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The ID of the repository.
	//
	// example:
	//
	// crr-gzsrlevmvoaq****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s ListRepoBuildRecordResponseBodyBuildRecordsImage) String() string {
	return dara.Prettify(s)
}

func (s ListRepoBuildRecordResponseBodyBuildRecordsImage) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordResponseBodyBuildRecordsImage) GetImageTag() *string {
	return s.ImageTag
}

func (s *ListRepoBuildRecordResponseBodyBuildRecordsImage) GetRepoId() *string {
	return s.RepoId
}

func (s *ListRepoBuildRecordResponseBodyBuildRecordsImage) GetRepoName() *string {
	return s.RepoName
}

func (s *ListRepoBuildRecordResponseBodyBuildRecordsImage) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *ListRepoBuildRecordResponseBodyBuildRecordsImage) SetImageTag(v string) *ListRepoBuildRecordResponseBodyBuildRecordsImage {
	s.ImageTag = &v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecordsImage) SetRepoId(v string) *ListRepoBuildRecordResponseBodyBuildRecordsImage {
	s.RepoId = &v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecordsImage) SetRepoName(v string) *ListRepoBuildRecordResponseBodyBuildRecordsImage {
	s.RepoName = &v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecordsImage) SetRepoNamespaceName(v string) *ListRepoBuildRecordResponseBodyBuildRecordsImage {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecordsImage) Validate() error {
	return dara.Validate(s)
}
