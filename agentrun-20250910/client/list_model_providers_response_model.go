// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelProvidersResponse
	GetStatusCode() *int32
	SetBody(v *ListModelProvidersResponseBody) *ListModelProvidersResponse
	GetBody() *ListModelProvidersResponseBody
}

type ListModelProvidersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListModelProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelProvidersResponse) GetBody() *ListModelProvidersResponseBody {
	return s.Body
}

func (s *ListModelProvidersResponse) SetHeaders(v map[string]*string) *ListModelProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListModelProvidersResponse) SetStatusCode(v int32) *ListModelProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelProvidersResponse) SetBody(v *ListModelProvidersResponseBody) *ListModelProvidersResponse {
	s.Body = v
	return s
}

func (s *ListModelProvidersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
