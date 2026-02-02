// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstancesSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceNames(v []*string) *ModifyInstancesSSLResponseBody
	GetInstanceNames() []*string
	SetRequestId(v string) *ModifyInstancesSSLResponseBody
	GetRequestId() *string
}

type ModifyInstancesSSLResponseBody struct {
	// The RDS Supabase instances whose SSL settings are modified.
	//
	// example:
	//
	// [
	//
	//     "ra-supabase-xxx",
	//
	//     "ra-supabase-xxx"
	//
	//   ]
	InstanceNames []*string `json:"InstanceNames,omitempty" xml:"InstanceNames,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 32DEFB4A-xxxx-ADD5-918E4FD7AB8C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstancesSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstancesSSLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstancesSSLResponseBody) GetInstanceNames() []*string {
	return s.InstanceNames
}

func (s *ModifyInstancesSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstancesSSLResponseBody) SetInstanceNames(v []*string) *ModifyInstancesSSLResponseBody {
	s.InstanceNames = v
	return s
}

func (s *ModifyInstancesSSLResponseBody) SetRequestId(v string) *ModifyInstancesSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstancesSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
