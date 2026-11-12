// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteFirewallTemplatesRequest
	GetClientToken() *string
	SetFirewallTemplateId(v []*string) *DeleteFirewallTemplatesRequest
	GetFirewallTemplateId() []*string
	SetRegionId(v string) *DeleteFirewallTemplatesRequest
	GetRegionId() *string
}

type DeleteFirewallTemplatesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the firewall templates.
	//
	// This parameter is required.
	FirewallTemplateId []*string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty" type:"Repeated"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteFirewallTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFirewallTemplatesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteFirewallTemplatesRequest) GetFirewallTemplateId() []*string {
	return s.FirewallTemplateId
}

func (s *DeleteFirewallTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteFirewallTemplatesRequest) SetClientToken(v string) *DeleteFirewallTemplatesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFirewallTemplatesRequest) SetFirewallTemplateId(v []*string) *DeleteFirewallTemplatesRequest {
	s.FirewallTemplateId = v
	return s
}

func (s *DeleteFirewallTemplatesRequest) SetRegionId(v string) *DeleteFirewallTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFirewallTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
