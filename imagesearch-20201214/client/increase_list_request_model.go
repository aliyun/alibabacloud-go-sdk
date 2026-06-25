// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncreaseListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *IncreaseListRequest
	GetBucketName() *string
	SetId(v int64) *IncreaseListRequest
	GetId() *int64
	SetInstanceName(v string) *IncreaseListRequest
	GetInstanceName() *string
	SetPageNumber(v int32) *IncreaseListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *IncreaseListRequest
	GetPageSize() *int32
	SetPath(v string) *IncreaseListRequest
	GetPath() *string
}

type IncreaseListRequest struct {
	// The name of the bucket.
	//
	// example:
	//
	// bucketName
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The batch task ID.
	//
	// >To obtain the batch task ID, call the [batch operation](https://help.aliyun.com/document_detail/377468.html) first and retrieve the ID from the response.
	//
	// example:
	//
	// 500
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// If you have purchased an Image Search instance, log on to the [Image Search console](https://imagesearch.console.aliyun.com/) to view the instance name.
	//
	// If you have not purchased an Image Search instance, refer to [Activate the service](https://help.aliyun.com/document_detail/179178.html) and [Create an instance](https://help.aliyun.com/document_detail/66569.html).
	//
	// >The instance name is not the instance ID. The instance name must be unique within the same region. Make sure you use the correct value.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The page number of the first page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The absolute path of the increment.meta file in OSS. The path must start with a forward slash (/) and must not end with a forward slash (/).
	//
	// example:
	//
	// /xxx/xxx
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s IncreaseListRequest) String() string {
	return dara.Prettify(s)
}

func (s IncreaseListRequest) GoString() string {
	return s.String()
}

func (s *IncreaseListRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *IncreaseListRequest) GetId() *int64 {
	return s.Id
}

func (s *IncreaseListRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *IncreaseListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *IncreaseListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *IncreaseListRequest) GetPath() *string {
	return s.Path
}

func (s *IncreaseListRequest) SetBucketName(v string) *IncreaseListRequest {
	s.BucketName = &v
	return s
}

func (s *IncreaseListRequest) SetId(v int64) *IncreaseListRequest {
	s.Id = &v
	return s
}

func (s *IncreaseListRequest) SetInstanceName(v string) *IncreaseListRequest {
	s.InstanceName = &v
	return s
}

func (s *IncreaseListRequest) SetPageNumber(v int32) *IncreaseListRequest {
	s.PageNumber = &v
	return s
}

func (s *IncreaseListRequest) SetPageSize(v int32) *IncreaseListRequest {
	s.PageSize = &v
	return s
}

func (s *IncreaseListRequest) SetPath(v string) *IncreaseListRequest {
	s.Path = &v
	return s
}

func (s *IncreaseListRequest) Validate() error {
	return dara.Validate(s)
}
