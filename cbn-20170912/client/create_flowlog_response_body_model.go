// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowlogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowLogId(v string) *CreateFlowlogResponseBody
	GetFlowLogId() *string
	SetRequestId(v string) *CreateFlowlogResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateFlowlogResponseBody
	GetSuccess() *string
}

type CreateFlowlogResponseBody struct {
	// The ID of the flow log.
	//
	// example:
	//
	// flowlog-m5evbtbpt****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B457
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
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

func (s CreateFlowlogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowlogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowlogResponseBody) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *CreateFlowlogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFlowlogResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateFlowlogResponseBody) SetFlowLogId(v string) *CreateFlowlogResponseBody {
	s.FlowLogId = &v
	return s
}

func (s *CreateFlowlogResponseBody) SetRequestId(v string) *CreateFlowlogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowlogResponseBody) SetSuccess(v string) *CreateFlowlogResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFlowlogResponseBody) Validate() error {
	return dara.Validate(s)
}
