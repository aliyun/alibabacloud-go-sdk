// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecordTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddRecordTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *AddRecordTemplateResponseBody
	GetTemplateId() *string
}

type AddRecordTemplateResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 76dasgb****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s AddRecordTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRecordTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddRecordTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRecordTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AddRecordTemplateResponseBody) SetRequestId(v string) *AddRecordTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRecordTemplateResponseBody) SetTemplateId(v string) *AddRecordTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *AddRecordTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
