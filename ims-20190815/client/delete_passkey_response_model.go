// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePasskeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePasskeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePasskeyResponse
	GetStatusCode() *int32
	SetBody(v *DeletePasskeyResponseBody) *DeletePasskeyResponse
	GetBody() *DeletePasskeyResponseBody
}

type DeletePasskeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePasskeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePasskeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePasskeyResponse) GoString() string {
	return s.String()
}

func (s *DeletePasskeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePasskeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePasskeyResponse) GetBody() *DeletePasskeyResponseBody {
	return s.Body
}

func (s *DeletePasskeyResponse) SetHeaders(v map[string]*string) *DeletePasskeyResponse {
	s.Headers = v
	return s
}

func (s *DeletePasskeyResponse) SetStatusCode(v int32) *DeletePasskeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePasskeyResponse) SetBody(v *DeletePasskeyResponseBody) *DeletePasskeyResponse {
	s.Body = v
	return s
}

func (s *DeletePasskeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
