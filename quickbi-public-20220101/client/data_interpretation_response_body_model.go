// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataInterpretationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DataInterpretationResponseBody
	GetRequestId() *string
	SetResult(v string) *DataInterpretationResponseBody
	GetResult() *string
	SetSuccess(v bool) *DataInterpretationResponseBody
	GetSuccess() *bool
}

type DataInterpretationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DataInterpretationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DataInterpretationResponseBody) GoString() string {
	return s.String()
}

func (s *DataInterpretationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DataInterpretationResponseBody) GetResult() *string {
	return s.Result
}

func (s *DataInterpretationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DataInterpretationResponseBody) SetRequestId(v string) *DataInterpretationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DataInterpretationResponseBody) SetResult(v string) *DataInterpretationResponseBody {
	s.Result = &v
	return s
}

func (s *DataInterpretationResponseBody) SetSuccess(v bool) *DataInterpretationResponseBody {
	s.Success = &v
	return s
}

func (s *DataInterpretationResponseBody) Validate() error {
	return dara.Validate(s)
}
