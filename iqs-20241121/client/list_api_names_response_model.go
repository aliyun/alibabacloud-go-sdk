// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiNamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApiNamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApiNamesResponse
	GetStatusCode() *int32
	SetBody(v *ListApiNamesResponseBody) *ListApiNamesResponse
	GetBody() *ListApiNamesResponseBody
}

type ListApiNamesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApiNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiNamesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApiNamesResponse) GoString() string {
	return s.String()
}

func (s *ListApiNamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApiNamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApiNamesResponse) GetBody() *ListApiNamesResponseBody {
	return s.Body
}

func (s *ListApiNamesResponse) SetHeaders(v map[string]*string) *ListApiNamesResponse {
	s.Headers = v
	return s
}

func (s *ListApiNamesResponse) SetStatusCode(v int32) *ListApiNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiNamesResponse) SetBody(v *ListApiNamesResponseBody) *ListApiNamesResponse {
	s.Body = v
	return s
}

func (s *ListApiNamesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
