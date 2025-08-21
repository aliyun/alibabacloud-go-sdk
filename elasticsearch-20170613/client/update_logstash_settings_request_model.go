// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogstashSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdateLogstashSettingsRequest
	GetBody() *string
	SetClientToken(v string) *UpdateLogstashSettingsRequest
	GetClientToken() *string
}

type UpdateLogstashSettingsRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF\\*\\*\\*\\*
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateLogstashSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogstashSettingsRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateLogstashSettingsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateLogstashSettingsRequest) SetBody(v string) *UpdateLogstashSettingsRequest {
	s.Body = &v
	return s
}

func (s *UpdateLogstashSettingsRequest) SetClientToken(v string) *UpdateLogstashSettingsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLogstashSettingsRequest) Validate() error {
	return dara.Validate(s)
}
