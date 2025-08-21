// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSimQuestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSimQuestionResponseBody
	GetRequestId() *string
}

type DeleteSimQuestionResponseBody struct {
	// example:
	//
	// 6419BA93-D111-5225-8998-13E63E6D3940
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSimQuestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSimQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSimQuestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSimQuestionResponseBody) SetRequestId(v string) *DeleteSimQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSimQuestionResponseBody) Validate() error {
	return dara.Validate(s)
}
