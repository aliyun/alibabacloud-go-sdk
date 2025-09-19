// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFixCheckWarningsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v int64) *FixCheckWarningsResponseBody
	GetBatchId() *int64
	SetRequestId(v string) *FixCheckWarningsResponseBody
	GetRequestId() *string
}

type FixCheckWarningsResponseBody struct {
	// The ID of the baseline risk item that has been fixed by using the Batch Repair feature.
	//
	// example:
	//
	// 52370
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 76C1D7FD-DB1E-45EA-B804-3FBD9A1DD9C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FixCheckWarningsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FixCheckWarningsResponseBody) GoString() string {
	return s.String()
}

func (s *FixCheckWarningsResponseBody) GetBatchId() *int64 {
	return s.BatchId
}

func (s *FixCheckWarningsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FixCheckWarningsResponseBody) SetBatchId(v int64) *FixCheckWarningsResponseBody {
	s.BatchId = &v
	return s
}

func (s *FixCheckWarningsResponseBody) SetRequestId(v string) *FixCheckWarningsResponseBody {
	s.RequestId = &v
	return s
}

func (s *FixCheckWarningsResponseBody) Validate() error {
	return dara.Validate(s)
}
