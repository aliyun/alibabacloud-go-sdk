// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAIServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAIServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAIServiceResponseBody) *DeleteAIServiceResponse
	GetBody() *DeleteAIServiceResponseBody
}

type DeleteAIServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAIServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAIServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteAIServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAIServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAIServiceResponse) GetBody() *DeleteAIServiceResponseBody {
	return s.Body
}

func (s *DeleteAIServiceResponse) SetHeaders(v map[string]*string) *DeleteAIServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteAIServiceResponse) SetStatusCode(v int32) *DeleteAIServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAIServiceResponse) SetBody(v *DeleteAIServiceResponseBody) *DeleteAIServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteAIServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
