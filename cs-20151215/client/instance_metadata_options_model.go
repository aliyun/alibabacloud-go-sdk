// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceMetadataOptions interface {
	dara.Model
	String() string
	GoString() string
	SetHttpTokens(v string) *InstanceMetadataOptions
	GetHttpTokens() *string
}

type InstanceMetadataOptions struct {
	// Configuration of the ECS instance metadata access mode. Valid values:
	//
	// 	- `optional`: Compatible with both NAT mode and hardened mode.
	//
	// 	- `required`: Enables hardened mode only (IMDSv2). After this mode is enabled, applications within edge zones cannot access ECS instance metadata through NAT mode.
	//
	// Default Value: `optional`.
	//
	// 	Notice:
	//
	// 	- This parameter is currently available only to users on the whitelist. To use it, submit a ticket to request access.
	//
	// 	- This parameter is supported only in ACK managed clusters of version 1.28 or later.
	//
	//
	//
	// > For more information about instance metadata access modes, see [Instance Metadata Access Mode](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"http_tokens,omitempty" xml:"http_tokens,omitempty"`
}

func (s InstanceMetadataOptions) String() string {
	return dara.Prettify(s)
}

func (s InstanceMetadataOptions) GoString() string {
	return s.String()
}

func (s *InstanceMetadataOptions) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *InstanceMetadataOptions) SetHttpTokens(v string) *InstanceMetadataOptions {
	s.HttpTokens = &v
	return s
}

func (s *InstanceMetadataOptions) Validate() error {
	return dara.Validate(s)
}
