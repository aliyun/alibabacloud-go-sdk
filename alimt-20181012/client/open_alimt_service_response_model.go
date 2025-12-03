// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAlimtServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenAlimtServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenAlimtServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenAlimtServiceResponseBody) *OpenAlimtServiceResponse
	GetBody() *OpenAlimtServiceResponseBody
}

type OpenAlimtServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenAlimtServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenAlimtServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenAlimtServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenAlimtServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenAlimtServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenAlimtServiceResponse) GetBody() *OpenAlimtServiceResponseBody {
	return s.Body
}

func (s *OpenAlimtServiceResponse) SetHeaders(v map[string]*string) *OpenAlimtServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenAlimtServiceResponse) SetStatusCode(v int32) *OpenAlimtServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenAlimtServiceResponse) SetBody(v *OpenAlimtServiceResponseBody) *OpenAlimtServiceResponse {
	s.Body = v
	return s
}

func (s *OpenAlimtServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
