// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetHotlineNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetHotlineNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetHotlineNumberResponse
	GetStatusCode() *int32
	SetBody(v *ResetHotlineNumberResponseBody) *ResetHotlineNumberResponse
	GetBody() *ResetHotlineNumberResponseBody
}

type ResetHotlineNumberResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetHotlineNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetHotlineNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetHotlineNumberResponse) GoString() string {
	return s.String()
}

func (s *ResetHotlineNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetHotlineNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetHotlineNumberResponse) GetBody() *ResetHotlineNumberResponseBody {
	return s.Body
}

func (s *ResetHotlineNumberResponse) SetHeaders(v map[string]*string) *ResetHotlineNumberResponse {
	s.Headers = v
	return s
}

func (s *ResetHotlineNumberResponse) SetStatusCode(v int32) *ResetHotlineNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetHotlineNumberResponse) SetBody(v *ResetHotlineNumberResponseBody) *ResetHotlineNumberResponse {
	s.Body = v
	return s
}

func (s *ResetHotlineNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
