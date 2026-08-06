// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePromptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePromptResponseBody
	GetRequestId() *string
}

type UpdatePromptResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePromptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePromptResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePromptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePromptResponseBody) SetRequestId(v string) *UpdatePromptResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePromptResponseBody) Validate() error {
	return dara.Validate(s)
}
