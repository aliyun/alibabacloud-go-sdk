// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *ResetDesktopsResponseBody) *ResetDesktopsResponse
	GetBody() *ResetDesktopsResponseBody
}

type ResetDesktopsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetDesktopsResponse) GoString() string {
	return s.String()
}

func (s *ResetDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetDesktopsResponse) GetBody() *ResetDesktopsResponseBody {
	return s.Body
}

func (s *ResetDesktopsResponse) SetHeaders(v map[string]*string) *ResetDesktopsResponse {
	s.Headers = v
	return s
}

func (s *ResetDesktopsResponse) SetStatusCode(v int32) *ResetDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetDesktopsResponse) SetBody(v *ResetDesktopsResponseBody) *ResetDesktopsResponse {
	s.Body = v
	return s
}

func (s *ResetDesktopsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
