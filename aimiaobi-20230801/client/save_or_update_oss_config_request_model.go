// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOrUpdateOssConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *SaveOrUpdateOssConfigRequest
	GetBucketName() *string
	SetEndPoint(v string) *SaveOrUpdateOssConfigRequest
	GetEndPoint() *string
	SetWorkspaceId(v string) *SaveOrUpdateOssConfigRequest
	GetWorkspaceId() *string
}

type SaveOrUpdateOssConfigRequest struct {
	// example:
	//
	// xxx
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SaveOrUpdateOssConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveOrUpdateOssConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveOrUpdateOssConfigRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *SaveOrUpdateOssConfigRequest) GetEndPoint() *string {
	return s.EndPoint
}

func (s *SaveOrUpdateOssConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SaveOrUpdateOssConfigRequest) SetBucketName(v string) *SaveOrUpdateOssConfigRequest {
	s.BucketName = &v
	return s
}

func (s *SaveOrUpdateOssConfigRequest) SetEndPoint(v string) *SaveOrUpdateOssConfigRequest {
	s.EndPoint = &v
	return s
}

func (s *SaveOrUpdateOssConfigRequest) SetWorkspaceId(v string) *SaveOrUpdateOssConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SaveOrUpdateOssConfigRequest) Validate() error {
	return dara.Validate(s)
}
