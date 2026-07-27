// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceStorageConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *DescribeInstanceStorageConfigResponseBody
	GetBranchName() *string
	SetConfigList(v []*DescribeInstanceStorageConfigResponseBodyConfigList) *DescribeInstanceStorageConfigResponseBody
	GetConfigList() []*DescribeInstanceStorageConfigResponseBodyConfigList
	SetInstanceName(v string) *DescribeInstanceStorageConfigResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *DescribeInstanceStorageConfigResponseBody
	GetRequestId() *string
}

type DescribeInstanceStorageConfigResponseBody struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The list of storage configurations.
	ConfigList []*DescribeInstanceStorageConfigResponseBodyConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// The instance ID of the AI application.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceStorageConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStorageConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStorageConfigResponseBody) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeInstanceStorageConfigResponseBody) GetConfigList() []*DescribeInstanceStorageConfigResponseBodyConfigList {
	return s.ConfigList
}

func (s *DescribeInstanceStorageConfigResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceStorageConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceStorageConfigResponseBody) SetBranchName(v string) *DescribeInstanceStorageConfigResponseBody {
	s.BranchName = &v
	return s
}

func (s *DescribeInstanceStorageConfigResponseBody) SetConfigList(v []*DescribeInstanceStorageConfigResponseBodyConfigList) *DescribeInstanceStorageConfigResponseBody {
	s.ConfigList = v
	return s
}

func (s *DescribeInstanceStorageConfigResponseBody) SetInstanceName(v string) *DescribeInstanceStorageConfigResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceStorageConfigResponseBody) SetRequestId(v string) *DescribeInstanceStorageConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceStorageConfigResponseBody) Validate() error {
	if s.ConfigList != nil {
		for _, item := range s.ConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceStorageConfigResponseBodyConfigList struct {
	// The name of the configuration item. Valid values:
	//
	// - **AWS_SESSION_TOKEN**: the temporary access token (Session Token) of OSS.
	//
	// - **AWS_ACCESS_KEY_ID**: the AccessKey ID of OSS.
	//
	// - **AWS_SECRET_ACCESS_KEY**: the AccessKey Secret of OSS.
	//
	// - **GLOBAL_S3_BUCKET**: the bucket name of OSS.
	//
	// - **TENANT_ID**: the tenant ID of the OSS prefix (prefix or directory).
	//
	// - **GLOBAL_S3_ENDPOINT**: the endpoint of OSS.
	//
	// - **REGION**: the region of OSS.
	//
	// example:
	//
	// REGION
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// cn-beijing
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceStorageConfigResponseBodyConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStorageConfigResponseBodyConfigList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStorageConfigResponseBodyConfigList) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceStorageConfigResponseBodyConfigList) GetValue() *string {
	return s.Value
}

func (s *DescribeInstanceStorageConfigResponseBodyConfigList) SetName(v string) *DescribeInstanceStorageConfigResponseBodyConfigList {
	s.Name = &v
	return s
}

func (s *DescribeInstanceStorageConfigResponseBodyConfigList) SetValue(v string) *DescribeInstanceStorageConfigResponseBodyConfigList {
	s.Value = &v
	return s
}

func (s *DescribeInstanceStorageConfigResponseBodyConfigList) Validate() error {
	return dara.Validate(s)
}
