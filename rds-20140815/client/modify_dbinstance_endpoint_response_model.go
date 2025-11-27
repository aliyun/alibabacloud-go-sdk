// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceEndpointResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceEndpointResponseBody) *ModifyDBInstanceEndpointResponse
	GetBody() *ModifyDBInstanceEndpointResponseBody
}

type ModifyDBInstanceEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceEndpointResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceEndpointResponse) GetBody() *ModifyDBInstanceEndpointResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceEndpointResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceEndpointResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceEndpointResponse) SetStatusCode(v int32) *ModifyDBInstanceEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceEndpointResponse) SetBody(v *ModifyDBInstanceEndpointResponseBody) *ModifyDBInstanceEndpointResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
