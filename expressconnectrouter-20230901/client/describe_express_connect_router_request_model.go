// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeExpressConnectRouterRequest
	GetClientToken() *string
	SetDryRun(v bool) *DescribeExpressConnectRouterRequest
	GetDryRun() *bool
	SetEcrId(v string) *DescribeExpressConnectRouterRequest
	GetEcrId() *string
	SetMaxResults(v int32) *DescribeExpressConnectRouterRequest
	GetMaxResults() *int32
	SetName(v string) *DescribeExpressConnectRouterRequest
	GetName() *string
	SetNextToken(v string) *DescribeExpressConnectRouterRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *DescribeExpressConnectRouterRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeExpressConnectRouterRequestTag) *DescribeExpressConnectRouterRequest
	GetTag() []*DescribeExpressConnectRouterRequestTag
	SetVersion(v string) *DescribeExpressConnectRouterRequest
	GetVersion() *string
}

type DescribeExpressConnectRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate a parameter value from your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The request is checked, but service resources are not modified.
	//
	// - **false*	- (default): sends a normal request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the Express Connect Router (ECR) instance.
	//
	// example:
	//
	// ecr-fu8rszhgv7623c****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The maximum number of entries to return. Valid values: 1 to 2147483647. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the Express Connect Router (ECR) instance. The name must be 0 to 128 characters in length.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token for the next query. Valid values:
	//
	// - For the first query or when no next query is required, leave this parameter empty.
	//
	// - If a next query is required, set the value to the NextToken value returned from the previous API call.
	//
	// example:
	//
	// e0a2dbeb69a8beeeb8194e92b702df3fd3e7bfe6ce7bfc16e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group to which the Express Connect Router (ECR) instance belongs.
	//
	// example:
	//
	// rg-aek2aq7f4va****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags used for filtering. You can specify up to 20 tags.
	Tag     []*DescribeExpressConnectRouterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	Version *string                                   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeExpressConnectRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeExpressConnectRouterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeExpressConnectRouterRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeExpressConnectRouterRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeExpressConnectRouterRequest) GetName() *string {
	return s.Name
}

func (s *DescribeExpressConnectRouterRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeExpressConnectRouterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeExpressConnectRouterRequest) GetTag() []*DescribeExpressConnectRouterRequestTag {
	return s.Tag
}

func (s *DescribeExpressConnectRouterRequest) GetVersion() *string {
	return s.Version
}

func (s *DescribeExpressConnectRouterRequest) SetClientToken(v string) *DescribeExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetDryRun(v bool) *DescribeExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetEcrId(v string) *DescribeExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetMaxResults(v int32) *DescribeExpressConnectRouterRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetName(v string) *DescribeExpressConnectRouterRequest {
	s.Name = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetNextToken(v string) *DescribeExpressConnectRouterRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetResourceGroupId(v string) *DescribeExpressConnectRouterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetTag(v []*DescribeExpressConnectRouterRequestTag) *DescribeExpressConnectRouterRequest {
	s.Tag = v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetVersion(v string) *DescribeExpressConnectRouterRequest {
	s.Version = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) Validate() error {
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

type DescribeExpressConnectRouterRequestTag struct {
	// The tag key of the instance. If you specify this parameter, it cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`. If you specify this parameter, it can be an empty string.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeExpressConnectRouterRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeExpressConnectRouterRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeExpressConnectRouterRequestTag) SetKey(v string) *DescribeExpressConnectRouterRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeExpressConnectRouterRequestTag) SetValue(v string) *DescribeExpressConnectRouterRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeExpressConnectRouterRequestTag) Validate() error {
	return dara.Validate(s)
}
