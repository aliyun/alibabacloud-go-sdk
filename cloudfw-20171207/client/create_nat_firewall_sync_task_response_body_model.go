// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatFirewallSyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateNatFirewallSyncTaskResponseBody
	GetRequestId() *string
}

type CreateNatFirewallSyncTaskResponseBody struct {
	// example:
	//
	// 3E048D45-A563-5F81-9D97-536B4A84****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNatFirewallSyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNatFirewallSyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNatFirewallSyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNatFirewallSyncTaskResponseBody) SetRequestId(v string) *CreateNatFirewallSyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNatFirewallSyncTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
