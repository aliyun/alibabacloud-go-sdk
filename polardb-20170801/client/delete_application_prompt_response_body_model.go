// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationPromptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPromptId(v string) *DeleteApplicationPromptResponseBody
	GetPromptId() *string
	SetRequestId(v string) *DeleteApplicationPromptResponseBody
	GetRequestId() *string
}

type DeleteApplicationPromptResponseBody struct {
	// example:
	//
	// papt-f9lajgw765f4fnrzn1
	PromptId *string `json:"PromptId,omitempty" xml:"PromptId,omitempty"`
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationPromptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationPromptResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationPromptResponseBody) GetPromptId() *string {
	return s.PromptId
}

func (s *DeleteApplicationPromptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationPromptResponseBody) SetPromptId(v string) *DeleteApplicationPromptResponseBody {
	s.PromptId = &v
	return s
}

func (s *DeleteApplicationPromptResponseBody) SetRequestId(v string) *DeleteApplicationPromptResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationPromptResponseBody) Validate() error {
	return dara.Validate(s)
}
