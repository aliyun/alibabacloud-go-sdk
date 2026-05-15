// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineAgentStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetHotlineAgentStatusRequest
	GetAccountName() *string
	SetInstanceId(v string) *GetHotlineAgentStatusRequest
	GetInstanceId() *string
}

type GetHotlineAgentStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123@****.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetHotlineAgentStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentStatusRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetHotlineAgentStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHotlineAgentStatusRequest) SetAccountName(v string) *GetHotlineAgentStatusRequest {
	s.AccountName = &v
	return s
}

func (s *GetHotlineAgentStatusRequest) SetInstanceId(v string) *GetHotlineAgentStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineAgentStatusRequest) Validate() error {
	return dara.Validate(s)
}
