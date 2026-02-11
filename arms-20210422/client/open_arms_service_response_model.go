// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenArmsServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenArmsServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenArmsServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenArmsServiceResponseBody) *OpenArmsServiceResponse
	GetBody() *OpenArmsServiceResponseBody
}

type OpenArmsServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenArmsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenArmsServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenArmsServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenArmsServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenArmsServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenArmsServiceResponse) GetBody() *OpenArmsServiceResponseBody {
	return s.Body
}

func (s *OpenArmsServiceResponse) SetHeaders(v map[string]*string) *OpenArmsServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenArmsServiceResponse) SetStatusCode(v int32) *OpenArmsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenArmsServiceResponse) SetBody(v *OpenArmsServiceResponseBody) *OpenArmsServiceResponse {
	s.Body = v
	return s
}

func (s *OpenArmsServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
