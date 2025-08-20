// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelStackOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelStackOperationResponseBody
	GetRequestId() *string
}

type CancelStackOperationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelStackOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelStackOperationResponseBody) GoString() string {
	return s.String()
}

func (s *CancelStackOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelStackOperationResponseBody) SetRequestId(v string) *CancelStackOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelStackOperationResponseBody) Validate() error {
	return dara.Validate(s)
}
