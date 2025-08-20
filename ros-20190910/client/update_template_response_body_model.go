// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *UpdateTemplateResponseBody
	GetTemplateId() *string
	SetTemplateVersion(v string) *UpdateTemplateResponseBody
	GetTemplateVersion() *string
}

type UpdateTemplateResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 8C5D90E1-66B6-496C-9371-3807F8DA80A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Template ID.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template version affected by this operation.
	//
	// example:
	//
	// v2
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s UpdateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateTemplateResponseBody) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetTemplateId(v string) *UpdateTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetTemplateVersion(v string) *UpdateTemplateResponseBody {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
