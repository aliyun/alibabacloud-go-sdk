// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDtsJobEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDtsJobEndpointResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDtsJobEndpointResponseBody) *ModifyDtsJobEndpointResponse
	GetBody() *ModifyDtsJobEndpointResponseBody
}

type ModifyDtsJobEndpointResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDtsJobEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDtsJobEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobEndpointResponse) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDtsJobEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDtsJobEndpointResponse) GetBody() *ModifyDtsJobEndpointResponseBody {
	return s.Body
}

func (s *ModifyDtsJobEndpointResponse) SetHeaders(v map[string]*string) *ModifyDtsJobEndpointResponse {
	s.Headers = v
	return s
}

func (s *ModifyDtsJobEndpointResponse) SetStatusCode(v int32) *ModifyDtsJobEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDtsJobEndpointResponse) SetBody(v *ModifyDtsJobEndpointResponseBody) *ModifyDtsJobEndpointResponse {
	s.Body = v
	return s
}

func (s *ModifyDtsJobEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
