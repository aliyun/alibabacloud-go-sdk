// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarFsMappingAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPolarFsMappingAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPolarFsMappingAuthResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPolarFsMappingAuthResponseBody) *ModifyPolarFsMappingAuthResponse
	GetBody() *ModifyPolarFsMappingAuthResponseBody
}

type ModifyPolarFsMappingAuthResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolarFsMappingAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolarFsMappingAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarFsMappingAuthResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolarFsMappingAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPolarFsMappingAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPolarFsMappingAuthResponse) GetBody() *ModifyPolarFsMappingAuthResponseBody {
	return s.Body
}

func (s *ModifyPolarFsMappingAuthResponse) SetHeaders(v map[string]*string) *ModifyPolarFsMappingAuthResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolarFsMappingAuthResponse) SetStatusCode(v int32) *ModifyPolarFsMappingAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolarFsMappingAuthResponse) SetBody(v *ModifyPolarFsMappingAuthResponseBody) *ModifyPolarFsMappingAuthResponse {
	s.Body = v
	return s
}

func (s *ModifyPolarFsMappingAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
