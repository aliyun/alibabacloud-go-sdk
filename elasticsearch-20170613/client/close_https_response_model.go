// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseHttpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseHttpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseHttpsResponse
	GetStatusCode() *int32
	SetBody(v *CloseHttpsResponseBody) *CloseHttpsResponse
	GetBody() *CloseHttpsResponseBody
}

type CloseHttpsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseHttpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseHttpsResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseHttpsResponse) GoString() string {
	return s.String()
}

func (s *CloseHttpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseHttpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseHttpsResponse) GetBody() *CloseHttpsResponseBody {
	return s.Body
}

func (s *CloseHttpsResponse) SetHeaders(v map[string]*string) *CloseHttpsResponse {
	s.Headers = v
	return s
}

func (s *CloseHttpsResponse) SetStatusCode(v int32) *CloseHttpsResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseHttpsResponse) SetBody(v *CloseHttpsResponseBody) *CloseHttpsResponse {
	s.Body = v
	return s
}

func (s *CloseHttpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
