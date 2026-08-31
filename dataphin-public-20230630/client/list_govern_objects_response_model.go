// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGovernObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGovernObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGovernObjectsResponse
	GetStatusCode() *int32
	SetBody(v *ListGovernObjectsResponseBody) *ListGovernObjectsResponse
	GetBody() *ListGovernObjectsResponseBody
}

type ListGovernObjectsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGovernObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGovernObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGovernObjectsResponse) GoString() string {
	return s.String()
}

func (s *ListGovernObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGovernObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGovernObjectsResponse) GetBody() *ListGovernObjectsResponseBody {
	return s.Body
}

func (s *ListGovernObjectsResponse) SetHeaders(v map[string]*string) *ListGovernObjectsResponse {
	s.Headers = v
	return s
}

func (s *ListGovernObjectsResponse) SetStatusCode(v int32) *ListGovernObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGovernObjectsResponse) SetBody(v *ListGovernObjectsResponseBody) *ListGovernObjectsResponse {
	s.Body = v
	return s
}

func (s *ListGovernObjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
