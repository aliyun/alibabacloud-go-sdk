// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *UpdateStatusShrinkRequest
	GetAppType() *string
	SetErrorLinesShrink(v string) *UpdateStatusShrinkRequest
	GetErrorLinesShrink() *string
	SetImportSequence(v string) *UpdateStatusShrinkRequest
	GetImportSequence() *string
	SetLanguage(v string) *UpdateStatusShrinkRequest
	GetLanguage() *string
	SetStatus(v string) *UpdateStatusShrinkRequest
	GetStatus() *string
	SetSystemToken(v string) *UpdateStatusShrinkRequest
	GetSystemToken() *string
}

type UpdateStatusShrinkRequest struct {
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType          *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	ErrorLinesShrink *string `json:"ErrorLines,omitempty" xml:"ErrorLines,omitempty"`
	// example:
	//
	// seq-123
	ImportSequence *string `json:"ImportSequence,omitempty" xml:"ImportSequence,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s UpdateStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStatusShrinkRequest) GetAppType() *string {
	return s.AppType
}

func (s *UpdateStatusShrinkRequest) GetErrorLinesShrink() *string {
	return s.ErrorLinesShrink
}

func (s *UpdateStatusShrinkRequest) GetImportSequence() *string {
	return s.ImportSequence
}

func (s *UpdateStatusShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *UpdateStatusShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateStatusShrinkRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *UpdateStatusShrinkRequest) SetAppType(v string) *UpdateStatusShrinkRequest {
	s.AppType = &v
	return s
}

func (s *UpdateStatusShrinkRequest) SetErrorLinesShrink(v string) *UpdateStatusShrinkRequest {
	s.ErrorLinesShrink = &v
	return s
}

func (s *UpdateStatusShrinkRequest) SetImportSequence(v string) *UpdateStatusShrinkRequest {
	s.ImportSequence = &v
	return s
}

func (s *UpdateStatusShrinkRequest) SetLanguage(v string) *UpdateStatusShrinkRequest {
	s.Language = &v
	return s
}

func (s *UpdateStatusShrinkRequest) SetStatus(v string) *UpdateStatusShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateStatusShrinkRequest) SetSystemToken(v string) *UpdateStatusShrinkRequest {
	s.SystemToken = &v
	return s
}

func (s *UpdateStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
