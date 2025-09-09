// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecursionRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRecursionRecordResponseBody
	GetRequestId() *string
}

type DeleteRecursionRecordResponseBody struct {
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRecursionRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecursionRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecursionRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRecursionRecordResponseBody) SetRequestId(v string) *DeleteRecursionRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecursionRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
