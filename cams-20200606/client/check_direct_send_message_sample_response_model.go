// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDirectSendMessageSampleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDirectSendMessageSampleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDirectSendMessageSampleResponse
	GetStatusCode() *int32
	SetBody(v *CheckDirectSendMessageSampleResponseBody) *CheckDirectSendMessageSampleResponse
	GetBody() *CheckDirectSendMessageSampleResponseBody
}

type CheckDirectSendMessageSampleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDirectSendMessageSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDirectSendMessageSampleResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDirectSendMessageSampleResponse) GoString() string {
	return s.String()
}

func (s *CheckDirectSendMessageSampleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDirectSendMessageSampleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDirectSendMessageSampleResponse) GetBody() *CheckDirectSendMessageSampleResponseBody {
	return s.Body
}

func (s *CheckDirectSendMessageSampleResponse) SetHeaders(v map[string]*string) *CheckDirectSendMessageSampleResponse {
	s.Headers = v
	return s
}

func (s *CheckDirectSendMessageSampleResponse) SetStatusCode(v int32) *CheckDirectSendMessageSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDirectSendMessageSampleResponse) SetBody(v *CheckDirectSendMessageSampleResponseBody) *CheckDirectSendMessageSampleResponse {
	s.Body = v
	return s
}

func (s *CheckDirectSendMessageSampleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
