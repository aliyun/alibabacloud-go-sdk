// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *CreateAppInstanceResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *CreateAppInstanceResponseBody
	GetRequestId() *string
}

type CreateAppInstanceResponseBody struct {
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

func (s CreateAppInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateAppInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppInstanceResponseBody) SetInstanceName(v string) *CreateAppInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *CreateAppInstanceResponseBody) SetRequestId(v string) *CreateAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
