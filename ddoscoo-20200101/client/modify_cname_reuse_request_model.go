// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCnameReuseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCname(v string) *ModifyCnameReuseRequest
	GetCname() *string
	SetDomain(v string) *ModifyCnameReuseRequest
	GetDomain() *string
	SetEnable(v int32) *ModifyCnameReuseRequest
	GetEnable() *int32
	SetResourceGroupId(v string) *ModifyCnameReuseRequest
	GetResourceGroupId() *string
}

type ModifyCnameReuseRequest struct {
	// The CNAME record that you want to reuse for the website.
	//
	// example:
	//
	// 4o6ep6q217k9****.aliyunddos0004.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to enable CNAME reuse. Valid values:
	//
	// 	- **0:*	- disabled
	//
	// 	- **1:*	- enabled
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

func (s ModifyCnameReuseRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCnameReuseRequest) GoString() string {
	return s.String()
}

func (s *ModifyCnameReuseRequest) GetCname() *string {
	return s.Cname
}

func (s *ModifyCnameReuseRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyCnameReuseRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *ModifyCnameReuseRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyCnameReuseRequest) SetCname(v string) *ModifyCnameReuseRequest {
	s.Cname = &v
	return s
}

func (s *ModifyCnameReuseRequest) SetDomain(v string) *ModifyCnameReuseRequest {
	s.Domain = &v
	return s
}

func (s *ModifyCnameReuseRequest) SetEnable(v int32) *ModifyCnameReuseRequest {
	s.Enable = &v
	return s
}

func (s *ModifyCnameReuseRequest) SetResourceGroupId(v string) *ModifyCnameReuseRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyCnameReuseRequest) Validate() error {
	return dara.Validate(s)
}
