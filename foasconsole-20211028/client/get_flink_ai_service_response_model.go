// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlinkAiServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFlinkAiServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFlinkAiServiceResponse
	GetStatusCode() *int32
	SetBody(v *GetFlinkAiServiceResponseBody) *GetFlinkAiServiceResponse
	GetBody() *GetFlinkAiServiceResponseBody
}

type GetFlinkAiServiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFlinkAiServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFlinkAiServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFlinkAiServiceResponse) GoString() string {
	return s.String()
}

func (s *GetFlinkAiServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFlinkAiServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFlinkAiServiceResponse) GetBody() *GetFlinkAiServiceResponseBody {
	return s.Body
}

func (s *GetFlinkAiServiceResponse) SetHeaders(v map[string]*string) *GetFlinkAiServiceResponse {
	s.Headers = v
	return s
}

func (s *GetFlinkAiServiceResponse) SetStatusCode(v int32) *GetFlinkAiServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlinkAiServiceResponse) SetBody(v *GetFlinkAiServiceResponseBody) *GetFlinkAiServiceResponse {
	s.Body = v
	return s
}

func (s *GetFlinkAiServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
