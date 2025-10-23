// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetRelationFromAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigSetRelationFromAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigSetRelationFromAddressResponse
	GetStatusCode() *int32
	SetBody(v *ConfigSetRelationFromAddressResponseBody) *ConfigSetRelationFromAddressResponse
	GetBody() *ConfigSetRelationFromAddressResponseBody
}

type ConfigSetRelationFromAddressResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigSetRelationFromAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigSetRelationFromAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetRelationFromAddressResponse) GoString() string {
	return s.String()
}

func (s *ConfigSetRelationFromAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigSetRelationFromAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigSetRelationFromAddressResponse) GetBody() *ConfigSetRelationFromAddressResponseBody {
	return s.Body
}

func (s *ConfigSetRelationFromAddressResponse) SetHeaders(v map[string]*string) *ConfigSetRelationFromAddressResponse {
	s.Headers = v
	return s
}

func (s *ConfigSetRelationFromAddressResponse) SetStatusCode(v int32) *ConfigSetRelationFromAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigSetRelationFromAddressResponse) SetBody(v *ConfigSetRelationFromAddressResponseBody) *ConfigSetRelationFromAddressResponse {
	s.Body = v
	return s
}

func (s *ConfigSetRelationFromAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
