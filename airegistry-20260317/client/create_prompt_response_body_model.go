// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePromptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreatePromptResponseBody
	GetData() *string
	SetRequestId(v string) *CreatePromptResponseBody
	GetRequestId() *string
}

type CreatePromptResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePromptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePromptResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePromptResponseBody) GetData() *string {
	return s.Data
}

func (s *CreatePromptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePromptResponseBody) SetData(v string) *CreatePromptResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePromptResponseBody) SetRequestId(v string) *CreatePromptResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePromptResponseBody) Validate() error {
	return dara.Validate(s)
}
