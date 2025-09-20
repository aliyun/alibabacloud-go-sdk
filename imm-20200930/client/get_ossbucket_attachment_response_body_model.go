// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOSSBucketAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetOSSBucketAttachmentResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetOSSBucketAttachmentResponseBody
	GetDescription() *string
	SetProjectName(v string) *GetOSSBucketAttachmentResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetOSSBucketAttachmentResponseBody
	GetRequestId() *string
	SetUpdateTime(v string) *GetOSSBucketAttachmentResponseBody
	GetUpdateTime() *string
}

type GetOSSBucketAttachmentResponseBody struct {
	// The time when the dataset was created.
	//
	// example:
	//
	// ""2023-12-19T17:29:34.790931971+08:00"
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the dataset.
	//
	// example:
	//
	// "Dataset"
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5F74C5C9-5AC0-49F9-914D-E01589D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the dataset was last updated.
	//
	// example:
	//
	// "2023-12-19T17:29:34.790931971+08:00"
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetOSSBucketAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOSSBucketAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetOSSBucketAttachmentResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetOSSBucketAttachmentResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetOSSBucketAttachmentResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetOSSBucketAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOSSBucketAttachmentResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetOSSBucketAttachmentResponseBody) SetCreateTime(v string) *GetOSSBucketAttachmentResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetOSSBucketAttachmentResponseBody) SetDescription(v string) *GetOSSBucketAttachmentResponseBody {
	s.Description = &v
	return s
}

func (s *GetOSSBucketAttachmentResponseBody) SetProjectName(v string) *GetOSSBucketAttachmentResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetOSSBucketAttachmentResponseBody) SetRequestId(v string) *GetOSSBucketAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOSSBucketAttachmentResponseBody) SetUpdateTime(v string) *GetOSSBucketAttachmentResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetOSSBucketAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
