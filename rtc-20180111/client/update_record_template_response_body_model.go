// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecordTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *UpdateRecordTemplateResponseBody
	GetTemplateId() *string
}

type UpdateRecordTemplateResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 76dasgb****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateRecordTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecordTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecordTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateRecordTemplateResponseBody) SetRequestId(v string) *UpdateRecordTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecordTemplateResponseBody) SetTemplateId(v string) *UpdateRecordTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *UpdateRecordTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
