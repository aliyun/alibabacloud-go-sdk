// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAgOneKeyDeleteTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbandonedDependency(v string) *ApplyAgOneKeyDeleteTaskRequest
	GetAbandonedDependency() *string
	SetAgAccountType(v string) *ApplyAgOneKeyDeleteTaskRequest
	GetAgAccountType() *string
	SetAppName(v string) *ApplyAgOneKeyDeleteTaskRequest
	GetAppName() *string
	SetMpk(v string) *ApplyAgOneKeyDeleteTaskRequest
	GetMpk() *string
	SetPk(v string) *ApplyAgOneKeyDeleteTaskRequest
	GetPk() *string
}

type ApplyAgOneKeyDeleteTaskRequest struct {
	AbandonedDependency *string `json:"AbandonedDependency,omitempty" xml:"AbandonedDependency,omitempty"`
	// This parameter is required.
	AgAccountType *string `json:"AgAccountType,omitempty" xml:"AgAccountType,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s ApplyAgOneKeyDeleteTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyAgOneKeyDeleteTaskRequest) GoString() string {
	return s.String()
}

func (s *ApplyAgOneKeyDeleteTaskRequest) GetAbandonedDependency() *string {
	return s.AbandonedDependency
}

func (s *ApplyAgOneKeyDeleteTaskRequest) GetAgAccountType() *string {
	return s.AgAccountType
}

func (s *ApplyAgOneKeyDeleteTaskRequest) GetAppName() *string {
	return s.AppName
}

func (s *ApplyAgOneKeyDeleteTaskRequest) GetMpk() *string {
	return s.Mpk
}

func (s *ApplyAgOneKeyDeleteTaskRequest) GetPk() *string {
	return s.Pk
}

func (s *ApplyAgOneKeyDeleteTaskRequest) SetAbandonedDependency(v string) *ApplyAgOneKeyDeleteTaskRequest {
	s.AbandonedDependency = &v
	return s
}

func (s *ApplyAgOneKeyDeleteTaskRequest) SetAgAccountType(v string) *ApplyAgOneKeyDeleteTaskRequest {
	s.AgAccountType = &v
	return s
}

func (s *ApplyAgOneKeyDeleteTaskRequest) SetAppName(v string) *ApplyAgOneKeyDeleteTaskRequest {
	s.AppName = &v
	return s
}

func (s *ApplyAgOneKeyDeleteTaskRequest) SetMpk(v string) *ApplyAgOneKeyDeleteTaskRequest {
	s.Mpk = &v
	return s
}

func (s *ApplyAgOneKeyDeleteTaskRequest) SetPk(v string) *ApplyAgOneKeyDeleteTaskRequest {
	s.Pk = &v
	return s
}

func (s *ApplyAgOneKeyDeleteTaskRequest) Validate() error {
	return dara.Validate(s)
}
