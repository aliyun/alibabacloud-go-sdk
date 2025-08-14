// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteSampleDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchDeleteSampleDataResponseBody
	GetRequestId() *string
}

type BatchDeleteSampleDataResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteSampleDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteSampleDataResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteSampleDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteSampleDataResponseBody) SetRequestId(v string) *BatchDeleteSampleDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteSampleDataResponseBody) Validate() error {
	return dara.Validate(s)
}
