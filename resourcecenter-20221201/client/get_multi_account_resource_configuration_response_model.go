// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountResourceConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultiAccountResourceConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultiAccountResourceConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetMultiAccountResourceConfigurationResponseBody) *GetMultiAccountResourceConfigurationResponse
	GetBody() *GetMultiAccountResourceConfigurationResponseBody
}

type GetMultiAccountResourceConfigurationResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiAccountResourceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiAccountResourceConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultiAccountResourceConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultiAccountResourceConfigurationResponse) GetBody() *GetMultiAccountResourceConfigurationResponseBody {
	return s.Body
}

func (s *GetMultiAccountResourceConfigurationResponse) SetHeaders(v map[string]*string) *GetMultiAccountResourceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponse) SetStatusCode(v int32) *GetMultiAccountResourceConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponse) SetBody(v *GetMultiAccountResourceConfigurationResponseBody) *GetMultiAccountResourceConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
