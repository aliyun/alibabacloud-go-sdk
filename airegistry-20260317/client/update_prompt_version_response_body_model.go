// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePromptVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdatePromptVersionResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdatePromptVersionResponseBody
	GetRequestId() *string
}

type UpdatePromptVersionResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePromptVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePromptVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePromptVersionResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdatePromptVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePromptVersionResponseBody) SetData(v bool) *UpdatePromptVersionResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePromptVersionResponseBody) SetRequestId(v string) *UpdatePromptVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePromptVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
