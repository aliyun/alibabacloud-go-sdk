// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePromptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdatePromptResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdatePromptResponseBody
	GetRequestId() *string
}

type UpdatePromptResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePromptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePromptResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePromptResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdatePromptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePromptResponseBody) SetData(v bool) *UpdatePromptResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePromptResponseBody) SetRequestId(v string) *UpdatePromptResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePromptResponseBody) Validate() error {
	return dara.Validate(s)
}
