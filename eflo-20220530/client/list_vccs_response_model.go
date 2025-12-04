// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVccsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVccsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVccsResponse
	GetStatusCode() *int32
	SetBody(v *ListVccsResponseBody) *ListVccsResponse
	GetBody() *ListVccsResponseBody
}

type ListVccsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVccsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVccsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVccsResponse) GoString() string {
	return s.String()
}

func (s *ListVccsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVccsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVccsResponse) GetBody() *ListVccsResponseBody {
	return s.Body
}

func (s *ListVccsResponse) SetHeaders(v map[string]*string) *ListVccsResponse {
	s.Headers = v
	return s
}

func (s *ListVccsResponse) SetStatusCode(v int32) *ListVccsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVccsResponse) SetBody(v *ListVccsResponseBody) *ListVccsResponse {
	s.Body = v
	return s
}

func (s *ListVccsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
