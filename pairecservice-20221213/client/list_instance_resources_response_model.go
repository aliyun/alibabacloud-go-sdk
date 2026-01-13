// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceResourcesResponseBody) *ListInstanceResourcesResponse
	GetBody() *ListInstanceResourcesResponseBody
}

type ListInstanceResourcesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceResourcesResponse) GetBody() *ListInstanceResourcesResponseBody {
	return s.Body
}

func (s *ListInstanceResourcesResponse) SetHeaders(v map[string]*string) *ListInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceResourcesResponse) SetStatusCode(v int32) *ListInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceResourcesResponse) SetBody(v *ListInstanceResourcesResponseBody) *ListInstanceResourcesResponse {
	s.Body = v
	return s
}

func (s *ListInstanceResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
