// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDifyInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDifyInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDifyInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListDifyInstancesResponseBody) *ListDifyInstancesResponse
	GetBody() *ListDifyInstancesResponseBody
}

type ListDifyInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDifyInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDifyInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDifyInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListDifyInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDifyInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDifyInstancesResponse) GetBody() *ListDifyInstancesResponseBody {
	return s.Body
}

func (s *ListDifyInstancesResponse) SetHeaders(v map[string]*string) *ListDifyInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListDifyInstancesResponse) SetStatusCode(v int32) *ListDifyInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDifyInstancesResponse) SetBody(v *ListDifyInstancesResponseBody) *ListDifyInstancesResponse {
	s.Body = v
	return s
}

func (s *ListDifyInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
