// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserAnalyzerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusiness(v string) *CreateUserAnalyzerRequest
	GetBusiness() *string
	SetBusinessAppGroupId(v string) *CreateUserAnalyzerRequest
	GetBusinessAppGroupId() *string
	SetBusinessType(v string) *CreateUserAnalyzerRequest
	GetBusinessType() *string
	SetName(v string) *CreateUserAnalyzerRequest
	GetName() *string
	SetType(v string) *CreateUserAnalyzerRequest
	GetType() *string
	SetDryRun(v bool) *CreateUserAnalyzerRequest
	GetDryRun() *bool
}

type CreateUserAnalyzerRequest struct {
	// The basic analyzer.
	//
	// example:
	//
	// chn_standard
	Business *string `json:"business,omitempty" xml:"business,omitempty"`
	// The application ID of the custom analyzer.
	//
	// example:
	//
	// 110123123
	BusinessAppGroupId *string `json:"businessAppGroupId,omitempty" xml:"businessAppGroupId,omitempty"`
	// The basic analyzer type. Valid values: AUTO, MODEL, SYSTEM, and USER.
	//
	// example:
	//
	// AUTO
	BusinessType *string `json:"businessType,omitempty" xml:"businessType,omitempty"`
	// The analyzer name.
	//
	// example:
	//
	// jmbon_analyzer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The engine type. Valid values: HA3 and ES.
	//
	// example:
	//
	// HA3
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Default value: false.
	//
	// Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateUserAnalyzerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserAnalyzerRequest) GoString() string {
	return s.String()
}

func (s *CreateUserAnalyzerRequest) GetBusiness() *string {
	return s.Business
}

func (s *CreateUserAnalyzerRequest) GetBusinessAppGroupId() *string {
	return s.BusinessAppGroupId
}

func (s *CreateUserAnalyzerRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateUserAnalyzerRequest) GetName() *string {
	return s.Name
}

func (s *CreateUserAnalyzerRequest) GetType() *string {
	return s.Type
}

func (s *CreateUserAnalyzerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateUserAnalyzerRequest) SetBusiness(v string) *CreateUserAnalyzerRequest {
	s.Business = &v
	return s
}

func (s *CreateUserAnalyzerRequest) SetBusinessAppGroupId(v string) *CreateUserAnalyzerRequest {
	s.BusinessAppGroupId = &v
	return s
}

func (s *CreateUserAnalyzerRequest) SetBusinessType(v string) *CreateUserAnalyzerRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateUserAnalyzerRequest) SetName(v string) *CreateUserAnalyzerRequest {
	s.Name = &v
	return s
}

func (s *CreateUserAnalyzerRequest) SetType(v string) *CreateUserAnalyzerRequest {
	s.Type = &v
	return s
}

func (s *CreateUserAnalyzerRequest) SetDryRun(v bool) *CreateUserAnalyzerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateUserAnalyzerRequest) Validate() error {
	return dara.Validate(s)
}
