// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSystemPropertyTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSystemPropertyTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *CreateSystemPropertyTemplateResponseBody
	GetTemplateId() *string
}

type CreateSystemPropertyTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the property template.
	//
	// example:
	//
	// ap-g6gyv4a4xlf****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateSystemPropertyTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSystemPropertyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSystemPropertyTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSystemPropertyTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateSystemPropertyTemplateResponseBody) SetRequestId(v string) *CreateSystemPropertyTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSystemPropertyTemplateResponseBody) SetTemplateId(v string) *CreateSystemPropertyTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateSystemPropertyTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
