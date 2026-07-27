// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceStorageConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *ModifyInstanceStorageConfigRequest
	GetBranchName() *string
	SetClientToken(v string) *ModifyInstanceStorageConfigRequest
	GetClientToken() *string
	SetConfigList(v []*ModifyInstanceStorageConfigRequestConfigList) *ModifyInstanceStorageConfigRequest
	GetConfigList() []*ModifyInstanceStorageConfigRequestConfigList
	SetInstanceName(v string) *ModifyInstanceStorageConfigRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyInstanceStorageConfigRequest
	GetRegionId() *string
}

type ModifyInstanceStorageConfigRequest struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, which ensures that the request is not repeated.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of storage configurations.
	ConfigList []*ModifyInstanceStorageConfigRequestConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// The instance ID of the AI application.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceStorageConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceStorageConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceStorageConfigRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *ModifyInstanceStorageConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceStorageConfigRequest) GetConfigList() []*ModifyInstanceStorageConfigRequestConfigList {
	return s.ConfigList
}

func (s *ModifyInstanceStorageConfigRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceStorageConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceStorageConfigRequest) SetBranchName(v string) *ModifyInstanceStorageConfigRequest {
	s.BranchName = &v
	return s
}

func (s *ModifyInstanceStorageConfigRequest) SetClientToken(v string) *ModifyInstanceStorageConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceStorageConfigRequest) SetConfigList(v []*ModifyInstanceStorageConfigRequestConfigList) *ModifyInstanceStorageConfigRequest {
	s.ConfigList = v
	return s
}

func (s *ModifyInstanceStorageConfigRequest) SetInstanceName(v string) *ModifyInstanceStorageConfigRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceStorageConfigRequest) SetRegionId(v string) *ModifyInstanceStorageConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceStorageConfigRequest) Validate() error {
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

type ModifyInstanceStorageConfigRequestConfigList struct {
	// The name of the configuration item. Valid values:
	//
	// - **AWS_SESSION_TOKEN*	- (optional): the temporary access token (Session Token) for OSS. If this parameter is not specified, AccessKey ID and AccessKey Secret are used for authentication.
	//
	// - **AWS_ACCESS_KEY_ID**: the AccessKey ID for OSS.
	//
	// - **AWS_SECRET_ACCESS_KEY**: the AccessKey Secret for OSS.
	//
	// - **GLOBAL_S3_BUCKET**: the bucket name of OSS.
	//
	// - **TENANT_ID**: the OSS directory name. You do not need to create it in advance.
	//
	// - **GLOBAL_S3_ENDPOINT**: the endpoint of OSS.
	//
	// - **REGION**: the region of OSS.
	//
	// example:
	//
	// TENANT_ID
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// test-prefix
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyInstanceStorageConfigRequestConfigList) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceStorageConfigRequestConfigList) GoString() string {
	return s.String()
}

func (s *ModifyInstanceStorageConfigRequestConfigList) GetName() *string {
	return s.Name
}

func (s *ModifyInstanceStorageConfigRequestConfigList) GetValue() *string {
	return s.Value
}

func (s *ModifyInstanceStorageConfigRequestConfigList) SetName(v string) *ModifyInstanceStorageConfigRequestConfigList {
	s.Name = &v
	return s
}

func (s *ModifyInstanceStorageConfigRequestConfigList) SetValue(v string) *ModifyInstanceStorageConfigRequestConfigList {
	s.Value = &v
	return s
}

func (s *ModifyInstanceStorageConfigRequestConfigList) Validate() error {
	return dara.Validate(s)
}
