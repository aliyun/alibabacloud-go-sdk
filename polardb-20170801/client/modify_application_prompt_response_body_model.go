// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationPromptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPromptId(v string) *ModifyApplicationPromptResponseBody
	GetPromptId() *string
	SetRequestId(v string) *ModifyApplicationPromptResponseBody
	GetRequestId() *string
}

type ModifyApplicationPromptResponseBody struct {
	// example:
	//
	// papt-f9lajgw765f4fnrzn1
	PromptId *string `json:"PromptId,omitempty" xml:"PromptId,omitempty"`
	// example:
	//
	// 7E2FE3BB-C677-5FF9-9FC5-9CF364BD6BE5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApplicationPromptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationPromptResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApplicationPromptResponseBody) GetPromptId() *string {
	return s.PromptId
}

func (s *ModifyApplicationPromptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApplicationPromptResponseBody) SetPromptId(v string) *ModifyApplicationPromptResponseBody {
	s.PromptId = &v
	return s
}

func (s *ModifyApplicationPromptResponseBody) SetRequestId(v string) *ModifyApplicationPromptResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApplicationPromptResponseBody) Validate() error {
	return dara.Validate(s)
}
