// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumePdnsServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumePdnsServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumePdnsServiceResponse
	GetStatusCode() *int32
	SetBody(v *ResumePdnsServiceResponseBody) *ResumePdnsServiceResponse
	GetBody() *ResumePdnsServiceResponseBody
}

type ResumePdnsServiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumePdnsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumePdnsServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumePdnsServiceResponse) GoString() string {
	return s.String()
}

func (s *ResumePdnsServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumePdnsServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumePdnsServiceResponse) GetBody() *ResumePdnsServiceResponseBody {
	return s.Body
}

func (s *ResumePdnsServiceResponse) SetHeaders(v map[string]*string) *ResumePdnsServiceResponse {
	s.Headers = v
	return s
}

func (s *ResumePdnsServiceResponse) SetStatusCode(v int32) *ResumePdnsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumePdnsServiceResponse) SetBody(v *ResumePdnsServiceResponseBody) *ResumePdnsServiceResponse {
	s.Body = v
	return s
}

func (s *ResumePdnsServiceResponse) Validate() error {
	return dara.Validate(s)
}
