// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDatasetExistedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckDatasetExistedResponseBody
	GetRequestId() *string
	SetResult(v bool) *CheckDatasetExistedResponseBody
	GetResult() *bool
	SetSuccess(v bool) *CheckDatasetExistedResponseBody
	GetSuccess() *bool
}

type CheckDatasetExistedResponseBody struct {
	// example:
	//
	// C67ABB*********682B0B214
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckDatasetExistedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDatasetExistedResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDatasetExistedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDatasetExistedResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CheckDatasetExistedResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckDatasetExistedResponseBody) SetRequestId(v string) *CheckDatasetExistedResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDatasetExistedResponseBody) SetResult(v bool) *CheckDatasetExistedResponseBody {
	s.Result = &v
	return s
}

func (s *CheckDatasetExistedResponseBody) SetSuccess(v bool) *CheckDatasetExistedResponseBody {
	s.Success = &v
	return s
}

func (s *CheckDatasetExistedResponseBody) Validate() error {
	return dara.Validate(s)
}
