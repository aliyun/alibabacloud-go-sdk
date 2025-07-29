// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOssBucketEndpoint(v string) *MigrateClusterRequest
	GetOssBucketEndpoint() *string
	SetOssBucketName(v string) *MigrateClusterRequest
	GetOssBucketName() *string
}

type MigrateClusterRequest struct {
	// The endpoint of the OSS bucket.
	//
	// example:
	//
	// *******.oss-cn-hangzhou.aliyuncs.com
	OssBucketEndpoint *string `json:"oss_bucket_endpoint,omitempty" xml:"oss_bucket_endpoint,omitempty"`
	// The name of the Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// bucket-****
	OssBucketName *string `json:"oss_bucket_name,omitempty" xml:"oss_bucket_name,omitempty"`
}

func (s MigrateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateClusterRequest) GoString() string {
	return s.String()
}

func (s *MigrateClusterRequest) GetOssBucketEndpoint() *string {
	return s.OssBucketEndpoint
}

func (s *MigrateClusterRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *MigrateClusterRequest) SetOssBucketEndpoint(v string) *MigrateClusterRequest {
	s.OssBucketEndpoint = &v
	return s
}

func (s *MigrateClusterRequest) SetOssBucketName(v string) *MigrateClusterRequest {
	s.OssBucketName = &v
	return s
}

func (s *MigrateClusterRequest) Validate() error {
	return dara.Validate(s)
}
