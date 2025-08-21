// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *DeleteObjectRequest
	GetBucketName() *string
	SetObjectKey(v string) *DeleteObjectRequest
	GetObjectKey() *string
}

type DeleteObjectRequest struct {
	// The name of the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// tets
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The name of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// image5
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
}

func (s DeleteObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteObjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteObjectRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *DeleteObjectRequest) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *DeleteObjectRequest) SetBucketName(v string) *DeleteObjectRequest {
	s.BucketName = &v
	return s
}

func (s *DeleteObjectRequest) SetObjectKey(v string) *DeleteObjectRequest {
	s.ObjectKey = &v
	return s
}

func (s *DeleteObjectRequest) Validate() error {
	return dara.Validate(s)
}
