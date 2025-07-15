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
	SetSuccess(v string) *CreateFlowLogResponseBody
	GetSuccess() *string
}

type CreateFlowLogResponseBody struct {
	// The ID of the flow log.
	//
	// example:
	//
	// fl-m5e8vhz2t21sel1nq****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B457
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazdjdhd****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *CreateFlowLogResponseBody) GetSuccess() *string {
	return s.Success
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

func (s *CreateFlowLogResponseBody) SetSuccess(v string) *CreateFlowLogResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFlowLogResponseBody) Validate() error {
	return dara.Validate(s)
}
