// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *StartInstanceRequest
	GetAppType() *string
	SetDepartmentId(v string) *StartInstanceRequest
	GetDepartmentId() *string
	SetFormDataJson(v string) *StartInstanceRequest
	GetFormDataJson() *string
	SetFormUuid(v string) *StartInstanceRequest
	GetFormUuid() *string
	SetLanguage(v string) *StartInstanceRequest
	GetLanguage() *string
	SetProcessCode(v string) *StartInstanceRequest
	GetProcessCode() *string
	SetProcessData(v string) *StartInstanceRequest
	GetProcessData() *string
	SetSystemToken(v string) *StartInstanceRequest
	GetSystemToken() *string
}

type StartInstanceRequest struct {
	// example:
	//
	// APP_PBxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 18295
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// example:
	//
	// {}
	FormDataJson *string `json:"FormDataJson,omitempty" xml:"FormDataJson,omitempty"`
	// example:
	//
	// FORM-EF6Yxxx
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// TPROC--EF6Y4xxx
	ProcessCode *string `json:"ProcessCode,omitempty" xml:"ProcessCode,omitempty"`
	ProcessData *string `json:"ProcessData,omitempty" xml:"ProcessData,omitempty"`
	// example:
	//
	// hexxyy
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) GetAppType() *string {
	return s.AppType
}

func (s *StartInstanceRequest) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *StartInstanceRequest) GetFormDataJson() *string {
	return s.FormDataJson
}

func (s *StartInstanceRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *StartInstanceRequest) GetLanguage() *string {
	return s.Language
}

func (s *StartInstanceRequest) GetProcessCode() *string {
	return s.ProcessCode
}

func (s *StartInstanceRequest) GetProcessData() *string {
	return s.ProcessData
}

func (s *StartInstanceRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *StartInstanceRequest) SetAppType(v string) *StartInstanceRequest {
	s.AppType = &v
	return s
}

func (s *StartInstanceRequest) SetDepartmentId(v string) *StartInstanceRequest {
	s.DepartmentId = &v
	return s
}

func (s *StartInstanceRequest) SetFormDataJson(v string) *StartInstanceRequest {
	s.FormDataJson = &v
	return s
}

func (s *StartInstanceRequest) SetFormUuid(v string) *StartInstanceRequest {
	s.FormUuid = &v
	return s
}

func (s *StartInstanceRequest) SetLanguage(v string) *StartInstanceRequest {
	s.Language = &v
	return s
}

func (s *StartInstanceRequest) SetProcessCode(v string) *StartInstanceRequest {
	s.ProcessCode = &v
	return s
}

func (s *StartInstanceRequest) SetProcessData(v string) *StartInstanceRequest {
	s.ProcessData = &v
	return s
}

func (s *StartInstanceRequest) SetSystemToken(v string) *StartInstanceRequest {
	s.SystemToken = &v
	return s
}

func (s *StartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
