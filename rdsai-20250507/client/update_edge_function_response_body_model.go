// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEdgeFunctionName(v string) *UpdateEdgeFunctionResponseBody
	GetEdgeFunctionName() *string
	SetInstanceName(v string) *UpdateEdgeFunctionResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *UpdateEdgeFunctionResponseBody
	GetRequestId() *string
}

type UpdateEdgeFunctionResponseBody struct {
	// fc-xxxx。
	//
	// example:
	//
	// ef-****
	EdgeFunctionName *string `json:"EdgeFunctionName,omitempty" xml:"EdgeFunctionName,omitempty"`
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Id of the request。
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEdgeFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEdgeFunctionResponseBody) GetEdgeFunctionName() *string {
	return s.EdgeFunctionName
}

func (s *UpdateEdgeFunctionResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateEdgeFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEdgeFunctionResponseBody) SetEdgeFunctionName(v string) *UpdateEdgeFunctionResponseBody {
	s.EdgeFunctionName = &v
	return s
}

func (s *UpdateEdgeFunctionResponseBody) SetInstanceName(v string) *UpdateEdgeFunctionResponseBody {
	s.InstanceName = &v
	return s
}

func (s *UpdateEdgeFunctionResponseBody) SetRequestId(v string) *UpdateEdgeFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEdgeFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}
