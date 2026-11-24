// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UntagResourcesRequest
	GetAll() *bool
	SetClientToken(v string) *UntagResourcesRequest
	GetClientToken() *string
	SetDryRun(v bool) *UntagResourcesRequest
	GetDryRun() *bool
	SetResourceId(v []*string) *UntagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UntagResourcesRequest
	GetTagKey() []*string
	SetVersion(v string) *UntagResourcesRequest
	GetVersion() *string
}

type UntagResourcesRequest struct {
	// Specifies whether to unbind all tags. This parameter is valid only when **TagKey*	- is empty. Valid values:
	//
	// - **true**: Unbinds all tags from the Express Connect Router (ECR) instance.
	//
	// - **false*	- (default): Does not unbind all tags from the Express Connect Router (ECR) instance.
	//
	// > You must specify either **TagKey*	- or **All**.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a parameter value from your client to ensure uniqueness across different requests. ClientToken supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - **true**: Performs a dry run without modifying EPS resources.
	//
	// - **false*	- (default): Sends the request normally.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The list of resource instance IDs. You can unbind tags from up to 50 resource instances at a time.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid values:
	//
	// - **EXPRESSCONNECTROUTER**: Express Connect Router (ECR) instance.
	//
	// - **EXPRESSCONNECTROUTERVPCASSOCIATION**: the associate relationship between an ECR and a VPC.
	//
	// - **EXPRESSCONNECTROUTERTRASSOCIATION**: the associate relationship between an ECR and a transit router.
	//
	// - **ECRFLOWLOG**: ECR flow log.
	//
	// This parameter is required.
	//
	// example:
	//
	// EXPRESSCONNECTROUTER
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys of the Express Connect Router (ECR) instance to unbind.
	TagKey  []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
	Version *string   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) GetAll() *bool {
	return s.All
}

func (s *UntagResourcesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UntagResourcesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UntagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UntagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UntagResourcesRequest) GetVersion() *string {
	return s.Version
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetClientToken(v string) *UntagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *UntagResourcesRequest) SetDryRun(v bool) *UntagResourcesRequest {
	s.DryRun = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesRequest) SetVersion(v string) *UntagResourcesRequest {
	s.Version = &v
	return s
}

func (s *UntagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
