// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProviderDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProviderVersion(v string) *GetProviderDocumentRequest
	GetProviderVersion() *string
	SetTerraformResourceType(v string) *GetProviderDocumentRequest
	GetTerraformResourceType() *string
}

type GetProviderDocumentRequest struct {
	ProviderVersion *string `json:"providerVersion,omitempty" xml:"providerVersion,omitempty"`
	// This parameter is required.
	TerraformResourceType *string `json:"terraformResourceType,omitempty" xml:"terraformResourceType,omitempty"`
}

func (s GetProviderDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProviderDocumentRequest) GoString() string {
	return s.String()
}

func (s *GetProviderDocumentRequest) GetProviderVersion() *string {
	return s.ProviderVersion
}

func (s *GetProviderDocumentRequest) GetTerraformResourceType() *string {
	return s.TerraformResourceType
}

func (s *GetProviderDocumentRequest) SetProviderVersion(v string) *GetProviderDocumentRequest {
	s.ProviderVersion = &v
	return s
}

func (s *GetProviderDocumentRequest) SetTerraformResourceType(v string) *GetProviderDocumentRequest {
	s.TerraformResourceType = &v
	return s
}

func (s *GetProviderDocumentRequest) Validate() error {
	return dara.Validate(s)
}
