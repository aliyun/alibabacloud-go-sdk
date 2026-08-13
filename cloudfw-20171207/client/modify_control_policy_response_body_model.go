// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *ModifyControlPolicyResponseBody
	GetDryRun() *bool
	SetRequestId(v string) *ModifyControlPolicyResponseBody
	GetRequestId() *string
}

type ModifyControlPolicyResponseBody struct {
	// Indicates whether this is a successful dry run response. A value of true indicates that only the dry run was completed and no actual modification was performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyResponseBody) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyControlPolicyResponseBody) SetDryRun(v bool) *ModifyControlPolicyResponseBody {
	s.DryRun = &v
	return s
}

func (s *ModifyControlPolicyResponseBody) SetRequestId(v string) *ModifyControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
