// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocExtractionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDocExtractionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDocExtractionTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDocExtractionTaskResponseBody) *SubmitDocExtractionTaskResponse
	GetBody() *SubmitDocExtractionTaskResponseBody
}

type SubmitDocExtractionTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocExtractionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDocExtractionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocExtractionTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocExtractionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDocExtractionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDocExtractionTaskResponse) GetBody() *SubmitDocExtractionTaskResponseBody {
	return s.Body
}

func (s *SubmitDocExtractionTaskResponse) SetHeaders(v map[string]*string) *SubmitDocExtractionTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocExtractionTaskResponse) SetStatusCode(v int32) *SubmitDocExtractionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocExtractionTaskResponse) SetBody(v *SubmitDocExtractionTaskResponseBody) *SubmitDocExtractionTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitDocExtractionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
