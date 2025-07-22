// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppRecordTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAppRecordTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *CreateAppRecordTemplateResponseBody
	GetTemplateId() *string
}

type CreateAppRecordTemplateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// lD7muaxx
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateAppRecordTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRecordTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppRecordTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppRecordTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateAppRecordTemplateResponseBody) SetRequestId(v string) *CreateAppRecordTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppRecordTemplateResponseBody) SetTemplateId(v string) *CreateAppRecordTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateAppRecordTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
