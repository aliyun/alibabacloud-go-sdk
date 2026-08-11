// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveRecordTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLiveRecordTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *CreateLiveRecordTemplateResponseBody
	GetTemplateId() *string
}

type CreateLiveRecordTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0622C702-41BE-467E-AF2E-883D4517962E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateLiveRecordTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveRecordTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveRecordTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLiveRecordTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateLiveRecordTemplateResponseBody) SetRequestId(v string) *CreateLiveRecordTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveRecordTemplateResponseBody) SetTemplateId(v string) *CreateLiveRecordTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateLiveRecordTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
