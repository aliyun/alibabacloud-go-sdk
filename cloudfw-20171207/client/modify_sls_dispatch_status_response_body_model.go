// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySlsDispatchStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySlsDispatchStatusResponseBody
	GetRequestId() *string
}

type ModifySlsDispatchStatusResponseBody struct {
	// example:
	//
	// CE901E31-4AE9-579D-AC37-D2F1BB43****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySlsDispatchStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySlsDispatchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySlsDispatchStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySlsDispatchStatusResponseBody) SetRequestId(v string) *ModifySlsDispatchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySlsDispatchStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
