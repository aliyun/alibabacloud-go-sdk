// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSignatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSignatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSignatureResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSignatureResponseBody) *DeleteSignatureResponse
	GetBody() *DeleteSignatureResponseBody
}

type DeleteSignatureResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSignatureResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSignatureResponse) GoString() string {
	return s.String()
}

func (s *DeleteSignatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSignatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSignatureResponse) GetBody() *DeleteSignatureResponseBody {
	return s.Body
}

func (s *DeleteSignatureResponse) SetHeaders(v map[string]*string) *DeleteSignatureResponse {
	s.Headers = v
	return s
}

func (s *DeleteSignatureResponse) SetStatusCode(v int32) *DeleteSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSignatureResponse) SetBody(v *DeleteSignatureResponseBody) *DeleteSignatureResponse {
	s.Body = v
	return s
}

func (s *DeleteSignatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
