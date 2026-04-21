// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoDisposeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetAutoDisposeConfigRequest
	GetLang() *string
	SetProductCode(v string) *GetAutoDisposeConfigRequest
	GetProductCode() *string
}

type GetAutoDisposeConfigRequest struct {
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

func (s GetAutoDisposeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutoDisposeConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAutoDisposeConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *GetAutoDisposeConfigRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetAutoDisposeConfigRequest) SetLang(v string) *GetAutoDisposeConfigRequest {
	s.Lang = &v
	return s
}

func (s *GetAutoDisposeConfigRequest) SetProductCode(v string) *GetAutoDisposeConfigRequest {
	s.ProductCode = &v
	return s
}

func (s *GetAutoDisposeConfigRequest) Validate() error {
	return dara.Validate(s)
}
