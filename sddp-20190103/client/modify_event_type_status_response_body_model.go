// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEventTypeStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyEventTypeStatusResponseBody
	GetRequestId() *string
}

type ModifyEventTypeStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 208B016D-4CB9-4A85-96A5-0B8ED1E*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEventTypeStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEventTypeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEventTypeStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEventTypeStatusResponseBody) SetRequestId(v string) *ModifyEventTypeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEventTypeStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
