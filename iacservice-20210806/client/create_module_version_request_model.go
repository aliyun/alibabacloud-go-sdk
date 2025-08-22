// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModuleVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateModuleVersionRequest
	GetClientToken() *string
	SetDescription(v string) *CreateModuleVersionRequest
	GetDescription() *string
	SetName(v string) *CreateModuleVersionRequest
	GetName() *string
}

type CreateModuleVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateModuleVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModuleVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateModuleVersionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateModuleVersionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateModuleVersionRequest) GetName() *string {
	return s.Name
}

func (s *CreateModuleVersionRequest) SetClientToken(v string) *CreateModuleVersionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModuleVersionRequest) SetDescription(v string) *CreateModuleVersionRequest {
	s.Description = &v
	return s
}

func (s *CreateModuleVersionRequest) SetName(v string) *CreateModuleVersionRequest {
	s.Name = &v
	return s
}

func (s *CreateModuleVersionRequest) Validate() error {
	return dara.Validate(s)
}
