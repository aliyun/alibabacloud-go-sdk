// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCriteriaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMachineTypes(v string) *DescribeCriteriaRequest
	GetMachineTypes() *string
	SetResourceDirectoryAccountId(v int64) *DescribeCriteriaRequest
	GetResourceDirectoryAccountId() *int64
	SetSupportAutoTag(v bool) *DescribeCriteriaRequest
	GetSupportAutoTag() *bool
	SetValue(v string) *DescribeCriteriaRequest
	GetValue() *string
}

type DescribeCriteriaRequest struct {
	// The Asset Type to query. Valid values:
	//
	// - **ecs**: queries all ECS servers.
	//
	// example:
	//
	// ecs
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	// The ID of the Alibaba Cloud account of the member accounts in the resource folder.
	//
	// >Invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// Specifies whether the fuzzy query field supports automatic matching. Default value: **false**. Valid values:
	//
	// - **true**: Supported.
	//
	// - **false**: Not supported.
	//
	// example:
	//
	// true
	SupportAutoTag *bool `json:"SupportAutoTag,omitempty" xml:"SupportAutoTag,omitempty"`
	// The fuzzy match value entered when querying assets.
	//
	// example:
	//
	// 47.96
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCriteriaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCriteriaRequest) GoString() string {
	return s.String()
}

func (s *DescribeCriteriaRequest) GetMachineTypes() *string {
	return s.MachineTypes
}

func (s *DescribeCriteriaRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeCriteriaRequest) GetSupportAutoTag() *bool {
	return s.SupportAutoTag
}

func (s *DescribeCriteriaRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeCriteriaRequest) SetMachineTypes(v string) *DescribeCriteriaRequest {
	s.MachineTypes = &v
	return s
}

func (s *DescribeCriteriaRequest) SetResourceDirectoryAccountId(v int64) *DescribeCriteriaRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeCriteriaRequest) SetSupportAutoTag(v bool) *DescribeCriteriaRequest {
	s.SupportAutoTag = &v
	return s
}

func (s *DescribeCriteriaRequest) SetValue(v string) *DescribeCriteriaRequest {
	s.Value = &v
	return s
}

func (s *DescribeCriteriaRequest) Validate() error {
	return dara.Validate(s)
}
