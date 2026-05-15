// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartChatWorkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *StartChatWorkRequest
	GetAccountName() *string
	SetInstanceId(v string) *StartChatWorkRequest
	GetInstanceId() *string
}

type StartChatWorkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123@123.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartChatWorkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartChatWorkRequest) GoString() string {
	return s.String()
}

func (s *StartChatWorkRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *StartChatWorkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartChatWorkRequest) SetAccountName(v string) *StartChatWorkRequest {
	s.AccountName = &v
	return s
}

func (s *StartChatWorkRequest) SetInstanceId(v string) *StartChatWorkRequest {
	s.InstanceId = &v
	return s
}

func (s *StartChatWorkRequest) Validate() error {
	return dara.Validate(s)
}
