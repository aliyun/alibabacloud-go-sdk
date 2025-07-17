// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePromptTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePromptTemplateResponseBody
	GetRequestId() *string
}

type UpdatePromptTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FE9B6CBF-47E6-5D76-9C5D-B814DD5ABxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdatePromptTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePromptTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePromptTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePromptTemplateResponseBody) SetRequestId(v string) *UpdatePromptTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePromptTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
