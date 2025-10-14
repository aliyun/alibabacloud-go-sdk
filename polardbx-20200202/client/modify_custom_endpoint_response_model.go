// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCustomEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCustomEndpointResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCustomEndpointResponseBody) *ModifyCustomEndpointResponse
	GetBody() *ModifyCustomEndpointResponseBody
}

type ModifyCustomEndpointResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCustomEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCustomEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomEndpointResponse) GoString() string {
	return s.String()
}

func (s *ModifyCustomEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCustomEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCustomEndpointResponse) GetBody() *ModifyCustomEndpointResponseBody {
	return s.Body
}

func (s *ModifyCustomEndpointResponse) SetHeaders(v map[string]*string) *ModifyCustomEndpointResponse {
	s.Headers = v
	return s
}

func (s *ModifyCustomEndpointResponse) SetStatusCode(v int32) *ModifyCustomEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCustomEndpointResponse) SetBody(v *ModifyCustomEndpointResponseBody) *ModifyCustomEndpointResponse {
	s.Body = v
	return s
}

func (s *ModifyCustomEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
