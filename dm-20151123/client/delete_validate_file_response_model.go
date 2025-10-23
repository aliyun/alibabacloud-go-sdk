// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteValidateFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteValidateFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteValidateFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteValidateFileResponseBody) *DeleteValidateFileResponse
	GetBody() *DeleteValidateFileResponseBody
}

type DeleteValidateFileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteValidateFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteValidateFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteValidateFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteValidateFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteValidateFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteValidateFileResponse) GetBody() *DeleteValidateFileResponseBody {
	return s.Body
}

func (s *DeleteValidateFileResponse) SetHeaders(v map[string]*string) *DeleteValidateFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteValidateFileResponse) SetStatusCode(v int32) *DeleteValidateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteValidateFileResponse) SetBody(v *DeleteValidateFileResponseBody) *DeleteValidateFileResponse {
	s.Body = v
	return s
}

func (s *DeleteValidateFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
