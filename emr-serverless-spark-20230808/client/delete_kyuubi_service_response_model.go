// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKyuubiServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKyuubiServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKyuubiServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKyuubiServiceResponseBody) *DeleteKyuubiServiceResponse
	GetBody() *DeleteKyuubiServiceResponseBody
}

type DeleteKyuubiServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKyuubiServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKyuubiServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKyuubiServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteKyuubiServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKyuubiServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKyuubiServiceResponse) GetBody() *DeleteKyuubiServiceResponseBody {
	return s.Body
}

func (s *DeleteKyuubiServiceResponse) SetHeaders(v map[string]*string) *DeleteKyuubiServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteKyuubiServiceResponse) SetStatusCode(v int32) *DeleteKyuubiServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKyuubiServiceResponse) SetBody(v *DeleteKyuubiServiceResponseBody) *DeleteKyuubiServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteKyuubiServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
