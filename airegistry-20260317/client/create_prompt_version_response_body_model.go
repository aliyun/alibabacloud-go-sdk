// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePromptVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreatePromptVersionResponseBody
	GetData() *string
	SetRequestId(v string) *CreatePromptVersionResponseBody
	GetRequestId() *string
}

type CreatePromptVersionResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePromptVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePromptVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePromptVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *CreatePromptVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePromptVersionResponseBody) SetData(v string) *CreatePromptVersionResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePromptVersionResponseBody) SetRequestId(v string) *CreatePromptVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePromptVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
