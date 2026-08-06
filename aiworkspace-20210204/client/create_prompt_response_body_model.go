// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePromptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPromptId(v string) *CreatePromptResponseBody
	GetPromptId() *string
	SetRequestId(v string) *CreatePromptResponseBody
	GetRequestId() *string
}

type CreatePromptResponseBody struct {
	// The prompt ID.
	//
	// example:
	//
	// pmt-axbxtc****xxx
	PromptId *string `json:"PromptId,omitempty" xml:"PromptId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePromptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePromptResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePromptResponseBody) GetPromptId() *string {
	return s.PromptId
}

func (s *CreatePromptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePromptResponseBody) SetPromptId(v string) *CreatePromptResponseBody {
	s.PromptId = &v
	return s
}

func (s *CreatePromptResponseBody) SetRequestId(v string) *CreatePromptResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePromptResponseBody) Validate() error {
	return dara.Validate(s)
}
