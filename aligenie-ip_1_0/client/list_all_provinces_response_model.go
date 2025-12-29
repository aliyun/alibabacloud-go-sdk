// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllProvincesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllProvincesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllProvincesResponse
	GetStatusCode() *int32
	SetBody(v *ListAllProvincesResponseBody) *ListAllProvincesResponse
	GetBody() *ListAllProvincesResponseBody
}

type ListAllProvincesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllProvincesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllProvincesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllProvincesResponse) GoString() string {
	return s.String()
}

func (s *ListAllProvincesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllProvincesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllProvincesResponse) GetBody() *ListAllProvincesResponseBody {
	return s.Body
}

func (s *ListAllProvincesResponse) SetHeaders(v map[string]*string) *ListAllProvincesResponse {
	s.Headers = v
	return s
}

func (s *ListAllProvincesResponse) SetStatusCode(v int32) *ListAllProvincesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllProvincesResponse) SetBody(v *ListAllProvincesResponseBody) *ListAllProvincesResponse {
	s.Body = v
	return s
}

func (s *ListAllProvincesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
