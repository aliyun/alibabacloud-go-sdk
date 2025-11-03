// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBusinessRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBusinessRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBusinessRegionsResponse
	GetStatusCode() *int32
	SetBody(v *ListBusinessRegionsResponseBody) *ListBusinessRegionsResponse
	GetBody() *ListBusinessRegionsResponseBody
}

type ListBusinessRegionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBusinessRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBusinessRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBusinessRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListBusinessRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBusinessRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBusinessRegionsResponse) GetBody() *ListBusinessRegionsResponseBody {
	return s.Body
}

func (s *ListBusinessRegionsResponse) SetHeaders(v map[string]*string) *ListBusinessRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListBusinessRegionsResponse) SetStatusCode(v int32) *ListBusinessRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBusinessRegionsResponse) SetBody(v *ListBusinessRegionsResponseBody) *ListBusinessRegionsResponse {
	s.Body = v
	return s
}

func (s *ListBusinessRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
