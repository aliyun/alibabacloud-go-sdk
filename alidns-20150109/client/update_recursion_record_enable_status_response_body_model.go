// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordEnableStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecursionRecordEnableStatusResponseBody
	GetRequestId() *string
}

type UpdateRecursionRecordEnableStatusResponseBody struct {
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecursionRecordEnableStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordEnableStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordEnableStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecursionRecordEnableStatusResponseBody) SetRequestId(v string) *UpdateRecursionRecordEnableStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecursionRecordEnableStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
