// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCleanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitCleanTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitCleanTaskResponseBody
	GetSuccess() *bool
}

type SubmitCleanTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DSSDF-SEWE-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitCleanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitCleanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCleanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitCleanTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitCleanTaskResponseBody) SetRequestId(v string) *SubmitCleanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCleanTaskResponseBody) SetSuccess(v bool) *SubmitCleanTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitCleanTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
