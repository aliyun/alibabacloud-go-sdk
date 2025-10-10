// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHealthCheckTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteHealthCheckTemplatesRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteHealthCheckTemplatesRequest
	GetDryRun() *bool
	SetHealthCheckTemplateIds(v []*string) *DeleteHealthCheckTemplatesRequest
	GetHealthCheckTemplateIds() []*string
}

type DeleteHealthCheckTemplatesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a **2xx HTTP*	- status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IDs of health check templates. You can specify at most 10 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-123
	HealthCheckTemplateIds []*string `json:"HealthCheckTemplateIds,omitempty" xml:"HealthCheckTemplateIds,omitempty" type:"Repeated"`
}

func (s DeleteHealthCheckTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHealthCheckTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteHealthCheckTemplatesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteHealthCheckTemplatesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteHealthCheckTemplatesRequest) GetHealthCheckTemplateIds() []*string {
	return s.HealthCheckTemplateIds
}

func (s *DeleteHealthCheckTemplatesRequest) SetClientToken(v string) *DeleteHealthCheckTemplatesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteHealthCheckTemplatesRequest) SetDryRun(v bool) *DeleteHealthCheckTemplatesRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteHealthCheckTemplatesRequest) SetHealthCheckTemplateIds(v []*string) *DeleteHealthCheckTemplatesRequest {
	s.HealthCheckTemplateIds = v
	return s
}

func (s *DeleteHealthCheckTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
