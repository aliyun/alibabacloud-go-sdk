// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReparseGroupSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReparseGroupSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReparseGroupSourceResponse
	GetStatusCode() *int32
	SetBody(v *ReparseGroupSourceResponseBody) *ReparseGroupSourceResponse
	GetBody() *ReparseGroupSourceResponseBody
}

type ReparseGroupSourceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReparseGroupSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReparseGroupSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReparseGroupSourceResponse) GoString() string {
	return s.String()
}

func (s *ReparseGroupSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReparseGroupSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReparseGroupSourceResponse) GetBody() *ReparseGroupSourceResponseBody {
	return s.Body
}

func (s *ReparseGroupSourceResponse) SetHeaders(v map[string]*string) *ReparseGroupSourceResponse {
	s.Headers = v
	return s
}

func (s *ReparseGroupSourceResponse) SetStatusCode(v int32) *ReparseGroupSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReparseGroupSourceResponse) SetBody(v *ReparseGroupSourceResponseBody) *ReparseGroupSourceResponse {
	s.Body = v
	return s
}

func (s *ReparseGroupSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
