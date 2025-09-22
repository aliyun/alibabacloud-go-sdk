// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVLExtractionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitVLExtractionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitVLExtractionTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitVLExtractionTaskResponseBody) *SubmitVLExtractionTaskResponse
	GetBody() *SubmitVLExtractionTaskResponseBody
}

type SubmitVLExtractionTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVLExtractionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVLExtractionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitVLExtractionTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitVLExtractionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitVLExtractionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitVLExtractionTaskResponse) GetBody() *SubmitVLExtractionTaskResponseBody {
	return s.Body
}

func (s *SubmitVLExtractionTaskResponse) SetHeaders(v map[string]*string) *SubmitVLExtractionTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitVLExtractionTaskResponse) SetStatusCode(v int32) *SubmitVLExtractionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVLExtractionTaskResponse) SetBody(v *SubmitVLExtractionTaskResponseBody) *SubmitVLExtractionTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitVLExtractionTaskResponse) Validate() error {
	return dara.Validate(s)
}
