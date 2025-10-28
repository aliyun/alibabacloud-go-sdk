// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContainerConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContainerConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContainerConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetContainerConfigurationResponseBody) *GetContainerConfigurationResponse
	GetBody() *GetContainerConfigurationResponseBody
}

type GetContainerConfigurationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContainerConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContainerConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContainerConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetContainerConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContainerConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContainerConfigurationResponse) GetBody() *GetContainerConfigurationResponseBody {
	return s.Body
}

func (s *GetContainerConfigurationResponse) SetHeaders(v map[string]*string) *GetContainerConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetContainerConfigurationResponse) SetStatusCode(v int32) *GetContainerConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContainerConfigurationResponse) SetBody(v *GetContainerConfigurationResponseBody) *GetContainerConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetContainerConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
