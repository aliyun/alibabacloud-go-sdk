// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoDisposeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDecisionStatus(v string) *UpdateAutoDisposeConfigRequest
	GetAutoDecisionStatus() *string
	SetLang(v string) *UpdateAutoDisposeConfigRequest
	GetLang() *string
	SetProductCode(v string) *UpdateAutoDisposeConfigRequest
	GetProductCode() *string
}

type UpdateAutoDisposeConfigRequest struct {
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

func (s UpdateAutoDisposeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoDisposeConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutoDisposeConfigRequest) GetAutoDecisionStatus() *string {
	return s.AutoDecisionStatus
}

func (s *UpdateAutoDisposeConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateAutoDisposeConfigRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *UpdateAutoDisposeConfigRequest) SetAutoDecisionStatus(v string) *UpdateAutoDisposeConfigRequest {
	s.AutoDecisionStatus = &v
	return s
}

func (s *UpdateAutoDisposeConfigRequest) SetLang(v string) *UpdateAutoDisposeConfigRequest {
	s.Lang = &v
	return s
}

func (s *UpdateAutoDisposeConfigRequest) SetProductCode(v string) *UpdateAutoDisposeConfigRequest {
	s.ProductCode = &v
	return s
}

func (s *UpdateAutoDisposeConfigRequest) Validate() error {
	return dara.Validate(s)
}
