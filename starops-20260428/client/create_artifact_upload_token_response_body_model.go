// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactUploadTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *CreateArtifactUploadTokenResponseBody
	GetAccessId() *string
	SetArtifactPath(v string) *CreateArtifactUploadTokenResponseBody
	GetArtifactPath() *string
	SetDir(v string) *CreateArtifactUploadTokenResponseBody
	GetDir() *string
	SetExpire(v int64) *CreateArtifactUploadTokenResponseBody
	GetExpire() *int64
	SetHost(v string) *CreateArtifactUploadTokenResponseBody
	GetHost() *string
	SetMaxSize(v int64) *CreateArtifactUploadTokenResponseBody
	GetMaxSize() *int64
	SetPolicy(v string) *CreateArtifactUploadTokenResponseBody
	GetPolicy() *string
	SetRequestId(v string) *CreateArtifactUploadTokenResponseBody
	GetRequestId() *string
	SetSignature(v string) *CreateArtifactUploadTokenResponseBody
	GetSignature() *string
	SetSuccessActionStatus(v string) *CreateArtifactUploadTokenResponseBody
	GetSuccessActionStatus() *string
}

type CreateArtifactUploadTokenResponseBody struct {
	// example:
	//
	// LTAI******
	AccessId *string `json:"accessId,omitempty" xml:"accessId,omitempty"`
	// example:
	//
	// upload/2026-05-25/
	ArtifactPath *string `json:"artifactPath,omitempty" xml:"artifactPath,omitempty"`
	// example:
	//
	// agents/123/sample-agent/home/starops/upload/2026-05-25/
	Dir *string `json:"dir,omitempty" xml:"dir,omitempty"`
	// example:
	//
	// 1770000000
	Expire *int64 `json:"expire,omitempty" xml:"expire,omitempty"`
	// example:
	//
	// https://example-bucket.oss-cn-shanghai.aliyuncs.com
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// example:
	//
	// 104857600
	MaxSize *int64 `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNi0wNS0yMVQwODowMDowMFoifQ==
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-1234567890AB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// abc123signature
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// example:
	//
	// 200
	SuccessActionStatus *string `json:"successActionStatus,omitempty" xml:"successActionStatus,omitempty"`
}

func (s CreateArtifactUploadTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactUploadTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArtifactUploadTokenResponseBody) GetAccessId() *string {
	return s.AccessId
}

func (s *CreateArtifactUploadTokenResponseBody) GetArtifactPath() *string {
	return s.ArtifactPath
}

func (s *CreateArtifactUploadTokenResponseBody) GetDir() *string {
	return s.Dir
}

func (s *CreateArtifactUploadTokenResponseBody) GetExpire() *int64 {
	return s.Expire
}

func (s *CreateArtifactUploadTokenResponseBody) GetHost() *string {
	return s.Host
}

func (s *CreateArtifactUploadTokenResponseBody) GetMaxSize() *int64 {
	return s.MaxSize
}

func (s *CreateArtifactUploadTokenResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *CreateArtifactUploadTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateArtifactUploadTokenResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *CreateArtifactUploadTokenResponseBody) GetSuccessActionStatus() *string {
	return s.SuccessActionStatus
}

func (s *CreateArtifactUploadTokenResponseBody) SetAccessId(v string) *CreateArtifactUploadTokenResponseBody {
	s.AccessId = &v
	return s
}

func (s *CreateArtifactUploadTokenResponseBody) SetArtifactPath(v string) *CreateArtifactUploadTokenResponseBody {
	s.ArtifactPath = &v
	return s
}

func (s *CreateArtifactUploadTokenResponseBody) SetDir(v string) *CreateArtifactUploadTokenResponseBody {
	s.Dir = &v
	return s
}

func (s *CreateArtifactUploadTokenResponseBody) SetExpire(v int64) *CreateArtifactUploadTokenResponseBody {
	s.Expire = &v
	return s
}

func (s *CreateArtifactUploadTokenResponseBody) SetHost(v string) *CreateArtifactUploadTokenResponseBody {
	s.Host = &v
	return s
}

func (s *CreateArtifactUploadTokenResponseBody) SetMaxSize(v int64) *CreateArtifactUploadTokenResponseBody {
	s.MaxSize = &v
	return s
}

func (s *CreateArtifactUploadTokenResponseBody) SetPolicy(v string) *CreateArtifactUploadTokenResponseBody {
	s.Policy = &v
	return s
}

func (s *CreateArtifactUploadTokenResponseBody) SetRequestId(v string) *CreateArtifactUploadTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateArtifactUploadTokenResponseBody) SetSignature(v string) *CreateArtifactUploadTokenResponseBody {
	s.Signature = &v
	return s
}

func (s *CreateArtifactUploadTokenResponseBody) SetSuccessActionStatus(v string) *CreateArtifactUploadTokenResponseBody {
	s.SuccessActionStatus = &v
	return s
}

func (s *CreateArtifactUploadTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
