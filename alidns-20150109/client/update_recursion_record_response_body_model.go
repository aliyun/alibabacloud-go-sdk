// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v string) *UpdateRecursionRecordResponseBody
	GetRecordId() *string
	SetRequestId(v string) *UpdateRecursionRecordResponseBody
	GetRequestId() *string
}

type UpdateRecursionRecordResponseBody struct {
	// example:
	//
	// 12*****
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecursionRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordResponseBody) GetRecordId() *string {
	return s.RecordId
}

func (s *UpdateRecursionRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecursionRecordResponseBody) SetRecordId(v string) *UpdateRecursionRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *UpdateRecursionRecordResponseBody) SetRequestId(v string) *UpdateRecursionRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecursionRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
