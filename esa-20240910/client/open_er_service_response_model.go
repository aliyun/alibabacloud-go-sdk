// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenErServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenErServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenErServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenErServiceResponseBody) *OpenErServiceResponse
	GetBody() *OpenErServiceResponseBody
}

type OpenErServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenErServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenErServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenErServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenErServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenErServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenErServiceResponse) GetBody() *OpenErServiceResponseBody {
	return s.Body
}

func (s *OpenErServiceResponse) SetHeaders(v map[string]*string) *OpenErServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenErServiceResponse) SetStatusCode(v int32) *OpenErServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenErServiceResponse) SetBody(v *OpenErServiceResponseBody) *OpenErServiceResponse {
	s.Body = v
	return s
}

func (s *OpenErServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
