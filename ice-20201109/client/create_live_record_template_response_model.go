// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveRecordTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLiveRecordTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLiveRecordTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateLiveRecordTemplateResponseBody) *CreateLiveRecordTemplateResponse
	GetBody() *CreateLiveRecordTemplateResponseBody
}

type CreateLiveRecordTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLiveRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLiveRecordTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveRecordTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveRecordTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLiveRecordTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLiveRecordTemplateResponse) GetBody() *CreateLiveRecordTemplateResponseBody {
	return s.Body
}

func (s *CreateLiveRecordTemplateResponse) SetHeaders(v map[string]*string) *CreateLiveRecordTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveRecordTemplateResponse) SetStatusCode(v int32) *CreateLiveRecordTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveRecordTemplateResponse) SetBody(v *CreateLiveRecordTemplateResponseBody) *CreateLiveRecordTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateLiveRecordTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
