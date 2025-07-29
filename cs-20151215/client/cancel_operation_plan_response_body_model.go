// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOperationPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelOperationPlanResponseBody
	GetRequestId() *string
}

type CancelOperationPlanResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// db82195b-75a8-40e5-9be4-16f1829dc624
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s CancelOperationPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelOperationPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOperationPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelOperationPlanResponseBody) SetRequestId(v string) *CancelOperationPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelOperationPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
