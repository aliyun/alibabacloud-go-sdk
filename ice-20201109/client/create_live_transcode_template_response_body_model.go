// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveTranscodeTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLiveTranscodeTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *CreateLiveTranscodeTemplateResponseBody
	GetTemplateId() *string
}

type CreateLiveTranscodeTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateLiveTranscodeTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveTranscodeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveTranscodeTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLiveTranscodeTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateLiveTranscodeTemplateResponseBody) SetRequestId(v string) *CreateLiveTranscodeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveTranscodeTemplateResponseBody) SetTemplateId(v string) *CreateLiveTranscodeTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateLiveTranscodeTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
