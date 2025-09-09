// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecursionRecordRemarkResponseBody
	GetRequestId() *string
}

type UpdateRecursionRecordRemarkResponseBody struct {
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecursionRecordRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecursionRecordRemarkResponseBody) SetRequestId(v string) *UpdateRecursionRecordRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecursionRecordRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
