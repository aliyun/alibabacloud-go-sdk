// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDDoSInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDDoSInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDDoSInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListDDoSInstancesResponseBody) *ListDDoSInstancesResponse
	GetBody() *ListDDoSInstancesResponseBody
}

type ListDDoSInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDDoSInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDDoSInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDDoSInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListDDoSInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDDoSInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDDoSInstancesResponse) GetBody() *ListDDoSInstancesResponseBody {
	return s.Body
}

func (s *ListDDoSInstancesResponse) SetHeaders(v map[string]*string) *ListDDoSInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListDDoSInstancesResponse) SetStatusCode(v int32) *ListDDoSInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDDoSInstancesResponse) SetBody(v *ListDDoSInstancesResponseBody) *ListDDoSInstancesResponse {
	s.Body = v
	return s
}

func (s *ListDDoSInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
