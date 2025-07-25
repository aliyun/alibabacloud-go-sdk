// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppKeyStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKeyId(v string) *UpdateAppKeyStateRequest
	GetAppKeyId() *string
	SetLang(v string) *UpdateAppKeyStateRequest
	GetLang() *string
	SetState(v string) *UpdateAppKeyStateRequest
	GetState() *string
}

type UpdateAppKeyStateRequest struct {
	AppKeyId *string `json:"AppKeyId,omitempty" xml:"AppKeyId,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	State    *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s UpdateAppKeyStateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppKeyStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppKeyStateRequest) GetAppKeyId() *string {
	return s.AppKeyId
}

func (s *UpdateAppKeyStateRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateAppKeyStateRequest) GetState() *string {
	return s.State
}

func (s *UpdateAppKeyStateRequest) SetAppKeyId(v string) *UpdateAppKeyStateRequest {
	s.AppKeyId = &v
	return s
}

func (s *UpdateAppKeyStateRequest) SetLang(v string) *UpdateAppKeyStateRequest {
	s.Lang = &v
	return s
}

func (s *UpdateAppKeyStateRequest) SetState(v string) *UpdateAppKeyStateRequest {
	s.State = &v
	return s
}

func (s *UpdateAppKeyStateRequest) Validate() error {
	return dara.Validate(s)
}
