// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeSaasServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeSaasServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeSaasServiceResponse
	GetStatusCode() *int32
	SetBody(v *ResumeSaasServiceResponseBody) *ResumeSaasServiceResponse
	GetBody() *ResumeSaasServiceResponseBody
}

type ResumeSaasServiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeSaasServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeSaasServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeSaasServiceResponse) GoString() string {
	return s.String()
}

func (s *ResumeSaasServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeSaasServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeSaasServiceResponse) GetBody() *ResumeSaasServiceResponseBody {
	return s.Body
}

func (s *ResumeSaasServiceResponse) SetHeaders(v map[string]*string) *ResumeSaasServiceResponse {
	s.Headers = v
	return s
}

func (s *ResumeSaasServiceResponse) SetStatusCode(v int32) *ResumeSaasServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeSaasServiceResponse) SetBody(v *ResumeSaasServiceResponseBody) *ResumeSaasServiceResponse {
	s.Body = v
	return s
}

func (s *ResumeSaasServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
