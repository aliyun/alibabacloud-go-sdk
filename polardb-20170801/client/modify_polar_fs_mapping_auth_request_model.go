// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarFsMappingAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *ModifyPolarFsMappingAuthRequest
	GetBucket() *string
	SetBucketAccessKeyId(v string) *ModifyPolarFsMappingAuthRequest
	GetBucketAccessKeyId() *string
	SetBucketAccessKeySecret(v string) *ModifyPolarFsMappingAuthRequest
	GetBucketAccessKeySecret() *string
	SetDBClusterId(v string) *ModifyPolarFsMappingAuthRequest
	GetDBClusterId() *string
	SetPath(v string) *ModifyPolarFsMappingAuthRequest
	GetPath() *string
	SetPolarFsInstanceId(v string) *ModifyPolarFsMappingAuthRequest
	GetPolarFsInstanceId() *string
}

type ModifyPolarFsMappingAuthRequest struct {
	// The bucket name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-xxx.oss-[regionId]-internal.aliyuncs.com
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The AccessKey ID for the storage bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	BucketAccessKeyId *string `json:"BucketAccessKeyId,omitempty" xml:"BucketAccessKeyId,omitempty"`
	// The AccessKey secret for the storage bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	BucketAccessKeySecret *string `json:"BucketAccessKeySecret,omitempty" xml:"BucketAccessKeySecret,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query information about all clusters in a specified region, including the cluster ID.
	//
	// example:
	//
	// pc-******************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The destination path.
	//
	// This parameter is required.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The PolarFS instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s ModifyPolarFsMappingAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarFsMappingAuthRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolarFsMappingAuthRequest) GetBucket() *string {
	return s.Bucket
}

func (s *ModifyPolarFsMappingAuthRequest) GetBucketAccessKeyId() *string {
	return s.BucketAccessKeyId
}

func (s *ModifyPolarFsMappingAuthRequest) GetBucketAccessKeySecret() *string {
	return s.BucketAccessKeySecret
}

func (s *ModifyPolarFsMappingAuthRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyPolarFsMappingAuthRequest) GetPath() *string {
	return s.Path
}

func (s *ModifyPolarFsMappingAuthRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *ModifyPolarFsMappingAuthRequest) SetBucket(v string) *ModifyPolarFsMappingAuthRequest {
	s.Bucket = &v
	return s
}

func (s *ModifyPolarFsMappingAuthRequest) SetBucketAccessKeyId(v string) *ModifyPolarFsMappingAuthRequest {
	s.BucketAccessKeyId = &v
	return s
}

func (s *ModifyPolarFsMappingAuthRequest) SetBucketAccessKeySecret(v string) *ModifyPolarFsMappingAuthRequest {
	s.BucketAccessKeySecret = &v
	return s
}

func (s *ModifyPolarFsMappingAuthRequest) SetDBClusterId(v string) *ModifyPolarFsMappingAuthRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyPolarFsMappingAuthRequest) SetPath(v string) *ModifyPolarFsMappingAuthRequest {
	s.Path = &v
	return s
}

func (s *ModifyPolarFsMappingAuthRequest) SetPolarFsInstanceId(v string) *ModifyPolarFsMappingAuthRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *ModifyPolarFsMappingAuthRequest) Validate() error {
	return dara.Validate(s)
}
