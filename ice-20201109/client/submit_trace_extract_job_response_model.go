// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTraceExtractJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitTraceExtractJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitTraceExtractJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitTraceExtractJobResponseBody) *SubmitTraceExtractJobResponse
	GetBody() *SubmitTraceExtractJobResponseBody
}

type SubmitTraceExtractJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTraceExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitTraceExtractJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceExtractJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitTraceExtractJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitTraceExtractJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitTraceExtractJobResponse) GetBody() *SubmitTraceExtractJobResponseBody {
	return s.Body
}

func (s *SubmitTraceExtractJobResponse) SetHeaders(v map[string]*string) *SubmitTraceExtractJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitTraceExtractJobResponse) SetStatusCode(v int32) *SubmitTraceExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTraceExtractJobResponse) SetBody(v *SubmitTraceExtractJobResponseBody) *SubmitTraceExtractJobResponse {
	s.Body = v
	return s
}

func (s *SubmitTraceExtractJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
