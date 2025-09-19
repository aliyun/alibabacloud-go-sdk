// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJenkinsImageScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDigest(v string) *CreateJenkinsImageScanTaskRequest
	GetDigest() *string
	SetImageCreate(v int64) *CreateJenkinsImageScanTaskRequest
	GetImageCreate() *int64
	SetImageId(v string) *CreateJenkinsImageScanTaskRequest
	GetImageId() *string
	SetImageSize(v int64) *CreateJenkinsImageScanTaskRequest
	GetImageSize() *int64
	SetImageUpdate(v int64) *CreateJenkinsImageScanTaskRequest
	GetImageUpdate() *int64
	SetJenkinsEnv(v string) *CreateJenkinsImageScanTaskRequest
	GetJenkinsEnv() *string
	SetNamespace(v string) *CreateJenkinsImageScanTaskRequest
	GetNamespace() *string
	SetRepoName(v string) *CreateJenkinsImageScanTaskRequest
	GetRepoName() *string
	SetSourceIp(v string) *CreateJenkinsImageScanTaskRequest
	GetSourceIp() *string
	SetTag(v string) *CreateJenkinsImageScanTaskRequest
	GetTag() *string
	SetToken(v string) *CreateJenkinsImageScanTaskRequest
	GetToken() *string
	SetUuid(v string) *CreateJenkinsImageScanTaskRequest
	GetUuid() *string
}

type CreateJenkinsImageScanTaskRequest struct {
	// The digest of the image.
	//
	// example:
	//
	// a8c9f3765684cd8d9053db9523eab58878e99a199217500efd9ae2a860a7e01e
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The time when the image was created.
	//
	// example:
	//
	// 1717430400000
	ImageCreate *int64 `json:"ImageCreate,omitempty" xml:"ImageCreate,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// b10ef7b245c34a0822055c74fc4a0e8b5baf0279306316b2c517a501ed250b1e
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The size of the image. Unit: bytes.
	//
	// example:
	//
	// 6120340
	ImageSize *int64 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The time when the image was updated.
	//
	// example:
	//
	// 1717430498600
	ImageUpdate *int64 `json:"ImageUpdate,omitempty" xml:"ImageUpdate,omitempty"`
	// The information about the Jenkins environment.
	//
	// example:
	//
	// release
	JenkinsEnv *string `json:"JenkinsEnv,omitempty" xml:"JenkinsEnv,omitempty"`
	// The namespace.
	//
	// example:
	//
	// lkl-zf-ss-ccss
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// sdk
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 119.136.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// 00f597223f-20210831-1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The token that is used to access the Jenkins image repository.
	//
	// example:
	//
	// c3de8326-273e-11fc-a0e3-d012435c****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The UUID of the image asset.
	//
	// example:
	//
	// 5b268326-273e-44fc-a0e3-9482435c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CreateJenkinsImageScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJenkinsImageScanTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateJenkinsImageScanTaskRequest) GetDigest() *string {
	return s.Digest
}

func (s *CreateJenkinsImageScanTaskRequest) GetImageCreate() *int64 {
	return s.ImageCreate
}

func (s *CreateJenkinsImageScanTaskRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateJenkinsImageScanTaskRequest) GetImageSize() *int64 {
	return s.ImageSize
}

func (s *CreateJenkinsImageScanTaskRequest) GetImageUpdate() *int64 {
	return s.ImageUpdate
}

func (s *CreateJenkinsImageScanTaskRequest) GetJenkinsEnv() *string {
	return s.JenkinsEnv
}

func (s *CreateJenkinsImageScanTaskRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateJenkinsImageScanTaskRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *CreateJenkinsImageScanTaskRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateJenkinsImageScanTaskRequest) GetTag() *string {
	return s.Tag
}

func (s *CreateJenkinsImageScanTaskRequest) GetToken() *string {
	return s.Token
}

func (s *CreateJenkinsImageScanTaskRequest) GetUuid() *string {
	return s.Uuid
}

func (s *CreateJenkinsImageScanTaskRequest) SetDigest(v string) *CreateJenkinsImageScanTaskRequest {
	s.Digest = &v
	return s
}

func (s *CreateJenkinsImageScanTaskRequest) SetImageCreate(v int64) *CreateJenkinsImageScanTaskRequest {
	s.ImageCreate = &v
	return s
}

func (s *CreateJenkinsImageScanTaskRequest) SetImageId(v string) *CreateJenkinsImageScanTaskRequest {
	s.ImageId = &v
	return s
}

func (s *CreateJenkinsImageScanTaskRequest) SetImageSize(v int64) *CreateJenkinsImageScanTaskRequest {
	s.ImageSize = &v
	return s
}

func (s *CreateJenkinsImageScanTaskRequest) SetImageUpdate(v int64) *CreateJenkinsImageScanTaskRequest {
	s.ImageUpdate = &v
	return s
}

func (s *CreateJenkinsImageScanTaskRequest) SetJenkinsEnv(v string) *CreateJenkinsImageScanTaskRequest {
	s.JenkinsEnv = &v
	return s
}

func (s *CreateJenkinsImageScanTaskRequest) SetNamespace(v string) *CreateJenkinsImageScanTaskRequest {
	s.Namespace = &v
	return s
}

func (s *CreateJenkinsImageScanTaskRequest) SetRepoName(v string) *CreateJenkinsImageScanTaskRequest {
	s.RepoName = &v
	return s
}

func (s *CreateJenkinsImageScanTaskRequest) SetSourceIp(v string) *CreateJenkinsImageScanTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateJenkinsImageScanTaskRequest) SetTag(v string) *CreateJenkinsImageScanTaskRequest {
	s.Tag = &v
	return s
}

func (s *CreateJenkinsImageScanTaskRequest) SetToken(v string) *CreateJenkinsImageScanTaskRequest {
	s.Token = &v
	return s
}

func (s *CreateJenkinsImageScanTaskRequest) SetUuid(v string) *CreateJenkinsImageScanTaskRequest {
	s.Uuid = &v
	return s
}

func (s *CreateJenkinsImageScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
