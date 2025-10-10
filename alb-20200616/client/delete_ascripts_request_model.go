// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAScriptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAScriptIds(v []*string) *DeleteAScriptsRequest
	GetAScriptIds() []*string
	SetClientToken(v string) *DeleteAScriptsRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteAScriptsRequest
	GetDryRun() *bool
}

type DeleteAScriptsRequest struct {
	// The AScript rule IDs.
	//
	// This parameter is required.
	AScriptIds []*string `json:"AScriptIds,omitempty" xml:"AScriptIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// f516e84e-fc0c-4c2d-a461-6cd774a84dbd
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**(default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s DeleteAScriptsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAScriptsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAScriptsRequest) GetAScriptIds() []*string {
	return s.AScriptIds
}

func (s *DeleteAScriptsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAScriptsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteAScriptsRequest) SetAScriptIds(v []*string) *DeleteAScriptsRequest {
	s.AScriptIds = v
	return s
}

func (s *DeleteAScriptsRequest) SetClientToken(v string) *DeleteAScriptsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAScriptsRequest) SetDryRun(v bool) *DeleteAScriptsRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteAScriptsRequest) Validate() error {
	return dara.Validate(s)
}
