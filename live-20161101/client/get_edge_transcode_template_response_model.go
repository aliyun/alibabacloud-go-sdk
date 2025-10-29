// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeTranscodeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeTranscodeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeTranscodeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeTranscodeTemplateResponseBody) *GetEdgeTranscodeTemplateResponse
	GetBody() *GetEdgeTranscodeTemplateResponseBody
}

type GetEdgeTranscodeTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeTranscodeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeTranscodeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeTranscodeTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeTranscodeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeTranscodeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeTranscodeTemplateResponse) GetBody() *GetEdgeTranscodeTemplateResponseBody {
	return s.Body
}

func (s *GetEdgeTranscodeTemplateResponse) SetHeaders(v map[string]*string) *GetEdgeTranscodeTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeTranscodeTemplateResponse) SetStatusCode(v int32) *GetEdgeTranscodeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeTranscodeTemplateResponse) SetBody(v *GetEdgeTranscodeTemplateResponseBody) *GetEdgeTranscodeTemplateResponse {
	s.Body = v
	return s
}

func (s *GetEdgeTranscodeTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
