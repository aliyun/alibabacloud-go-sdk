// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleTrialAuthInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetModuleTrialAuthInfoRequest
	GetLang() *string
	SetModuleCode(v string) *GetModuleTrialAuthInfoRequest
	GetModuleCode() *string
}

type GetModuleTrialAuthInfoRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
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

func (s GetModuleTrialAuthInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModuleTrialAuthInfoRequest) GoString() string {
	return s.String()
}

func (s *GetModuleTrialAuthInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *GetModuleTrialAuthInfoRequest) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetModuleTrialAuthInfoRequest) SetLang(v string) *GetModuleTrialAuthInfoRequest {
	s.Lang = &v
	return s
}

func (s *GetModuleTrialAuthInfoRequest) SetModuleCode(v string) *GetModuleTrialAuthInfoRequest {
	s.ModuleCode = &v
	return s
}

func (s *GetModuleTrialAuthInfoRequest) Validate() error {
	return dara.Validate(s)
}
