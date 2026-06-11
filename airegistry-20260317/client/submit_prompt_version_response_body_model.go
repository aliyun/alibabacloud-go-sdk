// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPromptVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SubmitPromptVersionResponseBody
	GetData() *string
	SetRequestId(v string) *SubmitPromptVersionResponseBody
	GetRequestId() *string
}

type SubmitPromptVersionResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitPromptVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitPromptVersionResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitPromptVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *SubmitPromptVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitPromptVersionResponseBody) SetData(v string) *SubmitPromptVersionResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitPromptVersionResponseBody) SetRequestId(v string) *SubmitPromptVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitPromptVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
