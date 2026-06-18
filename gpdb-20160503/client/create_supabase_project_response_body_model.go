// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSupabaseProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *CreateSupabaseProjectResponseBody
	GetOrderId() *string
	SetProjectId(v string) *CreateSupabaseProjectResponseBody
	GetProjectId() *string
	SetRequestId(v string) *CreateSupabaseProjectResponseBody
	GetRequestId() *string
}

type CreateSupabaseProjectResponseBody struct {
	// The ID of the associated order.
	//
	// example:
	//
	// 278880417310796
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The instance ID of the Supabase project.
	//
	// example:
	//
	// spb-xxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSupabaseProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSupabaseProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSupabaseProjectResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateSupabaseProjectResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateSupabaseProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSupabaseProjectResponseBody) SetOrderId(v string) *CreateSupabaseProjectResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateSupabaseProjectResponseBody) SetProjectId(v string) *CreateSupabaseProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateSupabaseProjectResponseBody) SetRequestId(v string) *CreateSupabaseProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSupabaseProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
