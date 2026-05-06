// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLlmTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLlmTemplateResponseBody
	GetRequestId() *string
}

type DeleteLlmTemplateResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLlmTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLlmTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLlmTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLlmTemplateResponseBody) SetRequestId(v string) *DeleteLlmTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLlmTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
