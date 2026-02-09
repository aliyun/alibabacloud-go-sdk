// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachment interface {
	dara.Model
	String() string
	GoString() string
	SetAttachResourceIds(v []*string) *Attachment
	GetAttachResourceIds() []*string
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
	// The resource IDs.
	AttachResourceIds []*string `json:"attachResourceIds,omitempty" xml:"attachResourceIds,omitempty" type:"Repeated"`
	// The supported mount point type. Valid values:
	//
	// 	- HttpApi: an HTTP API
	//
	// 	- Operation: an operation in an HTTP API
	//
	// 	- GatewayRoute: a gateway route
	//
	// 	- GatewayService: a gateway service
	//
	// 	- GatewayServicePort: a gateway service port
	//
	// 	- Domain: a gateway domain name
	//
	// 	- Gateway: a gateway instance
	//
	// example:
	//
	// HttpApi
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// The environment to which the mounted resource belongs. If an asterisk (\\*) is returned as the environment ID, the mounted resource is not related to the environment.
	//
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The instance to which the mounted resource belongs.
	//
	// example:
	//
	// gw-cpr4f9dlhtgq5ksfgmb0
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The mount ID.
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

func (s *Attachment) GetAttachResourceIds() []*string {
	return s.AttachResourceIds
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

func (s *Attachment) SetAttachResourceIds(v []*string) *Attachment {
	s.AttachResourceIds = v
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
