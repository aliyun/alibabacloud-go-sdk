// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseProjectResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifySupabaseProjectResourceResponseBody
	GetOrderId() *string
	SetProjectId(v string) *ModifySupabaseProjectResourceResponseBody
	GetProjectId() *string
	SetRequestId(v string) *ModifySupabaseProjectResourceResponseBody
	GetRequestId() *string
}

type ModifySupabaseProjectResourceResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// *********
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The Supabase project ID.
	//
	// example:
	//
	// sbp-tyarplz****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySupabaseProjectResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseProjectResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySupabaseProjectResourceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifySupabaseProjectResourceResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *ModifySupabaseProjectResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySupabaseProjectResourceResponseBody) SetOrderId(v string) *ModifySupabaseProjectResourceResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifySupabaseProjectResourceResponseBody) SetProjectId(v string) *ModifySupabaseProjectResourceResponseBody {
	s.ProjectId = &v
	return s
}

func (s *ModifySupabaseProjectResourceResponseBody) SetRequestId(v string) *ModifySupabaseProjectResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySupabaseProjectResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
