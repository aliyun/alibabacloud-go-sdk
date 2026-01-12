// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetInstancePasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *ResetInstancePasswordResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *ResetInstancePasswordResponseBody
	GetRequestId() *string
}

type ResetInstancePasswordResponseBody struct {
	// The ID of the RDS Supabase instance.
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

func (s ResetInstancePasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetInstancePasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetInstancePasswordResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ResetInstancePasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetInstancePasswordResponseBody) SetInstanceName(v string) *ResetInstancePasswordResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ResetInstancePasswordResponseBody) SetRequestId(v string) *ResetInstancePasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetInstancePasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
