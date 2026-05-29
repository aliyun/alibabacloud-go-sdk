// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProviderEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelProviderEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelProviderEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *ListModelProviderEndpointsResponseBody) *ListModelProviderEndpointsResponse
	GetBody() *ListModelProviderEndpointsResponseBody
}

type ListModelProviderEndpointsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelProviderEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelProviderEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListModelProviderEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelProviderEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelProviderEndpointsResponse) GetBody() *ListModelProviderEndpointsResponseBody {
	return s.Body
}

func (s *ListModelProviderEndpointsResponse) SetHeaders(v map[string]*string) *ListModelProviderEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListModelProviderEndpointsResponse) SetStatusCode(v int32) *ListModelProviderEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelProviderEndpointsResponse) SetBody(v *ListModelProviderEndpointsResponseBody) *ListModelProviderEndpointsResponse {
	s.Body = v
	return s
}

func (s *ListModelProviderEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
