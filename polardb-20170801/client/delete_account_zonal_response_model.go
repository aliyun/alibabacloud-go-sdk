// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccountZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccountZonalResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAccountZonalResponseBody) *DeleteAccountZonalResponse
	GetBody() *DeleteAccountZonalResponseBody
}

type DeleteAccountZonalResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccountZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccountZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountZonalResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccountZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccountZonalResponse) GetBody() *DeleteAccountZonalResponseBody {
	return s.Body
}

func (s *DeleteAccountZonalResponse) SetHeaders(v map[string]*string) *DeleteAccountZonalResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccountZonalResponse) SetStatusCode(v int32) *DeleteAccountZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccountZonalResponse) SetBody(v *DeleteAccountZonalResponseBody) *DeleteAccountZonalResponse {
	s.Body = v
	return s
}

func (s *DeleteAccountZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
