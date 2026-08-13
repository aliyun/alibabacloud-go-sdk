// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReparseSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReparseSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReparseSourceResponse
	GetStatusCode() *int32
	SetBody(v *ReparseSourceResponseBody) *ReparseSourceResponse
	GetBody() *ReparseSourceResponseBody
}

type ReparseSourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReparseSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReparseSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReparseSourceResponse) GoString() string {
	return s.String()
}

func (s *ReparseSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReparseSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReparseSourceResponse) GetBody() *ReparseSourceResponseBody {
	return s.Body
}

func (s *ReparseSourceResponse) SetHeaders(v map[string]*string) *ReparseSourceResponse {
	s.Headers = v
	return s
}

func (s *ReparseSourceResponse) SetStatusCode(v int32) *ReparseSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReparseSourceResponse) SetBody(v *ReparseSourceResponseBody) *ReparseSourceResponse {
	s.Body = v
	return s
}

func (s *ReparseSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
