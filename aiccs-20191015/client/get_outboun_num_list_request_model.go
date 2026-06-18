// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOutbounNumListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetOutbounNumListRequest
	GetAccountName() *string
	SetClientToken(v string) *GetOutbounNumListRequest
	GetClientToken() *string
	SetInstanceId(v string) *GetOutbounNumListRequest
	GetInstanceId() *string
}

type GetOutbounNumListRequest struct {
	// Agent account name (agent logon name).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123@****.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Unique ID of the customer request. Used for idempotency validation. You can generate it by using UUID.
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Instance ID.
	//
	// You can log on to the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview) and view the instance ID in **Instance Management**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetOutbounNumListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOutbounNumListRequest) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetOutbounNumListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetOutbounNumListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetOutbounNumListRequest) SetAccountName(v string) *GetOutbounNumListRequest {
	s.AccountName = &v
	return s
}

func (s *GetOutbounNumListRequest) SetClientToken(v string) *GetOutbounNumListRequest {
	s.ClientToken = &v
	return s
}

func (s *GetOutbounNumListRequest) SetInstanceId(v string) *GetOutbounNumListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOutbounNumListRequest) Validate() error {
	return dara.Validate(s)
}
