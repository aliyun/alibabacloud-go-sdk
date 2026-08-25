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
	// example:
	//
	// usr-123456
	AgentCoreUserId *string `json:"agentCoreUserId,omitempty" xml:"agentCoreUserId,omitempty"`
	// example:
	//
	// Example@2026
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
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
