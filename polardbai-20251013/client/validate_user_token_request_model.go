// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateUserTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ValidateUserTokenRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ValidateUserTokenRequest
	GetAuthType() *string
	SetDBClusterId(v string) *ValidateUserTokenRequest
	GetDBClusterId() *string
}

type ValidateUserTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yEtNjzuM21NVLVJTuL9Trw****
	AuthMessage *string `json:"AuthMessage,omitempty" xml:"AuthMessage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// token
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ValidateUserTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateUserTokenRequest) GoString() string {
	return s.String()
}

func (s *ValidateUserTokenRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ValidateUserTokenRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ValidateUserTokenRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ValidateUserTokenRequest) SetAuthMessage(v string) *ValidateUserTokenRequest {
	s.AuthMessage = &v
	return s
}

func (s *ValidateUserTokenRequest) SetAuthType(v string) *ValidateUserTokenRequest {
	s.AuthType = &v
	return s
}

func (s *ValidateUserTokenRequest) SetDBClusterId(v string) *ValidateUserTokenRequest {
	s.DBClusterId = &v
	return s
}

func (s *ValidateUserTokenRequest) Validate() error {
	return dara.Validate(s)
}
