// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveRecordTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveRecordTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveRecordTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveRecordTemplateResponseBody) *UpdateLiveRecordTemplateResponse
	GetBody() *UpdateLiveRecordTemplateResponseBody
}

type UpdateLiveRecordTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveRecordTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveRecordTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveRecordTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveRecordTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveRecordTemplateResponse) GetBody() *UpdateLiveRecordTemplateResponseBody {
	return s.Body
}

func (s *UpdateLiveRecordTemplateResponse) SetHeaders(v map[string]*string) *UpdateLiveRecordTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveRecordTemplateResponse) SetStatusCode(v int32) *UpdateLiveRecordTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveRecordTemplateResponse) SetBody(v *UpdateLiveRecordTemplateResponseBody) *UpdateLiveRecordTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveRecordTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
