// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *UpdateProductVersionRequest
	GetActive() *bool
	SetDescription(v string) *UpdateProductVersionRequest
	GetDescription() *string
	SetGuidance(v string) *UpdateProductVersionRequest
	GetGuidance() *string
	SetProductVersionId(v string) *UpdateProductVersionRequest
	GetProductVersionId() *string
	SetProductVersionName(v string) *UpdateProductVersionRequest
	GetProductVersionName() *string
}

type UpdateProductVersionRequest struct {
	// Specifies whether to enable the product version. Valid values:
	//
	// 	- true: enables the product version. This is the default value.
	//
	// 	- false: disables the product version.
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The description of the product version.
	//
	// The value must be 1 to 128 characters in length.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// The description of the product version.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The recommendation information. Valid values:
	//
	// 	- Default: No recommendation information is provided. This is the default value.
	//
	// 	- Recommended: the recommended version.
	//
	// 	- Latest: the latest version.
	//
	// 	- Deprecated: the version that is about to be discontinued.
	//
	// example:
	//
	// Default
	Guidance *string `json:"Guidance,omitempty" xml:"Guidance,omitempty"`
	// The ID of the product version.
	//
	// This parameter is required.
	//
	// example:
	//
	// pv-bp15e79d26****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The name of the product version.
	//
	// The value must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
}

func (s UpdateProductVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionRequest) GetActive() *bool {
	return s.Active
}

func (s *UpdateProductVersionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProductVersionRequest) GetGuidance() *string {
	return s.Guidance
}

func (s *UpdateProductVersionRequest) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *UpdateProductVersionRequest) GetProductVersionName() *string {
	return s.ProductVersionName
}

func (s *UpdateProductVersionRequest) SetActive(v bool) *UpdateProductVersionRequest {
	s.Active = &v
	return s
}

func (s *UpdateProductVersionRequest) SetDescription(v string) *UpdateProductVersionRequest {
	s.Description = &v
	return s
}

func (s *UpdateProductVersionRequest) SetGuidance(v string) *UpdateProductVersionRequest {
	s.Guidance = &v
	return s
}

func (s *UpdateProductVersionRequest) SetProductVersionId(v string) *UpdateProductVersionRequest {
	s.ProductVersionId = &v
	return s
}

func (s *UpdateProductVersionRequest) SetProductVersionName(v string) *UpdateProductVersionRequest {
	s.ProductVersionName = &v
	return s
}

func (s *UpdateProductVersionRequest) Validate() error {
	return dara.Validate(s)
}
