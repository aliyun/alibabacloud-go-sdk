// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppTokenServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppTokenServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppTokenServiceResponse
	GetStatusCode() *int32
	SetBody(v *GetAppTokenServiceResponseBody) *GetAppTokenServiceResponse
	GetBody() *GetAppTokenServiceResponseBody
}

type GetAppTokenServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppTokenServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppTokenServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppTokenServiceResponse) GoString() string {
	return s.String()
}

func (s *GetAppTokenServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppTokenServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppTokenServiceResponse) GetBody() *GetAppTokenServiceResponseBody {
	return s.Body
}

func (s *GetAppTokenServiceResponse) SetHeaders(v map[string]*string) *GetAppTokenServiceResponse {
	s.Headers = v
	return s
}

func (s *GetAppTokenServiceResponse) SetStatusCode(v int32) *GetAppTokenServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppTokenServiceResponse) SetBody(v *GetAppTokenServiceResponseBody) *GetAppTokenServiceResponse {
	s.Body = v
	return s
}

func (s *GetAppTokenServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
