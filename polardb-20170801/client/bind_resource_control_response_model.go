// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindResourceControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindResourceControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindResourceControlResponse
	GetStatusCode() *int32
	SetBody(v *BindResourceControlResponseBody) *BindResourceControlResponse
	GetBody() *BindResourceControlResponseBody
}

type BindResourceControlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindResourceControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindResourceControlResponse) String() string {
	return dara.Prettify(s)
}

func (s BindResourceControlResponse) GoString() string {
	return s.String()
}

func (s *BindResourceControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindResourceControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindResourceControlResponse) GetBody() *BindResourceControlResponseBody {
	return s.Body
}

func (s *BindResourceControlResponse) SetHeaders(v map[string]*string) *BindResourceControlResponse {
	s.Headers = v
	return s
}

func (s *BindResourceControlResponse) SetStatusCode(v int32) *BindResourceControlResponse {
	s.StatusCode = &v
	return s
}

func (s *BindResourceControlResponse) SetBody(v *BindResourceControlResponseBody) *BindResourceControlResponse {
	s.Body = v
	return s
}

func (s *BindResourceControlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
