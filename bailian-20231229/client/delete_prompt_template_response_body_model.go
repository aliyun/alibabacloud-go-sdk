// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePromptTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePromptTemplateResponseBody
	GetRequestId() *string
}

type DeletePromptTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FE9B6CBF-47E6-5D76-9C5D-B814DD5ABxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePromptTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePromptTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePromptTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePromptTemplateResponseBody) SetRequestId(v string) *DeletePromptTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePromptTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
