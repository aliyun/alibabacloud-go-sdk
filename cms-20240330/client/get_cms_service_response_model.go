// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCmsServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCmsServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCmsServiceResponse
	GetStatusCode() *int32
	SetBody(v *GetCmsServiceResponseBody) *GetCmsServiceResponse
	GetBody() *GetCmsServiceResponseBody
}

type GetCmsServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCmsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCmsServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCmsServiceResponse) GoString() string {
	return s.String()
}

func (s *GetCmsServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCmsServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCmsServiceResponse) GetBody() *GetCmsServiceResponseBody {
	return s.Body
}

func (s *GetCmsServiceResponse) SetHeaders(v map[string]*string) *GetCmsServiceResponse {
	s.Headers = v
	return s
}

func (s *GetCmsServiceResponse) SetStatusCode(v int32) *GetCmsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCmsServiceResponse) SetBody(v *GetCmsServiceResponseBody) *GetCmsServiceResponse {
	s.Body = v
	return s
}

func (s *GetCmsServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
