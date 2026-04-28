// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAgentRuntimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DisableAgentRuntimeResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *DisableAgentRuntimeResponseBody
	GetRequestId() *string
}

type DisableAgentRuntimeResponseBody struct {
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D984FD38-6C2D-55DF-B0D7-8BCAC2E1F8C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAgentRuntimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableAgentRuntimeResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAgentRuntimeResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DisableAgentRuntimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableAgentRuntimeResponseBody) SetInstanceName(v string) *DisableAgentRuntimeResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DisableAgentRuntimeResponseBody) SetRequestId(v string) *DisableAgentRuntimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAgentRuntimeResponseBody) Validate() error {
	return dara.Validate(s)
}
