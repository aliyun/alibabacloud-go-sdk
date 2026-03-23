// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApAddressByMacRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *GetApAddressByMacRequest
	GetAppCode() *string
	SetAppName(v string) *GetApAddressByMacRequest
	GetAppName() *string
	SetLanguage(v string) *GetApAddressByMacRequest
	GetLanguage() *string
	SetMac(v string) *GetApAddressByMacRequest
	GetMac() *string
}

type GetApAddressByMacRequest struct {
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
}

func (s GetApAddressByMacRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApAddressByMacRequest) GoString() string {
	return s.String()
}

func (s *GetApAddressByMacRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetApAddressByMacRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApAddressByMacRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetApAddressByMacRequest) GetMac() *string {
	return s.Mac
}

func (s *GetApAddressByMacRequest) SetAppCode(v string) *GetApAddressByMacRequest {
	s.AppCode = &v
	return s
}

func (s *GetApAddressByMacRequest) SetAppName(v string) *GetApAddressByMacRequest {
	s.AppName = &v
	return s
}

func (s *GetApAddressByMacRequest) SetLanguage(v string) *GetApAddressByMacRequest {
	s.Language = &v
	return s
}

func (s *GetApAddressByMacRequest) SetMac(v string) *GetApAddressByMacRequest {
	s.Mac = &v
	return s
}

func (s *GetApAddressByMacRequest) Validate() error {
	return dara.Validate(s)
}
