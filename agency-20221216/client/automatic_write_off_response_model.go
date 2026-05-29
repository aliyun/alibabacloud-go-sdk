// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutomaticWriteOffResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AutomaticWriteOffResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AutomaticWriteOffResponse
	GetStatusCode() *int32
	SetBody(v *AutomaticWriteOffResponseBody) *AutomaticWriteOffResponse
	GetBody() *AutomaticWriteOffResponseBody
}

type AutomaticWriteOffResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AutomaticWriteOffResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AutomaticWriteOffResponse) String() string {
	return dara.Prettify(s)
}

func (s AutomaticWriteOffResponse) GoString() string {
	return s.String()
}

func (s *AutomaticWriteOffResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AutomaticWriteOffResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AutomaticWriteOffResponse) GetBody() *AutomaticWriteOffResponseBody {
	return s.Body
}

func (s *AutomaticWriteOffResponse) SetHeaders(v map[string]*string) *AutomaticWriteOffResponse {
	s.Headers = v
	return s
}

func (s *AutomaticWriteOffResponse) SetStatusCode(v int32) *AutomaticWriteOffResponse {
	s.StatusCode = &v
	return s
}

func (s *AutomaticWriteOffResponse) SetBody(v *AutomaticWriteOffResponseBody) *AutomaticWriteOffResponse {
	s.Body = v
	return s
}

func (s *AutomaticWriteOffResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
