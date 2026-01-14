// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingPortMappingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomRoutingPortMappingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomRoutingPortMappingsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomRoutingPortMappingsResponseBody) *ListCustomRoutingPortMappingsResponse
	GetBody() *ListCustomRoutingPortMappingsResponseBody
}

type ListCustomRoutingPortMappingsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomRoutingPortMappingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomRoutingPortMappingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingPortMappingsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomRoutingPortMappingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomRoutingPortMappingsResponse) GetBody() *ListCustomRoutingPortMappingsResponseBody {
	return s.Body
}

func (s *ListCustomRoutingPortMappingsResponse) SetHeaders(v map[string]*string) *ListCustomRoutingPortMappingsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomRoutingPortMappingsResponse) SetStatusCode(v int32) *ListCustomRoutingPortMappingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponse) SetBody(v *ListCustomRoutingPortMappingsResponseBody) *ListCustomRoutingPortMappingsResponse {
	s.Body = v
	return s
}

func (s *ListCustomRoutingPortMappingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
