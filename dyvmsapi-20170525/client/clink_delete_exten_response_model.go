// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDeleteExtenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkDeleteExtenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkDeleteExtenResponse
	GetStatusCode() *int32
	SetBody(v *ClinkDeleteExtenResponseBody) *ClinkDeleteExtenResponse
	GetBody() *ClinkDeleteExtenResponseBody
}

type ClinkDeleteExtenResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkDeleteExtenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkDeleteExtenResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkDeleteExtenResponse) GoString() string {
	return s.String()
}

func (s *ClinkDeleteExtenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkDeleteExtenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkDeleteExtenResponse) GetBody() *ClinkDeleteExtenResponseBody {
	return s.Body
}

func (s *ClinkDeleteExtenResponse) SetHeaders(v map[string]*string) *ClinkDeleteExtenResponse {
	s.Headers = v
	return s
}

func (s *ClinkDeleteExtenResponse) SetStatusCode(v int32) *ClinkDeleteExtenResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkDeleteExtenResponse) SetBody(v *ClinkDeleteExtenResponseBody) *ClinkDeleteExtenResponse {
	s.Body = v
	return s
}

func (s *ClinkDeleteExtenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
