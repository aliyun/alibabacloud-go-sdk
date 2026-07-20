// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitQueryResponse
	GetStatusCode() *int32
	SetBody(v *SubmitQueryResponseBody) *SubmitQueryResponse
	GetBody() *SubmitQueryResponseBody
}

type SubmitQueryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitQueryResponse) GoString() string {
	return s.String()
}

func (s *SubmitQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitQueryResponse) GetBody() *SubmitQueryResponseBody {
	return s.Body
}

func (s *SubmitQueryResponse) SetHeaders(v map[string]*string) *SubmitQueryResponse {
	s.Headers = v
	return s
}

func (s *SubmitQueryResponse) SetStatusCode(v int32) *SubmitQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitQueryResponse) SetBody(v *SubmitQueryResponseBody) *SubmitQueryResponse {
	s.Body = v
	return s
}

func (s *SubmitQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
