// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRegisterInfoId(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest
	GetAgentRegisterInfoId() *string
	SetClientToken(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest
	GetClientToken() *string
}

type VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2074753647748672512
	AgentRegisterInfoId *string `json:"AgentRegisterInfoId,omitempty" xml:"AgentRegisterInfoId,omitempty"`
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest) GoString() string {
	return s.String()
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest) GetAgentRegisterInfoId() *string {
	return s.AgentRegisterInfoId
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest) SetAgentRegisterInfoId(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest {
	s.AgentRegisterInfoId = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest) SetClientToken(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest) Validate() error {
	return dara.Validate(s)
}
