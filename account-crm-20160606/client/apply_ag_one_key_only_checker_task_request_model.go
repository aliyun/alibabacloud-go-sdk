// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAgOneKeyOnlyCheckerTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgAccountType(v string) *ApplyAgOneKeyOnlyCheckerTaskRequest
	GetAgAccountType() *string
	SetAppName(v string) *ApplyAgOneKeyOnlyCheckerTaskRequest
	GetAppName() *string
	SetMpk(v string) *ApplyAgOneKeyOnlyCheckerTaskRequest
	GetMpk() *string
	SetPk(v string) *ApplyAgOneKeyOnlyCheckerTaskRequest
	GetPk() *string
}

type ApplyAgOneKeyOnlyCheckerTaskRequest struct {
	// This parameter is required.
	AgAccountType *string `json:"AgAccountType,omitempty" xml:"AgAccountType,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s ApplyAgOneKeyOnlyCheckerTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyAgOneKeyOnlyCheckerTaskRequest) GoString() string {
	return s.String()
}

func (s *ApplyAgOneKeyOnlyCheckerTaskRequest) GetAgAccountType() *string {
	return s.AgAccountType
}

func (s *ApplyAgOneKeyOnlyCheckerTaskRequest) GetAppName() *string {
	return s.AppName
}

func (s *ApplyAgOneKeyOnlyCheckerTaskRequest) GetMpk() *string {
	return s.Mpk
}

func (s *ApplyAgOneKeyOnlyCheckerTaskRequest) GetPk() *string {
	return s.Pk
}

func (s *ApplyAgOneKeyOnlyCheckerTaskRequest) SetAgAccountType(v string) *ApplyAgOneKeyOnlyCheckerTaskRequest {
	s.AgAccountType = &v
	return s
}

func (s *ApplyAgOneKeyOnlyCheckerTaskRequest) SetAppName(v string) *ApplyAgOneKeyOnlyCheckerTaskRequest {
	s.AppName = &v
	return s
}

func (s *ApplyAgOneKeyOnlyCheckerTaskRequest) SetMpk(v string) *ApplyAgOneKeyOnlyCheckerTaskRequest {
	s.Mpk = &v
	return s
}

func (s *ApplyAgOneKeyOnlyCheckerTaskRequest) SetPk(v string) *ApplyAgOneKeyOnlyCheckerTaskRequest {
	s.Pk = &v
	return s
}

func (s *ApplyAgOneKeyOnlyCheckerTaskRequest) Validate() error {
	return dara.Validate(s)
}
