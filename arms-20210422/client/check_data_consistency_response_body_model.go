// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataConsistencyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsDataConsistency(v bool) *CheckDataConsistencyResponseBody
	GetIsDataConsistency() *bool
	SetRequestId(v string) *CheckDataConsistencyResponseBody
	GetRequestId() *string
}

type CheckDataConsistencyResponseBody struct {
	IsDataConsistency *bool   `json:"IsDataConsistency,omitempty" xml:"IsDataConsistency,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDataConsistencyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDataConsistencyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDataConsistencyResponseBody) GetIsDataConsistency() *bool {
	return s.IsDataConsistency
}

func (s *CheckDataConsistencyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDataConsistencyResponseBody) SetIsDataConsistency(v bool) *CheckDataConsistencyResponseBody {
	s.IsDataConsistency = &v
	return s
}

func (s *CheckDataConsistencyResponseBody) SetRequestId(v string) *CheckDataConsistencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDataConsistencyResponseBody) Validate() error {
	return dara.Validate(s)
}
