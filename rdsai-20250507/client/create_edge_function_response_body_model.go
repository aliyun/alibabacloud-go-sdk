// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEdgeFunctionName(v string) *CreateEdgeFunctionResponseBody
	GetEdgeFunctionName() *string
	SetInstanceName(v string) *CreateEdgeFunctionResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *CreateEdgeFunctionResponseBody
	GetRequestId() *string
}

type CreateEdgeFunctionResponseBody struct {
	// example:
	//
	// ef-****
	EdgeFunctionName *string `json:"EdgeFunctionName,omitempty" xml:"EdgeFunctionName,omitempty"`
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEdgeFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEdgeFunctionResponseBody) GetEdgeFunctionName() *string {
	return s.EdgeFunctionName
}

func (s *CreateEdgeFunctionResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateEdgeFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEdgeFunctionResponseBody) SetEdgeFunctionName(v string) *CreateEdgeFunctionResponseBody {
	s.EdgeFunctionName = &v
	return s
}

func (s *CreateEdgeFunctionResponseBody) SetInstanceName(v string) *CreateEdgeFunctionResponseBody {
	s.InstanceName = &v
	return s
}

func (s *CreateEdgeFunctionResponseBody) SetRequestId(v string) *CreateEdgeFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEdgeFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}
