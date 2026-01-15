// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncreaseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *IncreaseInstanceRequest
	GetBucketName() *string
	SetCallbackAddress(v string) *IncreaseInstanceRequest
	GetCallbackAddress() *string
	SetInstanceName(v string) *IncreaseInstanceRequest
	GetInstanceName() *string
	SetPath(v string) *IncreaseInstanceRequest
	GetPath() *string
}

type IncreaseInstanceRequest struct {
	// The name of the Object Storage Service (OSS) bucket.
	//
	// >  The bucket must be in the same region as the Image Search instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// bucketName
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The callback address.
	//
	// example:
	//
	// http://xxxxx
	CallbackAddress *string `json:"CallbackAddress,omitempty" xml:"CallbackAddress,omitempty"`
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// imagesearchName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The absolute path to the increment.meta file in the bucket. The path must start with a forward slash (/) and cannot end with a forward slash (/).
	//
	// This parameter is required.
	//
	// example:
	//
	// /xxx/xxx
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s IncreaseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s IncreaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *IncreaseInstanceRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *IncreaseInstanceRequest) GetCallbackAddress() *string {
	return s.CallbackAddress
}

func (s *IncreaseInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *IncreaseInstanceRequest) GetPath() *string {
	return s.Path
}

func (s *IncreaseInstanceRequest) SetBucketName(v string) *IncreaseInstanceRequest {
	s.BucketName = &v
	return s
}

func (s *IncreaseInstanceRequest) SetCallbackAddress(v string) *IncreaseInstanceRequest {
	s.CallbackAddress = &v
	return s
}

func (s *IncreaseInstanceRequest) SetInstanceName(v string) *IncreaseInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *IncreaseInstanceRequest) SetPath(v string) *IncreaseInstanceRequest {
	s.Path = &v
	return s
}

func (s *IncreaseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
