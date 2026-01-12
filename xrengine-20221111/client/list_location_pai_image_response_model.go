// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLocationPaiImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLocationPaiImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLocationPaiImageResponse
	GetStatusCode() *int32
	SetBody(v *ListLocationPaiImageResponseBody) *ListLocationPaiImageResponse
	GetBody() *ListLocationPaiImageResponseBody
}

type ListLocationPaiImageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLocationPaiImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLocationPaiImageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLocationPaiImageResponse) GoString() string {
	return s.String()
}

func (s *ListLocationPaiImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLocationPaiImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLocationPaiImageResponse) GetBody() *ListLocationPaiImageResponseBody {
	return s.Body
}

func (s *ListLocationPaiImageResponse) SetHeaders(v map[string]*string) *ListLocationPaiImageResponse {
	s.Headers = v
	return s
}

func (s *ListLocationPaiImageResponse) SetStatusCode(v int32) *ListLocationPaiImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLocationPaiImageResponse) SetBody(v *ListLocationPaiImageResponseBody) *ListLocationPaiImageResponse {
	s.Body = v
	return s
}

func (s *ListLocationPaiImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
