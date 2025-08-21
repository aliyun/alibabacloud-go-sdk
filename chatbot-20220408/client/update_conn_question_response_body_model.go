// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnQuestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateConnQuestionResponseBody
	GetRequestId() *string
}

type UpdateConnQuestionResponseBody struct {
	// example:
	//
	// 004EB5C0-9DEB-53BF-A57A-0407A6D6B3C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConnQuestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnQuestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConnQuestionResponseBody) SetRequestId(v string) *UpdateConnQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConnQuestionResponseBody) Validate() error {
	return dara.Validate(s)
}
