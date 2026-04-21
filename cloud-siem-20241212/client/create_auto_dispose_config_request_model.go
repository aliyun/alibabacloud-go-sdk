// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoDisposeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDecisionStatus(v string) *CreateAutoDisposeConfigRequest
	GetAutoDecisionStatus() *string
	SetLang(v string) *CreateAutoDisposeConfigRequest
	GetLang() *string
	SetProductCode(v string) *CreateAutoDisposeConfigRequest
	GetProductCode() *string
}

type CreateAutoDisposeConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// enabled
	AutoDecisionStatus *string `json:"AutoDecisionStatus,omitempty" xml:"AutoDecisionStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alibaba_cloud_sas
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s CreateAutoDisposeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoDisposeConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoDisposeConfigRequest) GetAutoDecisionStatus() *string {
	return s.AutoDecisionStatus
}

func (s *CreateAutoDisposeConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateAutoDisposeConfigRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateAutoDisposeConfigRequest) SetAutoDecisionStatus(v string) *CreateAutoDisposeConfigRequest {
	s.AutoDecisionStatus = &v
	return s
}

func (s *CreateAutoDisposeConfigRequest) SetLang(v string) *CreateAutoDisposeConfigRequest {
	s.Lang = &v
	return s
}

func (s *CreateAutoDisposeConfigRequest) SetProductCode(v string) *CreateAutoDisposeConfigRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateAutoDisposeConfigRequest) Validate() error {
	return dara.Validate(s)
}
