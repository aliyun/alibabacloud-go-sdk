// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePromptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeletePromptResponseBody
	GetData() *bool
	SetRequestId(v string) *DeletePromptResponseBody
	GetRequestId() *string
}

type DeletePromptResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePromptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePromptResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePromptResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeletePromptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePromptResponseBody) SetData(v bool) *DeletePromptResponseBody {
	s.Data = &v
	return s
}

func (s *DeletePromptResponseBody) SetRequestId(v string) *DeletePromptResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePromptResponseBody) Validate() error {
	return dara.Validate(s)
}
