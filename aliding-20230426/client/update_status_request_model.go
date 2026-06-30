// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *UpdateStatusRequest
	GetAppType() *string
	SetErrorLines(v []*int32) *UpdateStatusRequest
	GetErrorLines() []*int32
	SetImportSequence(v string) *UpdateStatusRequest
	GetImportSequence() *string
	SetLanguage(v string) *UpdateStatusRequest
	GetLanguage() *string
	SetStatus(v string) *UpdateStatusRequest
	GetStatus() *string
	SetSystemToken(v string) *UpdateStatusRequest
	GetSystemToken() *string
}

type UpdateStatusRequest struct {
	AppType        *string  `json:"AppType,omitempty" xml:"AppType,omitempty"`
	ErrorLines     []*int32 `json:"ErrorLines,omitempty" xml:"ErrorLines,omitempty" type:"Repeated"`
	ImportSequence *string  `json:"ImportSequence,omitempty" xml:"ImportSequence,omitempty"`
	Language       *string  `json:"Language,omitempty" xml:"Language,omitempty"`
	Status         *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemToken    *string  `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s UpdateStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateStatusRequest) GetAppType() *string {
	return s.AppType
}

func (s *UpdateStatusRequest) GetErrorLines() []*int32 {
	return s.ErrorLines
}

func (s *UpdateStatusRequest) GetImportSequence() *string {
	return s.ImportSequence
}

func (s *UpdateStatusRequest) GetLanguage() *string {
	return s.Language
}

func (s *UpdateStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateStatusRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *UpdateStatusRequest) SetAppType(v string) *UpdateStatusRequest {
	s.AppType = &v
	return s
}

func (s *UpdateStatusRequest) SetErrorLines(v []*int32) *UpdateStatusRequest {
	s.ErrorLines = v
	return s
}

func (s *UpdateStatusRequest) SetImportSequence(v string) *UpdateStatusRequest {
	s.ImportSequence = &v
	return s
}

func (s *UpdateStatusRequest) SetLanguage(v string) *UpdateStatusRequest {
	s.Language = &v
	return s
}

func (s *UpdateStatusRequest) SetStatus(v string) *UpdateStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateStatusRequest) SetSystemToken(v string) *UpdateStatusRequest {
	s.SystemToken = &v
	return s
}

func (s *UpdateStatusRequest) Validate() error {
	return dara.Validate(s)
}
