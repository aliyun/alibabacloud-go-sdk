// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExplorerModuleAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateExplorerModuleAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateExplorerModuleAttributeRequest
	GetDescription() *string
	SetName(v string) *UpdateExplorerModuleAttributeRequest
	GetName() *string
}

type UpdateExplorerModuleAttributeRequest struct {
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateExplorerModuleAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExplorerModuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateExplorerModuleAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateExplorerModuleAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateExplorerModuleAttributeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateExplorerModuleAttributeRequest) SetClientToken(v string) *UpdateExplorerModuleAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateExplorerModuleAttributeRequest) SetDescription(v string) *UpdateExplorerModuleAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateExplorerModuleAttributeRequest) SetName(v string) *UpdateExplorerModuleAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateExplorerModuleAttributeRequest) Validate() error {
	return dara.Validate(s)
}
