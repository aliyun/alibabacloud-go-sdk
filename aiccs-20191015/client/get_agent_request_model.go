// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetAgentRequest
	GetAccountName() *string
	SetClientToken(v string) *GetAgentRequest
	GetClientToken() *string
	SetInstanceId(v string) *GetAgentRequest
	GetInstanceId() *string
}

type GetAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123@123.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetAgentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentRequest) SetAccountName(v string) *GetAgentRequest {
	s.AccountName = &v
	return s
}

func (s *GetAgentRequest) SetClientToken(v string) *GetAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *GetAgentRequest) SetInstanceId(v string) *GetAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentRequest) Validate() error {
	return dara.Validate(s)
}
