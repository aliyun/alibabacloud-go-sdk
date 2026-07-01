// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultStorageLocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *SetDefaultStorageLocationRequest
	GetBucket() *string
	SetPath(v string) *SetDefaultStorageLocationRequest
	GetPath() *string
	SetStorageType(v string) *SetDefaultStorageLocationRequest
	GetStorageType() *string
}

type SetDefaultStorageLocationRequest struct {
	// The name of the OSS bucket you created.
	//
	// example:
	//
	// oss-test-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// - When storageType is set to user_oss_bucket, temporary files are stored under this path. If path is empty or set to /, files are stored in the root directory.
	//
	// - This field does not take effect for VOD storage.
	//
	// example:
	//
	// ims/dir
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Storage type:
	//
	// - **vod_oss_bucket**: VOD-managed bucket.<br>
	//
	//   Supports adding buckets managed by the VOD system or OSS buckets added within the VOD system. If no active buckets are available, you can add a new bucket in the ApsaraVideo VOD console. After activating ApsaraVideo VOD, the system assigns a storage address in each storage region. You must enable this address before use. For details, see [Manage Storage Buckets](https://help.aliyun.com/document_detail/86097.html).
	//
	// - **user_oss_bucket**: User private bucket. Before adding an Object Storage address, you must activate Object Storage Service (OSS) and create a bucket. For details, see [Create a Bucket in the Console](https://help.aliyun.com/document_detail/31885.html).
	//
	// example:
	//
	// user_oss_bucket
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s SetDefaultStorageLocationRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultStorageLocationRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultStorageLocationRequest) GetBucket() *string {
	return s.Bucket
}

func (s *SetDefaultStorageLocationRequest) GetPath() *string {
	return s.Path
}

func (s *SetDefaultStorageLocationRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *SetDefaultStorageLocationRequest) SetBucket(v string) *SetDefaultStorageLocationRequest {
	s.Bucket = &v
	return s
}

func (s *SetDefaultStorageLocationRequest) SetPath(v string) *SetDefaultStorageLocationRequest {
	s.Path = &v
	return s
}

func (s *SetDefaultStorageLocationRequest) SetStorageType(v string) *SetDefaultStorageLocationRequest {
	s.StorageType = &v
	return s
}

func (s *SetDefaultStorageLocationRequest) Validate() error {
	return dara.Validate(s)
}
