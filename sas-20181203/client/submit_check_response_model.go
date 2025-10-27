// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitCheckResponse
	GetStatusCode() *int32
	SetBody(v *SubmitCheckResponseBody) *SubmitCheckResponse
	GetBody() *SubmitCheckResponseBody
}

type SubmitCheckResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitCheckResponse) GoString() string {
	return s.String()
}

func (s *SubmitCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitCheckResponse) GetBody() *SubmitCheckResponseBody {
	return s.Body
}

func (s *SubmitCheckResponse) SetHeaders(v map[string]*string) *SubmitCheckResponse {
	s.Headers = v
	return s
}

func (s *SubmitCheckResponse) SetStatusCode(v int32) *SubmitCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCheckResponse) SetBody(v *SubmitCheckResponseBody) *SubmitCheckResponse {
	s.Body = v
	return s
}

func (s *SubmitCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
