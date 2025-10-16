// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWuyingServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWuyingServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWuyingServerResponse
	GetStatusCode() *int32
	SetBody(v *ListWuyingServerResponseBody) *ListWuyingServerResponse
	GetBody() *ListWuyingServerResponseBody
}

type ListWuyingServerResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWuyingServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWuyingServerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerResponse) GoString() string {
	return s.String()
}

func (s *ListWuyingServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWuyingServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWuyingServerResponse) GetBody() *ListWuyingServerResponseBody {
	return s.Body
}

func (s *ListWuyingServerResponse) SetHeaders(v map[string]*string) *ListWuyingServerResponse {
	s.Headers = v
	return s
}

func (s *ListWuyingServerResponse) SetStatusCode(v int32) *ListWuyingServerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWuyingServerResponse) SetBody(v *ListWuyingServerResponseBody) *ListWuyingServerResponse {
	s.Body = v
	return s
}

func (s *ListWuyingServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
