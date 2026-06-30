// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AnalyzeLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AnalyzeLabelResponse
	GetStatusCode() *int32
	SetBody(v *AnalyzeLabelResponseBody) *AnalyzeLabelResponse
	GetBody() *AnalyzeLabelResponseBody
}

type AnalyzeLabelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnalyzeLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnalyzeLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeLabelResponse) GoString() string {
	return s.String()
}

func (s *AnalyzeLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AnalyzeLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AnalyzeLabelResponse) GetBody() *AnalyzeLabelResponseBody {
	return s.Body
}

func (s *AnalyzeLabelResponse) SetHeaders(v map[string]*string) *AnalyzeLabelResponse {
	s.Headers = v
	return s
}

func (s *AnalyzeLabelResponse) SetStatusCode(v int32) *AnalyzeLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *AnalyzeLabelResponse) SetBody(v *AnalyzeLabelResponseBody) *AnalyzeLabelResponse {
	s.Body = v
	return s
}

func (s *AnalyzeLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
