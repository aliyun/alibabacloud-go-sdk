// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnQuestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteConnQuestionResponseBody
	GetRequestId() *string
}

type DeleteConnQuestionResponseBody struct {
	// example:
	//
	// FC323352-3AD7-59A1-9088-A64470BAFC9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConnQuestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConnQuestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConnQuestionResponseBody) SetRequestId(v string) *DeleteConnQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConnQuestionResponseBody) Validate() error {
	return dara.Validate(s)
}
