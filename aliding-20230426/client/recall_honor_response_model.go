// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecallHonorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecallHonorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecallHonorResponse
	GetStatusCode() *int32
	SetBody(v *RecallHonorResponseBody) *RecallHonorResponse
	GetBody() *RecallHonorResponseBody
}

type RecallHonorResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecallHonorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecallHonorResponse) String() string {
	return dara.Prettify(s)
}

func (s RecallHonorResponse) GoString() string {
	return s.String()
}

func (s *RecallHonorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecallHonorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecallHonorResponse) GetBody() *RecallHonorResponseBody {
	return s.Body
}

func (s *RecallHonorResponse) SetHeaders(v map[string]*string) *RecallHonorResponse {
	s.Headers = v
	return s
}

func (s *RecallHonorResponse) SetStatusCode(v int32) *RecallHonorResponse {
	s.StatusCode = &v
	return s
}

func (s *RecallHonorResponse) SetBody(v *RecallHonorResponseBody) *RecallHonorResponse {
	s.Body = v
	return s
}

func (s *RecallHonorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
