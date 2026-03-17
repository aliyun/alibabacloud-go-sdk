// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowLogId(v string) *CreateFlowLogResponseBody
	GetFlowLogId() *string
	SetRequestId(v string) *CreateFlowLogResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateFlowLogResponseBody
	GetResourceGroupId() *string
}

type CreateFlowLogResponseBody struct {
	// The ID of the flow log.
	//
	// example:
	//
	// fl-7a56mar1kfw9vj****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 650CB9E8-20F3-4538-A4FC-1DA1B36E42D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the flow log belongs.
	//
	// example:
	//
	// rg-acfm2iu4f****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateFlowLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowLogResponseBody) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *CreateFlowLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFlowLogResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateFlowLogResponseBody) SetFlowLogId(v string) *CreateFlowLogResponseBody {
	s.FlowLogId = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetRequestId(v string) *CreateFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowLogResponseBody) SetResourceGroupId(v string) *CreateFlowLogResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateFlowLogResponseBody) Validate() error {
	return dara.Validate(s)
}
