// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishHotlineServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *FinishHotlineServiceRequest
	GetAccountName() *string
	SetClientToken(v string) *FinishHotlineServiceRequest
	GetClientToken() *string
	SetInstanceId(v string) *FinishHotlineServiceRequest
	GetInstanceId() *string
}

type FinishHotlineServiceRequest struct {
	// The agent account name, which is the mobile phone number or email address specified during account registration. This value is unique within the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The unique client request ID. Used for idempotence verification. You can use a UUID to generate this ID.
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The AICCS instance ID.
	//
	// You can obtain this ID from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s FinishHotlineServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s FinishHotlineServiceRequest) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *FinishHotlineServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *FinishHotlineServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FinishHotlineServiceRequest) SetAccountName(v string) *FinishHotlineServiceRequest {
	s.AccountName = &v
	return s
}

func (s *FinishHotlineServiceRequest) SetClientToken(v string) *FinishHotlineServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *FinishHotlineServiceRequest) SetInstanceId(v string) *FinishHotlineServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *FinishHotlineServiceRequest) Validate() error {
	return dara.Validate(s)
}
