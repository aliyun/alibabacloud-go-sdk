// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataProcessJobTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDataProcessJobTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *CreateDataProcessJobTemplateResponseBody
	GetTemplateId() *string
}

type CreateDataProcessJobTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the created task template.
	//
	// example:
	//
	// tpl-123456
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateDataProcessJobTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataProcessJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataProcessJobTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataProcessJobTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateDataProcessJobTemplateResponseBody) SetRequestId(v string) *CreateDataProcessJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataProcessJobTemplateResponseBody) SetTemplateId(v string) *CreateDataProcessJobTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateDataProcessJobTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
