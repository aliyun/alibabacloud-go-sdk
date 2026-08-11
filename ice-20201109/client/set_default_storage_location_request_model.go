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
	// The name of the OSS bucket that you created.
	//
	// example:
	//
	// oss-test-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// - When StorageType is set to user_oss_bucket, temporary files are stored in this path. If the path is empty or set to /, files are stored in the root directory.
	//
	// - This field does not take effect for VOD storage.
	//
	// example:
	//
	// ims/dir
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The storage type. Valid values:
	//
	// - vod_oss_bucket: VOD-managed bucket. You can add VOD system buckets and your own OSS buckets that have been added to the VOD system. If no bucket is available, you can create a bucket in the ApsaraVideo VOD console. The ApsaraVideo VOD system assigns a storage address in each storage region. After you activate ApsaraVideo VOD, you must enable the address before you can use it. For more information, see [Manage storage buckets](https://help.aliyun.com/document_detail/86097.html).
	//
	// - user_oss_bucket: user-owned private bucket. Before adding an OSS storage address, you must activate OSS and create a storage bucket. For more information, see [Create a bucket in the console](https://help.aliyun.com/document_detail/31885.html).
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
