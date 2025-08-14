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
	// example:
	//
	// oss-test-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// ims/dir
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
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
