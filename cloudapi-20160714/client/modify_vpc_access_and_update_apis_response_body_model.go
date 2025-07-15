// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcAccessAndUpdateApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *ModifyVpcAccessAndUpdateApisResponseBody
	GetOperationId() *string
	SetRequestId(v string) *ModifyVpcAccessAndUpdateApisResponseBody
	GetRequestId() *string
}

type ModifyVpcAccessAndUpdateApisResponseBody struct {
	// The ID of the asynchronous task.
	//
	// >
	//
	// 	- If the associated API is updated, you can use the task ID in the **DescribeUpdateVpcInfoTask*	- operation to query the update result.
	//
	// example:
	//
	// 7b6d0cb72b2e4215b0129f675c889746
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C2CEC6EA-EEBA-5FD6-8BD9-2CF01980FE39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcAccessAndUpdateApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcAccessAndUpdateApisResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcAccessAndUpdateApisResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *ModifyVpcAccessAndUpdateApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpcAccessAndUpdateApisResponseBody) SetOperationId(v string) *ModifyVpcAccessAndUpdateApisResponseBody {
	s.OperationId = &v
	return s
}

func (s *ModifyVpcAccessAndUpdateApisResponseBody) SetRequestId(v string) *ModifyVpcAccessAndUpdateApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpcAccessAndUpdateApisResponseBody) Validate() error {
	return dara.Validate(s)
}
