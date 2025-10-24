// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDevelopServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DevelopServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DevelopServiceResponse
	GetStatusCode() *int32
	SetBody(v *DevelopServiceResponseBody) *DevelopServiceResponse
	GetBody() *DevelopServiceResponseBody
}

type DevelopServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DevelopServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DevelopServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DevelopServiceResponse) GoString() string {
	return s.String()
}

func (s *DevelopServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DevelopServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DevelopServiceResponse) GetBody() *DevelopServiceResponseBody {
	return s.Body
}

func (s *DevelopServiceResponse) SetHeaders(v map[string]*string) *DevelopServiceResponse {
	s.Headers = v
	return s
}

func (s *DevelopServiceResponse) SetStatusCode(v int32) *DevelopServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DevelopServiceResponse) SetBody(v *DevelopServiceResponseBody) *DevelopServiceResponse {
	s.Body = v
	return s
}

func (s *DevelopServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
