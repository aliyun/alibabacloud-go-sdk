// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterEndpointZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterEndpointZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterEndpointZonalResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterEndpointZonalResponseBody) *ModifyDBClusterEndpointZonalResponse
	GetBody() *ModifyDBClusterEndpointZonalResponseBody
}

type ModifyDBClusterEndpointZonalResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterEndpointZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterEndpointZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterEndpointZonalResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterEndpointZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterEndpointZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterEndpointZonalResponse) GetBody() *ModifyDBClusterEndpointZonalResponseBody {
	return s.Body
}

func (s *ModifyDBClusterEndpointZonalResponse) SetHeaders(v map[string]*string) *ModifyDBClusterEndpointZonalResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterEndpointZonalResponse) SetStatusCode(v int32) *ModifyDBClusterEndpointZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalResponse) SetBody(v *ModifyDBClusterEndpointZonalResponseBody) *ModifyDBClusterEndpointZonalResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterEndpointZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
