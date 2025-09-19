// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBinarySecurityPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteBinarySecurityPolicyRequest
	GetName() *string
	SetResourceOwnerId(v int64) *DeleteBinarySecurityPolicyRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DeleteBinarySecurityPolicyRequest
	GetSourceIp() *string
}

type DeleteBinarySecurityPolicyRequest struct {
	// The name of the binary security policy.
	//
	// example:
	//
	// policy-auto-bfu7pm
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 42.120.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteBinarySecurityPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBinarySecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteBinarySecurityPolicyRequest) GetName() *string {
	return s.Name
}

func (s *DeleteBinarySecurityPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteBinarySecurityPolicyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteBinarySecurityPolicyRequest) SetName(v string) *DeleteBinarySecurityPolicyRequest {
	s.Name = &v
	return s
}

func (s *DeleteBinarySecurityPolicyRequest) SetResourceOwnerId(v int64) *DeleteBinarySecurityPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteBinarySecurityPolicyRequest) SetSourceIp(v string) *DeleteBinarySecurityPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteBinarySecurityPolicyRequest) Validate() error {
	return dara.Validate(s)
}
