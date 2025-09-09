// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordWeightEnableStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecursionRecordWeightEnableStatusResponseBody
	GetRequestId() *string
}

type UpdateRecursionRecordWeightEnableStatusResponseBody struct {
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecursionRecordWeightEnableStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordWeightEnableStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordWeightEnableStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecursionRecordWeightEnableStatusResponseBody) SetRequestId(v string) *UpdateRecursionRecordWeightEnableStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
