// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortPlaybookExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AbortPlaybookExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *AbortPlaybookExecutionResponseBody
	GetSuccess() *string
}

type AbortPlaybookExecutionResponseBody struct {
	// example:
	//
	// 0727DAC8-****-51EC-****-14999C62B502
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AbortPlaybookExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbortPlaybookExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *AbortPlaybookExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbortPlaybookExecutionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *AbortPlaybookExecutionResponseBody) SetRequestId(v string) *AbortPlaybookExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortPlaybookExecutionResponseBody) SetSuccess(v string) *AbortPlaybookExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *AbortPlaybookExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
