// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdvancedSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdateAdvancedSettingRequest
	GetBody() *string
	SetClientToken(v string) *UpdateAdvancedSettingRequest
	GetClientToken() *string
}

type UpdateAdvancedSettingRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateAdvancedSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdvancedSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdvancedSettingRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateAdvancedSettingRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAdvancedSettingRequest) SetBody(v string) *UpdateAdvancedSettingRequest {
	s.Body = &v
	return s
}

func (s *UpdateAdvancedSettingRequest) SetClientToken(v string) *UpdateAdvancedSettingRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAdvancedSettingRequest) Validate() error {
	return dara.Validate(s)
}
