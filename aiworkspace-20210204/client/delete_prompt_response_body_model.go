// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePromptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePromptResponseBody
	GetRequestId() *string
}

type DeletePromptResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A0F049F0-8D69-5BAC-8F10-B******A34C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePromptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePromptResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePromptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePromptResponseBody) SetRequestId(v string) *DeletePromptResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePromptResponseBody) Validate() error {
	return dara.Validate(s)
}
