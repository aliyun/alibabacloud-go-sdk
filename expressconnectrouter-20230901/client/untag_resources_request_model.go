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
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags. This parameter is valid only when the **TagKey*	- parameter is not specified. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >  You must specify one of **TagKey*	- and **All**.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Set the value to **EXPRESSCONNECTROUTER**.
	//
	// This parameter is required.
	//
	// example:
	//
	// EXPRESSCONNECTROUTER
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The keys of the tags that you want to remove from the ECRs.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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

func (s *UntagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
