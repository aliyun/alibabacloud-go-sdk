// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSasModuleTrialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *AddSasModuleTrialRequest
	GetLang() *string
	SetModuleCode(v string) *AddSasModuleTrialRequest
	GetModuleCode() *string
}

type AddSasModuleTrialRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The code of the feature. Valid values:
	//
	// 	- **vulFix**: vulnerability fixing.
	//
	// 	- **cloudSiem**: threat analysis and response.
	//
	// example:
	//
	// vulFix
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
}

func (s AddSasModuleTrialRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSasModuleTrialRequest) GoString() string {
	return s.String()
}

func (s *AddSasModuleTrialRequest) GetLang() *string {
	return s.Lang
}

func (s *AddSasModuleTrialRequest) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *AddSasModuleTrialRequest) SetLang(v string) *AddSasModuleTrialRequest {
	s.Lang = &v
	return s
}

func (s *AddSasModuleTrialRequest) SetModuleCode(v string) *AddSasModuleTrialRequest {
	s.ModuleCode = &v
	return s
}

func (s *AddSasModuleTrialRequest) Validate() error {
	return dara.Validate(s)
}
