// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAIServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAIServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAIServiceResponse
	GetStatusCode() *int32
	SetBody(v *SetAIServiceResponseBody) *SetAIServiceResponse
	GetBody() *SetAIServiceResponseBody
}

type SetAIServiceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAIServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAIServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAIServiceResponse) GoString() string {
	return s.String()
}

func (s *SetAIServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAIServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAIServiceResponse) GetBody() *SetAIServiceResponseBody {
	return s.Body
}

func (s *SetAIServiceResponse) SetHeaders(v map[string]*string) *SetAIServiceResponse {
	s.Headers = v
	return s
}

func (s *SetAIServiceResponse) SetStatusCode(v int32) *SetAIServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAIServiceResponse) SetBody(v *SetAIServiceResponseBody) *SetAIServiceResponse {
	s.Body = v
	return s
}

func (s *SetAIServiceResponse) Validate() error {
	return dara.Validate(s)
}
