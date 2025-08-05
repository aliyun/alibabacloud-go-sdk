// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyObjectGroupOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyObjectGroupOperationResponseBody
	GetRequestId() *string
}

type ModifyObjectGroupOperationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CB32593D************775F41D6ED84
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyObjectGroupOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyObjectGroupOperationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyObjectGroupOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyObjectGroupOperationResponseBody) SetRequestId(v string) *ModifyObjectGroupOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyObjectGroupOperationResponseBody) Validate() error {
	return dara.Validate(s)
}
