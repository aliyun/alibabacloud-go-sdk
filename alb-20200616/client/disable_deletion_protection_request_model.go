// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDeletionProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisableDeletionProtectionRequest
	GetClientToken() *string
	SetDryRun(v bool) *DisableDeletionProtectionRequest
	GetDryRun() *bool
	SetResourceId(v string) *DisableDeletionProtectionRequest
	GetResourceId() *string
}

type DisableDeletionProtectionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ALB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// re-atstuj3rtop****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DisableDeletionProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *DisableDeletionProtectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableDeletionProtectionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DisableDeletionProtectionRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DisableDeletionProtectionRequest) SetClientToken(v string) *DisableDeletionProtectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableDeletionProtectionRequest) SetDryRun(v bool) *DisableDeletionProtectionRequest {
	s.DryRun = &v
	return s
}

func (s *DisableDeletionProtectionRequest) SetResourceId(v string) *DisableDeletionProtectionRequest {
	s.ResourceId = &v
	return s
}

func (s *DisableDeletionProtectionRequest) Validate() error {
	return dara.Validate(s)
}
