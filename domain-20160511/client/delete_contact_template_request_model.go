// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactTemplateId(v int64) *DeleteContactTemplateRequest
	GetContactTemplateId() *int64
	SetLang(v string) *DeleteContactTemplateRequest
	GetLang() *string
	SetUserClientIp(v string) *DeleteContactTemplateRequest
	GetUserClientIp() *string
}

type DeleteContactTemplateRequest struct {
	// This parameter is required.
	ContactTemplateId *int64  `json:"ContactTemplateId,omitempty" xml:"ContactTemplateId,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp      *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DeleteContactTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactTemplateRequest) GetContactTemplateId() *int64 {
	return s.ContactTemplateId
}

func (s *DeleteContactTemplateRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteContactTemplateRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DeleteContactTemplateRequest) SetContactTemplateId(v int64) *DeleteContactTemplateRequest {
	s.ContactTemplateId = &v
	return s
}

func (s *DeleteContactTemplateRequest) SetLang(v string) *DeleteContactTemplateRequest {
	s.Lang = &v
	return s
}

func (s *DeleteContactTemplateRequest) SetUserClientIp(v string) *DeleteContactTemplateRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteContactTemplateRequest) Validate() error {
	return dara.Validate(s)
}
