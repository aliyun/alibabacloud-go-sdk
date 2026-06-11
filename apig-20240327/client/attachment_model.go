// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachment interface {
	dara.Model
	String() string
	GoString() string
	SetAttachResourceId(v string) *Attachment
	GetAttachResourceId() *string
	SetAttachResourceIds(v []*string) *Attachment
	GetAttachResourceIds() []*string
	SetAttachResourceParentIds(v []*string) *Attachment
	GetAttachResourceParentIds() []*string
	SetAttachResourceType(v string) *Attachment
	GetAttachResourceType() *string
	SetEnvironmentId(v string) *Attachment
	GetEnvironmentId() *string
	SetGatewayId(v string) *Attachment
	GetGatewayId() *string
	SetPolicyAttachmentId(v string) *Attachment
	GetPolicyAttachmentId() *string
}

type Attachment struct {
	// The attachment point ID.
	AttachResourceId *string `json:"attachResourceId,omitempty" xml:"attachResourceId,omitempty"`
	// A list of attached resource IDs.
	AttachResourceIds []*string `json:"attachResourceIds,omitempty" xml:"attachResourceIds,omitempty" type:"Repeated"`
	// A list of parent resource IDs.
	AttachResourceParentIds []*string `json:"attachResourceParentIds,omitempty" xml:"attachResourceParentIds,omitempty" type:"Repeated"`
	// The supported attachment point types for the policy.
	//
	// - `HttpApi`: An HTTP API.
	//
	// - `Operation`: An operation of an HTTP API.
	//
	// - `GatewayRoute`: A gateway route.
	//
	// - `GatewayService`: A gateway service.
	//
	// - `GatewayServicePort`: A gateway service port.
	//
	// - `Domain`: A gateway domain.
	//
	// - `Gateway`: A gateway.
	//
	// example:
	//
	// HttpApi
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// The ID of the environment for the attached resource. An asterisk (`*`) indicates that the policy attachment is not environment-specific.
	//
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The ID of the gateway for the attached resource.
	//
	// example:
	//
	// gw-cpr4f9dlhtgq5ksfgmb0
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The policy attachment ID.
	//
	// example:
	//
	// pr-cq7l5s5lhtgi6qasrdc0
	PolicyAttachmentId *string `json:"policyAttachmentId,omitempty" xml:"policyAttachmentId,omitempty"`
}

func (s Attachment) String() string {
	return dara.Prettify(s)
}

func (s Attachment) GoString() string {
	return s.String()
}

func (s *Attachment) GetAttachResourceId() *string {
	return s.AttachResourceId
}

func (s *Attachment) GetAttachResourceIds() []*string {
	return s.AttachResourceIds
}

func (s *Attachment) GetAttachResourceParentIds() []*string {
	return s.AttachResourceParentIds
}

func (s *Attachment) GetAttachResourceType() *string {
	return s.AttachResourceType
}

func (s *Attachment) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *Attachment) GetGatewayId() *string {
	return s.GatewayId
}

func (s *Attachment) GetPolicyAttachmentId() *string {
	return s.PolicyAttachmentId
}

func (s *Attachment) SetAttachResourceId(v string) *Attachment {
	s.AttachResourceId = &v
	return s
}

func (s *Attachment) SetAttachResourceIds(v []*string) *Attachment {
	s.AttachResourceIds = v
	return s
}

func (s *Attachment) SetAttachResourceParentIds(v []*string) *Attachment {
	s.AttachResourceParentIds = v
	return s
}

func (s *Attachment) SetAttachResourceType(v string) *Attachment {
	s.AttachResourceType = &v
	return s
}

func (s *Attachment) SetEnvironmentId(v string) *Attachment {
	s.EnvironmentId = &v
	return s
}

func (s *Attachment) SetGatewayId(v string) *Attachment {
	s.GatewayId = &v
	return s
}

func (s *Attachment) SetPolicyAttachmentId(v string) *Attachment {
	s.PolicyAttachmentId = &v
	return s
}

func (s *Attachment) Validate() error {
	return dara.Validate(s)
}
