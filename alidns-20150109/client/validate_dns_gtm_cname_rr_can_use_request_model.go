// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateDnsGtmCnameRrCanUseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCnameMode(v string) *ValidateDnsGtmCnameRrCanUseRequest
	GetCnameMode() *string
	SetCnameRr(v string) *ValidateDnsGtmCnameRrCanUseRequest
	GetCnameRr() *string
	SetCnameType(v string) *ValidateDnsGtmCnameRrCanUseRequest
	GetCnameType() *string
	SetCnameZone(v string) *ValidateDnsGtmCnameRrCanUseRequest
	GetCnameZone() *string
	SetInstanceId(v string) *ValidateDnsGtmCnameRrCanUseRequest
	GetInstanceId() *string
	SetLang(v string) *ValidateDnsGtmCnameRrCanUseRequest
	GetLang() *string
}

type ValidateDnsGtmCnameRrCanUseRequest struct {
	// This parameter is required.
	CnameMode *string `json:"CnameMode,omitempty" xml:"CnameMode,omitempty"`
	// This parameter is required.
	CnameRr *string `json:"CnameRr,omitempty" xml:"CnameRr,omitempty"`
	// This parameter is required.
	CnameType *string `json:"CnameType,omitempty" xml:"CnameType,omitempty"`
	// This parameter is required.
	CnameZone *string `json:"CnameZone,omitempty" xml:"CnameZone,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ValidateDnsGtmCnameRrCanUseRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateDnsGtmCnameRrCanUseRequest) GoString() string {
	return s.String()
}

func (s *ValidateDnsGtmCnameRrCanUseRequest) GetCnameMode() *string {
	return s.CnameMode
}

func (s *ValidateDnsGtmCnameRrCanUseRequest) GetCnameRr() *string {
	return s.CnameRr
}

func (s *ValidateDnsGtmCnameRrCanUseRequest) GetCnameType() *string {
	return s.CnameType
}

func (s *ValidateDnsGtmCnameRrCanUseRequest) GetCnameZone() *string {
	return s.CnameZone
}

func (s *ValidateDnsGtmCnameRrCanUseRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ValidateDnsGtmCnameRrCanUseRequest) GetLang() *string {
	return s.Lang
}

func (s *ValidateDnsGtmCnameRrCanUseRequest) SetCnameMode(v string) *ValidateDnsGtmCnameRrCanUseRequest {
	s.CnameMode = &v
	return s
}

func (s *ValidateDnsGtmCnameRrCanUseRequest) SetCnameRr(v string) *ValidateDnsGtmCnameRrCanUseRequest {
	s.CnameRr = &v
	return s
}

func (s *ValidateDnsGtmCnameRrCanUseRequest) SetCnameType(v string) *ValidateDnsGtmCnameRrCanUseRequest {
	s.CnameType = &v
	return s
}

func (s *ValidateDnsGtmCnameRrCanUseRequest) SetCnameZone(v string) *ValidateDnsGtmCnameRrCanUseRequest {
	s.CnameZone = &v
	return s
}

func (s *ValidateDnsGtmCnameRrCanUseRequest) SetInstanceId(v string) *ValidateDnsGtmCnameRrCanUseRequest {
	s.InstanceId = &v
	return s
}

func (s *ValidateDnsGtmCnameRrCanUseRequest) SetLang(v string) *ValidateDnsGtmCnameRrCanUseRequest {
	s.Lang = &v
	return s
}

func (s *ValidateDnsGtmCnameRrCanUseRequest) Validate() error {
	return dara.Validate(s)
}
