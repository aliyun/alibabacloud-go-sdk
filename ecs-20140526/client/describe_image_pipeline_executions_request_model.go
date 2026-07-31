// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagePipelineExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionId(v string) *DescribeImagePipelineExecutionsRequest
	GetExecutionId() *string
	SetImagePipelineId(v string) *DescribeImagePipelineExecutionsRequest
	GetImagePipelineId() *string
	SetMaxResults(v int32) *DescribeImagePipelineExecutionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeImagePipelineExecutionsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeImagePipelineExecutionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeImagePipelineExecutionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeImagePipelineExecutionsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeImagePipelineExecutionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeImagePipelineExecutionsRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeImagePipelineExecutionsRequest
	GetStatus() *string
	SetTag(v []*DescribeImagePipelineExecutionsRequestTag) *DescribeImagePipelineExecutionsRequest
	GetTag() []*DescribeImagePipelineExecutionsRequestTag
}

type DescribeImagePipelineExecutionsRequest struct {
	// The ID of the image build task.
	//
	// example:
	//
	// exec-5fb8facb8ed7427c****
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The ID of the image template.
	//
	// example:
	//
	// ip-2ze5tsl5bp6nf2b3****
	ImagePipelineId *string `json:"ImagePipelineId,omitempty" xml:"ImagePipelineId,omitempty"`
	// The maximum number of entries per page for paging. Valid values: 1 to 500.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the value of NextToken returned in the previous call. You do not need to set this parameter for the first request.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the image build task. You can specify multiple values at the same time. Separate multiple values with commas (,). Example: `BUILDING,DISTRIBUTING`. Valid values:
	//
	// - PREPARING: The task is being prepared. Resources such as the temporary intermediate instance are being created.
	//
	// - REPAIRING: The task is being repaired. The source image is being repaired.
	//
	// - BUILDING: The task is being built. Custom commands are being run and the image is being created.
	//
	// - TESTING: The task is being tested. Custom test commands are being run.
	//
	// - DISTRIBUTING: The task is being distributed. Image copying and sharing are being performed.
	//
	// - RELEASING: Resources are being reclaimed. Temporary resources generated during the build process are being released.
	//
	// - SUCCESS: The task succeeded.
	//
	// - PARTITION_SUCCESS: The task partially succeeded. The image was built, but exceptions may have occurred during distribution or resource cleanup.
	//
	// - FAILED: The task failed.
	//
	// - TEST_FAILED: The test failed. The image was created, but the test failed.
	//
	// - CANCELLING: The task is being canceled.
	//
	// - CANCELLED: The task was canceled.
	//
	// > If this parameter is empty, image build tasks in all states are queried.
	//
	// example:
	//
	// BUILDING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tag []*DescribeImagePipelineExecutionsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeImagePipelineExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelineExecutionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelineExecutionsRequest) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *DescribeImagePipelineExecutionsRequest) GetImagePipelineId() *string {
	return s.ImagePipelineId
}

func (s *DescribeImagePipelineExecutionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeImagePipelineExecutionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeImagePipelineExecutionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeImagePipelineExecutionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeImagePipelineExecutionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImagePipelineExecutionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeImagePipelineExecutionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeImagePipelineExecutionsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeImagePipelineExecutionsRequest) GetTag() []*DescribeImagePipelineExecutionsRequestTag {
	return s.Tag
}

func (s *DescribeImagePipelineExecutionsRequest) SetExecutionId(v string) *DescribeImagePipelineExecutionsRequest {
	s.ExecutionId = &v
	return s
}

func (s *DescribeImagePipelineExecutionsRequest) SetImagePipelineId(v string) *DescribeImagePipelineExecutionsRequest {
	s.ImagePipelineId = &v
	return s
}

func (s *DescribeImagePipelineExecutionsRequest) SetMaxResults(v int32) *DescribeImagePipelineExecutionsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeImagePipelineExecutionsRequest) SetNextToken(v string) *DescribeImagePipelineExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeImagePipelineExecutionsRequest) SetOwnerAccount(v string) *DescribeImagePipelineExecutionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeImagePipelineExecutionsRequest) SetOwnerId(v int64) *DescribeImagePipelineExecutionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeImagePipelineExecutionsRequest) SetRegionId(v string) *DescribeImagePipelineExecutionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImagePipelineExecutionsRequest) SetResourceOwnerAccount(v string) *DescribeImagePipelineExecutionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeImagePipelineExecutionsRequest) SetResourceOwnerId(v int64) *DescribeImagePipelineExecutionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeImagePipelineExecutionsRequest) SetStatus(v string) *DescribeImagePipelineExecutionsRequest {
	s.Status = &v
	return s
}

func (s *DescribeImagePipelineExecutionsRequest) SetTag(v []*DescribeImagePipelineExecutionsRequestTag) *DescribeImagePipelineExecutionsRequest {
	s.Tag = v
	return s
}

func (s *DescribeImagePipelineExecutionsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImagePipelineExecutionsRequestTag struct {
	// The key of the tag. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImagePipelineExecutionsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelineExecutionsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelineExecutionsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeImagePipelineExecutionsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeImagePipelineExecutionsRequestTag) SetKey(v string) *DescribeImagePipelineExecutionsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeImagePipelineExecutionsRequestTag) SetValue(v string) *DescribeImagePipelineExecutionsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeImagePipelineExecutionsRequestTag) Validate() error {
	return dara.Validate(s)
}
