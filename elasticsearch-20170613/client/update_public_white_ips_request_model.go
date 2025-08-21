// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicWhiteIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdatePublicWhiteIpsRequest
	GetBody() *string
	SetClientToken(v string) *UpdatePublicWhiteIpsRequest
	GetClientToken() *string
	SetModifyMode(v string) *UpdatePublicWhiteIpsRequest
	GetModifyMode() *string
}

type UpdatePublicWhiteIpsRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The results that are returned.
	//
	// example:
	//
	// Cover
	ModifyMode *string `json:"modifyMode,omitempty" xml:"modifyMode,omitempty"`
}

func (s UpdatePublicWhiteIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicWhiteIpsRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublicWhiteIpsRequest) GetBody() *string {
	return s.Body
}

func (s *UpdatePublicWhiteIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdatePublicWhiteIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *UpdatePublicWhiteIpsRequest) SetBody(v string) *UpdatePublicWhiteIpsRequest {
	s.Body = &v
	return s
}

func (s *UpdatePublicWhiteIpsRequest) SetClientToken(v string) *UpdatePublicWhiteIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePublicWhiteIpsRequest) SetModifyMode(v string) *UpdatePublicWhiteIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *UpdatePublicWhiteIpsRequest) Validate() error {
	return dara.Validate(s)
}
