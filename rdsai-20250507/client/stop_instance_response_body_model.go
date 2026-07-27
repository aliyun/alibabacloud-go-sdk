// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *StopInstanceResponseBody
	GetBranchName() *string
	SetInstanceName(v string) *StopInstanceResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *StopInstanceResponseBody
	GetRequestId() *string
}

type StopInstanceResponseBody struct {
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

func (s StopInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) GetBranchName() *string {
	return s.BranchName
}

func (s *StopInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *StopInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopInstanceResponseBody) SetBranchName(v string) *StopInstanceResponseBody {
	s.BranchName = &v
	return s
}

func (s *StopInstanceResponseBody) SetInstanceName(v string) *StopInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
