// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAIServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAIServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateAIServiceResponseBody) *CreateAIServiceResponse
	GetBody() *CreateAIServiceResponseBody
}

type CreateAIServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAIServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAIServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAIServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateAIServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAIServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAIServiceResponse) GetBody() *CreateAIServiceResponseBody {
	return s.Body
}

func (s *CreateAIServiceResponse) SetHeaders(v map[string]*string) *CreateAIServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateAIServiceResponse) SetStatusCode(v int32) *CreateAIServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAIServiceResponse) SetBody(v *CreateAIServiceResponseBody) *CreateAIServiceResponse {
	s.Body = v
	return s
}

func (s *CreateAIServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
