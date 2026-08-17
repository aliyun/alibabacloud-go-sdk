// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenFlinkAiServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenFlinkAiServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenFlinkAiServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenFlinkAiServiceResponseBody) *OpenFlinkAiServiceResponse
	GetBody() *OpenFlinkAiServiceResponseBody
}

type OpenFlinkAiServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenFlinkAiServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenFlinkAiServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenFlinkAiServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenFlinkAiServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenFlinkAiServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenFlinkAiServiceResponse) GetBody() *OpenFlinkAiServiceResponseBody {
	return s.Body
}

func (s *OpenFlinkAiServiceResponse) SetHeaders(v map[string]*string) *OpenFlinkAiServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenFlinkAiServiceResponse) SetStatusCode(v int32) *OpenFlinkAiServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenFlinkAiServiceResponse) SetBody(v *OpenFlinkAiServiceResponseBody) *OpenFlinkAiServiceResponse {
	s.Body = v
	return s
}

func (s *OpenFlinkAiServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
