// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRtcTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetRtcTokenRequest
	GetAccountName() *string
	SetInstanceId(v string) *GetRtcTokenRequest
	GetInstanceId() *string
}

type GetRtcTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agent_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRtcTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRtcTokenRequest) GoString() string {
	return s.String()
}

func (s *GetRtcTokenRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetRtcTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRtcTokenRequest) SetAccountName(v string) *GetRtcTokenRequest {
	s.AccountName = &v
	return s
}

func (s *GetRtcTokenRequest) SetInstanceId(v string) *GetRtcTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRtcTokenRequest) Validate() error {
	return dara.Validate(s)
}
