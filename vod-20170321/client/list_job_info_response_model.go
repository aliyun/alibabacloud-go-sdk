// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListJobInfoResponseBody) *ListJobInfoResponse
	GetBody() *ListJobInfoResponseBody
}

type ListJobInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobInfoResponse) GoString() string {
	return s.String()
}

func (s *ListJobInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobInfoResponse) GetBody() *ListJobInfoResponseBody {
	return s.Body
}

func (s *ListJobInfoResponse) SetHeaders(v map[string]*string) *ListJobInfoResponse {
	s.Headers = v
	return s
}

func (s *ListJobInfoResponse) SetStatusCode(v int32) *ListJobInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobInfoResponse) SetBody(v *ListJobInfoResponseBody) *ListJobInfoResponse {
	s.Body = v
	return s
}

func (s *ListJobInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
