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
	// The name of the bucket.
	//
	// > Only a bucket in the same region as the instance is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// bucketName
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The callback URL of the task.
	//
	// example:
	//
	// http://xxxxx
	CallbackAddress *string `json:"CallbackAddress,omitempty" xml:"CallbackAddress,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// If you have purchased an Image Search instance, log on to the [Image Search console](https://imagesearch.console.aliyun.com/) to view the instance name.
	//
	// If you have not purchased an Image Search instance, refer to [Activate the service](https://help.aliyun.com/document_detail/179178.html) and [Create an instance](https://help.aliyun.com/document_detail/66569.html).
	//
	// >The instance name is not the instance ID. Make sure that you distinguish between the two.
	//
	// This parameter is required.
	//
	// example:
	//
	// imagesearchName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The absolute path of the increment.meta file in OSS. The path must start with a forward slash (/) and must not end with a forward slash (/).
	//
	// > Prepare the increment.meta file in advance. For more information, see [Batch operations](https://help.aliyun.com/document_detail/66580.html).
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
