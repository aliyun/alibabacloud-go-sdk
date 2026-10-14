// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateC2paBucketSignTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *CreateC2paBucketSignTaskRequest
	GetBucketName() *string
	SetBucketRegionId(v string) *CreateC2paBucketSignTaskRequest
	GetBucketRegionId() *string
	SetSoftwareAgent(v string) *CreateC2paBucketSignTaskRequest
	GetSoftwareAgent() *string
	SetSourcePrefix(v string) *CreateC2paBucketSignTaskRequest
	GetSourcePrefix() *string
	SetSourceType(v string) *CreateC2paBucketSignTaskRequest
	GetSourceType() *string
}

type CreateC2paBucketSignTaskRequest struct {
	// The name of the user OSS bucket. You must authorize the service-linked role AliyunServiceRoleForAIDeepSign to access the bucket in advance.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The region of the user OSS bucket. It must be consistent with the actual region of the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BucketRegionId *string `json:"BucketRegionId,omitempty" xml:"BucketRegionId,omitempty"`
	// The optional name of the software that generated the content. This parameter is not supported in the current version.
	//
	// example:
	//
	// Adobe Photoshop 25.0
	SoftwareAgent *string `json:"SoftwareAgent,omitempty" xml:"SoftwareAgent,omitempty"`
	// The source object prefix. It must be `ai-deepsign/` or its subdirectory. The service recursively processes objects under this prefix.
	//
	// This parameter is required.
	//
	// example:
	//
	// ai-deepsign/images/
	SourcePrefix *string `json:"SourcePrefix,omitempty" xml:"SourcePrefix,omitempty"`
	// The optional content source type. This parameter is not supported in the current version.
	//
	// example:
	//
	// digitalCapture
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s CreateC2paBucketSignTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateC2paBucketSignTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateC2paBucketSignTaskRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateC2paBucketSignTaskRequest) GetBucketRegionId() *string {
	return s.BucketRegionId
}

func (s *CreateC2paBucketSignTaskRequest) GetSoftwareAgent() *string {
	return s.SoftwareAgent
}

func (s *CreateC2paBucketSignTaskRequest) GetSourcePrefix() *string {
	return s.SourcePrefix
}

func (s *CreateC2paBucketSignTaskRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateC2paBucketSignTaskRequest) SetBucketName(v string) *CreateC2paBucketSignTaskRequest {
	s.BucketName = &v
	return s
}

func (s *CreateC2paBucketSignTaskRequest) SetBucketRegionId(v string) *CreateC2paBucketSignTaskRequest {
	s.BucketRegionId = &v
	return s
}

func (s *CreateC2paBucketSignTaskRequest) SetSoftwareAgent(v string) *CreateC2paBucketSignTaskRequest {
	s.SoftwareAgent = &v
	return s
}

func (s *CreateC2paBucketSignTaskRequest) SetSourcePrefix(v string) *CreateC2paBucketSignTaskRequest {
	s.SourcePrefix = &v
	return s
}

func (s *CreateC2paBucketSignTaskRequest) SetSourceType(v string) *CreateC2paBucketSignTaskRequest {
	s.SourceType = &v
	return s
}

func (s *CreateC2paBucketSignTaskRequest) Validate() error {
	return dara.Validate(s)
}
