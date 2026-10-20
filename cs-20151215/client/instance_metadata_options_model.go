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
	// The ECS instance metadata access mode configuration. Valid values:
	//
	// 	- `optional`: Compatible with both normal mode and hardened mode.
	//
	// 	- `required`: Enables hardened mode only (IMDSv2). After this mode is enabled, applications in the node cannot access ECS instance metadata in normal mode.
	//
	// Default value: `optional`.
	//
	// 	Notice:
	//
	// 	- This parameter is currently available only by whitelist. To use this parameter, submit a ticket.
	//
	// 	- Only ACK managed clusters of version 1.28 and later support this parameter.
	//
	//
	//
	// > For more information about instance metadata access modes, see [Instance metadata access modes](https://help.aliyun.com/document_detail/108460.html).
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
