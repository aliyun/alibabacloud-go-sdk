// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppGroupCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *CreateAppGroupCredentialsRequest
	GetType() *string
	SetDryRun(v bool) *CreateAppGroupCredentialsRequest
	GetDryRun() *bool
}

type CreateAppGroupCredentialsRequest struct {
	// example:
	//
	// api-token
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
	DryRun *bool   `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateAppGroupCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppGroupCredentialsRequest) GoString() string {
	return s.String()
}

func (s *CreateAppGroupCredentialsRequest) GetType() *string {
	return s.Type
}

func (s *CreateAppGroupCredentialsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateAppGroupCredentialsRequest) SetType(v string) *CreateAppGroupCredentialsRequest {
	s.Type = &v
	return s
}

func (s *CreateAppGroupCredentialsRequest) SetDryRun(v bool) *CreateAppGroupCredentialsRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAppGroupCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
