// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteServerGroupRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteServerGroupRequest
	GetDryRun() *bool
	SetServerGroupId(v string) *DeleteServerGroupRequest
	GetServerGroupId() *string
}

type DeleteServerGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the server group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-cige6j****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s DeleteServerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteServerGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteServerGroupRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *DeleteServerGroupRequest) SetClientToken(v string) *DeleteServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServerGroupRequest) SetDryRun(v bool) *DeleteServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteServerGroupRequest) SetServerGroupId(v string) *DeleteServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *DeleteServerGroupRequest) Validate() error {
	return dara.Validate(s)
}
