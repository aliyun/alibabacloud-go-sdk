// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *CreateCredentialsRequest
	GetType() *string
	SetDryRun(v bool) *CreateCredentialsRequest
	GetDryRun() *bool
}

type CreateCredentialsRequest struct {
	// The credential type. Valid types:
	//
	// - api-token
	//
	// example:
	//
	// api-token
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Specifies whether to only validate the request parameters. Default value: false.
	//
	// Valid values:
	//
	// - **true**: Only validates the request parameters.
	//
	// - **false**: Validates the request parameters and creates the attribution configuration.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialsRequest) GoString() string {
	return s.String()
}

func (s *CreateCredentialsRequest) GetType() *string {
	return s.Type
}

func (s *CreateCredentialsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateCredentialsRequest) SetType(v string) *CreateCredentialsRequest {
	s.Type = &v
	return s
}

func (s *CreateCredentialsRequest) SetDryRun(v bool) *CreateCredentialsRequest {
	s.DryRun = &v
	return s
}

func (s *CreateCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
