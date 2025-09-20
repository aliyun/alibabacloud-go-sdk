// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachOSSBucketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *AttachOSSBucketRequest
	GetDescription() *string
	SetOSSBucket(v string) *AttachOSSBucketRequest
	GetOSSBucket() *string
	SetProjectName(v string) *AttachOSSBucketRequest
	GetProjectName() *string
}

type AttachOSSBucketRequest struct {
	// The description of the binding. The description must be 1 to 128 characters in length. By default, no description is applied.
	//
	// example:
	//
	// test-attachment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the OSS bucket in the same region as the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// examplebucket
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// The name of the project. For information about how to create a project, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s AttachOSSBucketRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachOSSBucketRequest) GoString() string {
	return s.String()
}

func (s *AttachOSSBucketRequest) GetDescription() *string {
	return s.Description
}

func (s *AttachOSSBucketRequest) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *AttachOSSBucketRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *AttachOSSBucketRequest) SetDescription(v string) *AttachOSSBucketRequest {
	s.Description = &v
	return s
}

func (s *AttachOSSBucketRequest) SetOSSBucket(v string) *AttachOSSBucketRequest {
	s.OSSBucket = &v
	return s
}

func (s *AttachOSSBucketRequest) SetProjectName(v string) *AttachOSSBucketRequest {
	s.ProjectName = &v
	return s
}

func (s *AttachOSSBucketRequest) Validate() error {
	return dara.Validate(s)
}
