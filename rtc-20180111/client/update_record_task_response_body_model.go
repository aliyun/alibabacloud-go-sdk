// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecordTaskResponseBody
	GetRequestId() *string
}

type UpdateRecordTaskResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecordTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecordTaskResponseBody) SetRequestId(v string) *UpdateRecordTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecordTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
