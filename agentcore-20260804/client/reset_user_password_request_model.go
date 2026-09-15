// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ResetUserPasswordRequestBody) *ResetUserPasswordRequest
	GetBody() *ResetUserPasswordRequestBody
}

type ResetUserPasswordRequest struct {
	// The request body for resetting the user password.
	Body *ResetUserPasswordRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s ResetUserPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordRequest) GetBody() *ResetUserPasswordRequestBody {
	return s.Body
}

func (s *ResetUserPasswordRequest) SetBody(v *ResetUserPasswordRequestBody) *ResetUserPasswordRequest {
	s.Body = v
	return s
}

func (s *ResetUserPasswordRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetUserPasswordRequestBody struct {
	// The user ID. At least one of agentCoreUserId and username must be specified. If both are specified, agentCoreUserId takes precedence.
	//
	// example:
	//
	// usr-123456
	AgentCoreUserId *string `json:"agentCoreUserId,omitempty" xml:"agentCoreUserId,omitempty"`
	// The new password after the reset. The password must be 8 to 32 characters in length and must contain uppercase letters, lowercase letters, digits, and special characters. The password cannot contain the username. If this parameter is not specified, the server generates a random password.
	//
	// example:
	//
	// Example@2026
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The username. At least one of username and agentCoreUserId must be specified.
	//
	// example:
	//
	// user-01
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ResetUserPasswordRequestBody) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordRequestBody) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordRequestBody) GetAgentCoreUserId() *string {
	return s.AgentCoreUserId
}

func (s *ResetUserPasswordRequestBody) GetPassword() *string {
	return s.Password
}

func (s *ResetUserPasswordRequestBody) GetUsername() *string {
	return s.Username
}

func (s *ResetUserPasswordRequestBody) SetAgentCoreUserId(v string) *ResetUserPasswordRequestBody {
	s.AgentCoreUserId = &v
	return s
}

func (s *ResetUserPasswordRequestBody) SetPassword(v string) *ResetUserPasswordRequestBody {
	s.Password = &v
	return s
}

func (s *ResetUserPasswordRequestBody) SetUsername(v string) *ResetUserPasswordRequestBody {
	s.Username = &v
	return s
}

func (s *ResetUserPasswordRequestBody) Validate() error {
	return dara.Validate(s)
}
