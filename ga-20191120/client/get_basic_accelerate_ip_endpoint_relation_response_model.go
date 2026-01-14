// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicAccelerateIpEndpointRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBasicAccelerateIpEndpointRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBasicAccelerateIpEndpointRelationResponse
	GetStatusCode() *int32
	SetBody(v *GetBasicAccelerateIpEndpointRelationResponseBody) *GetBasicAccelerateIpEndpointRelationResponse
	GetBody() *GetBasicAccelerateIpEndpointRelationResponseBody
}

type GetBasicAccelerateIpEndpointRelationResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBasicAccelerateIpEndpointRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBasicAccelerateIpEndpointRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAccelerateIpEndpointRelationResponse) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpEndpointRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBasicAccelerateIpEndpointRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBasicAccelerateIpEndpointRelationResponse) GetBody() *GetBasicAccelerateIpEndpointRelationResponseBody {
	return s.Body
}

func (s *GetBasicAccelerateIpEndpointRelationResponse) SetHeaders(v map[string]*string) *GetBasicAccelerateIpEndpointRelationResponse {
	s.Headers = v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponse) SetStatusCode(v int32) *GetBasicAccelerateIpEndpointRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponse) SetBody(v *GetBasicAccelerateIpEndpointRelationResponseBody) *GetBasicAccelerateIpEndpointRelationResponse {
	s.Body = v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
