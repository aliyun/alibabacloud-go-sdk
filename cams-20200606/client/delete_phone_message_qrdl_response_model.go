// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePhoneMessageQrdlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePhoneMessageQrdlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePhoneMessageQrdlResponse
	GetStatusCode() *int32
	SetBody(v *DeletePhoneMessageQrdlResponseBody) *DeletePhoneMessageQrdlResponse
	GetBody() *DeletePhoneMessageQrdlResponseBody
}

type DeletePhoneMessageQrdlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePhoneMessageQrdlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePhoneMessageQrdlResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePhoneMessageQrdlResponse) GoString() string {
	return s.String()
}

func (s *DeletePhoneMessageQrdlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePhoneMessageQrdlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePhoneMessageQrdlResponse) GetBody() *DeletePhoneMessageQrdlResponseBody {
	return s.Body
}

func (s *DeletePhoneMessageQrdlResponse) SetHeaders(v map[string]*string) *DeletePhoneMessageQrdlResponse {
	s.Headers = v
	return s
}

func (s *DeletePhoneMessageQrdlResponse) SetStatusCode(v int32) *DeletePhoneMessageQrdlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePhoneMessageQrdlResponse) SetBody(v *DeletePhoneMessageQrdlResponseBody) *DeletePhoneMessageQrdlResponse {
	s.Body = v
	return s
}

func (s *DeletePhoneMessageQrdlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
