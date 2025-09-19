// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDockerhubImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageList(v []*ListDockerhubImageResponseBodyImageList) *ListDockerhubImageResponseBody
	GetImageList() []*ListDockerhubImageResponseBodyImageList
	SetRequestId(v string) *ListDockerhubImageResponseBody
	GetRequestId() *string
}

type ListDockerhubImageResponseBody struct {
	// The information about the images.
	ImageList []*ListDockerhubImageResponseBodyImageList `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDockerhubImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDockerhubImageResponseBody) GoString() string {
	return s.String()
}

func (s *ListDockerhubImageResponseBody) GetImageList() []*ListDockerhubImageResponseBodyImageList {
	return s.ImageList
}

func (s *ListDockerhubImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDockerhubImageResponseBody) SetImageList(v []*ListDockerhubImageResponseBodyImageList) *ListDockerhubImageResponseBody {
	s.ImageList = v
	return s
}

func (s *ListDockerhubImageResponseBody) SetRequestId(v string) *ListDockerhubImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDockerhubImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDockerhubImageResponseBodyImageList struct {
	// The digest value of the image.
	//
	// example:
	//
	// 5ffded22661b0f1e9c7fcccb0d488cff*****f8c52a819bd7179ef3e4a041988
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The number of baseline risks.
	//
	// example:
	//
	// 0
	HcCount *int32 `json:"HcCount,omitempty" xml:"HcCount,omitempty"`
	// The image ID.
	//
	// example:
	//
	// d943de1933650d74b415d3ae8b37c064a0e0c700574d7a949c26db3291******
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The size of the image. Unit: bytes.
	//
	// example:
	//
	// 1024
	ImageSize *int64 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// python
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// python
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The risk details of the image.
	//
	// example:
	//
	// {"vul":0}
	RiskLevelDetail *string `json:"RiskLevelDetail,omitempty" xml:"RiskLevelDetail,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// 3.9
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The UUID of the image.
	//
	// example:
	//
	// a9b50827-801f-414c-900d-c4a223d*****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The number of detected vulnerabilities.
	//
	// example:
	//
	// 0
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
}

func (s ListDockerhubImageResponseBodyImageList) String() string {
	return dara.Prettify(s)
}

func (s ListDockerhubImageResponseBodyImageList) GoString() string {
	return s.String()
}

func (s *ListDockerhubImageResponseBodyImageList) GetDigest() *string {
	return s.Digest
}

func (s *ListDockerhubImageResponseBodyImageList) GetHcCount() *int32 {
	return s.HcCount
}

func (s *ListDockerhubImageResponseBodyImageList) GetImageId() *string {
	return s.ImageId
}

func (s *ListDockerhubImageResponseBodyImageList) GetImageSize() *int64 {
	return s.ImageSize
}

func (s *ListDockerhubImageResponseBodyImageList) GetRepoName() *string {
	return s.RepoName
}

func (s *ListDockerhubImageResponseBodyImageList) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *ListDockerhubImageResponseBodyImageList) GetRiskLevelDetail() *string {
	return s.RiskLevelDetail
}

func (s *ListDockerhubImageResponseBodyImageList) GetTag() *string {
	return s.Tag
}

func (s *ListDockerhubImageResponseBodyImageList) GetUuid() *string {
	return s.Uuid
}

func (s *ListDockerhubImageResponseBodyImageList) GetVulCount() *int32 {
	return s.VulCount
}

func (s *ListDockerhubImageResponseBodyImageList) SetDigest(v string) *ListDockerhubImageResponseBodyImageList {
	s.Digest = &v
	return s
}

func (s *ListDockerhubImageResponseBodyImageList) SetHcCount(v int32) *ListDockerhubImageResponseBodyImageList {
	s.HcCount = &v
	return s
}

func (s *ListDockerhubImageResponseBodyImageList) SetImageId(v string) *ListDockerhubImageResponseBodyImageList {
	s.ImageId = &v
	return s
}

func (s *ListDockerhubImageResponseBodyImageList) SetImageSize(v int64) *ListDockerhubImageResponseBodyImageList {
	s.ImageSize = &v
	return s
}

func (s *ListDockerhubImageResponseBodyImageList) SetRepoName(v string) *ListDockerhubImageResponseBodyImageList {
	s.RepoName = &v
	return s
}

func (s *ListDockerhubImageResponseBodyImageList) SetRepoNamespace(v string) *ListDockerhubImageResponseBodyImageList {
	s.RepoNamespace = &v
	return s
}

func (s *ListDockerhubImageResponseBodyImageList) SetRiskLevelDetail(v string) *ListDockerhubImageResponseBodyImageList {
	s.RiskLevelDetail = &v
	return s
}

func (s *ListDockerhubImageResponseBodyImageList) SetTag(v string) *ListDockerhubImageResponseBodyImageList {
	s.Tag = &v
	return s
}

func (s *ListDockerhubImageResponseBodyImageList) SetUuid(v string) *ListDockerhubImageResponseBodyImageList {
	s.Uuid = &v
	return s
}

func (s *ListDockerhubImageResponseBodyImageList) SetVulCount(v int32) *ListDockerhubImageResponseBodyImageList {
	s.VulCount = &v
	return s
}

func (s *ListDockerhubImageResponseBodyImageList) Validate() error {
	return dara.Validate(s)
}
