// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareSampleConsistencyJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CompareSampleConsistencyJobResponseBody
	GetRequestId() *string
}

type CompareSampleConsistencyJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompareSampleConsistencyJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompareSampleConsistencyJobResponseBody) GoString() string {
	return s.String()
}

func (s *CompareSampleConsistencyJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompareSampleConsistencyJobResponseBody) SetRequestId(v string) *CompareSampleConsistencyJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareSampleConsistencyJobResponseBody) Validate() error {
	return dara.Validate(s)
}
