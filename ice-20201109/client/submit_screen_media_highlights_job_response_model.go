// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitScreenMediaHighlightsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitScreenMediaHighlightsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitScreenMediaHighlightsJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitScreenMediaHighlightsJobResponseBody) *SubmitScreenMediaHighlightsJobResponse
	GetBody() *SubmitScreenMediaHighlightsJobResponseBody
}

type SubmitScreenMediaHighlightsJobResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitScreenMediaHighlightsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitScreenMediaHighlightsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitScreenMediaHighlightsJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitScreenMediaHighlightsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitScreenMediaHighlightsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitScreenMediaHighlightsJobResponse) GetBody() *SubmitScreenMediaHighlightsJobResponseBody {
	return s.Body
}

func (s *SubmitScreenMediaHighlightsJobResponse) SetHeaders(v map[string]*string) *SubmitScreenMediaHighlightsJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitScreenMediaHighlightsJobResponse) SetStatusCode(v int32) *SubmitScreenMediaHighlightsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitScreenMediaHighlightsJobResponse) SetBody(v *SubmitScreenMediaHighlightsJobResponseBody) *SubmitScreenMediaHighlightsJobResponse {
	s.Body = v
	return s
}

func (s *SubmitScreenMediaHighlightsJobResponse) Validate() error {
	return dara.Validate(s)
}
