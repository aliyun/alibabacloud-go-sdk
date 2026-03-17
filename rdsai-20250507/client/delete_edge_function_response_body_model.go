// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEdgeFunctionName(v string) *DeleteEdgeFunctionResponseBody
	GetEdgeFunctionName() *string
	SetInstanceName(v string) *DeleteEdgeFunctionResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *DeleteEdgeFunctionResponseBody
	GetRequestId() *string
}

type DeleteEdgeFunctionResponseBody struct {
	// fc-xxxx
	//
	// example:
	//
	// ef-****
	EdgeFunctionName *string `json:"EdgeFunctionName,omitempty" xml:"EdgeFunctionName,omitempty"`
	// example:
	//
	// ra-supabase-8moov5lxba***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEdgeFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEdgeFunctionResponseBody) GetEdgeFunctionName() *string {
	return s.EdgeFunctionName
}

func (s *DeleteEdgeFunctionResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DeleteEdgeFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEdgeFunctionResponseBody) SetEdgeFunctionName(v string) *DeleteEdgeFunctionResponseBody {
	s.EdgeFunctionName = &v
	return s
}

func (s *DeleteEdgeFunctionResponseBody) SetInstanceName(v string) *DeleteEdgeFunctionResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DeleteEdgeFunctionResponseBody) SetRequestId(v string) *DeleteEdgeFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEdgeFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}
