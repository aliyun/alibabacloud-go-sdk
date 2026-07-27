// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *ModifyInstanceSSLResponseBody
	GetBranchName() *string
	SetInstanceName(v string) *ModifyInstanceSSLResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *ModifyInstanceSSLResponseBody
	GetRequestId() *string
}

type ModifyInstanceSSLResponseBody struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The instance ID of the AI application.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSSLResponseBody) GetBranchName() *string {
	return s.BranchName
}

func (s *ModifyInstanceSSLResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceSSLResponseBody) SetBranchName(v string) *ModifyInstanceSSLResponseBody {
	s.BranchName = &v
	return s
}

func (s *ModifyInstanceSSLResponseBody) SetInstanceName(v string) *ModifyInstanceSSLResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceSSLResponseBody) SetRequestId(v string) *ModifyInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
