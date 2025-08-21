// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimQuestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSimQuestionResponseBody
	GetRequestId() *string
	SetSimQuestionId(v int64) *CreateSimQuestionResponseBody
	GetSimQuestionId() *int64
}

type CreateSimQuestionResponseBody struct {
	// example:
	//
	// 16AC1B3C-66E0-438B-BB7C-71B692407B67
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1000002788
	SimQuestionId *int64 `json:"SimQuestionId,omitempty" xml:"SimQuestionId,omitempty"`
}

func (s CreateSimQuestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSimQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSimQuestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSimQuestionResponseBody) GetSimQuestionId() *int64 {
	return s.SimQuestionId
}

func (s *CreateSimQuestionResponseBody) SetRequestId(v string) *CreateSimQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSimQuestionResponseBody) SetSimQuestionId(v int64) *CreateSimQuestionResponseBody {
	s.SimQuestionId = &v
	return s
}

func (s *CreateSimQuestionResponseBody) Validate() error {
	return dara.Validate(s)
}
