// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationPromptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPromptId(v string) *CreateApplicationPromptResponseBody
	GetPromptId() *string
	SetRequestId(v string) *CreateApplicationPromptResponseBody
	GetRequestId() *string
}

type CreateApplicationPromptResponseBody struct {
	// example:
	//
	// papt-f9lajgw765f4fnrzn1
	PromptId *string `json:"PromptId,omitempty" xml:"PromptId,omitempty"`
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationPromptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationPromptResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationPromptResponseBody) GetPromptId() *string {
	return s.PromptId
}

func (s *CreateApplicationPromptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationPromptResponseBody) SetPromptId(v string) *CreateApplicationPromptResponseBody {
	s.PromptId = &v
	return s
}

func (s *CreateApplicationPromptResponseBody) SetRequestId(v string) *CreateApplicationPromptResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationPromptResponseBody) Validate() error {
	return dara.Validate(s)
}
