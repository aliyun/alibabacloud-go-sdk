// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactiveFlowLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeactiveFlowLogResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeactiveFlowLogResponseBody
	GetSuccess() *string
}

type DeactiveFlowLogResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F7DDDC17-FA06-4AC2-8F35-59D2470FCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s DeactiveFlowLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeactiveFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *DeactiveFlowLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeactiveFlowLogResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeactiveFlowLogResponseBody) SetRequestId(v string) *DeactiveFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeactiveFlowLogResponseBody) SetSuccess(v string) *DeactiveFlowLogResponseBody {
	s.Success = &v
	return s
}

func (s *DeactiveFlowLogResponseBody) Validate() error {
	return dara.Validate(s)
}
