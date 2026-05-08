// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalUserId(v string) *GetAccessTokenRequest
	GetExternalUserId() *string
	SetTemplateId(v string) *GetAccessTokenRequest
	GetTemplateId() *string
}

type GetAccessTokenRequest struct {
	// example:
	//
	// "user-38764"
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	// example:
	//
	// 1600112233445566
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetAccessTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAccessTokenRequest) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *GetAccessTokenRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetAccessTokenRequest) SetExternalUserId(v string) *GetAccessTokenRequest {
	s.ExternalUserId = &v
	return s
}

func (s *GetAccessTokenRequest) SetTemplateId(v string) *GetAccessTokenRequest {
	s.TemplateId = &v
	return s
}

func (s *GetAccessTokenRequest) Validate() error {
	return dara.Validate(s)
}
