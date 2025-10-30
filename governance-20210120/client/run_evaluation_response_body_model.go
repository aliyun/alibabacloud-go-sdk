// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunEvaluationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RunEvaluationResponseBody
	GetRequestId() *string
}

type RunEvaluationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2D3E2A3A-F2B8-578D-9659-3195F94A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunEvaluationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunEvaluationResponseBody) GoString() string {
	return s.String()
}

func (s *RunEvaluationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunEvaluationResponseBody) SetRequestId(v string) *RunEvaluationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunEvaluationResponseBody) Validate() error {
	return dara.Validate(s)
}
