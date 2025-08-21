// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHttp2EnableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ModifyHttp2EnableRequest
	GetDomain() *string
	SetEnable(v int32) *ModifyHttp2EnableRequest
	GetEnable() *int32
	SetResourceGroupId(v string) *ModifyHttp2EnableRequest
	GetResourceGroupId() *string
}

type ModifyHttp2EnableRequest struct {
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
	// Specifies whether to enable HTTP/2. Valid values:
	//
	// 	- **0**: disables HTTP/2.
	//
	// 	- **1**: enables HTTP/2.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyHttp2EnableRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHttp2EnableRequest) GoString() string {
	return s.String()
}

func (s *ModifyHttp2EnableRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyHttp2EnableRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *ModifyHttp2EnableRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyHttp2EnableRequest) SetDomain(v string) *ModifyHttp2EnableRequest {
	s.Domain = &v
	return s
}

func (s *ModifyHttp2EnableRequest) SetEnable(v int32) *ModifyHttp2EnableRequest {
	s.Enable = &v
	return s
}

func (s *ModifyHttp2EnableRequest) SetResourceGroupId(v string) *ModifyHttp2EnableRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyHttp2EnableRequest) Validate() error {
	return dara.Validate(s)
}
