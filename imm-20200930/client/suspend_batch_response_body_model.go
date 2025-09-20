// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SuspendBatchResponseBody
	GetRequestId() *string
}

type SuspendBatchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC564A9A-BA5C-4499-A087-D9B9E76E*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SuspendBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendBatchResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendBatchResponseBody) SetRequestId(v string) *SuspendBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
