// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordWeightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecursionRecordWeightResponseBody
	GetRequestId() *string
}

type UpdateRecursionRecordWeightResponseBody struct {
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecursionRecordWeightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordWeightResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordWeightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecursionRecordWeightResponseBody) SetRequestId(v string) *UpdateRecursionRecordWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecursionRecordWeightResponseBody) Validate() error {
	return dara.Validate(s)
}
