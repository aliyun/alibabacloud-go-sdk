// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterEndpointResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterEndpointResponseBody) *ModifyDBClusterEndpointResponse
	GetBody() *ModifyDBClusterEndpointResponseBody
}

type ModifyDBClusterEndpointResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterEndpointResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterEndpointResponse) GetBody() *ModifyDBClusterEndpointResponseBody {
	return s.Body
}

func (s *ModifyDBClusterEndpointResponse) SetHeaders(v map[string]*string) *ModifyDBClusterEndpointResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterEndpointResponse) SetStatusCode(v int32) *ModifyDBClusterEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterEndpointResponse) SetBody(v *ModifyDBClusterEndpointResponseBody) *ModifyDBClusterEndpointResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
