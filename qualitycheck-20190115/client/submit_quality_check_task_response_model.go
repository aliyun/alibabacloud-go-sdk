// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitQualityCheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitQualityCheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitQualityCheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitQualityCheckTaskResponseBody) *SubmitQualityCheckTaskResponse
	GetBody() *SubmitQualityCheckTaskResponseBody
}

type SubmitQualityCheckTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitQualityCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitQualityCheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitQualityCheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitQualityCheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitQualityCheckTaskResponse) GetBody() *SubmitQualityCheckTaskResponseBody {
	return s.Body
}

func (s *SubmitQualityCheckTaskResponse) SetHeaders(v map[string]*string) *SubmitQualityCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitQualityCheckTaskResponse) SetStatusCode(v int32) *SubmitQualityCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitQualityCheckTaskResponse) SetBody(v *SubmitQualityCheckTaskResponseBody) *SubmitQualityCheckTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitQualityCheckTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
