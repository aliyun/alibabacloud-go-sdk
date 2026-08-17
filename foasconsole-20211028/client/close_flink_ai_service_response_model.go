// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseFlinkAiServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseFlinkAiServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseFlinkAiServiceResponse
	GetStatusCode() *int32
	SetBody(v *CloseFlinkAiServiceResponseBody) *CloseFlinkAiServiceResponse
	GetBody() *CloseFlinkAiServiceResponseBody
}

type CloseFlinkAiServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseFlinkAiServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseFlinkAiServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseFlinkAiServiceResponse) GoString() string {
	return s.String()
}

func (s *CloseFlinkAiServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseFlinkAiServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseFlinkAiServiceResponse) GetBody() *CloseFlinkAiServiceResponseBody {
	return s.Body
}

func (s *CloseFlinkAiServiceResponse) SetHeaders(v map[string]*string) *CloseFlinkAiServiceResponse {
	s.Headers = v
	return s
}

func (s *CloseFlinkAiServiceResponse) SetStatusCode(v int32) *CloseFlinkAiServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseFlinkAiServiceResponse) SetBody(v *CloseFlinkAiServiceResponseBody) *CloseFlinkAiServiceResponse {
	s.Body = v
	return s
}

func (s *CloseFlinkAiServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
