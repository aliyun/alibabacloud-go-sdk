// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteExtenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudDeleteExtenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudDeleteExtenResponse
	GetStatusCode() *int32
	SetBody(v *CloudDeleteExtenResponseBody) *CloudDeleteExtenResponse
	GetBody() *CloudDeleteExtenResponseBody
}

type CloudDeleteExtenResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudDeleteExtenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudDeleteExtenResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteExtenResponse) GoString() string {
	return s.String()
}

func (s *CloudDeleteExtenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudDeleteExtenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudDeleteExtenResponse) GetBody() *CloudDeleteExtenResponseBody {
	return s.Body
}

func (s *CloudDeleteExtenResponse) SetHeaders(v map[string]*string) *CloudDeleteExtenResponse {
	s.Headers = v
	return s
}

func (s *CloudDeleteExtenResponse) SetStatusCode(v int32) *CloudDeleteExtenResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudDeleteExtenResponse) SetBody(v *CloudDeleteExtenResponseBody) *CloudDeleteExtenResponse {
	s.Body = v
	return s
}

func (s *CloudDeleteExtenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
