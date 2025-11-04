// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveRecordTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLiveRecordTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLiveRecordTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetLiveRecordTemplateResponseBody) *GetLiveRecordTemplateResponse
	GetBody() *GetLiveRecordTemplateResponseBody
}

type GetLiveRecordTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLiveRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLiveRecordTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLiveRecordTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetLiveRecordTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLiveRecordTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLiveRecordTemplateResponse) GetBody() *GetLiveRecordTemplateResponseBody {
	return s.Body
}

func (s *GetLiveRecordTemplateResponse) SetHeaders(v map[string]*string) *GetLiveRecordTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetLiveRecordTemplateResponse) SetStatusCode(v int32) *GetLiveRecordTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveRecordTemplateResponse) SetBody(v *GetLiveRecordTemplateResponseBody) *GetLiveRecordTemplateResponse {
	s.Body = v
	return s
}

func (s *GetLiveRecordTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
