// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCacheModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ModifyWebCacheModeRequest
	GetDomain() *string
	SetMode(v string) *ModifyWebCacheModeRequest
	GetMode() *string
	SetResourceGroupId(v string) *ModifyWebCacheModeRequest
	GetResourceGroupId() *string
}

type ModifyWebCacheModeRequest struct {
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name, and the domain name must be associated with an instance that uses the Enhanced function plan. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The cache mode of the Static Page Caching policy. Valid values:
	//
	// 	- **standard**: uses the standard cache mode.
	//
	// 	- **aggressive**: uses the enhanced cache mode.
	//
	// 	- **bypass**: caches no data.
	//
	// This parameter is required.
	//
	// example:
	//
	// standard
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyWebCacheModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCacheModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheModeRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyWebCacheModeRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifyWebCacheModeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyWebCacheModeRequest) SetDomain(v string) *ModifyWebCacheModeRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebCacheModeRequest) SetMode(v string) *ModifyWebCacheModeRequest {
	s.Mode = &v
	return s
}

func (s *ModifyWebCacheModeRequest) SetResourceGroupId(v string) *ModifyWebCacheModeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebCacheModeRequest) Validate() error {
	return dara.Validate(s)
}
