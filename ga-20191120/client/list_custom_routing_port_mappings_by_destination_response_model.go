// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingPortMappingsByDestinationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomRoutingPortMappingsByDestinationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomRoutingPortMappingsByDestinationResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomRoutingPortMappingsByDestinationResponseBody) *ListCustomRoutingPortMappingsByDestinationResponse
	GetBody() *ListCustomRoutingPortMappingsByDestinationResponseBody
}

type ListCustomRoutingPortMappingsByDestinationResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomRoutingPortMappingsByDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomRoutingPortMappingsByDestinationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingPortMappingsByDestinationResponse) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsByDestinationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomRoutingPortMappingsByDestinationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomRoutingPortMappingsByDestinationResponse) GetBody() *ListCustomRoutingPortMappingsByDestinationResponseBody {
	return s.Body
}

func (s *ListCustomRoutingPortMappingsByDestinationResponse) SetHeaders(v map[string]*string) *ListCustomRoutingPortMappingsByDestinationResponse {
	s.Headers = v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponse) SetStatusCode(v int32) *ListCustomRoutingPortMappingsByDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponse) SetBody(v *ListCustomRoutingPortMappingsByDestinationResponseBody) *ListCustomRoutingPortMappingsByDestinationResponse {
	s.Body = v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
