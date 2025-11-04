// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHighlightExtractionJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitHighlightExtractionJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitHighlightExtractionJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitHighlightExtractionJobResponseBody) *SubmitHighlightExtractionJobResponse
	GetBody() *SubmitHighlightExtractionJobResponseBody
}

type SubmitHighlightExtractionJobResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitHighlightExtractionJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitHighlightExtractionJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitHighlightExtractionJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitHighlightExtractionJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitHighlightExtractionJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitHighlightExtractionJobResponse) GetBody() *SubmitHighlightExtractionJobResponseBody {
	return s.Body
}

func (s *SubmitHighlightExtractionJobResponse) SetHeaders(v map[string]*string) *SubmitHighlightExtractionJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitHighlightExtractionJobResponse) SetStatusCode(v int32) *SubmitHighlightExtractionJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitHighlightExtractionJobResponse) SetBody(v *SubmitHighlightExtractionJobResponseBody) *SubmitHighlightExtractionJobResponse {
	s.Body = v
	return s
}

func (s *SubmitHighlightExtractionJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
