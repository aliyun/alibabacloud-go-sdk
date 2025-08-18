// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceIpWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *ModifyInstanceIpWhitelistResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *ModifyInstanceIpWhitelistResponseBody
	GetRequestId() *string
}

type ModifyInstanceIpWhitelistResponseBody struct {
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 87249A6F-xxx-804C-E1E0AD1FAD90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceIpWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceIpWhitelistResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceIpWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceIpWhitelistResponseBody) SetInstanceName(v string) *ModifyInstanceIpWhitelistResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceIpWhitelistResponseBody) SetRequestId(v string) *ModifyInstanceIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceIpWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
