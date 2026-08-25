// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPredefinedModelProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPredefinedModelProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPredefinedModelProvidersResponse
	GetStatusCode() *int32
	SetBody(v *ListPredefinedModelProvidersResponseBody) *ListPredefinedModelProvidersResponse
	GetBody() *ListPredefinedModelProvidersResponseBody
}

type ListPredefinedModelProvidersResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPredefinedModelProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPredefinedModelProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPredefinedModelProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListPredefinedModelProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPredefinedModelProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPredefinedModelProvidersResponse) GetBody() *ListPredefinedModelProvidersResponseBody {
	return s.Body
}

func (s *ListPredefinedModelProvidersResponse) SetHeaders(v map[string]*string) *ListPredefinedModelProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListPredefinedModelProvidersResponse) SetStatusCode(v int32) *ListPredefinedModelProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPredefinedModelProvidersResponse) SetBody(v *ListPredefinedModelProvidersResponseBody) *ListPredefinedModelProvidersResponse {
	s.Body = v
	return s
}

func (s *ListPredefinedModelProvidersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
