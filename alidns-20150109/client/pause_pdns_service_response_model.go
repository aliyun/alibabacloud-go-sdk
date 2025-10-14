// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPausePdnsServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PausePdnsServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PausePdnsServiceResponse
	GetStatusCode() *int32
	SetBody(v *PausePdnsServiceResponseBody) *PausePdnsServiceResponse
	GetBody() *PausePdnsServiceResponseBody
}

type PausePdnsServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PausePdnsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PausePdnsServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s PausePdnsServiceResponse) GoString() string {
	return s.String()
}

func (s *PausePdnsServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PausePdnsServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PausePdnsServiceResponse) GetBody() *PausePdnsServiceResponseBody {
	return s.Body
}

func (s *PausePdnsServiceResponse) SetHeaders(v map[string]*string) *PausePdnsServiceResponse {
	s.Headers = v
	return s
}

func (s *PausePdnsServiceResponse) SetStatusCode(v int32) *PausePdnsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *PausePdnsServiceResponse) SetBody(v *PausePdnsServiceResponseBody) *PausePdnsServiceResponse {
	s.Body = v
	return s
}

func (s *PausePdnsServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
