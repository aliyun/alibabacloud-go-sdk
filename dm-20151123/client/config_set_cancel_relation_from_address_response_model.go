// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetCancelRelationFromAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigSetCancelRelationFromAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigSetCancelRelationFromAddressResponse
	GetStatusCode() *int32
	SetBody(v *ConfigSetCancelRelationFromAddressResponseBody) *ConfigSetCancelRelationFromAddressResponse
	GetBody() *ConfigSetCancelRelationFromAddressResponseBody
}

type ConfigSetCancelRelationFromAddressResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigSetCancelRelationFromAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigSetCancelRelationFromAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetCancelRelationFromAddressResponse) GoString() string {
	return s.String()
}

func (s *ConfigSetCancelRelationFromAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigSetCancelRelationFromAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigSetCancelRelationFromAddressResponse) GetBody() *ConfigSetCancelRelationFromAddressResponseBody {
	return s.Body
}

func (s *ConfigSetCancelRelationFromAddressResponse) SetHeaders(v map[string]*string) *ConfigSetCancelRelationFromAddressResponse {
	s.Headers = v
	return s
}

func (s *ConfigSetCancelRelationFromAddressResponse) SetStatusCode(v int32) *ConfigSetCancelRelationFromAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigSetCancelRelationFromAddressResponse) SetBody(v *ConfigSetCancelRelationFromAddressResponseBody) *ConfigSetCancelRelationFromAddressResponse {
	s.Body = v
	return s
}

func (s *ConfigSetCancelRelationFromAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
