// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUmodelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateUmodelRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateUmodelRequest
	GetDryRun() *bool
}

type CreateUmodelRequest struct {
	// The description of the Umodel.
	//
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - true: performs a dry run. The system checks the request but does not make actual changes.
	//
	// - false: sends a normal request.
	//
	// Default value: false.
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateUmodelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUmodelRequest) GoString() string {
	return s.String()
}

func (s *CreateUmodelRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUmodelRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateUmodelRequest) SetDescription(v string) *CreateUmodelRequest {
	s.Description = &v
	return s
}

func (s *CreateUmodelRequest) SetDryRun(v bool) *CreateUmodelRequest {
	s.DryRun = &v
	return s
}

func (s *CreateUmodelRequest) Validate() error {
	return dara.Validate(s)
}
