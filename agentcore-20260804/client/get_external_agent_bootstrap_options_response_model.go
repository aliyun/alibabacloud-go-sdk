// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExternalAgentBootstrapOptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExternalAgentBootstrapOptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExternalAgentBootstrapOptionsResponse
	GetStatusCode() *int32
	SetBody(v *GetExternalAgentBootstrapOptionsResponseBody) *GetExternalAgentBootstrapOptionsResponse
	GetBody() *GetExternalAgentBootstrapOptionsResponseBody
}

type GetExternalAgentBootstrapOptionsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExternalAgentBootstrapOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExternalAgentBootstrapOptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentBootstrapOptionsResponse) GoString() string {
	return s.String()
}

func (s *GetExternalAgentBootstrapOptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExternalAgentBootstrapOptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExternalAgentBootstrapOptionsResponse) GetBody() *GetExternalAgentBootstrapOptionsResponseBody {
	return s.Body
}

func (s *GetExternalAgentBootstrapOptionsResponse) SetHeaders(v map[string]*string) *GetExternalAgentBootstrapOptionsResponse {
	s.Headers = v
	return s
}

func (s *GetExternalAgentBootstrapOptionsResponse) SetStatusCode(v int32) *GetExternalAgentBootstrapOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExternalAgentBootstrapOptionsResponse) SetBody(v *GetExternalAgentBootstrapOptionsResponseBody) *GetExternalAgentBootstrapOptionsResponse {
	s.Body = v
	return s
}

func (s *GetExternalAgentBootstrapOptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
