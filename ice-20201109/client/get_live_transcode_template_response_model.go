// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveTranscodeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLiveTranscodeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLiveTranscodeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetLiveTranscodeTemplateResponseBody) *GetLiveTranscodeTemplateResponse
	GetBody() *GetLiveTranscodeTemplateResponseBody
}

type GetLiveTranscodeTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLiveTranscodeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLiveTranscodeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLiveTranscodeTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetLiveTranscodeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLiveTranscodeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLiveTranscodeTemplateResponse) GetBody() *GetLiveTranscodeTemplateResponseBody {
	return s.Body
}

func (s *GetLiveTranscodeTemplateResponse) SetHeaders(v map[string]*string) *GetLiveTranscodeTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetLiveTranscodeTemplateResponse) SetStatusCode(v int32) *GetLiveTranscodeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveTranscodeTemplateResponse) SetBody(v *GetLiveTranscodeTemplateResponseBody) *GetLiveTranscodeTemplateResponse {
	s.Body = v
	return s
}

func (s *GetLiveTranscodeTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
