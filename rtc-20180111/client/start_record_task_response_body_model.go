// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRecordTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartRecordTaskResponseBody
	GetRequestId() *string
}

type StartRecordTaskResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartRecordTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRecordTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartRecordTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRecordTaskResponseBody) SetRequestId(v string) *StartRecordTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRecordTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
