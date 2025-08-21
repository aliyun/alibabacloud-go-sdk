// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSimQuestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSimQuestionResponseBody
	GetRequestId() *string
}

type UpdateSimQuestionResponseBody struct {
	// example:
	//
	// DFB71B34-4188-4EA2-9988-EF3014E75910
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSimQuestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSimQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSimQuestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSimQuestionResponseBody) SetRequestId(v string) *UpdateSimQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSimQuestionResponseBody) Validate() error {
	return dara.Validate(s)
}
