// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefaultStorageLocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *GetDefaultStorageLocationResponseBody
	GetBucket() *string
	SetPath(v string) *GetDefaultStorageLocationResponseBody
	GetPath() *string
	SetRequestId(v string) *GetDefaultStorageLocationResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetDefaultStorageLocationResponseBody
	GetStatus() *string
	SetStorageType(v string) *GetDefaultStorageLocationResponseBody
	GetStorageType() *string
}

type GetDefaultStorageLocationResponseBody struct {
	// OSS bucket name
	//
	// example:
	//
	// oss-test-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// Path
	//
	// example:
	//
	// ice/dir
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Request ID
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Bucket status:
	//
	// - Normal: The bucket is in normal status.
	//
	// - notExist: The bucket does not exist.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Storage class
	//
	// example:
	//
	// user_oss_bucket
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s GetDefaultStorageLocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultStorageLocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultStorageLocationResponseBody) GetBucket() *string {
	return s.Bucket
}

func (s *GetDefaultStorageLocationResponseBody) GetPath() *string {
	return s.Path
}

func (s *GetDefaultStorageLocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDefaultStorageLocationResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetDefaultStorageLocationResponseBody) GetStorageType() *string {
	return s.StorageType
}

func (s *GetDefaultStorageLocationResponseBody) SetBucket(v string) *GetDefaultStorageLocationResponseBody {
	s.Bucket = &v
	return s
}

func (s *GetDefaultStorageLocationResponseBody) SetPath(v string) *GetDefaultStorageLocationResponseBody {
	s.Path = &v
	return s
}

func (s *GetDefaultStorageLocationResponseBody) SetRequestId(v string) *GetDefaultStorageLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDefaultStorageLocationResponseBody) SetStatus(v string) *GetDefaultStorageLocationResponseBody {
	s.Status = &v
	return s
}

func (s *GetDefaultStorageLocationResponseBody) SetStorageType(v string) *GetDefaultStorageLocationResponseBody {
	s.StorageType = &v
	return s
}

func (s *GetDefaultStorageLocationResponseBody) Validate() error {
	return dara.Validate(s)
}
