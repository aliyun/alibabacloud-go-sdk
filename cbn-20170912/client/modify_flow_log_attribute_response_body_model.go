// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFlowLogAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyFlowLogAttributeResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyFlowLogAttributeResponseBody
	GetSuccess() *string
}

type ModifyFlowLogAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9A411874-2EC8-4633-9D30-9012DCE89DD1
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

func (s ModifyFlowLogAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFlowLogAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFlowLogAttributeResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyFlowLogAttributeResponseBody) SetRequestId(v string) *ModifyFlowLogAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) SetSuccess(v string) *ModifyFlowLogAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyFlowLogAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
