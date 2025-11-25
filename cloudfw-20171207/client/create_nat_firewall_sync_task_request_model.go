// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatFirewallSyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateNatFirewallSyncTaskRequest
	GetLang() *string
}

type CreateNatFirewallSyncTaskRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s CreateNatFirewallSyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNatFirewallSyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateNatFirewallSyncTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateNatFirewallSyncTaskRequest) SetLang(v string) *CreateNatFirewallSyncTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateNatFirewallSyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
