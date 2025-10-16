// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *RebootDesktopsResponseBody) *RebootDesktopsResponse
	GetBody() *RebootDesktopsResponseBody
}

type RebootDesktopsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RebootDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootDesktopsResponse) GetBody() *RebootDesktopsResponseBody {
	return s.Body
}

func (s *RebootDesktopsResponse) SetHeaders(v map[string]*string) *RebootDesktopsResponse {
	s.Headers = v
	return s
}

func (s *RebootDesktopsResponse) SetStatusCode(v int32) *RebootDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootDesktopsResponse) SetBody(v *RebootDesktopsResponseBody) *RebootDesktopsResponse {
	s.Body = v
	return s
}

func (s *RebootDesktopsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
