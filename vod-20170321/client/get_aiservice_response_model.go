// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAIServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAIServiceResponse
	GetStatusCode() *int32
	SetBody(v *GetAIServiceResponseBody) *GetAIServiceResponse
	GetBody() *GetAIServiceResponseBody
}

type GetAIServiceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAIServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAIServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAIServiceResponse) GoString() string {
	return s.String()
}

func (s *GetAIServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAIServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAIServiceResponse) GetBody() *GetAIServiceResponseBody {
	return s.Body
}

func (s *GetAIServiceResponse) SetHeaders(v map[string]*string) *GetAIServiceResponse {
	s.Headers = v
	return s
}

func (s *GetAIServiceResponse) SetStatusCode(v int32) *GetAIServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAIServiceResponse) SetBody(v *GetAIServiceResponseBody) *GetAIServiceResponse {
	s.Body = v
	return s
}

func (s *GetAIServiceResponse) Validate() error {
	return dara.Validate(s)
}
