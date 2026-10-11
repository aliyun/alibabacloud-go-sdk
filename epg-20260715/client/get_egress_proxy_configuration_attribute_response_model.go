// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEgressProxyConfigurationAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEgressProxyConfigurationAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEgressProxyConfigurationAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetEgressProxyConfigurationAttributeResponseBody) *GetEgressProxyConfigurationAttributeResponse
	GetBody() *GetEgressProxyConfigurationAttributeResponseBody
}

type GetEgressProxyConfigurationAttributeResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEgressProxyConfigurationAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEgressProxyConfigurationAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyConfigurationAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetEgressProxyConfigurationAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEgressProxyConfigurationAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEgressProxyConfigurationAttributeResponse) GetBody() *GetEgressProxyConfigurationAttributeResponseBody {
	return s.Body
}

func (s *GetEgressProxyConfigurationAttributeResponse) SetHeaders(v map[string]*string) *GetEgressProxyConfigurationAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponse) SetStatusCode(v int32) *GetEgressProxyConfigurationAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponse) SetBody(v *GetEgressProxyConfigurationAttributeResponseBody) *GetEgressProxyConfigurationAttributeResponse {
	s.Body = v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
